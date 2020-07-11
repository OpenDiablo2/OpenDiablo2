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
func NewPosition(x, y float64) Position {
	p := Position{NewVector(x, y)}
	p.checkValues()

	return p
}

// EntityPosition returns a Position struct based on the given entity spawn point.
// The value given should be the one set in d2mapstamp.Stamp.Entities:
// (tileOffsetX*5)+object.X, (tileOffsetY*5)+object.Y
// TODO: This probably doesn't support positions in between sub tiles so will only be suitable for spawning entities from map generation, not for multiplayer syncing.
func EntityPosition(x, y float64) Position {
	return NewPosition(x/5, y/5)
}

// Set sets this position to the given x and y world position.
func (p *Position) Set(x, y float64) {
	p.x, p.y = x, y
	p.checkValues()
}

// TODO: test this
// SetSubWorld sets this position to the given x and y sub tile coordinates.
func (p *Position) SetSubWorld(x, y float64) {
	p.x, p.y = x/5, y/5
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

// World is the position, where 1 = one map tile. This is a pointer to the Position value and must be cloned with
// Clone() if the position is not to be changed.
func (p *Position) World() *Vector {
	return &p.Vector
}

// Tile is the tile position, always a whole number. (tileX, tileY)
func (p *Position) Tile() *Vector {
	c := p.World().Clone()
	return c.Floor()
}

// TileOffset is the offset from the tile position, always < 1.
// unused
func (p *Position) TileOffset() *Vector {
	c := p.World().Clone()
	return c.Subtract(p.Tile())
}

// WorldSubTile is the position, where 5 = one map tile. (locationX, locationY)
func (p *Position) WorldSubTile() *Vector {
	c := p.World().Clone()
	return c.Scale(subTilesPerTile)
}

// TileSubTile is the tile position in sub tiles, always a multiple of 5.
// unused
func (p *Position) TileSubTile() *Vector {
	return p.Tile().Scale(subTilesPerTile)
}

// SubTileOffset is the offset from the sub tile position in sub tiles, always < 1.
// unused
func (p *Position) SubTileOffset() *Vector {
	return p.WorldSubTile().Subtract(p.TileSubTile())
}

// This original value here was always zero. It is never assigned to but it is used. (offsetX, offsetY)
func (p *Position) Offset() *Vector {
	v := VectorZero()
	return &v
}

// TODO: understand this and maybe improve/remove it
// SubTileOffset() + 1. It's used for rendering. It seems to always do this:
// 	v.offsetX+int((v.subcellX-v.subcellY)*16),
//	v.offsetY+int(((v.subcellX+v.subcellY)*8)-5),
// ^ Maybe similar to Viewport.OrthoToWorld? (subCellX, subCellY)
func (p *Position) SubCell() *Vector {
	return p.SubTileOffset().AddScalar(1)
}

func (p Position) String() string {
	return fmt.Sprintf("World: %s\nTile: %s\nSubWorld: %s\nSubCell: %s", p.World(), p.Tile(), p.WorldSubTile(), p.SubCell())
}
