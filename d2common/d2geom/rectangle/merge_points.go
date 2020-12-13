package rectangle

import (
	"math"

	"github.com/gravestench/pho/geom/point"
)

// Merges a Rectangle with a list of points by repositioning and/or resizing
// it such that all points are located on or within its bounds.
func MergePoints(r *Rectangle, points []*point.Point) *Rectangle {
	minX, maxX, minY, maxY := r.X, r.Right(), r.Y, r.Bottom()

	for idx := range points {
		minX, maxX = math.Min(minX, points[idx].X), math.Max(maxX, points[idx].X)
		minY, maxY = math.Min(minY, points[idx].Y), math.Max(maxY, points[idx].Y)
	}

	r.X, r.Y = minX, minY
	r.Width, r.Height = maxX-minX, maxY-minY

	return r
}
