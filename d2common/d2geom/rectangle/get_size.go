package rectangle

import "github.com/gravestench/pho/geom/point"

// GetSize returns the size of the Rectangle, expressed as a Point object.
// With the value of the `width` as the `x` property and the `height` as the `y` property.
func GetSize(r *Rectangle) *point.Point {
	return point.New(r.Width, r.Height)
}
