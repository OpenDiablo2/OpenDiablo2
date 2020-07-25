// Package d2object implements objects placed on the map and their functionality
package d2object

import (
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// Object represents a composite of animations that can be projected onto the map.
type Object struct {
	Position  d2vector.Position
	composite *d2asset.Composite
	highlight bool
	// nameLabel    d2ui.Label
	objectRecord *d2datadict.ObjectRecord
	drawLayer    int
	name         string
}

// CreateObject creates an instance of AnimatedComposite
func CreateObject(x, y int, objectRec *d2datadict.ObjectRecord, palettePath string) (*Object, error) {
	locX, locY := float64(x), float64(y)
	entity := &Object{
		objectRecord: objectRec,
		Position:     d2vector.NewPosition(locX, locY),
		name:         d2common.TranslateString(objectRec.Name),
	}
	objectType := &d2datadict.ObjectTypes[objectRec.Index]

	composite, err := d2asset.LoadComposite(d2enum.ObjectTypeItem, objectType.Token,
		d2resource.PaletteUnits)
	if err != nil {
		return nil, err
	}

	entity.composite = composite

	entity.setMode(d2enum.ObjectAnimationModeNeutral, 0, false)

	initObject(entity)

	return entity, nil
}

// setMode changes the graphical mode of this animated entity
func (ob *Object) setMode(animationMode d2enum.ObjectAnimationMode, direction int, randomFrame bool) error {
	err := ob.composite.SetMode(animationMode, "HTH")
	if err != nil {
		return err
	}

	ob.composite.SetDirection(direction)

	// mode := d2enum.ObjectAnimationModeFromString(animationMode)
	ob.drawLayer = ob.objectRecord.OrderFlag[d2enum.ObjectAnimationModeNeutral]

	// For objects their txt record entry overrides animationdata
	speed := ob.objectRecord.FrameDelta[animationMode]
	if speed != 0 {
		ob.composite.SetAnimSpeed(speed)
	}

	frameCount := ob.objectRecord.FrameCount[animationMode]

	if frameCount != 0 {
		ob.composite.SetSubLoop(0, frameCount)
	}

	ob.composite.SetPlayLoop(ob.objectRecord.CycleAnimation[animationMode])
	ob.composite.SetCurrentFrame(ob.objectRecord.StartFrame[animationMode])

	if randomFrame {
		n := rand.Intn(frameCount)
		ob.composite.SetCurrentFrame(n)
	}

	return err
}

// Highlight sets the entity highlighted flag to true.
func (ob *Object) Highlight() {
	ob.highlight = true
}

func (ob *Object) Selectable() bool {
	mode := ob.composite.ObjectAnimationMode()
	return ob.objectRecord.Selectable[mode]
}

// Render draws this animated entity onto the target
func (ob *Object) Render(target d2interface.Surface) {
	renderOffset := ob.Position.RenderOffset()
	target.PushTranslation(
		int((renderOffset.X()-renderOffset.Y())*16),
		int(((renderOffset.X() + renderOffset.Y()) * 8)),
	)

	if ob.highlight {
		target.PushBrightness(2)
		defer target.Pop()
	}

	defer target.Pop()
	ob.composite.Render(target)
	ob.highlight = false
}

// Advance updates the animation
func (ob *Object) Advance(elapsed float64) {
	ob.composite.Advance(elapsed)
}

// GetLayer returns which layer of the map the object is drawn
func (ob *Object) GetLayer() int {
	return ob.drawLayer
}

// GetPositionF of the object but differently
func (ob *Object) GetPositionF() (x, y float64) {
	w := ob.Position.World()
	return w.X(), w.Y()
}

// Name gets the name of the object
func (ob *Object) Name() string {
	return ob.name
}

// GetPosition returns the object's position
func (ob *Object) GetPosition() d2vector.Position {
	return ob.Position
}

// GetVelocity returns the object's velocity vector
func (ob *Object) GetVelocity() d2vector.Vector {
	return *d2vector.VectorZero()
}
