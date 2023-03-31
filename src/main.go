package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"tripko.local/godoku/gdk"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", gdk.GenerateGrid).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
