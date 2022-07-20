package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	// Initialize connection constants.
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "postgres"
	DBNAME   = "mutants"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) DBConnection(user, password, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db, nil
}

func (a *App) Initialize(user, password, dbname string) {
	r := mux.NewRouter()
	a.Router = r
	a.initializeRoutes()
	db, _ := a.DBConnection(user, password, dbname)
	a.DB = db
}

func (a *App) Run(addr string) {
	log.Printf("Starting server on http://localhost%s", addr)
	err := http.ListenAndServe(addr, a.Router)
	log.Fatal(err)

}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/mutant", isMutantHandler).Methods(http.MethodPost)
}
