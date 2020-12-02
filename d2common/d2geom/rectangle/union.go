package rectangle

import "math"

// Union creates a new Rectangle or repositions and/or resizes an existing Rectangle so that it
// encompasses the two given Rectangles, i.e. calculates their union.
func Union(a, b, out *Rectangle) *Rectangle {
	x, y := math.Min(a.X, b.X), math.Min(a.Y, b.Y)
	w, h := math.Min(a.Width, b.Width), math.Min(a.Height, b.Height)

	return out.SetTo(x, y, w, h)
}
