package rectangle

// Contains checks if the given x, y is inside the Rectangle's bounds.
func Contains(r *Rectangle, x, y float64) bool {
	if r.Width <= 0 || r.Height <= 0 {
		return false
	}

	return r.X <= x && r.X+r.Width >= x && r.Y <= y && r.Y+r.Height >= y
}
