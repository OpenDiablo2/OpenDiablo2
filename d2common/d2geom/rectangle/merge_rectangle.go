package rectangle

// MergeRectangle merges the source rectangle into the target rectangle and returns the target.
// Neither rectangle should have a negative width or height.
func MergeRectangle(target, source *Rectangle) *Rectangle {
	return MergePoints(target, source.Deconstruct(nil))
}
