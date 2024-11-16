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

// Init the connections to AWS services with the sdk config.
func init() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	dbClient = *dynamodb.NewFromConfig(sdkConfig)
}

// Generate a string that represents a game id.
func generateGameId() string {
	code := ""
	for i := 0; i < 6; i++ {
		digit := rand.Intn(10)
		code += strconv.Itoa(digit)
	}

	return code
}

// This function is what is run by lambda, the ctx parameter is the context which provides information about the invocation, function, and execution
// the req parameter has information about the requestion that was made
func handleRequest(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (response.Response, error) {
	// Get information from the url query strings about the connection type and the username of the player
	username := req.QueryStringParameters["username"]
	connectionType := req.QueryStringParameters["connectiontype"]

	// create a struct to hold the initial game information based on the connection type will
	var game models.Game
	if connectionType == "create" { // create a new game
		game = models.Game{
			GameId: generateGameId(),
			Players: []models.Player{
				{
					Username:     username,
					ConnectionId: req.RequestContext.ConnectionID,
					Connected:    true,
				},
			},
		}
	} else if connectionType == "join" { // join an existing game
		id := req.QueryStringParameters["id"]

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

	// save the game to dynamodb
	_, err := db.InsertGame(ctx, &dbClient, game)
	if err != nil {
		return response.InternalSeverErrorResponse(), err
	}

	// save connection information to dynamodb
	connection := models.Connection{
		ConnectionId: req.RequestContext.ConnectionID,
		GameId:       game.GameId,
	}
	_, err = db.InsertConnection(ctx, &dbClient, connection)
	if err != nil {
		return response.InternalSeverErrorResponse(), err
	}

	return response.OkResponseWithBody(game.GameId), nil
}

// start the lambda function with the handleRequest() function
func main() {
	lambda.Start(handleRequest)
}
