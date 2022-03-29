package spotController

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

type SpotController struct {
	interactor entity.SpotInteractorInteface
}

func NewSpotController(DBClient *dynamodb.Client) *SpotController {
	return &SpotController{
		interactor: usecase.NewSpotInteractor(repository.NewSpotRepository(DBClient)),
	}
}

func (controller *SpotController) ShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	spotId := vars["spotid"]
	startWork := "spot"
	table := os.Getenv("TABLE_NAME")

	data, err := controller.interactor.DetailSpot(table, user, startWork, spotId)

	if err != nil {
		errorHandler.ControllerError(err, &w)
		return
	}

	result, err := json.Marshal(data)

	if err != nil {
		errorHandler.ControllerError(err, &w)
		return
	}

	fmt.Fprintf(w, "%s\n", result)
}

func (controller *SpotController) NewHandler(w http.ResponseWriter, r *http.Request) {
	type SpotCreateHandlerRequestBody struct {
		User      string `json:"User"`
		Name      string `json:"Name"`
		Salaly    int    `json:"Salaly"`
		CutOffDay string `json:"CutOffDay"`
		PayDay    string `json:"PayDay"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorHandler.ControllerError(err, &w)
		return
	}

	spotNewHandlerRequestBody := SpotCreateHandlerRequestBody{}
	err = json.Unmarshal(body, &spotNewHandlerRequestBody)
	if err != nil {
		errorHandler.ControllerError(err, &w)
		return
	}

	user := spotNewHandlerRequestBody.User
	name := spotNewHandlerRequestBody.Name
	salaly := spotNewHandlerRequestBody.Salaly
	cutOffDay := spotNewHandlerRequestBody.CutOffDay
	payDay := spotNewHandlerRequestBody.PayDay
	table := os.Getenv("TABLE_NAME")

	data, err := controller.interactor.CreateSpot(table, user, name, salaly, cutOffDay, payDay)
	if err != nil {
		errorHandler.ControllerError(err, &w)
		return
	}
	fmt.Fprintf(w, "%s\n", data)
}
