package rectangle

import "github.com/gravestench/pho/phomath"

// Equals compares the `x`, `y`, `width` and `height` properties of two rectangles.
func Equals(a, b *Rectangle) bool {
	dx := (a.X - b.X) * (a.X - b.X)
	dy := (a.Y - b.Y) * (a.Y - b.Y)
	dw := (a.Width - b.Width) * (a.Width - b.Width)
	dh := (a.Height - b.Height) * (a.Height - b.Height)

	return dx < phomath.Epsilon &&
		dy < phomath.Epsilon &&
		dw < phomath.Epsilon &&
		dh < phomath.Epsilon
}
