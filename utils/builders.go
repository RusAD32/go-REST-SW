package utils

import (
	"SW1/handlers"
	"github.com/gorilla/mux"
)

func BuildBookResource(router *mux.Router) {
	router.HandleFunc("/solve", handlers.SolveEquation).Methods("GET")
	router.HandleFunc("/grab", handlers.AddEquation).Methods("POST")
}

