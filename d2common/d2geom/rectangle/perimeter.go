package rectangle

// Perimeter calculates the perimeter of a Rectangle.
func Perimeter(r *Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}
