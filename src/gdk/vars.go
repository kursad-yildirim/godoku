package gdk

var valueArray []int
var indexArray []int

type propsType struct {
	BlockSize int `json:"blokSize"`
	GridSize int `json:"gridSize"`
	Difficulty int `json:"difficulty"`
	GridPopulation int `json: gridPopulation`
}

type cellType struct {
	value   int
	display bool
}

type gridType struct {
	members [][]cellType
}

var props propsType
var grid gridType
