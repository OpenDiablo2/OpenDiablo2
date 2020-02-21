package d2map

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/beefsack/go-astar"
	"math"
)

// AnimatedEntity represents an entity on the map that can be animated
type AnimatedEntity struct {
	LocationX          float64
	LocationY          float64
	TileX, TileY       int     // Coordinates of the tile the unit is within
	subcellX, subcellY float64 // Subcell coordinates within the current tile
	weaponClass        string
	direction          int
	offsetX, offsetY   int32
	TargetX            float64
	TargetY            float64
	action             int32
	repetitions        int
	path               []astar.Pather

	animation *d2asset.Animation
}

// CreateAnimatedEntity creates an instance of AnimatedEntity
func CreateAnimatedEntity(x, y int32, animation *d2asset.Animation) (*AnimatedEntity, error) {
	entity := &AnimatedEntity{animation: animation}
	entity.LocationX = float64(x)
	entity.LocationY = float64(y)
	entity.TargetX = entity.LocationX
	entity.TargetY = entity.LocationY
	entity.path = []astar.Pather{}

	entity.TileX = int(entity.LocationX / 5)
	entity.TileY = int(entity.LocationY / 5)
	entity.subcellX = 1 + math.Mod(entity.LocationX, 5)
	entity.subcellY = 1 + math.Mod(entity.LocationY, 5)

	return entity, nil
}

func (v *AnimatedEntity) SetPath(path []astar.Pather) {
	v.path = path
}

// Render draws this animated entity onto the target
func (v *AnimatedEntity) Render(target d2render.Surface) {
	target.PushTranslation(
		int(v.offsetX)+int((v.subcellX-v.subcellY)*16),
		int(v.offsetY)+int(((v.subcellX+v.subcellY)*8)-5),
	)
	defer target.Pop()
	v.animation.Render(target)
}

func (v AnimatedEntity) GetDirection() int {
	return v.direction
}

func (v *AnimatedEntity) getStepLength(tickTime float64) (float64, float64) {
	speed := 6.0
	length := tickTime * speed

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

func (v *AnimatedEntity) Step(tickTime float64) {
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
		v.SetTarget(v.path[0].(*PathTile).X*5, v.path[0].(*PathTile).Y*5)

		if len(v.path) > 1 {
			v.path = v.path[1:]
		} else {
			v.path = []astar.Pather{}
		}
		return
	}

}

func (v *AnimatedEntity) HasPathFinding() bool {
	return len(v.path) > 0
}

// SetTarget sets target coordinates and changes animation based on proximity and direction
func (v *AnimatedEntity) SetTarget(tx, ty float64) {
	angle := 359 - d2common.GetAngleBetween(
		v.LocationX,
		v.LocationY,
		tx,
		ty,
	)

	v.TargetX, v.TargetY = tx, ty
	v.direction = angleToDirection(float64(angle), v.animation.GetDirectionCount())
	v.animation.SetDirection(v.direction)

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

func (v *AnimatedEntity) Advance(elapsed float64) {
	v.animation.Advance(elapsed)
}

func (v *AnimatedEntity) GetPosition() (float64, float64) {
	return float64(v.TileX), float64(v.TileY)
}
