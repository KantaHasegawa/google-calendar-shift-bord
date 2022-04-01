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
	spotController := controllers.NewSpotController(DBClient)
	shiftController := controllers.NewShiftController(DBClient)
	router.HandleFunc("/spots/{user}", spotController.IndexHandler).Methods("GET")
	router.HandleFunc("/spot/show/{user}/{spotid}", spotController.ShowHandler).Methods("GET")
	router.HandleFunc("/spot/new", spotController.NewHandler).Methods("POST")
	router.HandleFunc("/spot/delete/{user}/{spotid}", spotController.DeleteHandler).Methods("DELETE")
	router.HandleFunc("/shift/new", shiftController.NewHandler).Methods("POST")
	return router
}
