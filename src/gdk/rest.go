package gdk

import (
	"encoding/json"
	"net/http"
)

func GenerateGrid(w http.ResponseWriter, r *http.Request) {
	displayMasked := true
	for !GodokuGrid.FillGrid() {
	}
	for !GodokuGrid.Mask() {
	}
	GodokuGrid.Print(displayMasked)
	w.Header().Set("Content-Type", "application/json")

	jsonBytes, err := json.Marshal(httpSuccess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonBytes)
}

// post request take difficulty send response
