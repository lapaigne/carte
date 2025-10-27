package main

import (
	"bytes"
	"carte/ui"
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var gtfs *text.GoTextFaceSource

func (a *App) Update() error {
	if a.CurrentScene == nil {
		return errors.New("current scene is nil")
	}

	a.CurrentScene.Update()
	return nil
}

func (a *App) Layout(outW, outH int) (int, int) {
	return 1280, 720
}

func (a *App) Draw(screen *ebiten.Image) {
	if a.CurrentScene == nil {
		return
	}

	a.CurrentScene.Draw(screen)
}

func main() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	gtfs = s

	b := &ui.Button{
		X:         100,
		Y:         120,
		Text:      "BUTTON",
		Margin:    10,
		Width:     400,
		LockWidth: true,
	}

	app := App{}

	ms := []*ebiten.MonitorType{}
	ms = ebiten.AppendMonitors(ms)

	settings := Settings{
		Width:   1280,
		Height:  720,
		Title:   "La Carte",
		Monitor: ms[0],
	}

	app.CurrentScene = ui.NewMenu(s, settings.Width, settings.Height)

	b.Init(gtfs)
	b.CenterHor(settings.Width)

	ebiten.SetMonitor(settings.Monitor)
	ebiten.SetWindowSize(settings.Width, settings.Height)
	ebiten.SetWindowTitle(settings.Title)
	if err := ebiten.RunGame(&app); err != nil {
		log.Fatal(err)
	}
}
