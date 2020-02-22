package d2map

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

// AnimatedComposite represents a composite of animations that can be projected onto the map.
type AnimatedComposite struct {
	mapEntity
	animationMode string
	composite     *d2asset.Composite
	direction     int
}

// CreateAnimatedComposite creates an instance of AnimatedComposite
func CreateAnimatedComposite(x, y int32, object *d2datadict.ObjectLookupRecord, palettePath string) (*AnimatedComposite, error) {
	composite, err := d2asset.LoadComposite(object, palettePath)
	if err != nil {
		return nil, err
	}

	entity := &AnimatedComposite{
		mapEntity: createMapEntity(x, y),
		composite: composite,
	}
	entity.mapEntity.directioner = entity.setDirection
	return entity, nil
}

func (v *AnimatedComposite) SetAnimationMode(animationMode string) error {
	v.animationMode = animationMode
	return v.composite.SetMode(animationMode, v.weaponClass, v.direction)
}

// SetMode changes the graphical mode of this animated entity
func (v *AnimatedComposite) SetMode(animationMode, weaponClass string, direction int) error {
	v.animationMode = animationMode
	v.direction = direction

	err := v.composite.SetMode(animationMode, weaponClass, direction)
	if err != nil {
		err = v.composite.SetMode(animationMode, "HTH", direction)
		v.weaponClass = "HTH"
	}

	return err
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

// setDirection changes animation based on proximity and direction
func (v *AnimatedComposite) setDirection(angle float64) {
	// TODO: Check if is in town and if is player.
	newAnimationMode := v.animationMode
	if !v.IsAtTarget() {
		newAnimationMode = d2enum.AnimationModeMonsterWalk.String()
	}

	if newAnimationMode != v.animationMode {
		v.SetMode(newAnimationMode, v.weaponClass, v.direction)
	}

	newDirection := angleToDirection(angle, v.composite.GetDirectionCount())
	if newDirection != v.direction {
		v.SetMode(v.animationMode, v.weaponClass, newDirection)
	}
}

func (v *AnimatedComposite) Advance(elapsed float64) {
	v.composite.Advance(elapsed)
}
