package controllers

import (
	"shiftboard/src/entity"
	"shiftboard/src/repository"
	"shiftboard/src/usecase"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ShiftController struct {
	interactor entity.ShiftInteractorInteface
}

func NewShiftController(DBClient *dynamodb.Client) *ShiftController {
	return &ShiftController{
		interactor: usecase.NewShiftInteractor(repository.NewShiftRepository(DBClient)),
	}
}
