package ui

import (
	"fmt"
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

func (b *Button) Click() {
	fmt.Println("clicked")
}

func (b *Button) Draw(screen *ebiten.Image, source *text.GoTextFaceSource) {
	tf := &text.GoTextFace{
		Source: source,
		Size:   24,
	}
	w, h := text.Measure(b.Text, tf, b.DrawOpts.LineSpacing)

	b.w = int(w) + 2*b.Margin
	b.h = int(h) + 2*b.Margin

	text.Draw(screen, b.Text, tf, b.DrawOpts)
	vector.StrokeRect(screen, float32(b.X), float32(b.Y), float32(b.w), float32(b.h), 1, color.White, false)
}

func (b *Button) In(x, y int) bool {
	return x >= b.X && x <= b.w+b.X && y >= b.Y && y <= b.h+b.Y
}
