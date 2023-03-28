package gdk

const gdkDimension = 3
const gdkArraySize = 9
const gdkNumCount = 81
const displayCount = 35

var valueArray []int
var indexArray []int
var indexDispArray []int

var emptyValue int = 0

type gdkCellType struct {
	value   int
	display bool
}

type gdkMatrixType struct {
	gdkMatrix [gdkArraySize][gdkArraySize]gdkCellType
}

var Gdkm gdkMatrixType
