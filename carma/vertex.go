package carma

type Vectorer interface {
	Length() float32
}

type Vec2 struct {
	X float32
	Y float32
}

func FromInt(x, y int) Vec2 {
	return Vec2{float32(x), float32(y)}
}

func Add(v, u Vec2) Vec2 {
	return Vec2{v.X + u.X, v.Y + u.Y}
}

func Sub(v, u Vec2) Vec2 {
	return Vec2{v.X - u.X, v.Y - u.Y}
}
