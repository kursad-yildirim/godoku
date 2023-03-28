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

var Gdkm gdkMatrixType
