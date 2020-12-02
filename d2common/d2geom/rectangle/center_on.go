package rectangle

// CenterOn moves the top-left corner of a Rectangle so that its center is at the given coordinates.
func CenterOn(r *Rectangle, x, y float64) *Rectangle {
	return r.SetCenterX(x).SetCenterY(y)
}
