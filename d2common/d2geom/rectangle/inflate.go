package rectangle

// Inflate increases the size of a Rectangle by a specified amount.
// The center of the Rectangle stays the same. The amounts are added to each side,
// so the actual increase in width or height is two times bigger than the respective argument.
func Inflate(r *Rectangle, x, y float64) *Rectangle {
	cx, cy := r.CenterX(), r.CenterY()
	r.Width, r.Height = r.Width+(2*x), r.Height+(2*y)

	return r.CenterOn(cx, cy)
}
