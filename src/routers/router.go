package routers

import (
	"shiftboard/src/controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	spotController := spotController.NewSpotController()
	router.HandleFunc("/spots", spotController.GetAll)

	return router
}
