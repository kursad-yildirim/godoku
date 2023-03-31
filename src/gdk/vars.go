package gdk

var gdkDimension int
var valueArray []int
var indexArray []int

type godokuProps struct {
	Size       int `json:"size"`
	Difficulty int `json:"difficulty"`
	Population int
}

type cell struct {
	value   int
	display bool
}

type godokuGrid struct {
	mainGrid [][]cell
}

var props godokuProps
var GodokuGrid godokuGrid
