package gdk

import (
	"fmt"
	"math/rand"
)

func (gdkm *gdkMatrixType) Empty() {
	for i, row := range gdkm.gdkMatrix {
		for j := range row {
			gdkm.gdkMatrix[i][j].value = emptyValue
		}
	}
}

func (gdkm *gdkMatrixType) Print() {
	fmt.Print("\n===========================\n")
	for i, row := range gdkm.gdkMatrix {
		fmt.Println()
		for j := range row {
			fmt.Print(" ")
			fmt.Print(gdkm.gdkMatrix[i][j].value)
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Print("===========================\n\n")
}

func genRandom(values []int) (cellRandom int) {
	cellRandom = values[rand.Intn(len(values))]

	return
}

func (gdkm *gdkMatrixType) FillGrid() {
	cellValue := -1
	for i, row := range gdkm.gdkMatrix {
		valueArray = fillValueArray()
		for j := range row {
			//	getColumnAsArray(j, *gdkm)
			if len(valueArray) == 1 {
				cellValue = valueArray[0]
			} else {
				cellValue = genRandom(valueArray)
				//				fmt.Println(valueArray)
				//				fmt.Println(cellValue)
				valueArray = truncateArray(cellValue, valueArray)
				//				fmt.Println(valueArray)

			}
			gdkm.gdkMatrix[i][j].value = cellValue
		}
	}
}

/*func arrayContains(gdkArray [gdkArraySize]gdkCellType, num int) (numExists bool) {
	numExists = false
	for j := range gdkArray {
		if gdkArray[j].value == num {
			numExists = true
		}
	}

	return
}

func getColumnAsArray(j int, gdkm gdkMatrixType) (gdkArray [gdkArraySize]gdkCellType) {
	for i := range gdkm.gdkMatrix[j] {
		gdkArray[i] = gdkm.gdkMatrix[i][j]
	}
	return
}*/

func truncateArray(takeOut int, valueArray []int) (truncatedValueArray []int) {
	for i := range valueArray {
		if valueArray[i] != takeOut {
			truncatedValueArray = append(truncatedValueArray, valueArray[i])
		}
	}

	return
}

func fillValueArray() (valueArray []int) {
	for i := 0; i < gdkArraySize; i++ {
		valueArray = append(valueArray, i+1)
	}

	return
}
