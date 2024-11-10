package main

import (
	"context"
	"log"

	"trivia-cloud/backend/lib/db"
	"trivia-cloud/backend/lib/models"
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
	// TODO: add submitting on connection
	game := models.Game{
		GameID: 100,
		Players: []models.Player{
			{
				Username:     "Testing",
				ConnectionID: req.RequestContext.ConnectionID,
				Connected:    true,
			},
		},
	}

	_, err := db.InsertItem(ctx, &dbClient, game)
	if err != nil {
		return response.InternalSeverErrorResponse(), err
	}

	return response.OkReponse(), nil
}

func main() {
	lambda.Start(handleRequest)
}
