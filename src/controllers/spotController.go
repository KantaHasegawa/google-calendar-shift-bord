package spotController

import (
	"fmt"
	"net/http"
)

type SpotControllerInterface interface {}

type SpotController struct {}

func NewSpotController() *SpotController {
	return &SpotController{}
}

func (controller *SpotController) GetAll(w http.ResponseWriter, r *http.Request) {
	response := "spot controller get"
	fmt.Fprint(w, response)
}
