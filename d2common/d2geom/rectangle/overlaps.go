package rectangle

// Checks if two Rectangles overlap. If a Rectangle is within another Rectangle,
// the two will be considered overlapping. Thus, the Rectangles are treated as "solid".
func Overlaps(a, b *Rectangle) bool {
	return a.X < b.Right() &&
		a.Right() > b.X &&
		a.Y < b.Bottom() &&
		a.Bottom() > b.Y
}
