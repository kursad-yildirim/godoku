package gdk

import (
	"math/rand"
)

func genRandom(values []int) (cellRandom int) {
	cellRandom = values[rand.Intn(len(values))]

	return
}

func removeCurrentRowElements(i, j int, valueArray *[]int, gdkm *gdkMatrixType) {
	for subJ := 0; subJ < j; subJ++ {
		*valueArray = truncateArray(gdkm.gdkMatrix[i][subJ].value, *valueArray)
	}
}

func removeCurrentColumnElements(i, j int, valueArray *[]int, gdkm *gdkMatrixType) {
	for subI := 0; subI < i; subI++ {
		*valueArray = truncateArray(gdkm.gdkMatrix[subI][j].value, *valueArray)
	}
}

func removeSubMatrixElements(i, j int, valueArray *[]int, gdkm *gdkMatrixType) {
	matrixI := i / gdkDimension
	matrixJ := j / gdkDimension
	for i := 0; i < gdkDimension; i++ {
		for j := 0; j < gdkDimension; j++ {
			*valueArray = truncateArray(gdkm.gdkMatrix[matrixI*gdkDimension+i][matrixJ*gdkDimension+j].value, *valueArray)
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
