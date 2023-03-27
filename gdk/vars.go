package gdk

const gdkDimension = 4
const gdkArraySize = 16
const gdkNumCount = 256
const displayCount = 111

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
