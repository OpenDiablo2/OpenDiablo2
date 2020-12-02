package rectangle

// ContainsRectangle checks if one rectangle fully contains another.
func ContainsRectangle(a, b *Rectangle) bool {
	if b.Area() > a.Area() {
		return false
	}

	return (b.X > a.X && b.X < a.Right()) &&
		(b.Right() > a.X && b.Right() < a.Right()) &&
		(b.Y > a.Y && b.Y < a.Bottom()) &&
		(b.Bottom() > a.Y && b.Bottom() < a.Bottom())
}
