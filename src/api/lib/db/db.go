package db

import (
	"context"
	"errors"
	"fmt"
	"os"

	"trivia-cloud/backend/lib/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go"
)

var (
	DataTableName   = os.Getenv("DATA_TABLE")
	PlayerTableName = os.Getenv("PLAYER_TABLE")
)

// read a game from dynamodb
func GetGame(ctx context.Context, db *dynamodb.Client, gameId string) (*models.Game, error) {
	key, err := attributevalue.Marshal(gameId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(DataTableName),
		Key: map[string]types.AttributeValue{
			"gameId": key,
		},
	}

	result, err := db.GetItem(ctx, input)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	game := new(models.Game)
	err = attributevalue.UnmarshalMap(result.Item, game)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return game, nil
}

// create a game in dynamodb
func InsertGame(ctx context.Context, db *dynamodb.Client, game models.Game) (*models.Game, error) {
	item, err := attributevalue.MarshalMap(game)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(DataTableName),
		Item:      item,
	}

	res, err := db.PutItem(ctx, input)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = attributevalue.UnmarshalMap(res.Attributes, &game)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &game, nil
}

// update a game in dynamodb
func UpdateGame(ctx context.Context, db *dynamodb.Client, gameId string, updateGame models.Game) (*models.Game, error) {
	key, err := attributevalue.Marshal(gameId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	expr, err := expression.NewBuilder().WithUpdate(
		expression.Set(
			expression.Name("players"),
			expression.Value(updateGame.Players),
		),
	).WithCondition(
		expression.Equal(
			expression.Name("gameId"),
			expression.Value(gameId),
		),
	).Build()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	input := &dynamodb.UpdateItemInput{
		Key: map[string]types.AttributeValue{
			"id": key,
		},
		TableName:                 aws.String(DataTableName),
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ConditionExpression:       expr.Condition(),
		ReturnValues:              types.ReturnValue(*aws.String("ALL_NEW")),
	}

	res, err := db.UpdateItem(ctx, input)
	if err != nil {
		var smErr *smithy.OperationError
		if errors.As(err, &smErr) {
			var condCheckFailed *types.ConditionalCheckFailedException
			if errors.As(err, &condCheckFailed) {
				return nil, nil
			}
		}

		fmt.Println(nil, err)
		return nil, err
	}

	if res.Attributes == nil {
		return nil, nil
	}

	game := new(models.Game)
	err = attributevalue.UnmarshalMap(res.Attributes, game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

// get a connection from the player table
func GetConnection(ctx context.Context, db *dynamodb.Client, connectionId string) (*models.Connection, error) {
	key, err := attributevalue.Marshal(connectionId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(PlayerTableName),
		Key: map[string]types.AttributeValue{
			"connectionId": key,
		},
	}

	result, err := db.GetItem(ctx, input)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	connection := new(models.Connection)
	err = attributevalue.UnmarshalMap(result.Item, connection)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return connection, nil
}

// create a connection in the player table
func InsertConnection(ctx context.Context, db *dynamodb.Client, connection models.Connection) (*models.Connection, error) {
	item, err := attributevalue.MarshalMap(connection)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(PlayerTableName),
		Item:      item,
	}

	res, err := db.PutItem(ctx, input)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = attributevalue.UnmarshalMap(res.Attributes, &connection)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &connection, nil
}
