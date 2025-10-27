package math

type KeepRat int

const (
	None KeepRat = iota
	Width
	Height
)

type Projector struct {
	Screen    Vec2Int
	Camera    *Camera
	KeepRatio KeepRat
}

func (p *Projector) WorldToScreen(world Vec2Int) (screen Vec2Int) {
	w, h := p.Screen.X, p.Screen.Y
	u, v := world.X, world.Y
	l, r, t, b := p.Camera.L, p.Camera.R, p.Camera.T, p.Camera.B

	screen.X = (w / (l - r)) * (u - l)
	screen.Y = (h / (b - t)) * (v - r)
	return screen
}

func (p *Projector) ScreenToWorld(screen Vec2Int) (world Vec2Int) {
	w, h := p.Screen.X, p.Screen.Y
	u, v := screen.X, screen.Y
	l, r, t, b := p.Camera.L, p.Camera.R, p.Camera.T, p.Camera.B

	world.X = u*(r-l)/w + l
	world.Y = v*(b-t)/h + r
	return world
}
