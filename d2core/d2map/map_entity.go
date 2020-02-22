package d2map

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/beefsack/go-astar"
	"math"
)

// mapEntity represents an entity on the map that can be animated
type mapEntity struct {
	LocationX          float64
	LocationY          float64
	TileX, TileY       int     // Coordinates of the tile the unit is within
	subcellX, subcellY float64 // Subcell coordinates within the current tile
	weaponClass        string
	offsetX, offsetY   int32
	TargetX            float64
	TargetY            float64
	Speed              float64
	path               []astar.Pather

	done        func()
	directioner func(angle float64)
}

// createMapEntity creates an instance of mapEntity
func createMapEntity(x, y int32) mapEntity {
	locX, locY := float64(x), float64(y)
	return mapEntity{
		LocationX: locX,
		LocationY: locY,
		TargetX:   locX,
		TargetY:   locY,
		TileX:     int(locX / 5),
		TileY:     int(locY / 5),
		subcellX:  1 + math.Mod(locX, 5),
		subcellY:  1 + math.Mod(locY, 5),
		Speed:     6,
		path:      []astar.Pather{},
	}
}

func (v *mapEntity) SetPath(path []astar.Pather, done func()) {
	v.path = path
	v.done = done
}

func (v *mapEntity) getStepLength(tickTime float64) (float64, float64) {
	length := tickTime * v.Speed

	angle := 359 - d2common.GetAngleBetween(
		v.LocationX,
		v.LocationY,
		v.TargetX,
		v.TargetY,
	)
	radians := (math.Pi / 180.0) * float64(angle)
	oneStepX := length * math.Cos(radians)
	oneStepY := length * math.Sin(radians)
	return oneStepX, oneStepY
}

func (v *mapEntity) IsAtTarget() bool {
	return v.LocationX == v.TargetX && v.LocationY == v.TargetY && !v.HasPathFinding()
}

func (v *mapEntity) Step(tickTime float64) {
	if v.IsAtTarget() {
		if v.done != nil {
			v.done()
			v.done = nil
		}
		return
	}

	stepX, stepY := v.getStepLength(tickTime)

	if d2common.AlmostEqual(v.LocationX, v.TargetX, stepX) {
		v.LocationX = v.TargetX
	}
	if d2common.AlmostEqual(v.LocationY, v.TargetY, stepY) {
		v.LocationY = v.TargetY
	}
	if v.LocationX != v.TargetX {
		v.LocationX += stepX
	}
	if v.LocationY != v.TargetY {
		v.LocationY += stepY
	}

	v.subcellX = 1 + math.Mod(v.LocationX, 5)
	v.subcellY = 1 + math.Mod(v.LocationY, 5)
	v.TileX = int(v.LocationX / 5)
	v.TileY = int(v.LocationY / 5)

	if (v.LocationX != v.TargetX) || (v.LocationY != v.TargetY) {
		return
	}

	if len(v.path) > 0 {
		v.SetTarget(v.path[0].(*PathTile).X*5, v.path[0].(*PathTile).Y*5, v.done)

		if len(v.path) > 1 {
			v.path = v.path[1:]
		} else {
			v.path = []astar.Pather{}
		}
		return
	}

}

func (v *mapEntity) HasPathFinding() bool {
	return len(v.path) > 0
}

// SetTarget sets target coordinates and changes animation based on proximity and direction
func (v *mapEntity) SetTarget(tx, ty float64, done func()) {
	v.TargetX, v.TargetY = tx, ty
	v.done = done

	if v.directioner != nil {
		angle := 359 - d2common.GetAngleBetween(
			v.LocationX,
			v.LocationY,
			tx,
			ty,
		)
		v.directioner(float64(angle))
	}
}

func angleToDirection(angle float64, numberOfDirections int) int {
	if numberOfDirections == 0 {
		return 0
	}

	degreesPerDirection := 360.0 / float64(numberOfDirections)
	offset := 45.0 - (degreesPerDirection / 2)
	newDirection := int((angle - offset) / degreesPerDirection)
	if newDirection >= numberOfDirections {
		newDirection = newDirection - numberOfDirections
	} else if newDirection < 0 {
		newDirection = numberOfDirections + newDirection
	}

	return newDirection
}

func (v *mapEntity) GetPosition() (float64, float64) {
	return float64(v.TileX), float64(v.TileY)
}
