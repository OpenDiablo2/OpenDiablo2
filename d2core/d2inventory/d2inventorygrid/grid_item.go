package d2inventorygrid

const OutOfBounds int = -1

func NewGridItem(width, height int) *GridItem {
	return &GridItem{
		x:      OutOfBounds,
		y:      OutOfBounds,
		width:  width,
		height: height,
	}
}

type GridItem struct {
	x, y int
	width, height int
}

func (gi *GridItem) Size() (width, height int) {
	return gi.width, gi.height
}
