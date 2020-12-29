package d2geom

// Rectangle represents a rectangle
type Rectangle struct {
	Left   int
	Top    int
	Width  int
	Height int
}

// Bottom returns y of the bottom point of the rectangle
func (v *Rectangle) Bottom() int {
	return v.Top + v.Height
}

// Right returns x of the right point of the rectangle
func (v *Rectangle) Right() int {
	return v.Left + v.Width
}

// IsInRect returns if the given position is in the rectangle or not
func (v *Rectangle) IsInRect(x, y int) bool {
	return x >= v.Left && x < v.Left+v.Width && y >= v.Top && y < v.Top+v.Height
}
