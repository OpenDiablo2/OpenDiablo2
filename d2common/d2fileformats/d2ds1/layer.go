package d2ds1

// layerStreamType represents a layer stream type
type layerStreamType int

// Layer stream types
const (
	layerStreamWall1 layerStreamType = iota
	layerStreamWall2
	layerStreamWall3
	layerStreamWall4
	layerStreamOrientation1
	layerStreamOrientation2
	layerStreamOrientation3
	layerStreamOrientation4
	layerStreamFloor1
	layerStreamFloor2
	layerStreamShadow1
	layerStreamSubstitute1
)

type (
	tileRow  []Tile    // index is x coordinate
	tileGrid []tileRow // index is y coordinate
)

// Layer is an abstraction of a tile grid with some helper methods
type Layer struct {
	tiles tileGrid
}

// Tile returns the tile at the given x,y coordinate in the grid, or nil if empty.
func (l *Layer) Tile(x, y int) *Tile {
	if l.Width() < x || l.Height() < y {
		return nil
	}

	return &l.tiles[y][x]
}

// SetTile sets the tile at the given x,y coordinate in the tile grid
func (l *Layer) SetTile(x, y int, t *Tile) {
	if l.Width() > x || l.Height() > y {
		return
	}

	l.tiles[y][x] = *t
}

// Width returns the width of the tile grid
func (l *Layer) Width() int {
	if len(l.tiles[0]) < 1 {
		l.SetWidth(1)
	}

	return len(l.tiles[0])
}

// SetWidth sets the width of the tile grid, minimum of 1
func (l *Layer) SetWidth(w int) *Layer {
	if w < 1 {
		w = 1
	}

	// ensure at least one row
	if len(l.tiles) < 1 {
		l.tiles = make(tileGrid, 1)
	}

	// create/copy tiles as required to satisfy width
	for y := range l.tiles {
		if (w - len(l.tiles[y])) == 0 { // if requested width same as row width
			continue
		}

		tmpRow := make(tileRow, w)

		for x := range tmpRow {
			if x < len(l.tiles[y]) { // if tile exists
				tmpRow[x] = l.tiles[y][x] // copy it
			}
		}

		l.tiles[y] = tmpRow
	}

	return l
}

// Height returns the height of the tile grid
func (l *Layer) Height() int {
	if len(l.tiles) < 1 {
		l.SetHeight(1)
	}

	return len(l.tiles)
}

// SetHeight sets the height of the tile grid, minimum of 1
func (l *Layer) SetHeight(h int) *Layer {
	if h < 1 {
		h = 1
	}

	// make tmpGrid to move existing tiles into
	tmpGrid := make(tileGrid, h)

	for y := range tmpGrid {
		tmpGrid[y] = make(tileRow, l.Width())
	}

	// move existing tiles over
	for y := range l.tiles {
		if y >= len(tmpGrid) {
			continue
		}

		for x := range l.tiles[y] {
			if x >= len(tmpGrid[y]) {
				continue
			}

			tmpGrid[y][x] = l.tiles[y][x]
		}
	}

	l.tiles = tmpGrid

	return l
}

// Size returns the width and height of the tile grid
func (l *Layer) Size() (w, h int) {
	return l.Width(), l.Height()
}

// SetSize sets the width and height of the tile grid
func (l *Layer) SetSize(w, h int) *Layer {
	return l.SetWidth(w).SetHeight(h)
}
