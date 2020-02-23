package d2map

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

// AnimatedEntity represents an animation that can be projected onto the map.
type AnimatedEntity struct {
	mapEntity
	direction   int
	action      int
	repetitions int

	animation *d2asset.Animation
}

// CreateAnimatedEntity creates an instance of AnimatedEntity
func CreateAnimatedEntity(x, y int, animation *d2asset.Animation) *AnimatedEntity {
	entity := &AnimatedEntity{
		mapEntity: createMapEntity(x, y),
		animation: animation,
	}
	entity.mapEntity.directioner = entity.rotate
	return entity
}

// Render draws this animated entity onto the target
func (ae *AnimatedEntity) Render(target d2render.Surface) {
	target.PushTranslation(
		ae.offsetX+int((ae.subcellX-ae.subcellY)*16),
		ae.offsetY+int(((ae.subcellX+ae.subcellY)*8)-5),
	)
	defer target.Pop()
	ae.animation.Render(target)
}

func (ae AnimatedEntity) GetDirection() int {
	return ae.direction
}

// rotate sets direction and changes animation
func (ae *AnimatedEntity) rotate(angle float64) {
	ae.direction = angleToDirection(angle, ae.animation.GetDirectionCount())

	var layerDirection int
	switch ae.animation.GetDirectionCount() {
	case 4:
		layerDirection = d2dcc.CofToDir4[ae.direction]
	case 8:
		layerDirection = d2dcc.CofToDir8[ae.direction]
	case 16:
		layerDirection = d2dcc.CofToDir16[ae.direction]
	case 32:
		layerDirection = d2dcc.CofToDir32[ae.direction]
	}

	ae.animation.SetDirection(layerDirection)
}

func (ae *AnimatedEntity) Advance(elapsed float64) {
	ae.animation.Advance(elapsed)
}
