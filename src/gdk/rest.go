package gdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GenerateGrid(w http.ResponseWriter, r *http.Request) {
	if err := ReadEnvironment(); err != nil {
		fmt.Println(err)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&Props)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	grid.FillGrid()
	grid.DigHoles()
	json.NewEncoder(w).Encode(GenerateJSON(grid))
}
