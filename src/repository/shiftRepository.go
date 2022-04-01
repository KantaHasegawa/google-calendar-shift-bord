package repository

import (
	"context"
	"fmt"
	"shiftboard/src/entity"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ShiftRepository struct {
	DBClient *dynamodb.Client
}

func NewShiftRepository(DBClient *dynamodb.Client) *ShiftRepository {
	return &ShiftRepository{DBClient}
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
