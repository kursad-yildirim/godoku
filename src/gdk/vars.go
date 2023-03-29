package gdk

var gridNumTotal, gdkDimension, displayCount, difficulty, gdkArraySize int
var valueArray []int
var indexArray []int
var indexDispArray []int

type gdkCellType struct {
	value   int
	display bool
}

type gdkMatrixType struct {
	gdkMatrix [][]gdkCellType
}

type httpSuccessType struct {
	Difficulty int `json:"difficulty"`
}

var Gdkm gdkMatrixType

var httpSuccess httpSuccessType
