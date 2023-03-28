package gdk

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func (gdkm *gdkMatrixType) FillGrid() (finished bool) {
	finished = true
	gdkm.initGDK()
	cellValue := -1
	for i, row := range gdkm.gdkMatrix {
		for j := range row {
			valueArray = fillArray(gdkArraySize, false)
			if len(valueArray) == 1 {
				cellValue = valueArray[0]
			} else {
				removeCurrentRowElements(i, j, &valueArray, gdkm)
				removeCurrentColumnElements(i, j, &valueArray, gdkm)
				removeSubMatrixElements(i, j, &valueArray, gdkm)
				if len(valueArray) < 1 {
					finished = false
					return
				}
				cellValue = genRandom(valueArray)
			}
			gdkm.gdkMatrix[i][j].value = cellValue
		}
	}

	return
}

func (gdkm *gdkMatrixType) initGDK() {
	gdkArraySize, _ = strconv.Atoi(os.Getenv("GODOKU_SIZE"))
	difficulty, _ = strconv.Atoi(os.Getenv("GODOKU_DIFFICULTY"))
	gridNumTotal = gdkArraySize * gdkArraySize
	gdkDimension = int(math.Sqrt(float64(gdkArraySize)))
	displayCount = gridNumTotal * difficulty / 100
	emptyCell := gdkCellType{
		value:   0,
		display: false,
	}
	gdkm.gdkMatrix = gdkm.gdkMatrix[:0]
	var emptyRow []gdkCellType
	for i := 0; i < gdkArraySize; i++ {
		for j := 0; j < gdkArraySize; j++ {
			emptyRow = append(emptyRow, emptyCell)
		}
		gdkm.gdkMatrix = append(gdkm.gdkMatrix, emptyRow)
		emptyRow = nil
	}
}

func (gdkm *gdkMatrixType) Print(masked bool) {
	for i, row := range gdkm.gdkMatrix {
		fmt.Println()
		for j := range row {
			fmt.Print(" ")
			if masked && !gdkm.gdkMatrix[i][j].display {
				fmt.Print("_")
			} else {
				fmt.Print(gdkm.gdkMatrix[i][j].value)
			}
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n")
}

func (gdkm *gdkMatrixType) Mask() {
	indexArray = fillArray(gridNumTotal, true)
	i, j := 0, 0
	for d := 0; d < displayCount; d++ {
		indexDispArray = append(indexDispArray, genRandom(indexArray))
		indexArray = truncateArray(indexDispArray[d], indexArray)
		i = indexDispArray[d] / gdkArraySize
		j = indexDispArray[d] - i*gdkArraySize
		gdkm.gdkMatrix[i][j].display = true
	}
}
