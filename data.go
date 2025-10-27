package main

import (
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

	Mode int
	X    int
	Y    int
}

type Settings struct {
	Width   int
	Height  int
	Title   string
	Monitor *ebiten.MonitorType
}
