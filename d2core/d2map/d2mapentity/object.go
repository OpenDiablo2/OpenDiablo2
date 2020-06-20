package d2mapentity

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

// Object represents a composite of animations that can be projected onto the map.
type Object struct {
	mapEntity
	composite    *d2asset.Composite
	direction    int
	highlight    bool
	nameLabel    d2ui.Label
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
		nameLabel:    d2ui.CreateLabel(d2resource.FontFormal11, d2resource.PaletteStatic),
	}
	entity.mapEntity.directioner = entity.rotate
	entity.objectRecord = d2datadict.Objects[object.ObjectsTxtId]
	entity.drawLayer = entity.objectRecord.OrderFlag[d2enum.AnimationModeObjectNeutral]
	entity.SetMode(object.Mode, object.Class, 0)

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

	mode := d2enum.ObjectAnimationModeFromString(animationMode)
	ob.mapEntity.drawLayer = ob.objectRecord.OrderFlag[mode]

	// For objects their txt record entry overrides animationdata
	speed := ob.objectRecord.FrameDelta[mode]
	if speed != 0 {
		ob.composite.SetSpeed(speed)
	}

	return err
}

func (ob *Object) Highlight() {
	ob.highlight = true
}

func (ob *Object) Selectable() bool {
	mode := d2enum.ObjectAnimationModeFromString(ob.composite.GetAnimationMode())
	return ob.objectRecord.Selectable[mode]
}

// Render draws this animated entity onto the target
func (ob *Object) Render(target d2interface.Surface) {
	target.PushTranslation(
		ob.offsetX+int((ob.subcellX-ob.subcellY)*16),
		ob.offsetY+int(((ob.subcellX+ob.subcellY)*8)-5),
	)
	if ob.highlight {
		ob.nameLabel.SetText(d2common.TranslateString(ob.objectRecord.Name))
		ob.nameLabel.SetPosition(-50, -50)
		ob.nameLabel.Render(target)
		target.PushBrightness(2)
		defer target.Pop()
	}
	defer target.Pop()
	ob.composite.Render(target)
	ob.highlight = false
}

// rotate sets direction and changes animation
func (ob *Object) rotate(direction int) {
}

func (ob *Object) Advance(elapsed float64) {
	ob.composite.Advance(elapsed)
}
