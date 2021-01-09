// Package d2mapentity implements map entities
package d2mapentity

import (
	"fmt"
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// Object represents a composite of animations that can be projected onto the map.
type Object struct {
	uuid      string
	Position  d2vector.Position
	composite *d2asset.Composite
	highlight bool
	// nameLabel    d2ui.Label
	objectRecord *d2records.ObjectDetailsRecord
	drawLayer    int
	name         string
}

// setMode changes the graphical mode of this animated entity
// nolint:unparam // direction may not always be passed 0 in the future
func (ob *Object) setMode(animationMode d2enum.ObjectAnimationMode, direction int, randomFrame bool) error {
	err := ob.composite.SetMode(animationMode, "HTH")
	if err != nil {
		return err
	}

	ob.composite.SetDirection(direction)

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
		// nolint:gosec // not concerned with crypto-strong randomness
		n := rand.Intn(frameCount)
		ob.composite.SetCurrentFrame(n)
	}

	return err
}

// ID returns the object uuid
func (ob *Object) ID() string {
	return ob.uuid
}

// Highlight sets the entity highlighted flag to true.
func (ob *Object) Highlight() {
	ob.highlight = true
}

// Selectable returns if the object is selectable or not
func (ob *Object) Selectable() bool {
	mode := ob.composite.ObjectAnimationMode()
	return ob.objectRecord.Selectable[mode]
}

// Render draws this animated entity onto the target
func (ob *Object) Render(target d2interface.Surface) {
	renderOffset := ob.Position.RenderOffset()
	target.PushTranslation(
		int((renderOffset.X()-renderOffset.Y())*subtileWidth),
		int((renderOffset.X()+renderOffset.Y())*subtileHeight),
	)

	if ob.highlight {
		target.PushBrightness(highlightBrightness)
		defer target.Pop()
	}

	defer target.Pop()

	if err := ob.composite.Render(target); err != nil {
		fmt.Printf("failed to render composite animation, err: %v\n", err)
	}

	ob.highlight = false
}

// Advance updates the animation
func (ob *Object) Advance(elapsed float64) {
	if err := ob.composite.Advance(elapsed); err != nil {
		fmt.Printf("failed to advance composiste animation, err: %v\n", err)
	}
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

// Label gets the name of the object
func (ob *Object) Label() string {
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

// GetSize returns the current frame size
func (ob *Object) GetSize() (width, height int) {
	return ob.composite.GetSize()
}
