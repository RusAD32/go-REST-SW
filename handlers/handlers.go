package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"SW1/models"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}


func SolveEquation(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Getting a solution...")
	equation := models.GetEquation()
	writer.WriteHeader(200)
	err := json.NewEncoder(writer).Encode(equation)
	if err != nil {
		log.Fatal("can't encode solution, something very wrong")
	}
}

func AddEquation(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Creating new equation ....")
	var equation models.Equation

	err := json.NewDecoder(request.Body).Decode(&equation)
	if err != nil {
		msg := models.Message{Message: "provideed json file is invalid"}
		writer.WriteHeader(400)
		err = json.NewEncoder(writer).Encode(msg)
		if err != nil {
			log.Fatal("can't encode message, something very wrong")
		}
		return
	}
	models.AddEquation(equation)
	writer.WriteHeader(201)
}
