package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/mutant", isMutantHandler).Methods(http.MethodPost)
	log.Print("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)

	log.Fatal(err)
}
