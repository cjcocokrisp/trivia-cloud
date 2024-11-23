package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"trivia-cloud/backend/lib/apigw"
	"trivia-cloud/backend/lib/db"
	"trivia-cloud/backend/lib/models"
	"trivia-cloud/backend/lib/response"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"

	//"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"

	"github.com/aws/aws-sdk-go-v2/config"

	//"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
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
	// Get infromation about the players connection and the game they are connected to
	connection, err := db.GetConnection(ctx, &dbClient, req.RequestContext.ConnectionID)
	if err != nil {
		fmt.Println(err)
		return response.InternalSeverErrorResponse(), nil
	}

	game, err := db.GetGame(ctx, &dbClient, connection.GameId)
	if err != nil {
		fmt.Println(err)
		return response.InternalSeverErrorResponse(), nil
	}

	// Get the username of the player by searching by connection id and compile list of users
	var users []string
	var newUser string
	for _, element := range game.Players {
		if element.ConnectionId == req.RequestContext.ConnectionID {
			newUser = element.Username
		}
		if element.Connected {
			users = append(users, element.Username)
		}
	}

	// connect to endpoint
	endpointClient := apigw.ResolveApiEndpoint(&apiClient, req.RequestContext.DomainName, req.RequestContext.Stage)

	// send information about the player connecting to all players that are connected
	var wg sync.WaitGroup
	for i := 0; i < len(game.Players); i++ {
		wg.Add(1)

		go func(client string, connected bool) {
			defer wg.Done()
			var message any
			if client != req.RequestContext.ConnectionID && connected {
				// encode new user and send to message to client
				message = models.Message[string]{
					Type:    "new_connection",
					Content: newUser,
				}
			} else if connected {
				// this will be the new player so return the player list and game id
				message = models.Message[models.GameInformation]{
					Type: "connected",
					Content: models.GameInformation{
						GameId:       game.GameId,
						Players:      users,
						NumQuestions: game.NumQuestions,
						Category:     game.Category,
					},
				}
			}
			encodedMessage, err := json.Marshal(message)
			if err != nil {
				log.Fatal(err)
			}

			_, err = endpointClient.PostToConnection(ctx, &apigatewaymanagementapi.PostToConnectionInput{
				ConnectionId: &client,
				Data:         encodedMessage,
			})
			if err != nil {
				log.Fatal(err)
			}

		}(game.Players[i].ConnectionId, game.Players[i].Connected)

	}
	wg.Wait()

	return response.OkReponse(), nil
}

// start the lambda function with the handleRequest() function
func main() {
	lambda.Start(handleRequest)
}
