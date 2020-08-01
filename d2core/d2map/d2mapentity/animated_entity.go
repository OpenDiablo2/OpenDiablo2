package d2mapentity

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// AnimatedEntity represents an animation that can be projected onto the map.
type AnimatedEntity struct {
	mapEntity
	animation d2interface.Animation

	direction   int
	action      int
	repetitions int

	highlight bool
}

// Render draws this animated entity onto the target
func (ae *AnimatedEntity) Render(target d2interface.Surface) {
	renderOffset := ae.Position.RenderOffset()
	ox, oy := renderOffset.X(), renderOffset.Y()
	tx, ty := int((ox-oy)*16), int((ox+oy)*8)-5

	target.PushTranslation(tx, ty)

	defer target.Pop()

	if ae.highlight {
		target.PushBrightness(2)
		defer target.Pop()

		ae.highlight = false
	}

	if err := ae.animation.Render(target); err != nil {
		fmt.Printf("failed to render animated entity, err: %v\n", err)
	}
}

// GetDirection returns the current facing direction of this entity.
func (ae *AnimatedEntity) GetDirection() int {
	return ae.direction
}

// rotate sets direction and changes animation
func (ae *AnimatedEntity) rotate(direction int) {
	ae.direction = direction

	if err := ae.animation.SetDirection(ae.direction); err != nil {
		fmt.Printf("failed to update the animation direction, err: %v\n", err)
	}
}

// Advance is called once per frame and processes a
// single game tick.
func (ae *AnimatedEntity) Advance(elapsed float64) {
	if err := ae.animation.Advance(elapsed); err != nil {
		fmt.Printf("failed to advance the animation, err: %v\n", err)
	}
}

// SetHighlight sets the highlight state of the animated entity
func (ae *AnimatedEntity) SetHighlight(set bool) {
	ae.highlight = set
}
