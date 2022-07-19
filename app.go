package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	r := mux.NewRouter()
	a.Router = r
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Printf("Starting server on http://localhost%s", addr)
	err := http.ListenAndServe(addr, a.Router)
	log.Fatal(err)

}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/mutant", isMutantHandler).Methods(http.MethodPost)
}
