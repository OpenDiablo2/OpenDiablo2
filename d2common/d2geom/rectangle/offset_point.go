package rectangle

import "github.com/gravestench/pho/geom/point"

// OffsetPoint nudges (translates) the top left corner of a Rectangle by the coordinates of a point.
func OffsetPoint(r *Rectangle, p *point.Point) *Rectangle {
	r.X += p.X
	r.Y += p.Y

	return r
}
