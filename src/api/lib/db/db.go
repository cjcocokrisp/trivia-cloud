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
	TableName = os.Getenv("DATA_TABLE")
)

func GetItem(ctx context.Context, db *dynamodb.Client, gameID int) (*models.Game, error) {
	key, err := attributevalue.Marshal(gameID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"gameID": key,
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

func InsertItem(ctx context.Context, db *dynamodb.Client, game models.Game) (*models.Game, error) {
	item, err := attributevalue.MarshalMap(game)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
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

func UpdateItem(ctx context.Context, db *dynamodb.Client, gameID int, updateGame models.Game) (*models.Game, error) {
	key, err := attributevalue.Marshal(gameID)
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
			expression.Name("gameID"),
			expression.Value(gameID),
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
		TableName:                 aws.String(TableName),
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
