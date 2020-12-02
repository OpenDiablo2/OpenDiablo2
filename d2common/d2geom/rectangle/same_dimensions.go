package rectangle

import "github.com/gravestench/pho/phomath"

// SameDimensions determines if the two objects (either Rectangles or Rectangle-like) have the same
// width and height values under strict equality.
func SameDimensions(a, b *Rectangle) bool {
	dw := (a.Width - b.Width) * (a.Width - b.Width)
	dh := (a.Height - b.Height) * (a.Height - b.Height)

	return dw < phomath.Epsilon && dh < phomath.Epsilon
}
