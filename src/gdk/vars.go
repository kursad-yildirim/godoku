package gdk

var gridNumTotal, gdkDimension, displayCount, difficulty, gdkArraySize int
var valueArray []int
var indexArray []int

type cell struct {
	value   int  `json:"value"`
	display bool `json:"display"`
}

type godokuGrid struct {
	mainGrid [][]cell
}

type httpSuccessType struct {
	Difficulty int `json:"difficulty"`
}

var GodokuGrid godokuGrid

var httpSuccess httpSuccessType
