package gdk

import (
	"encoding/json"
	"net/http"
)

func GenerateGrid(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&Props)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Props.GridPopulation = Props.GridSize * Props.GridSize
	grid.FillGrid()
	grid.DigHoles()
	json.NewEncoder(w).Encode(GenerateJSON(grid))
}
