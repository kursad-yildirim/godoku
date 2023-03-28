package main

import (
	"tripko.local/godoku/gdk"
)

func main() {
	gdk.Gdkm.Empty()
	for !gdk.Gdkm.FillGrid() {
	}
	gdk.Gdkm.Mask()
	gdk.Gdkm.Print(true)
}
