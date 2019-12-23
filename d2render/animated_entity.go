package d2render

import (
	"math"
	"math/rand"

	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"
	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"
	"github.com/OpenDiablo2/D2Shared/d2helper"
	"github.com/OpenDiablo2/OpenDiablo2/d2asset"
	"github.com/hajimehoshi/ebiten"
)

// AnimatedEntity represents an entity on the map that can be animated
type AnimatedEntity struct {
	LocationX          float64
	LocationY          float64
	TileX, TileY       int     // Coordinates of the tile the unit is within
	subcellX, subcellY float64 // Subcell coordinates within the current tile
	animationMode      string
	weaponClass        string
	direction          int
	offsetX, offsetY   int32
	TargetX            float64
	TargetY            float64
	action             int32
	repetitions        int

	composite *d2asset.Composite
}

// CreateAnimatedEntity creates an instance of AnimatedEntity
func CreateAnimatedEntity(x, y int32, object *d2datadict.ObjectLookupRecord, palettePath string) (*AnimatedEntity, error) {
	composite, err := d2asset.LoadComposite(object, palettePath)
	if err != nil {
		return nil, err
	}

	entity := &AnimatedEntity{composite: composite}
	entity.LocationX = float64(x)
	entity.LocationY = float64(y)
	entity.TargetX = entity.LocationX
	entity.TargetY = entity.LocationY

	entity.TileX = int(entity.LocationX / 5)
	entity.TileY = int(entity.LocationY / 5)
	entity.subcellX = 1 + math.Mod(entity.LocationX, 5)
	entity.subcellY = 1 + math.Mod(entity.LocationY, 5)

	return entity, nil
}

// SetMode changes the graphical mode of this animated entity
func (v *AnimatedEntity) SetMode(animationMode, weaponClass string, direction int) error {
	v.animationMode = animationMode
	v.direction = direction

	err := v.composite.SetMode(animationMode, weaponClass, direction)
	if err != nil {
		err = v.composite.SetMode(animationMode, "HTH", direction)
	}

	return err
}

// If an npc has a path to pause at each location.
// Waits for animation to end and all repetitions to be exhausted.
func (v AnimatedEntity) Wait() bool {
	return v.composite.GetPlayedCount() > v.repetitions
}

// Render draws this animated entity onto the target
func (v *AnimatedEntity) Render(target *ebiten.Image, offsetX, offsetY int) {
	localX := (v.subcellX - v.subcellY) * 16
	localY := ((v.subcellX + v.subcellY) * 8) - 5
	v.composite.Render(
		target,
		int(v.offsetX)+offsetX+int(localX),
		int(v.offsetY)+offsetY+int(localY),
	)
}

func (v AnimatedEntity) GetDirection() int {
	return v.direction
}

func (v *AnimatedEntity) getStepLength(tickTime float64) (float64, float64) {
	speed := 6.0
	length := tickTime * speed

	angle := 359 - d2helper.GetAngleBetween(
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

	if d2helper.AlmostEqual(v.LocationX, v.TargetX, stepX) {
		v.LocationX = v.TargetX
	}
	if d2helper.AlmostEqual(v.LocationY, v.TargetY, stepY) {
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

	if v.LocationX == v.TargetX && v.LocationY == v.TargetY {
		v.repetitions = 3 + rand.Intn(5)
		newAnimationMode := d2enum.AnimationModeObjectNeutral
		// TODO: Figure out what 1-3 are for, 4 is correct.
		switch v.action {
		case 1:
			newAnimationMode = d2enum.AnimationModeMonsterNeutral
		case 2:
			newAnimationMode = d2enum.AnimationModeMonsterNeutral
		case 3:
			newAnimationMode = d2enum.AnimationModeMonsterNeutral
		case 4:
			newAnimationMode = d2enum.AnimationModeMonsterSkill1
			v.repetitions = 0
		}

		v.composite.ResetPlayedCount()
		if v.animationMode != newAnimationMode.String() {
			v.SetMode(newAnimationMode.String(), v.weaponClass, v.direction)
		}
	}
}

// SetTarget sets target coordinates and changes animation based on proximity and direction
func (v *AnimatedEntity) SetTarget(tx, ty float64, action int32) {
	angle := 359 - d2helper.GetAngleBetween(
		v.LocationX,
		v.LocationY,
		tx,
		ty,
	)

	v.action = action
	// TODO: Check if is in town and if is player.
	newAnimationMode := d2enum.AnimationModeMonsterWalk.String()
	if tx != v.LocationX || ty != v.LocationY {
		v.TargetX, v.TargetY = tx, ty
		newAnimationMode = d2enum.AnimationModeMonsterWalk.String()
	}

	newDirection := angleToDirection(float64(angle), v.composite.GetDirectionCount())
	if newDirection != v.GetDirection() || newAnimationMode != v.animationMode {
		v.SetMode(newAnimationMode, v.weaponClass, newDirection)
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

func (v *AnimatedEntity) Advance(elapsed float64) {
	v.composite.Advance(elapsed)
}

func (v *AnimatedEntity) GetPosition() (float64, float64) {
	return float64(v.TileX), float64(v.TileY)
}
