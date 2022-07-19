package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	code := m.Run()

	os.Exit(code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func TestIsMutantDiagonal(t *testing.T) {

	jsonStr := []byte(`{"dna": ["ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"]}`)
	req, _ := http.NewRequest("POST", "/mutant", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestIsMutantHorizontal(t *testing.T) {

	jsonStr := []byte(`{"dna": ["TTGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"]}`)
	req, _ := http.NewRequest("POST", "/mutant", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestIsMutantVerticual(t *testing.T) {

	jsonStr := []byte(`{"dna": ["TTGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCATA", "TCACTG"]}`)
	req, _ := http.NewRequest("POST", "/mutant", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestIsNotMutant(t *testing.T) {
	jsonStr := []byte(`{"dna": ["TTGGTA", "CAGTGC", "TTATGT", "AGAAGG", "ACCCTA", "TCACTG"]}`)
	req, _ := http.NewRequest("POST", "/mutant", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusForbidden, response.Code)
}

func TestIsNotMutantInvaildDNA(t *testing.T) {
	jsonStr := []byte(`{"dna": ["ZTGGTA", "CAGTGC", "TTATGT", "AGAAGG", "ACCCTA", "TCACTG"]}`)
	req, _ := http.NewRequest("POST", "/mutant", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusForbidden, response.Code)
}
