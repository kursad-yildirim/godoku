package main

import (
	"log"
	"net/http"

	"tripko.local/godoku/gdk"
)

func main() {

	http.HandleFunc("/generate", gdk.GenerateGrid)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
