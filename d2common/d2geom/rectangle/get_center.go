package rectangle

import "github.com/gravestench/pho/geom/point"

// GetCenter returns the center of the Rectangle as a Point.
func GetCenter(r *Rectangle) *point.Point {
	return point.New(r.CenterX(), r.CenterY())
}
