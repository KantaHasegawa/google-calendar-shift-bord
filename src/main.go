package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"shiftboard/src/routers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
)

var gorillaMuxLambda *gorillamux.GorillaMuxAdapter

func handleHello(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Message string `json:"message"`
	}
	result, err := json.Marshal(response{Message: "hello"})
	if err != nil {
		log.Println(err.Error())
		fmt.Fprint(w, "sorry server error...")
		return
	}
	fmt.Fprintf(w, "%s", result)
}

func init() {
	log.Printf("cold start")
	router := routers.InitRouter()
	router.HandleFunc("/", handleHello)
	gorillaMuxLambda = gorillamux.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return gorillaMuxLambda.ProxyWithContext(ctx, req)
}

func main() {
	if os.Getenv("ENV") == "dev" {
		router := routers.InitRouter()
		router.HandleFunc("/", handleHello)
		http.Handle("/", router)
		http.ListenAndServe(":8070", nil)
	} else {
		lambda.Start(Handler)
	}
}
