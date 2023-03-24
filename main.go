package main

import (
	"tripko.local/godoku/gdk"
)

func main() {
	gdk.Gdkm.Empty()
	gdk.Gdkm.FillGrid()
	gdk.Gdkm.Print()
	gdk.Gdkm.PrintSubMatrix(0, 0)
}
