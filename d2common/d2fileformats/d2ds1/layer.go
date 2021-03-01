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

type tileRow []Tile     // index is x coordinate
type tileGrid []tileRow // index is y coordinate

type grid interface {
	Width() int
	SetWidth(w int) grid

	Height() int
	SetHeight(h int) grid

	Size() (w, h int)
	SetSize(w, h int) grid

	Tile(x, y int) *Tile
	SetTile(x, y int, t *Tile)
}

type layer struct {
	tiles tileGrid
}

func (l *layer) Tile(x, y int) *Tile {
	if l.Width() < x || l.Height() < y {
		return nil
	}

	return &l.tiles[y][x]
}

func (l *layer) SetTile(x, y int, t *Tile) {
	if l.Width() > x || l.Height() > y {
		return
	}

	l.tiles[y][x] = *t
}

func (l *layer) Width() int {
	if len(l.tiles[0]) < 1 {
		l.SetWidth(1)
	}

	return len(l.tiles[0])
}

func (l *layer) SetWidth(w int) grid {
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

func (l *layer) Height() int {
	if len(l.tiles) < 1 {
		l.SetHeight(1)
	}

	return len(l.tiles)
}

func (l *layer) SetHeight(h int) grid {
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

func (l *layer) Size() (w, h int) {
	return l.Width(), l.Height()
}

func (l *layer) SetSize(w, h int) grid {
	return l.SetWidth(w).SetHeight(h)
}
