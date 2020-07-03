package d2mapentity

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// Object represents a composite of animations that can be projected onto the map.
type Object struct {
	mapEntity
	composite *d2asset.Composite
	direction int
	highlight bool
	// nameLabel    d2ui.Label
	objectLookup *d2datadict.ObjectLookupRecord
	objectRecord *d2datadict.ObjectRecord
	objectType   *d2datadict.ObjectTypeRecord
}

// CreateObject creates an instance of AnimatedComposite
func CreateObject(x, y int, objectRec *d2datadict.ObjectRecord, palettePath string) (*Object, error) {
	entity := &Object{
		mapEntity:    createMapEntity(x, y),
		objectRecord: objectRec,
		objectType:   &d2datadict.ObjectTypes[objectRec.Index],
		// nameLabel:    d2ui.CreateLabel(renderer, d2resource.FontFormal11, d2resource.PaletteStatic),
	}

	equipment := &[d2enum.CompositeTypeMax]string{}
	for i := range equipment {
		equipment[i] = "LIT"
	}

	composite, err := d2asset.LoadComposite("/Data/Global/Objects", entity.objectType.Token,
		d2resource.PaletteUnits, equipment)
	if err != nil {
		return nil, err
	}

	entity.composite = composite

	entity.mapEntity.directioner = entity.rotate
	entity.drawLayer = entity.objectRecord.OrderFlag[d2enum.AnimationModeObjectNeutral]

	entity.setMode("NU", 0)

	// stop torches going crazy for now
	// need initFunc handling to set objects up properly
	if objectRec.HasAnimationMode[d2enum.AnimationModeObjectOpened] {
		entity.setMode("ON", 0)
	}

	return entity, nil
}

// setMode changes the graphical mode of this animated entity
func (ob *Object) setMode(animationMode string, direction int) error {
	ob.direction = direction

	err := ob.composite.SetMode(animationMode, "HTH")
	if err != nil {
		return err
	}

	ob.composite.SetDirection(direction)

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
		// ob.nameLabel.SetText(d2common.TranslateString(ob.objectRecord.Name))
		// ob.nameLabel.SetPosition(-50, -50)
		// ob.nameLabel.Render(target)
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
