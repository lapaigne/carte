package ui

import (
	"carte/math"
	"carte/world"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Map struct {
	Path      vector.Path
	Projector math.Projector
	World     *world.World
	Dotted    bool
}

func NewMap() *Map {
	proj := math.Projector{}
	proj.Camera = &math.Camera{L: -16, R: 16, T: -9, B: 9}
	proj.Screen = math.Vec2{1280, 720}

	m := &Map{}
	m.World = &world.World{Path: []math.Vec2{}}
	m.Projector = proj

	for range 5 {
		m.World.Path = append(m.World.Path, math.Vec2{
			X: 10*rand.Float32() - 5,
			Y: 10*rand.Float32() - 5,
		})
	}
	m.Path = m.Projector.ScreenPath(m.World.Path)

	m.Dotted = true

	return m
}

func (m *Map) Load()   {}
func (m *Map) Unload() {}

func (m *Map) Name() string {
	return "Map Scene"
}

func (m *Map) Update() error {
	if _, dy := ebiten.Wheel(); dy != 0 {
		f := float32(dy) * 0.3
		m.Projector.Camera.Zoom(f)
		m.Path = m.Projector.ScreenPath(m.World.Path)
	}

	return nil
}

func (m *Map) Draw(screen *ebiten.Image) {
	vector.StrokePath(screen, &m.Path, &vector.StrokeOptions{Width: 4}, &vector.DrawPathOptions{AntiAlias: false})

	if m.Dotted {
		xmin, xmax := m.Projector.Camera.L, m.Projector.Camera.R
		ymin, ymax := m.Projector.Camera.T, m.Projector.Camera.B
		for i := xmin; i < xmax; i++ {
			for j := ymin; j < ymax; j++ {
				s := m.Projector.WorldToScreen(math.Vec2{X: i, Y: j})
				vector.StrokeCircle(screen, s.X, s.Y, 1, 1, color.White, false)
			}
		}
	}
}
