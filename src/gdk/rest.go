package gdk

import (
	"encoding/json"
	"net/http"
)

func GenerateGrid(w http.ResponseWriter, r *http.Request) {
	displayMasked := true
	for !Gdkm.FillGrid() {
	}
	for !Gdkm.Mask() {
	}
	Gdkm.Print(displayMasked)
	// Set the response header to JSON format
	w.Header().Set("Content-Type", "application/json")

	// Convert the 'people' slice to JSON
	jsonBytes, err := json.Marshal(httpSuccess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Write(jsonBytes)
}
