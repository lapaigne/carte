package ui

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Menu struct {
	Buttons []*Button
}

func (m *Menu) AddButton(b *Button) *Menu {
	m.Buttons = append(m.Buttons, b)
	return m
}

func NewMenu(src *text.GoTextFaceSource, width, height int) *Menu {
	m := &Menu{}

	m.Buttons = make([]*Button, 0)

	margin := 10
	vertSpace := 25

	m.AddButton(&Button{Y: 100, Text: "NEW", Width: 400})
	m.AddButton(&Button{Text: "LOAD"})
	m.AddButton(&Button{Text: "SETTINGS"})

	exitBtn := &Button{Text: "EXIT"}
	m.AddButton(exitBtn)

	var nw int
	for _, b := range m.Buttons {
		b.Init(src)
		w, _ := b.MinDims()
		if b.Width > w {
			w = b.Width
		}
		if w > nw {
			nw = w
		}
	}

	for i, b := range m.Buttons {
		b.Click = func() {
			fmt.Printf("clicked: %s\n", b.Text)
		}
		b.Margin = margin
		b.Width = nw
		b.LockWidth = true
		if i > 0 {
			prev := m.Buttons[i-1]
			b.Y = prev.Y + prev.Height + vertSpace
		}
		b.adjust()
		b.CenterHor(width)
	}

	exitBtn.Click = func() {
		fmt.Printf("clicked: %s\n", exitBtn.Text)
		os.Exit(0)
	}

	return m
}
