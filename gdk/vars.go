package gdk

const gdkDimension = 3
const gdkArraySize = 9

var valueArray = [gdkArraySize]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

var emptyValue int = 0

type gdkCellType struct {
	value int
}

type gdkMatrixType struct {
	gdkMatrix [gdkArraySize][gdkArraySize]gdkCellType
}

var Gdkm gdkMatrixType
