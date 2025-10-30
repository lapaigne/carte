package main

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Name() string
	Load()
	Unload()
	Update() error
	Draw(screen *ebiten.Image)
}
