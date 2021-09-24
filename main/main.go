package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	router := routes.getRouter()
	methods := handlers.AllowedMethods([]string{"GET"})
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(methods)(router)))
}
