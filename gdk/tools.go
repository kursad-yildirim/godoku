package gdk

import (
	"fmt"
	"math/rand"
)

func (gdkm *gdkMatrixType) Empty() {
	emptyArray := [9]int32{emptyValue}
	for iFull, rowFull := range gdkm.gdkMatrix {
		for jFull := range rowFull {
			for iSub, rowSub := range gdkm.gdkMatrix[iFull][jFull].gdkSubMatrix {
				for jSub := range rowSub {
					gdkm.gdkMatrix[iFull][jFull].gdkSubMatrix[iSub][jSub].value = emptyValue
					gdkm.gdkMatrix[iFull][jFull].gdkSubMatrix[iSub][jSub].decided = false
					gdkm.gdkMatrix[iFull][jFull].gdkSubMatrix[iSub][jSub].possibilities = emptyArray
					gdkm.gdkMatrix[iFull][jFull].gdkSubMatrix[iSub][jSub].excluded = emptyArray
				}
			}
		}
	}
}

func (gdkm *gdkMatrixType) Print() {
	fmt.Println("=======================================")
	for iFull, rowFull := range gdkm.gdkMatrix {
		for jFull := range rowFull {
			for iSub, rowSub := range gdkm.gdkMatrix[iFull][jFull].gdkSubMatrix {
				fmt.Print("| ")
				for jSub := range rowSub {
					fmt.Print(" ")
					fmt.Print(gdkm.gdkMatrix[iFull][jFull].gdkSubMatrix[iSub][jSub].value)
					fmt.Print(" ")
				}
				fmt.Print(" |")
			}
			fmt.Println()
		}
		fmt.Println("---------------------------------------")
	}
	//fmt.Println("=======================================")
}

func genRandom(truncatedValueArray []int32) (cellRandom int32) {
	newLen := len(truncatedValueArray) - 1
	cellRandom = truncatedValueArray[rand.Intn(newLen)]

	return
}

// print uzerinden submatrix indexlerini duzelt
func (gdkm *gdkMatrixType) FillGrid() {
	var cellValue int32 = 0
	var truncatedValueArray []int32 = valueArray[:]
	for iFull, rowFull := range gdkm.gdkMatrix {
		for jFull := range rowFull {
			for iSub, rowSub := range gdkm.gdkMatrix[iFull][jFull].gdkSubMatrix {
				for jSub := range rowSub {
					numberFound := false
					for (len(truncatedValueArray) > 1) && (!numberFound) {
						cellValue = genRandom(truncatedValueArray)
						numberFound = gdkm.CheckSubMatrix(iFull, jFull, cellValue)
						fmt.Println(cellValue, numberFound)
						truncatedValueArray = truncatedValueArray[:len(truncatedValueArray)-1]
						truncatedValueArray = truncateValueArray(cellValue, truncatedValueArray)
					}
					gdkm.gdkMatrix[iFull][jFull].gdkSubMatrix[iSub][jSub].value = cellValue
				}
			}
		}
	}
}

func (gdkm *gdkMatrixType) CheckSubMatrix(subI, subJ int, cellValue int32) (numberFit bool) {
	numberFit = true
	for i := range gdkm.gdkMatrix[subI][subJ].gdkSubMatrix {
		for j := range gdkm.gdkMatrix[subI][subJ].gdkSubMatrix[i] {
			if cellValue == gdkm.gdkMatrix[subI][i].gdkSubMatrix[subJ][j].value {
				numberFit = false
			}
		}
	}
	return
}

func truncateValueArray(takeOut int32, valueArray []int32) (shortValueArray []int32) {
	//	fmt.Println(valueArray)
	//	fmt.Println(takeOut)
	for _, num := range valueArray {
		if num != takeOut {
			shortValueArray = append(shortValueArray, num)
			//		fmt.Println(shortValueArray)
		}
	}

	return
}
