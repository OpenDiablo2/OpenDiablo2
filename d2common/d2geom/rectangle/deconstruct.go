package rectangle

import "github.com/gravestench/pho/geom/point"

// Deconstruct creates a slice of points for each corner of a Rectangle.
// If a slice is specified, each point object will be added to the end of the slice,
// otherwise a new slice will be created.
func Deconstruct(r *Rectangle, to []*point.Point) []*point.Point {
	if to == nil {
		to = make([]*point.Point, 0)
	}

	to = append(to, point.New(r.X, r.Y))
	to = append(to, point.New(r.Right(), r.Y))
	to = append(to, point.New(r.Right(), r.Bottom()))
	to = append(to, point.New(r.X, r.Bottom()))

	return to
}
