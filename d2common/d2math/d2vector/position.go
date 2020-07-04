package d2vector

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

// Position is a vector in world space. The stored
// value is  the one returned by Position.World()
type Position struct {
	d2interface.Vector
}

// World is the position, where 1 = one map
// tile.
func (p *Position) World() d2interface.Vector {
	return p.Vector
}

// Tile is the tile position, always a whole
// number.
func (p *Position) Tile() d2interface.Vector {
	c := p.World().Clone()
	return c.Floor()
}

// TileOffset is the offset from the tile position,
// always < 1.
func (p *Position) TileOffset() d2interface.Vector {
	c := p.World().Clone()
	return c.Subtract(p.Tile())
}

// SubWorld is the position, where 5 = one map
// tile.
func (p *Position) SubWorld() d2interface.Vector {
	c := p.World().Clone()
	return c.Multiply(New(5, 5))
}

// SubTile is the tile position in sub tiles,
// always a multiple of 5.
func (p *Position) SubTile() d2interface.Vector {
	c := p.Tile().Clone()
	return c.Multiply(New(5, 5))
}

// SubTileOffset is the offset from the sub tile
// position in sub tiles, always < 1.
func (p *Position) SubTileOffset() d2interface.Vector {
	c := p.SubWorld().Clone()
	return c.Subtract(p.SubTile())
}
