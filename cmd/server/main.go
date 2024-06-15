package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"restapi/handlers"
)

func main() {
	r := mux.NewRouter()
	h := handlers.NewHandler()

	r.HandleFunc("/objects", h.GetObjects).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}