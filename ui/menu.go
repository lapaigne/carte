package ui

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Menu struct {
	Buttons []*Button
}

func (m *Menu) addButton(b *Button) *Menu {
	m.Buttons = append(m.Buttons, b)
	return m
}

func (m *Menu) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		for _, b := range m.Buttons {
			if b.In(ebiten.CursorPosition()) {
				b.Click()
			}
		}
	}

	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	for _, b := range m.Buttons {
		b.Draw(screen)
	}
}

func (m *Menu) Load()   {}
func (m *Menu) Unload() {}

func NewMenu(src *text.GoTextFaceSource, sWidth, sHeight int) *Menu {
	m := &Menu{}

	m.Buttons = make([]*Button, 0)

	margin := 10
	vertSpace := 25
	w := 400

	m.addButton(&Button{Y: 100, Text: "NEW"})
	m.addButton(&Button{Text: "LOAD"})
	m.addButton(&Button{Text: "SETTINGS"})

	exitBtn := &Button{Text: "EXIT"}
	m.addButton(exitBtn)

	for i, b := range m.Buttons {
		b.Click = func() {
			fmt.Printf("clicked: %s\n", b.Text)
		}
		b.Margin = margin
		b.Width = w
		b.LockWidth = true
		if i > 0 {
			prev := m.Buttons[i-1]
			b.Y = prev.Y + prev.Height + vertSpace
		}

		b.Init(src)
		b.CenterHor(sWidth)
	}

	exitBtn.Click = func() {
		fmt.Printf("clicked: %s\n", exitBtn.Text)
		os.Exit(0)
	}

	return m
}
