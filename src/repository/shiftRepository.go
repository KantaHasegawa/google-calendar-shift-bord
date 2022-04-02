package repository

import (
	"context"
	"fmt"
	"shiftboard/src/entity"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ShiftRepository struct {
	DBClient *dynamodb.Client
}

func NewShiftRepository(DBClient *dynamodb.Client) *ShiftRepository {
	return &ShiftRepository{DBClient}
}

func (repository *ShiftRepository) Get(table string, user string, year string, month string)([]entity.TShift, error){
	item := []entity.TShift{}
	data, err := repository.DBClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                &table,
		KeyConditionExpression:   aws.String("#U=:userval and begins_with(StartWork, :startworkval)"),
		ExpressionAttributeNames: map[string]string{"#U": "User"},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":userval":      &types.AttributeValueMemberS{Value: user},
			":startworkval": &types.AttributeValueMemberS{Value: year + month},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return item, err
	}

	if data.Items == nil {
		return item, err
	}

	err = attributevalue.UnmarshalListOfMaps(data.Items, &item)
	if err != nil {
		return item, err
	}

	return item, nil
}


func (repository *ShiftRepository) Post(table string, user string, startWork string, finishWork string, spotId string)(error){

		params := entity.TShift{
			User: user,
			StartWork: startWork,
			FinishWork: finishWork,
			SpotId: spotId,
		}

	attributeValue, err := attributevalue.MarshalMap(params)

	if err != nil {
		return fmt.Errorf("MarshalParams: %v", err)
	}

	_, err = repository.DBClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: &table,
		Item:      attributeValue,
	})

	if err != nil {
		return fmt.Errorf("PutItem: %v", err)
	}

	return nil
}

func (repository *ShiftRepository) Delete(table string, user string, startWork string)(error){
		_, err := repository.DBClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"User": &types.AttributeValueMemberS{Value: user},
			"StartWork": &types.AttributeValueMemberS{Value: startWork},
		},
	})
	if err != nil {
		return fmt.Errorf("DeleteItem: %v", err)
	}
	return nil
}
