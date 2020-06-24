package d2mapentity

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

// AnimatedComposite represents a composite of animations that can be projected onto the map.
type AnimatedComposite struct {
	mapEntity
	composite    *d2asset.Composite
	direction    int
	player       *Player
	objectLookup *d2datadict.ObjectLookupRecord
}

// CreateAnimatedComposite creates an instance of AnimatedComposite
func CreateAnimatedComposite(x, y int, object *d2datadict.ObjectLookupRecord, palettePath string) (*AnimatedComposite, error) {
	composite, err := d2asset.LoadComposite(object, palettePath)
	if err != nil {
		return nil, err
	}

	entity := &AnimatedComposite{
		mapEntity:    createMapEntity(x, y),
		composite:    composite,
		objectLookup: object,
	}
	entity.mapEntity.directioner = entity.rotate
	return entity, nil
}

// SetMode changes the graphical mode of this animated entity
func (ac *AnimatedComposite) SetMode(animationMode, weaponClass string, direction int) error {
	ac.composite.SetMode(animationMode, weaponClass, direction)
	ac.direction = direction
	ac.weaponClass = weaponClass

	err := ac.composite.SetMode(animationMode, weaponClass, direction)
	if err != nil {
		err = ac.composite.SetMode(animationMode, "HTH", direction)
		ac.weaponClass = "HTH"
	}

	return err
}

// Render draws this animated entity onto the target
func (ac *AnimatedComposite) Render(target d2render.Surface) {
	target.PushTranslation(
		ac.offsetX+int((ac.subcellX-ac.subcellY)*16),
		ac.offsetY+int(((ac.subcellX+ac.subcellY)*8)-5),
	)
	defer target.Pop()
	ac.composite.Render(target)
}

// rotate sets direction and changes animation
func (ac *AnimatedComposite) rotate(direction int) {
}

func (ac *AnimatedComposite) Advance(elapsed float64) {
	ac.composite.Advance(elapsed)
}
