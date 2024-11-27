package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

	"trivia-cloud/backend/lib/db"
	"trivia-cloud/backend/lib/models"
	"trivia-cloud/backend/lib/response"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go-v2/config"
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
		questionNum := req.QueryStringParameters["questionnum"]
		categoryNum := req.QueryStringParameters["categorynum"]
		categoryName := req.QueryStringParameters["categoryname"]

		// prepare url for api request
		apiURL := url.URL{
			Scheme: "https",
			Host:   "opentdb.com",
			Path:   "api.php",
		}
		//var query url.Values
		query := apiURL.Query()
		query.Add("amount", questionNum)
		if categoryNum != "none" {
			query.Add("category", categoryNum)
		}
		query.Add("type", "multiple")
		apiURL.RawQuery = query.Encode()

		fmt.Println(apiURL.String())
		// get questions for the game
		res, err := http.Get(apiURL.String())
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		var questions models.OpenTriviaDBResponse
		json.Unmarshal(data, &questions)

		game = models.Game{
			GameId:          generateGameId(),
			Started:         false,
			CurrentQuestion: 0,
			NumQuestions:    questionNum,
			Questions:       questions.Results,
			Category:        categoryName,
			Players: []models.Player{
				{
					Username:     username,
					ConnectionId: req.RequestContext.ConnectionID,
					Connected:    true,
					Score:        0,
					Submitted:    false,
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
			Score:        0,
			Submitted:    false,
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
