package gdk

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func (grid *gridType) FillGrid() (finished bool) {
	rand.Seed(time.Now().UnixNano())
	finished = true
	grid.initialize()
	cellValue := -1
	for i, row := range grid.members {
		for j := range row {
			valueArray = fillArray(props.GridSize, false)
			if len(valueArray) == 1 {
				cellValue = valueArray[0]
			} else {
				removeExistingRowElements(i, j, &valueArray, *grid)
				removeExistingColumnElements(i, j, &valueArray, *grid)
				removeExistingBlockElements(i, j, &valueArray, *grid)
				if len(valueArray) < 1 {
					finished = false
					return
				}
				cellValue = generateRandomDigit(valueArray)
			}
			grid.members[i][j].value = cellValue
		}
	}

	return
}

func (grid *gridType) initialize() {
	props.BlockSize = int(math.Sqrt(float64(props.GridSize)))
	// make a custom error function here if props.BlockSize is not an integer
	emptyCell := cellType{
		value:   0,
		display: false,
	}
	grid.members = grid.members[:0] // try grid.members = nil
	var emptyRow []cellType
	for i := 0; i < props.GridSize; i++ {
		for j := 0; j < props.GridSize; j++ {
			emptyRow = append(emptyRow, emptyCell)
		}
		grid.members = append(grid.members, emptyRow)
		emptyRow = nil
	}
}

func (grid *gridType) Print(masked bool) {
	for i, row := range grid.members {
		fmt.Println()
		for j := range row {
			fmt.Print(" ")
			if masked && !grid.members[i][j].display {
				fmt.Print("_")
			} else {
				fmt.Print(grid.members[i][j].value)
			}
			fmt.Print(" ")
		}
		fmt.Print("")
	}
	fmt.Print("\n\n")
}

func (grid *gridType) Mask() (finished bool) {
	for i, row := range grid.members {
		for j := range row {
			grid.members[i][j].display = false
		}
	}
	finished = false
	i, j := 0, 0
	givenCount := props.GridSize * props.Difficulty / 100
	givenIndex := 0
	gridIndex := 0
	given := 0
	for blockIndex := 0; blockIndex < props.GridSize; blockIndex++ {
		indexArray = fillArray(props.GridSize, true)
		givenCount = rand.Intn(givenCount) + 2
		for g := 0; g < givenCount; g++ {
			given = generateRandomDigit(indexArray)
			indexArray = truncateArray(given, indexArray)
			gridIndex = given + (blockIndex / 3 ) * 27 + ( given / 3 ) * 6 + int(math.Mod(float64(blockIndex), 3 ) ) * 3
			i = gridIndex / props.GridSize
			j = gridIndex - i * props.GridSize
			grid.members[i][j].display = true
			givenIndex++
			if givenIndex == (props.GridPopulation * props.Difficulty / 100 - 1) {
				finished = true
				return
			}
		}
		givenCount = props.GridSize * props.Difficulty / 100
	}

	return
}

// solveGRid function
