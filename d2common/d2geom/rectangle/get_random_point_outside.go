package rectangle

import (
	"math/rand"

	"github.com/gravestench/pho/geom/point"
	"github.com/gravestench/pho/phomath"
)

// Calculates a random point that lies within the `outer` Rectangle, but outside of the `inner`
// Rectangle. The inner Rectangle must be fully contained within the outer rectangle.
func GetRandomPointOutside(outer, inner *Rectangle, out *point.Point) *point.Point {
	if out == nil {
		out = point.New(0, 0)
	}

	if !ContainsRectangle(outer, inner) {
		return out
	}

	const (
		top = iota
		right
		bottom
		left
	)

	switch phomath.Between(top, left) {
	case top:
		out.X = outer.X + (rand.Float64() * (inner.Right() - outer.X))
		out.Y = outer.Y + (rand.Float64() * (inner.Top() - outer.Y))
	case right:
		out.X = inner.X + (rand.Float64() * (outer.Right() - inner.X))
		out.Y = inner.Y + (rand.Float64() * (outer.Bottom() - inner.Bottom()))
	case bottom:
		out.X = outer.X + (rand.Float64() * (inner.X - outer.X))
		out.Y = inner.Y + (rand.Float64() * (outer.Bottom() - inner.Y))
	case left:
		out.X = inner.Right() + (rand.Float64() * (outer.Right() - inner.Right()))
		out.Y = outer.Y + (rand.Float64() * (inner.Bottom() - outer.Y))
	}

	return out
}
