package spotController

import (
	"encoding/json"
	"fmt"
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

	if err != nil{
		errorHandler.ControllerError(err, &w)
		return
	}

	result, err := json.Marshal(data)

	if err != nil{
		errorHandler.ControllerError(err, &w)
		return
	}

	fmt.Fprintf(w, "%s\n", result)
}
