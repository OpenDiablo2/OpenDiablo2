package rectangle

import (
	"math/rand"

	"github.com/gravestench/pho/geom/point"
)

// GetRandomPoint returns a random point within the Rectangle's bounds.
func GetRandomPoint(r *Rectangle, p *point.Point) *point.Point {
	if p == nil {
		p = point.New(0, 0)
	}

	p.X = r.X + (rand.Float64() * r.Width)
	p.Y = r.Y + (rand.Float64() * r.Height)

	return p
}
