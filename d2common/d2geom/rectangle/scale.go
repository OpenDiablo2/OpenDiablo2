package rectangle

// Scale the width and height of this Rectangle by the given amounts.
func Scale(r *Rectangle, x, y float64) *Rectangle {
	r.Width *= x
	r.Height *= y

	return r
}
