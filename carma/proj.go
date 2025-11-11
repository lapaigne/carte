package carma

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type KeepRat int

const (
	None KeepRat = iota
	Width
	Height
)

type Projector struct {
	Screen    Vec2
	Camera    *Camera
	KeepRatio KeepRat
}

func (p *Projector) WorldToScreen(world Vec2) (screen Vec2) {
	w, h := p.Screen.X, p.Screen.Y
	wx, wy := world.X, world.Y
	l, r, t, b := p.Camera.L, p.Camera.R, p.Camera.T, p.Camera.B

	screen.X = (w / (l - r)) * (l - wx)
	screen.Y = (h / (b - t)) * (b - wy)
	return screen
}

func (p *Projector) ScreenToWorld(screen Vec2) (world Vec2) {
	w, h := p.Screen.X, p.Screen.Y
	sx, sy := screen.X, screen.Y
	l, r, t, b := p.Camera.L, p.Camera.R, p.Camera.T, p.Camera.B

	world.X = sx*(r-l)/w + l
	world.Y = sy*(t-b)/h + b
	return world
}

func (p *Projector) ScreenPath(worldPath []Vec2) (screenPath vector.Path) {
	screenPath.Reset()

	for _, v := range worldPath {
		s := p.WorldToScreen(v)
		screenPath.LineTo(s.X, s.Y)
	}

	return screenPath
}
