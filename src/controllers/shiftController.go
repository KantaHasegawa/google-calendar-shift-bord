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
	"github.com/gorilla/mux"
)

type ShiftController struct {
	interactor entity.ShiftInteractorInteface
}

func NewShiftController(DBClient *dynamodb.Client) *ShiftController {
	return &ShiftController{
		interactor: usecase.NewShiftInteractor(repository.NewShiftRepository(DBClient)),
	}
}

func (controller *ShiftController) IndexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	year := vars["year"]
	month := vars["month"]
	table := os.Getenv("TABLE_NAME")

	data, err := controller.interactor.IndexShift(table, user, year, month)
	if err != nil {
		errorHandler.ControllerError(err, &w)
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		errorHandler.ControllerError(err, &w)
		return
	}

	fmt.Fprintf(w, "%s", result)
}

func (controller *ShiftController) NewHandler(w http.ResponseWriter, r *http.Request) {
	type ShiftNewHandlerRequestBody struct {
		User       string `json:"User"`
		StartWork  string `json:"StartWork"`
		FinishWork string `json:"FinishWork"`
		SpotId     string `json:"SpotId"`
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

func (controller *ShiftController) DeleteHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	user := vars["user"]
	startWork := vars["startwork"]
	table := os.Getenv("TABLE_NAME")
	err := controller.interactor.DeleteShift(table, user, startWork)
	if err != nil {
		errorHandler.ControllerError(err, &w)
		return
	}
	fmt.Fprint(w, "Delete Success")
}
