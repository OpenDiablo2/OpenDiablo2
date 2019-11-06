package Common

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
