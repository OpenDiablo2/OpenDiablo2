package rectangle

// Adjusts rectangle `a`, changing its width, height and position,
// so that it fits inside the area of rectangle `b`, while maintaining its original
// aspect ratio.
func FitInside(a, b *Rectangle) *Rectangle {
	aRatio := GetAspectRatio(a)
	bRatio := GetAspectRatio(b)

	if aRatio < bRatio {
		a.SetSize(b.Height*aRatio, b.Height)
	} else {
		a.SetSize(b.Width, b.Width/aRatio)
	}

	return a.CenterOn(b.CenterX(), b.CenterY())
}
