package main

import (
	"carte/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type App struct {
	Mode   int
	X      int
	Y      int
	Button *ui.Button
}

type Settings struct {
	Width   int
	Height  int
	Title   string
	Monitor *ebiten.MonitorType
}
