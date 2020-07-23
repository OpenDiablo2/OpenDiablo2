package d2mapentity

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// AnimatedEntity represents an animation that can be projected onto the map.
type AnimatedEntity struct {
	mapEntity
	direction   int
	action      int
	repetitions int

	animation d2interface.Animation
}

// CreateAnimatedEntity creates an instance of AnimatedEntity
func CreateAnimatedEntity(x, y int, animation d2interface.Animation) *AnimatedEntity {
	entity := &AnimatedEntity{
		mapEntity: newMapEntity(x, y),
		animation: animation,
	}
	entity.mapEntity.directioner = entity.rotate

	return entity
}

// Render draws this animated entity onto the target
func (ae *AnimatedEntity) Render(target d2interface.Surface) {
	renderOffset := ae.Position.RenderOffset()
	target.PushTranslation(
		int((renderOffset.X()-renderOffset.Y())*16),
		int(((renderOffset.X()+renderOffset.Y())*8)-5),
	)

	defer target.Pop()

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
