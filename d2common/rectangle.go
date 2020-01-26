package d2common

type Rectangle struct {
	Left   int
	Top    int
	Width  int
	Height int
}

func (v *Rectangle) Bottom() int {
	return v.Top + v.Height
}

func (v *Rectangle) Right() int {
	return v.Left + v.Width
}

func (v *Rectangle) IsInRect(x, y int) bool {
	return x >= v.Left && x < v.Left+v.Width && y >= v.Top && y < v.Top+v.Height
}
