package gdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func GenerateGrid(w http.ResponseWriter, r *http.Request) {
	var err error
	props.GridSize, err = strconv.Atoi(os.Getenv("GODOKU_SIZE"))
	if err != nil {
		fmt.Println("GODOKU_SIZE not found in the environment")
		return
	}
	props.Difficulty, err = strconv.Atoi(os.Getenv("GODOKU_DIFFICULTY"))
	if err != nil {
		fmt.Println("GODOKU_DIFFICULTY not found in the environment")
		return
	}
	err = json.NewDecoder(r.Body).Decode(&props)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	props.GridPopulation = props.GridSize * props.GridSize

	for !grid.FillGrid() {
	}
	for !grid.Mask() {
	}
	json.NewEncoder(w).Encode(GenerateJSON(grid))
}
