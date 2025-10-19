package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func (a *App) Update() error {
	return nil
}

func (a *App) Layout(outW, outH int) (int, int) {
	return 320, 240
}

func (a *App) Draw(screen *ebiten.Image) {
}

func main() {
	app := App{}
	settings := Settings{
		Width:  1280,
		Height: 720,
		Title:  "La Carte",
	}

	ebiten.SetWindowSize(settings.Width, settings.Height)
	ebiten.SetWindowTitle(settings.Title)
	if err := ebiten.RunGame(&app); err != nil {
		log.Fatal(err)
	}
}
