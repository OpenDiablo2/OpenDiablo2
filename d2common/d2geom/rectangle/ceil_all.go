package rectangle

import "math"

// CeilAll rounds a Rectangle's position and size up to the smallest
// integer greater than or equal to each respective value.
func CeilAll(r *Rectangle) *Rectangle {
	r.X = math.Ceil(r.X)
	r.Y = math.Ceil(r.Y)
	r.Width = math.Ceil(r.Width)
	r.Height = math.Ceil(r.Height)

	return r
}
