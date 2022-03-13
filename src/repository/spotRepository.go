package repository

import (
	"context"
	"fmt"

	"shiftboard/src/entity"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

type SpotRepositoryInterface interface {
	Get(table string, user string, startWork string) (entity.TSpot, error)
}

type SpotRepository struct {
	DBClient *dynamodb.Client
}

func (repository *SpotRepository) Get(table string, user string, startWork string) (entity.TSpot, error) {
	item := entity.TSpot{}

	data, err := repository.DBClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: &table,
		Key: map[string]types.AttributeValue{
			"User": &types.AttributeValueMemberS{Value: user},
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
