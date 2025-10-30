package math

type Vectorer interface {
	Length() float32
}

type Vec2 struct {
	X float32
	Y float32
}
