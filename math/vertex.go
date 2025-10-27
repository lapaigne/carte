package math

type Vectorer interface {
	Length() float32
}

type Vec2Int struct {
	X int
	Y int
}
