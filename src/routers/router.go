package routers

import (
	"os"
	"shiftboard/src/controllers"
	"shiftboard/src/database"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	DBClient := database.NewDatabaseClient(os.Getenv("ENV"))
	router := mux.NewRouter()
	spotController := spotController.NewSpotController(DBClient)
	router.HandleFunc("/spot/{user}/{spotid}", spotController.ShowHandler)

	return router
}
