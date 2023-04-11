package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"tripko.local/godoku/gdk"
)

func main() {
	// godoku main
	if err := gdk.ReadEnvironment(); err != nil {
		fmt.Println(err)
		return
	}
	// Define REST interface
	router := mux.NewRouter()
	router.HandleFunc("/", gdk.GenerateGrid).Methods("POST")
	// Start Server
	log.Fatal(http.ListenAndServe(":"+gdk.Props.PortNumber, router))
}
