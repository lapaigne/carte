package math

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
