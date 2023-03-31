package gdk

import (
	"encoding/json"
	"net/http"
)

func GenerateGrid(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&props)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	props.Population = props.Size * props.Size

	for !GodokuGrid.FillGrid() {
	}
	for !GodokuGrid.Mask() {
	}
	json.NewEncoder(w).Encode(GenerateJSON(GodokuGrid))
}
