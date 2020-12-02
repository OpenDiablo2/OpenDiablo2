package rectangle

// CopyFrom copies the values of one Rectangle to a destination Rectangle.
func CopyFrom(source, dest *Rectangle) *Rectangle {
	return dest.SetTo(source.X, source.Y, source.Width, source.Height)
}
