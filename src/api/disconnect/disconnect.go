package main

import (
	"context"
	"log"

	//"trivia-cloud/backend/lib/db"
	//"trivia-cloud/backend/lib/models"
	"trivia-cloud/backend/lib/db"
	"trivia-cloud/backend/lib/response"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	//"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	//"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	dbClient dynamodb.Client
)

func init() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	dbClient = *dynamodb.NewFromConfig(sdkConfig)
}

func handleRequest(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (response.Response, error) {
	connection, err := db.GetConnection(ctx, &dbClient, req.RequestContext.ConnectionID)
	if connection != nil {
		response.InternalSeverErrorResponse()
	}

	game, err := db.GetGame(ctx, &dbClient, connection.GameId)
	if err != nil {
		return response.InternalSeverErrorResponse(), nil
	}

	for index, element := range game.Players {
		if element.ConnectionId == req.RequestContext.ConnectionID {
			game.Players[index].Connected = false
		}
	}

	_, err = db.InsertGame(ctx, &dbClient, *game)
	if err != nil {
		return response.InternalSeverErrorResponse(), nil
	}

	return response.OkReponse(), nil
}

func main() {
	lambda.Start(handleRequest)
}
