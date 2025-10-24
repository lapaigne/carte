package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type ButtonState int

const (
	Disabled ButtonState = iota - 1
	Off
	On
)

type Button struct {
	X int
	Y int
	w int
	h int

	Text       string
	LayoutOpts *text.LayoutOptions
	DrawOpts   *text.DrawOptions
	State      ButtonState
	Margin     int
}

func (b *Button) Init() {
	draw := &text.DrawOptions{}
	draw.ColorScale.ScaleWithColor(color.White)
	draw.GeoM.Translate(float64(b.X+b.Margin), float64(b.Y+b.Margin))
	draw.Filter = ebiten.FilterNearest
	b.DrawOpts = draw
}

func (b *Button) Click() {}

func (b *Button) Draw(screen *ebiten.Image, source *text.GoTextFaceSource) {
	tf := &text.GoTextFace{
		Source: source,
		Size:   24,
	}
	text.Draw(screen, b.Text, tf, b.DrawOpts)
	w, h := text.Measure(b.Text, tf, b.DrawOpts.LineSpacing)
	vector.StrokeRect(screen, float32(b.X), float32(b.Y), float32(w)+float32(2*b.Margin), float32(h)+float32(2*b.Margin), 1, color.White, false)
}
