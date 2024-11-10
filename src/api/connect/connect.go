package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"

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

func generateGameId() string {
	code := ""
	for i := 0; i < 6; i++ {
		digit := rand.Intn(10)
		code += strconv.Itoa(digit)
	}

	return code
}

func handleRequest(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (response.Response, error) {
	username := req.QueryStringParameters["username"]
	connectionType := req.QueryStringParameters["connectiontype"]

	var id string
	var game models.Game
	if connectionType == "create" {
		id = generateGameId()

		game = models.Game{
			GameId: id,
			Players: []models.Player{
				{
					Username:     username,
					ConnectionId: req.RequestContext.ConnectionID,
					Connected:    true,
				},
			},
		}
	} else if connectionType == "join" {
		id = req.QueryStringParameters["id"]

		res, err := db.GetGame(ctx, &dbClient, id)
		if res == nil || err != nil {
			fmt.Println(res, err)
			return response.InternalSeverErrorResponse(), nil
		}
		game = *res

		player := models.Player{
			Username:     username,
			ConnectionId: req.RequestContext.ConnectionID,
			Connected:    true,
		}
		game.Players = append(game.Players, player)
	}

	_, err := db.InsertGame(ctx, &dbClient, game)
	if err != nil {
		return response.InternalSeverErrorResponse(), err
	}

	connection := models.Connection{
		ConnectionId: req.RequestContext.ConnectionID,
		GameId:       id,
	}
	_, err = db.InsertConnection(ctx, &dbClient, connection)
	if err != nil {
		return response.InternalSeverErrorResponse(), err
	}

	return response.OkReponse(), nil
}

func main() {
	lambda.Start(handleRequest)
}
