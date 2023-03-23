package gdk

const gdkBoundary = 3

var valueArray = [9]int32{1, 2, 3, 4, 5, 6, 7, 8, 9}

var emptyValue int32 = 0

type gdkCellType struct {
	value         int32
	possibilities [9]int32
	excluded      [9]int32
	decided       bool
}

type gdkSubMatrixType struct {
	gdkSubMatrix [gdkBoundary][gdkBoundary]gdkCellType
}

type gdkMatrixType struct {
	gdkMatrix [gdkBoundary][gdkBoundary]gdkSubMatrixType
}

var Gdkm gdkMatrixType
