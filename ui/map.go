package ui

import (
	"carte/carma"
	"carte/world"
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	Hand ToolMode = iota
	Free
)

type ToolMode int

type Map struct {
	Path      vector.Path
	Projector carma.Projector
	World     *world.World
	Mode      ToolMode
	Hand      bool
	Dragged   bool
	Dotted    bool
	Initial   carma.Vec2
}

func NewMap() *Map {
	proj := carma.Projector{}
	proj.Camera = &carma.Camera{L: -16, R: 16, T: -9, B: 9}
	proj.Screen = carma.Vec2{1280, 720}

	m := &Map{}
	m.World = &world.World{Path: []carma.Vec2{}}
	m.Projector = proj

	for range 5 {
		m.World.Path = append(m.World.Path, carma.Vec2{
			X: float32(rand.Intn(11) - 5),
			Y: float32(rand.Intn(11) - 5),
		})
	}
	m.Path = m.Projector.ScreenPath(m.World.Path)

	m.Hand = true
	m.Dotted = true

	return m
}

func (m *Map) Load()   {}
func (m *Map) Unload() {}

func (m *Map) Name() string {
	return "Map Scene"
}

func (m *Map) Update() error {
	if m.Hand {
		if m.Dragged {
			cur := carma.FromInt(ebiten.CursorPosition())
			curs := m.Projector.ScreenToWorld(cur)
			inits := m.Projector.ScreenToWorld(m.Initial)
			pan := carma.Sub(inits, curs)

			m.Initial = cur

			m.Projector.Camera.Pan(pan)

			if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
				m.Dragged = false
			}
		}
		if _, dy := ebiten.Wheel(); dy != 0 {
			f := float32(dy) * 0.3
			m.Projector.Camera.Zoom(f)
		}

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			m.Dragged = true
			m.Initial = carma.FromInt(ebiten.CursorPosition())
		}

	}

	m.Path = m.Projector.ScreenPath(m.World.Path)
	return nil
}

func (m *Map) Draw(screen *ebiten.Image) {
	vector.StrokePath(screen, &m.Path, &vector.StrokeOptions{Width: 4}, &vector.DrawPathOptions{AntiAlias: false})

	if m.Dotted {
		xn, xx, yn, yx := m.Projector.Camera.Dims64Rounded()
		ebitenutil.DebugPrint(screen, fmt.Sprintf("L: %.3f\tR: %.3f\tT: %.3f\tB: %.3f", xn, xx, yn, yx))
		for i := xn; i <= xx; i++ {
			for j := yn; j <= yx; j++ {
				s := m.Projector.WorldToScreen(carma.Vec2{X: float32(i), Y: float32(j)})
				vector.StrokeCircle(screen, s.X, s.Y, 1, 1, color.White, false)
			}
		}
	}
}
