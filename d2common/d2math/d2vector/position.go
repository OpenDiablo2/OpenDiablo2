package d2vector

import (
	"fmt"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

const (
	subTilesPerTile          float64 = 5  // The number of sub tiles that make up one map tile.
	entityDirectionCount     float64 = 64 // The Diablo equivalent of 360 degrees when dealing with entity rotation.
	entityDirectionIncrement float64 = 8  // One 8th of 64. There 8 possible facing directions for entities.
	// Note: there should be 16 facing directions for entities. See Position.DirectionTo()
)

// Position is a vector in world space. The stored value is the sub tile position.
type Position struct {
	Vector
}

// NewPosition returns a Position struct with the given sub tile coordinates where 1 = 1 sub tile, with a fractional
// offset.
func NewPosition(x, y float64) Position {
	p := Position{*NewVector(x, y)}
	p.checkValues()

	return p
}

// NewPositionTile returns a Position struct with the given tile coordinates where 1 = 1 tile, with a fractional offset.
func NewPositionTile(x, y float64) Position {
	p := Position{*NewVector(x*subTilesPerTile, y*subTilesPerTile)}
	p.checkValues()

	return p
}

// Set sets this position to the given sub tile coordinates where 1 = 1 sub tile, with a fractional offset.
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

// World is the exact position where 1 = one map tile and 0.2 = one sub tile.
func (p *Position) World() *Vector {
	c := p.Clone()
	return c.DivideScalar(subTilesPerTile)
}

// Tile is the position of the current map tile. It is the floor of World(), always a whole number.
func (p *Position) Tile() *Vector {
	return p.World().Floor()
}

// RenderOffset is the offset in sub tiles from the curren tile, + 1. This places the vector at the bottom vertex of an
// isometric diamond visually representing one sub tile. Sub tile indices increase to the lower right diagonal ('down')
// and to the lower left diagonal ('left') of the isometric grid. This renders the target one index above which visually
// is one tile below.
func (p *Position) RenderOffset() *Vector {
	return p.SubTileOffset().AddScalar(1)
}

// SubTileOffset is the offset from the current map tile in sub tiles.
func (p *Position) SubTileOffset() *Vector {
	t := p.Tile().Scale(subTilesPerTile)
	c := p.Clone()

	return c.Subtract(t)
}

// DirectionTo returns the entity direction from this vector to the given vector.
func (v *Vector) DirectionTo(target Vector) int {
	direction := target.Clone()
	direction.Subtract(v)

	angle := direction.SignedAngle(VectorRight())
	radiansPerDirection := d2math.RadFull / entityDirectionCount

	// Note: The direction is always one increment out so we must subtract one increment.
	// This might not work when we implement all 16 directions (as of this writing entities can only face one of 8
	// directions). See entityDirectionIncrement.
	newDirection := int((angle / radiansPerDirection) - entityDirectionIncrement)

	return d2math.WrapInt(newDirection, int(entityDirectionCount))
}
