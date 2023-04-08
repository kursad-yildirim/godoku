package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"tripko.local/godoku/gdk"
)

func main() {
	gdk.ReadEnvironment()
	// Define REST interface
	router := mux.NewRouter()
	router.HandleFunc("/", gdk.GenerateGrid).Methods("POST")
	// Start Server
	log.Fatal(http.ListenAndServe(":"+gdk.Props.PortNumber, router))
}
