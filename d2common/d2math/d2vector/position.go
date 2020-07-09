package d2vector

import (
	"fmt"
	"math"
)

const subTilesPerTile float64 = 5

// Position is a vector in world space. The stored value is  the one returned by Position.World()
type Position struct {
	Vector
}

// NewPosition creates a new Position at the given float64 world position.
func NewPosition(x, y float64) *Position {
	p := &Position{NewVector(x, y)}
	p.checkValues()

	return p
}

// Set sets this position to the given x and y values.
func (p *Position) Set(x, y float64) {
	p.x, p.y = x, y
	p.checkValues()
}

func (p *Position) checkValues() {
	if math.IsNaN(p.x) || math.IsNaN(p.y) {
		panic(fmt.Sprintf("float value is NaN: %s", p.Vector))
	}

	if math.IsInf(p.x, 0) || math.IsInf(p.y, 0) {
		panic(fmt.Sprintf("float value is Inf: %s", p.Vector))
	}
}

// World is the position, where 1 = one map tile.
func (p *Position) World() *Vector {
	return &p.Vector
}

// Tile is the tile position, always a whole number.
func (p *Position) Tile() *Vector {
	c := p.World().Clone()
	return c.Floor()
}

// TileOffset is the offset from the tile position, always < 1.
func (p *Position) TileOffset() *Vector {
	c := p.World().Clone()
	return c.Subtract(p.Tile())
}

// SubWorld is the position, where 5 = one map tile.
func (p *Position) SubWorld() *Vector {
	c := p.World().Clone()
	return c.Scale(subTilesPerTile)
}

// SubTile is the tile position in sub tiles, always a multiple of 5.
func (p *Position) SubTile() *Vector {
	c := p.Tile().Clone()
	return c.Scale(subTilesPerTile)
}

// SubTileOffset is the offset from the sub tile position in sub tiles, always < 1.
func (p *Position) SubTileOffset() *Vector {
	c := p.SubWorld().Clone()
	return c.Subtract(p.SubTile())
}
