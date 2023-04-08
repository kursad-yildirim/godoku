package gdk

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func ReadEnvironment() error {
	var err error
	var present bool
	// define an interface for different data types and try to use a map instead of struct for props
	/*
		for _, myVar := range envVars {
			if myVar["type"] == "int" {
				fmt.Println(myVar["name"] + " is an integer")
			} else if myVar["type"] == "string" {
				fmt.Println(myVar["name"] + " is a string")
			}
		}
	*/
	Props.BlockSize, err = strconv.Atoi(os.Getenv(envVars["blockSize"]["name"]))
	if err != nil {
		gdkError.errMsg = envVars["blockSize"]["name"] + errors["missingEnv"]
		return &gdkError
	}
	Props.Difficulty, err = strconv.Atoi(os.Getenv(envVars["difficulty"]["name"]))
	if err != nil {
		gdkError.errMsg = envVars["difficulty"]["name"] + errors["missingEnv"]
		return &gdkError
	}
	Props.PortNumber, present = os.LookupEnv(envVars["portNumber"]["name"])
	if !present {
		gdkError.errMsg = envVars["portNumber"]["name"] + errors["missingEnv"]
		return &gdkError
	}

	return nil
}

func generateRandomDigit(values []int) (out int) {
	out = values[rand.Intn(len(values))]

	return
}

func removeExistingRowElements(cellRow, cellColumn int, valueArray *[]int, grid gridType) {
	for j := 0; j < cellColumn; j++ {
		*valueArray = truncateArray(grid.members[cellRow][j].value, *valueArray)
	}
}

func removeExistingColumnElements(cellRow, cellColumn int, valueArray *[]int, grid gridType) {
	for i := 0; i < cellRow; i++ {
		*valueArray = truncateArray(grid.members[i][cellColumn].value, *valueArray)
	}
}

func removeExistingBlockElements(cellRow, cellColumn int, valueArray *[]int, grid gridType) {
	blockRow := cellRow / Props.BlockSize
	blockColumn := cellColumn / Props.BlockSize
	for i := 0; i < Props.BlockSize; i++ {
		for j := 0; j < Props.BlockSize; j++ {
			*valueArray = truncateArray(grid.members[blockRow*Props.BlockSize+i][blockColumn*Props.BlockSize+j].value, *valueArray)
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

func fillArray(arraySize int, isIndex bool) (valueArray []int) {
	for i := 0; i < arraySize; i++ {
		if isIndex {
			valueArray = append(valueArray, i)
		} else {
			valueArray = append(valueArray, i+1)
		}
	}

	return
}

func GenerateJSON(grid gridType) [][]string {
	jsonGrid := make([][]string, Props.GridSize)
	for i := range jsonGrid {
		jsonGrid[i] = make([]string, Props.GridSize)
		for j := range jsonGrid[i] {
			if grid.members[i][j].display {
				jsonGrid[i][j] = fmt.Sprint(grid.members[i][j].value)
			} else {
				jsonGrid[i][j] = "_"
			}
		}
	}
	return jsonGrid
}

func (err *errorType) Error() string {
	return fmt.Sprintf("ERROR: " + err.errMsg)
}
