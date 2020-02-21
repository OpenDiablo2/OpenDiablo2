package d2map

import (
	"github.com/beefsack/go-astar"
	"math"
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

// AnimatedComposite represents an entity on the map that can be animated
type AnimatedComposite struct {
	AnimatedEntity

	animationMode string
	composite     *d2asset.Composite
}

// CreateAnimatedComposite creates an instance of AnimatedComposite
func CreateAnimatedComposite(x, y int32, object *d2datadict.ObjectLookupRecord, palettePath string) (*AnimatedComposite, error) {
	composite, err := d2asset.LoadComposite(object, palettePath)
	if err != nil {
		return nil, err
	}

	entity := &AnimatedComposite{composite: composite}
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

// SetMode changes the graphical mode of this animated entity
func (v *AnimatedComposite) SetMode(animationMode, weaponClass string, direction int) error {
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
func (v AnimatedComposite) Wait() bool {
	return v.composite.GetPlayedCount() > v.repetitions
}

// Render draws this animated entity onto the target
func (v *AnimatedComposite) Render(target d2render.Surface) {
	target.PushTranslation(
		int(v.offsetX)+int((v.subcellX-v.subcellY)*16),
		int(v.offsetY)+int(((v.subcellX+v.subcellY)*8)-5),
	)
	defer target.Pop()
	v.composite.Render(target)
}

func (v *AnimatedComposite) Step(tickTime float64) {
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
		v.SetTarget(v.path[0].(*PathTile).X*5, v.path[0].(*PathTile).Y*5, 1)

		if len(v.path) > 1 {
			v.path = v.path[1:]
		} else {
			v.path = []astar.Pather{}
		}
		return
	}

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

// SetTarget sets target coordinates and changes animation based on proximity and direction
func (v *AnimatedComposite) SetTarget(tx, ty float64, action int32) {
	angle := 359 - d2common.GetAngleBetween(
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

	if newAnimationMode != v.animationMode {
		v.SetMode(newAnimationMode, v.weaponClass, v.direction)
	}

	newDirection := angleToDirection(float64(angle), v.composite.GetDirectionCount())

	if newDirection != v.GetDirection() {
		v.SetMode(v.animationMode, v.weaponClass, newDirection)
	}
}

func (v *AnimatedComposite) Advance(elapsed float64) {
	v.composite.Advance(elapsed)
}

func (v *AnimatedComposite) GetPosition() (float64, float64) {
	return float64(v.TileX), float64(v.TileY)
}
