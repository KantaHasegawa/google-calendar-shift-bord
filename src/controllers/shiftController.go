package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"shiftboard/src/entity"
	"shiftboard/src/errorHandler"
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

func (controller *ShiftController) NewHandler(w http.ResponseWriter, r *http.Request){
	type ShiftNewHandlerRequestBody struct {
		User string `json:"User"`
		StartWork string `json:"StartWork"`
		FinishWork string `json:"FinishWork"`
		SpotId string `json:"SpotId"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorHandler.ControllerError(err, &w)
		return
	}

	spotNewHandlerRequestBody := ShiftNewHandlerRequestBody{}
	err = json.Unmarshal(body, &spotNewHandlerRequestBody)
	if err != nil {
		errorHandler.ControllerError(err, &w)
		return
	}

	user := spotNewHandlerRequestBody.User
	startWork := spotNewHandlerRequestBody.StartWork + "_" + spotNewHandlerRequestBody.SpotId
	finishWork := spotNewHandlerRequestBody.FinishWork
	spotId := spotNewHandlerRequestBody.SpotId
	table := os.Getenv("TABLE_NAME")

	err = controller.interactor.CreateShift(table, user, startWork, finishWork, spotId)
	if err != nil {
		errorHandler.ControllerError(err, &w)
		return
	}

	fmt.Fprint(w, "Create Success")
}
