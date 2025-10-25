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
	X      int
	Y      int
	Width  int
	Height int

	LockWidth  bool
	LockHeight bool
	Text       string
	LayoutOpts *text.LayoutOptions
	DrawOpts   *text.DrawOptions
	State      ButtonState
	Margin     int
	TextFace   *text.GoTextFace
}

func (b *Button) adjust() {
	_ftw, _fth := text.Measure(b.Text, b.TextFace, b.DrawOpts.LineSpacing)
	tw, th := int(_ftw), int(_fth)
	w, h := b.Width, b.Height

	b.DrawOpts.GeoM.Reset()
	b.DrawOpts.GeoM.Translate(
		float64(b.X+(w-tw)/2),
		float64(b.Y+(h-th)/2),
	)

	if !b.LockHeight {
		b.Height = 2*b.Margin + th
	}

	if !b.LockWidth {
		b.Width = 2*b.Margin + tw
	}
}

func (b *Button) Init(source *text.GoTextFaceSource) {
	draw := &text.DrawOptions{}
	draw.ColorScale.ScaleWithColor(color.White)
	draw.Filter = ebiten.FilterNearest
	b.DrawOpts = draw

	tf := &text.GoTextFace{
		Source: source,
		Size:   24,
	}

	b.TextFace = tf
	b.adjust()
}

func (b *Button) CenterVer(height int) {
	b.Y = height/2 - b.Height/2
	b.adjust()
}

func (b *Button) CenterHor(width int) {
	b.X = width/2 - b.Width/2
	b.adjust()
}

func (b *Button) Click() {
	fmt.Println("clicked")
}

func (b *Button) Draw(screen *ebiten.Image) {
	text.Draw(screen, b.Text, b.TextFace, b.DrawOpts)
	vector.StrokeRect(screen, float32(b.X), float32(b.Y), float32(b.Width), float32(b.Height), 1, color.White, false)
}

func (b *Button) In(x, y int) bool {
	return x >= b.X && x <= b.Width+b.X && y >= b.Y && y <= b.Height+b.Y
}
