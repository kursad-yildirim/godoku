package gdk

import (
	"fmt"
	"math/rand"
)

func genRandom(values []int) (cellRandom int) {
	cellRandom = values[rand.Intn(len(values))]

	return
}

func removeCurrentRowElements(i, j int, valueArray *[]int, grid *godokuGrid) {
	for subJ := 0; subJ < j; subJ++ {
		*valueArray = truncateArray(grid.mainGrid[i][subJ].value, *valueArray)
	}
}

func removeCurrentColumnElements(i, j int, valueArray *[]int, grid *godokuGrid) {
	for subI := 0; subI < i; subI++ {
		*valueArray = truncateArray(grid.mainGrid[subI][j].value, *valueArray)
	}
}

func removeSubMatrixElements(i, j int, valueArray *[]int, grid *godokuGrid) {
	matrixI := i / gdkDimension
	matrixJ := j / gdkDimension
	for i := 0; i < gdkDimension; i++ {
		for j := 0; j < gdkDimension; j++ {
			*valueArray = truncateArray(grid.mainGrid[matrixI*gdkDimension+i][matrixJ*gdkDimension+j].value, *valueArray)
		}
	}
}

func truncateArray(takeOut int, valueArray []int) (truncatedValueArray []int) {
	for i := range valueArray {
		if valueArray[i] != takeOut {
			truncatedValueArray = append(truncatedValueArray, valueArray[i])
		}
	}

	return
}

func fillArray(s int, isIndex bool) (valueArray []int) {
	for i := 0; i < s; i++ {
		if isIndex {
			valueArray = append(valueArray, i)
		} else {
			valueArray = append(valueArray, i+1)
		}
	}

	return
}

func GenerateJSON(grid godokuGrid) [][]string {
	jsonGrid := make([][]string, len(grid.mainGrid))
	for i := range jsonGrid {
		jsonGrid[i] = make([]string, len(grid.mainGrid[i]))
		for j := range jsonGrid[i] {
			if grid.mainGrid[i][j].display {
				jsonGrid[i][j] = fmt.Sprint(grid.mainGrid[i][j].value)
			} else {
				jsonGrid[i][j] = "_"
			}
		}
	}
	return jsonGrid
}
