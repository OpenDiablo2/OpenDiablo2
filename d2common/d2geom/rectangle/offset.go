package rectangle

// Offset nudges (translates) the top left corner of a Rectangle by a given offset.
func Offset(r *Rectangle, x, y float64) *Rectangle {
	r.X += x
	r.Y += y

	return r
}
