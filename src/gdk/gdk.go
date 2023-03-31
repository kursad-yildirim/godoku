package gdk

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func (grid *godokuGrid) FillGrid() (finished bool) {
	rand.Seed(time.Now().UnixNano())
	finished = true
	grid.initGDK()
	cellValue := -1
	for i, row := range grid.mainGrid {
		for j := range row {
			valueArray = fillArray(gdkArraySize, false)
			if len(valueArray) == 1 {
				cellValue = valueArray[0]
			} else {
				removeCurrentRowElements(i, j, &valueArray, grid)
				removeCurrentColumnElements(i, j, &valueArray, grid)
				removeSubMatrixElements(i, j, &valueArray, grid)
				if len(valueArray) < 1 {
					finished = false
					return
				}
				cellValue = genRandom(valueArray)
			}
			grid.mainGrid[i][j].value = cellValue
		}
	}

	return
}

func (grid *godokuGrid) initGDK() {
	gdkArraySize, _ = strconv.Atoi(os.Getenv("GODOKU_SIZE"))
	difficulty, _ = strconv.Atoi(os.Getenv("GODOKU_DIFFICULTY"))
	httpSuccess.Difficulty = difficulty
	gridNumTotal = gdkArraySize * gdkArraySize
	gdkDimension = int(math.Sqrt(float64(gdkArraySize)))
	displayCount = gridNumTotal * difficulty / 100
	emptyCell := cell{
		value:   0,
		display: false,
	}
	grid.mainGrid = grid.mainGrid[:0]
	var emptyRow []cell
	for i := 0; i < gdkArraySize; i++ {
		for j := 0; j < gdkArraySize; j++ {
			emptyRow = append(emptyRow, emptyCell)
		}
		grid.mainGrid = append(grid.mainGrid, emptyRow)
		emptyRow = nil
	}
}

func (grid *godokuGrid) Print(masked bool) {
	for i, row := range grid.mainGrid {
		fmt.Println()
		for j := range row {
			fmt.Print(" ")
			if masked && !grid.mainGrid[i][j].display {
				fmt.Print("_")
			} else {
				fmt.Print(grid.mainGrid[i][j].value)
			}
			fmt.Print(" ")
		}
		fmt.Print("")
	}
	fmt.Print("\n\n")
}

func (grid *godokuGrid) Mask() (finished bool) {
	for i, row := range grid.mainGrid {
		for j := range row {
			grid.mainGrid[i][j].display = false
		}
	}
	finished = false
	i, j := 0, 0
	cueCount := gdkArraySize * difficulty / 100
	cue := 0
	gridIndex := 0
	cueNumber := 0
	for matrixIndex := 0; matrixIndex < gdkArraySize; matrixIndex++ {
		indexArray = fillArray(gdkArraySize, true)
		cueCount = rand.Intn(cueCount) + 2
		for cueIndex := 0; cueIndex < cueCount; cueIndex++ {
			cueNumber = genRandom(indexArray)
			indexArray = truncateArray(cueNumber, indexArray)
			gridIndex = cueNumber + (matrixIndex/3)*27 + (cueNumber/3)*6 + int(math.Mod(float64(matrixIndex), 3))*3
			i = gridIndex / gdkArraySize
			j = gridIndex - i*gdkArraySize
			grid.mainGrid[i][j].display = true
			cue++
			if cue == (gdkArraySize*gdkArraySize*difficulty/100 - 1) {
				finished = true
				return
			}
		}
		cueCount = gdkArraySize * difficulty / 100
	}

	return
}

// solveGRid function
