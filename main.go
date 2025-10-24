package main

import (
	"bytes"
	"carte/ui"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var gtfs *text.GoTextFaceSource

func (a *App) Update() error {
	return nil
}

func (a *App) Layout(outW, outH int) (int, int) {
	return 1280, 720
}

func (a *App) Draw(screen *ebiten.Image) {
	a.Button.Init()
	a.Button.Draw(screen, gtfs)
}

func main() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	gtfs = s

	b := &ui.Button{X: 100, Y: 120, Text: "BUTTON", Margin: 10}

	app := App{Button: b}

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
