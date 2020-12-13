package rectangle

import "github.com/gravestench/pho/geom/point"

// GetRectangleFromPoints calculates the Axis Aligned Bounding Box ( or aabb) from an array of
// points.
func GetRectangleFromPoints(points []*point.Point) *Rectangle {
	return New(0, 0, 0, 0).MergePoints(points)
}
