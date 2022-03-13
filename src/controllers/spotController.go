package spotController

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"shiftboard/src/repository"
	"shiftboard/src/errorHandler"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gorilla/mux"
)

type SpotControllerInterface interface {}

type SpotController struct {
	repository repository.SpotRepositoryInterface
}

func NewSpotController(DBClient *dynamodb.Client) *SpotController {
	return &SpotController{
		repository: &repository.SpotRepository{DBClient: DBClient},
	}
}

func (controller *SpotController) ShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	spotId := vars["spotid"]
	startWork := "spot"
	table := os.Getenv("TABLE_NAME")

	data, err := controller.repository.Get(table, user, startWork + "_" + spotId)

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
