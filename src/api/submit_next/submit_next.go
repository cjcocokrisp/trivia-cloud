package main

import (
	"context"
	"encoding/json"
	"log"

	"trivia-cloud/backend/lib/apigw"
	"trivia-cloud/backend/lib/db"
	"trivia-cloud/backend/lib/models"
	"trivia-cloud/backend/lib/response"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	//"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	//"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	dbClient  dynamodb.Client
	apiClient apigatewaymanagementapi.Client
)

// Init the connections to AWS services with the sdk config.
func init() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	dbClient = *dynamodb.NewFromConfig(sdkConfig)
	apiClient = *apigatewaymanagementapi.NewFromConfig(sdkConfig)
}

// This function is what is run by lambda, the ctx parameter is the context which provides information about the invocation, function, and execution
// the req parameter has information about the requestion that was made
func handleRequest(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (response.Response, error) {
	// retrieve information about the players connection and pull the game that they are connected to
	connection, err := db.GetConnection(ctx, &dbClient, req.RequestContext.ConnectionID)
	if connection != nil {
		response.InternalSeverErrorResponse()
	}

	game, err := db.GetGame(ctx, &dbClient, connection.GameId)
	if err != nil {
		return response.InternalSeverErrorResponse(), nil
	}

	var player models.Player
	for _, element := range game.Players {
		if element.ConnectionId == req.RequestContext.ConnectionID {
			player = element
			break
		}
	}

	player.Submitted = true

	for index, element := range game.Players {
		if element.ConnectionId == req.RequestContext.ConnectionID {
			game.Players[index] = player
			break
		}
	}

	allSubmitted := true
	for _, element := range game.Players {
		if element.Connected && !element.Submitted {
			allSubmitted = false
			break
		}
	}

	res := models.Message[bool]{
		Type:    "next_submission",
		Content: allSubmitted,
	}

	encodedRes, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.InsertGame(ctx, &dbClient, *game)
	if err != nil {
		log.Fatal(err)
	}

	endpoint := apigw.ResolveApiEndpoint(&apiClient, req.RequestContext.DomainName, req.RequestContext.Stage)
	_, err = endpoint.PostToConnection(ctx, &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: &req.RequestContext.ConnectionID,
		Data:         encodedRes,
	})
	if err != nil {
		log.Fatal(err)
	}

	return response.OkReponse(), nil
}

// start the lambda function with the handleRequest() function
func main() {
	lambda.Start(handleRequest)
}
