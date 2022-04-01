package repository

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type ShiftRepository struct {
	DBClient *dynamodb.Client
}

func NewShiftRepository(DBClient *dynamodb.Client) *ShiftRepository {
	return &ShiftRepository{DBClient}
}
