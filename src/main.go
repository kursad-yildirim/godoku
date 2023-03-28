package main

import (
	"tripko.local/godoku/gdk"
)

func main() {
	displayMasked := true
	for !gdk.Gdkm.FillGrid() {
	}
	gdk.Gdkm.Mask()
	gdk.Gdkm.Print(displayMasked)
}
