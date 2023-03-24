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
	fmt.Println("\n===========================")
	for i, row := range gdkm.gdkMatrix {
		fmt.Println()
		for j := range row {
			fmt.Print(" ")
			fmt.Print(gdkm.gdkMatrix[i][j].value)
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println("===========================\n")
}

func genRandom(truncatedValueArray []int) (cellRandom int) {
	newLen := len(truncatedValueArray) - 1
	cellRandom = truncatedValueArray[rand.Intn(newLen)]

	return
}

func (gdkm *gdkMatrixType) FillGrid() {
	var truncatedValueArray []int = valueArray[:]
	for i, row := range gdkm.gdkMatrix {
		for j := range row {
			gdkm.gdkMatrix[i][j].value = genRandom(truncatedValueArray)
		}
	}
}
