package main

import (
	"carte/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Load()
	Unload()
	Update() error
	Draw(screen *ebiten.Image)
}

type App struct {
	CurrentScene Scene
	Scenes       []Scene

	Menu   *ui.Menu
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
