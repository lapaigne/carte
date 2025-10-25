package ui

import (
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
	m.AddButton(&Button{Text: "EXIT"})

	var nw int
	for _, btn := range m.Buttons {
		btn.Init(src)
		w, _ := btn.MinDims()
		if btn.Width > w {
			w = btn.Width
		}
		if w > nw {
			nw = w
		}
	}

	for i, btn := range m.Buttons {
		btn.Margin = margin
		btn.Width = nw
		btn.LockWidth = true
		if i > 0 {
			prev := m.Buttons[i-1]
			btn.Y = prev.Y + prev.Height + vertSpace
		}
		btn.adjust()
		btn.CenterHor(width)
	}

	return m
}
