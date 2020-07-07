package d2vector

// Position is a vector in world space. The stored
// value is  the one returned by Position.World()
type Position struct {
	Vector
}

// NewPosition creates a new Position at the given
// float64 world position.
func NewPosition(x, y float64) *Position {
	// TODO: BigFloat dependency
	return &Position{NewVector(x, y)}
}

// World is the position, where 1 = one map
// tile.
func (p *Position) World() *Vector {
	return &p.Vector
}

// Tile is the tile position, always a whole
// number.
func (p *Position) Tile() *Vector {
	c := p.World().Clone()
	return c.Floor()
}

// TileOffset is the offset from the tile position,
// always < 1.
func (p *Position) TileOffset() *Vector {
	c := p.World().Clone()
	return c.Subtract(p.Tile())
}

// SubWorld is the position, where 5 = one map
// tile.
func (p *Position) SubWorld() *Vector {
	c := p.World().Clone()
	// TODO: BigFloat dependency
	return c.Scale(5)
}

// SubTile is the tile position in sub tiles,
// always a multiple of 5.
func (p *Position) SubTile() *Vector {
	c := p.Tile().Clone()
	// TODO: BigFloat dependency
	return c.Scale(5)
}

// SubTileOffset is the offset from the sub tile
// position in sub tiles, always < 1.
func (p *Position) SubTileOffset() *Vector {
	c := p.SubWorld().Clone()
	return c.Subtract(p.SubTile())
}
