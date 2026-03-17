package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/forecast", getForecast).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getForecast(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Aviation weather forecast: clear skies")
}