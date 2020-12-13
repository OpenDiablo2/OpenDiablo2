package rectangle

import "github.com/gravestench/pho/geom/point"

// MergeXY merges a Rectangle with a point by repositioning and/or resizing it
// so that the point is on or within its bounds.
func MergeXY(r *Rectangle, x, y float64) *Rectangle {
	return r.MergePoints([]*point.Point{point.New(x, y)})
}
