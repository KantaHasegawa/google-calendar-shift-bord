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

type SpotRepository struct {
	DBClient *dynamodb.Client
}

func NewSpotRepository(DBClient *dynamodb.Client) *SpotRepository {
	return &SpotRepository{DBClient}
}

func (repository *SpotRepository) GetAll(table string, user string) ([]entity.TSpot, error) {
	item := []entity.TSpot{}
	data, err := repository.DBClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                &table,
		KeyConditionExpression:   aws.String("#U=:userval and begins_with(StartWork, :startworkval)"),
		ExpressionAttributeNames: map[string]string{"#U": "User"},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":userval":      &types.AttributeValueMemberS{Value: user},
			":startworkval": &types.AttributeValueMemberS{Value: "spot_"},
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

func (repository *SpotRepository) Get(table string, user string, startWork string) (entity.TSpot, error) {
	item := entity.TSpot{}

	data, err := repository.DBClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"User":      &types.AttributeValueMemberS{Value: user},
			"StartWork": &types.AttributeValueMemberS{Value: startWork},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return item, err
	}

	if data.Item == nil {
		return item, err
	}

	err = attributevalue.UnmarshalMap(data.Item, &item)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (repository *SpotRepository) Post(table string, user string, spotId string, name string, salaly int, cutOffDay string, payDay string) (string, error) {

	params := entity.TSpot{
		User:      user,
		StartWork: "spot_" + spotId,
		SpotId:    spotId,
		SpotData: entity.TSpotData{
			Name:      name,
			Salary:    salaly,
			CutOffDay: cutOffDay,
			PayDay:    payDay,
		},
	}

	attributeValue, err := attributevalue.MarshalMap(params)

	if err != nil {
		return "", fmt.Errorf("MarshalParams: %v", err)
	}

	_, err = repository.DBClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: &table,
		Item:      attributeValue,
	})

	if err != nil {
		return "", fmt.Errorf("PutItem: %v", err)
	}

	return spotId, nil
}

func (repository *SpotRepository) Delete(table string, user string, startWork string, spotId string)(string, error){
		_, err := repository.DBClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"User": &types.AttributeValueMemberS{Value: user},
			"StartWork": &types.AttributeValueMemberS{Value: startWork + "_" + spotId},
		},
	})
	if err != nil {
		return "", fmt.Errorf("DeleteItem: %v", err)
	}
	return spotId, nil
}
