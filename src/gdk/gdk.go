package gdk

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func (grid *gridType) initialize() {
	Props.GridSize = Props.BlockSize * Props.BlockSize
	Props.GridPopulation = Props.GridSize * Props.GridSize

	emptyCell := cellType{
		value:   0,
		display: false,
	}
	grid.members = nil
	var emptyRow []cellType
	for i := 0; i < Props.GridSize; i++ {
		for j := 0; j < Props.GridSize; j++ {
			emptyRow = append(emptyRow, emptyCell)
		}
		grid.members = append(grid.members, emptyRow)
		emptyRow = nil
	}
}

func (grid *gridType) FillGrid() {
	rand.Seed(time.Now().UnixNano())
	grid.initialize()
	cellValue := -1
	for i, row := range grid.members {
		for j := range row {
			valueArray = fillArray(Props.GridSize, false)
			if len(valueArray) == 1 {
				cellValue = valueArray[0]
			} else {
				removeExistingRowElements(i, j, &valueArray, *grid)
				removeExistingColumnElements(i, j, &valueArray, *grid)
				removeExistingBlockElements(i, j, &valueArray, *grid)
				if len(valueArray) < 1 {
					return
				}
				cellValue = generateRandomDigit(valueArray)
			}
			grid.members[i][j].value = cellValue
		}
	}
	grid.FillGrid()
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

func (grid *gridType) DigHoles() {
	for i, row := range grid.members {
		for j := range row {
			grid.members[i][j].display = false
		}
	}
	i, j := 0, 0
	givenCount := Props.GridSize * Props.Difficulty / 100
	givenIndex := 0
	gridIndex := 0
	given := 0
	for blockIndex := 0; blockIndex < Props.GridSize; blockIndex++ {
		indexArray = fillArray(Props.GridSize, true)
		givenCount = rand.Intn(givenCount) + 2
		for g := 0; g < givenCount; g++ {
			given = generateRandomDigit(indexArray)
			indexArray = truncateArray(given, indexArray)
			gridIndex = given + (blockIndex/3)*27 + (given/3)*6 + int(math.Mod(float64(blockIndex), 3))*3
			i = gridIndex / Props.GridSize
			j = gridIndex - i*Props.GridSize
			grid.members[i][j].display = true
			givenIndex++
			if givenIndex == (Props.GridPopulation*Props.Difficulty/100 - 1) {
				return
			}
		}
		givenCount = Props.GridSize * Props.Difficulty / 100
	}
	grid.DigHoles()
}
