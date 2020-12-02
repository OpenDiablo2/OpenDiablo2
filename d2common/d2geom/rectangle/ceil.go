package rectangle

import "math"

// Ceil rounds a Rectangle's position up to the smallest integer greater than or equal to each
// current coordinate.
func Ceil(r *Rectangle) *Rectangle {
	r.X = math.Ceil(r.X)
	r.Y = math.Ceil(r.Y)

	return r
}
