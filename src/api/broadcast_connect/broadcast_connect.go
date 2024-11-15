package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"sync"

	"trivia-cloud/backend/lib/db"
	"trivia-cloud/backend/lib/response"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"

	//"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"

	//"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	dbClient  dynamodb.Client
	apiClient apigatewaymanagementapi.Client
	region    *string
)

func init() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	dbClient = *dynamodb.NewFromConfig(sdkConfig)
	apiClient = *apigatewaymanagementapi.NewFromConfig(sdkConfig)
	region = aws.String(sdkConfig.Region)

}

func handleRequest(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (response.Response, error) {
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

	var newUser string
	for _, element := range game.Players {
		if element.ConnectionId == req.RequestContext.ConnectionID {
			newUser = element.Username
		}
	}

	var endpoint url.URL
	endpoint.Scheme = "https"
	endpoint.Host = req.RequestContext.DomainName
	endpoint.Path = req.RequestContext.Stage
	endpointStr := aws.String(endpoint.String())

	cp := apiClient.Options().Copy()
	cp.BaseEndpoint = endpointStr
	endpointClient := apigatewaymanagementapi.New(cp)

	var wg sync.WaitGroup
	for i := 0; i < len(game.Players); i++ {
		wg.Add(1)

		go func(client string) {
			defer wg.Done()
			if client != req.RequestContext.ConnectionID {
				_, err := endpointClient.PostToConnection(ctx, &apigatewaymanagementapi.PostToConnectionInput{
					Data:         []byte(newUser + " has connected!"),
					ConnectionId: aws.String(client),
				})

				if err != nil {
					log.Fatal(err)
				}
			}

		}(game.Players[i].ConnectionId)

	}
	wg.Wait()

	return response.OkReponse(), nil
}

func main() {
	lambda.Start(handleRequest)
}
