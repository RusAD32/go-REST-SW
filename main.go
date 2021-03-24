package main

import (
	"log"
	"net/http"
	"os"

	"SW1/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	port                    string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not find .env file:", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()

	utils.BuildBookResource(router)

	log.Println("Router initalizing successfully. Ready to go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
