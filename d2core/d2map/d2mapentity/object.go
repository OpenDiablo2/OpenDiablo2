package d2mapentity

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

// Object represents a composite of animations that can be projected onto the map.
type Object struct {
	mapEntity
	composite    *d2asset.Composite
	direction    int
	objectRecord *d2datadict.ObjectRecord
	objectLookup *d2datadict.ObjectLookupRecord
}

// CreateObject creates an instance of AnimatedComposite
func CreateObject(x, y int, object *d2datadict.ObjectLookupRecord, palettePath string) (*Object, error) {
	composite, err := d2asset.LoadComposite(object, palettePath)
	if err != nil {
		return nil, err
	}

	entity := &Object{
		mapEntity:    createMapEntity(x, y),
		composite:    composite,
		objectLookup: object,
	}
	entity.mapEntity.directioner = entity.rotate
	entity.objectRecord = d2datadict.Objects[object.ObjectsTxtId]
	return entity, nil
}

// SetMode changes the graphical mode of this animated entity
func (ob *Object) SetMode(animationMode, weaponClass string, direction int) error {
	ob.composite.SetMode(animationMode, weaponClass, direction)
	ob.direction = direction
	ob.weaponClass = weaponClass

	err := ob.composite.SetMode(animationMode, weaponClass, direction)
	if err != nil {
		err = ob.composite.SetMode(animationMode, "HTH", direction)
		ob.weaponClass = "HTH"
	}

	return err
}

// Render draws this animated entity onto the target
func (ob *Object) Render(target d2render.Surface) {
	target.PushTranslation(
		ob.offsetX+int((ob.subcellX-ob.subcellY)*16),
		ob.offsetY+int(((ob.subcellX+ob.subcellY)*8)-5),
	)
	defer target.Pop()
	ob.composite.Render(target)
}

// rotate sets direction and changes animation
func (ob *Object) rotate(direction int) {
}

func (ob *Object) Advance(elapsed float64) {
	ob.composite.Advance(elapsed)
}
