package carma

import (
	"math"
)

type Camera struct {
	L float32
	R float32
	T float32
	B float32
}

func (c *Camera) Zoom(factor float32) {
	if factor < 0 {
		factor = -1 / factor
	}
	if factor > 10 {
		factor = 10
	}
	if factor < .1 {
		factor = .1
	}

	f := float32(math.Sqrt(float64(factor)))

	c.L *= f
	c.R *= f
	c.T *= f
	c.B *= f
}

// dir is vector in world coords
func (c *Camera) Pan(dir Vec2) {
	c.L += dir.X
	c.R += dir.X
	c.T += dir.Y
	c.B += dir.Y
}

func (c Camera) Dims32() (L, R, T, B float32) {
	return c.L, c.R, c.T, c.B
}

func (c Camera) Dims64() (L, R, T, B float64) {
	return float64(c.L), float64(c.R), float64(c.T), float64(c.B)
}

func (c Camera) Dims64Rounded() (L, R, T, B float64) {
	return math.Round(float64(c.L)),
		math.Round(float64(c.R)),
		math.Round(float64(c.T)),
		math.Round(float64(c.B))
}
