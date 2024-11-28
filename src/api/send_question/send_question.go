package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"sync"

	//"trivia-cloud/backend/lib/db"
	//"trivia-cloud/backend/lib/models"
	"trivia-cloud/backend/lib"
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

	for index := range game.Players {
		game.Players[index].Submitted = false
		game.Players[index].Correct = false
	}

	numQuestions, err := strconv.Atoi(game.NumQuestions)
	if err != nil {
		log.Fatal(err)
	}

	var gameover bool
	game.CurrentQuestion++
	if game.CurrentQuestion == numQuestions {
		gameover = true
	}

	_, err = db.InsertGame(ctx, &dbClient, *game)
	if err != nil {
		log.Fatal(err)
	}

	endpointClient := apigw.ResolveApiEndpoint(&apiClient, req.RequestContext.DomainName, req.RequestContext.Stage)
	var wg sync.WaitGroup
	for i := 0; i < len(game.Players); i++ {
		wg.Add(1)
		go func(client string, connected bool) {
			defer wg.Done()
			if connected {
				var message any
				if gameover {
					message = models.Message[string]{
						Type:    "gameover",
						Content: "TODO: SEND SORTED LIST OF PLAYERS BY SCORE TO SHOW WINNER",
					}
				} else {
					question := lib.PrepareQuestionInfo(*game)
					message = models.Message[models.NextQuestion]{
						Type: "next_question",
						Content: models.NextQuestion{
							QuestionNum:         game.CurrentQuestion,
							QuestionInformation: question,
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
