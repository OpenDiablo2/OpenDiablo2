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
	action      int32
	repetitions int

	animation *d2asset.Animation
}

// CreateAnimatedEntity creates an instance of AnimatedEntity
func CreateAnimatedEntity(x, y int32, animation *d2asset.Animation) *AnimatedEntity {
	entity := &AnimatedEntity{
		mapEntity: createMapEntity(x, y),
		animation: animation,
	}
	entity.mapEntity.directioner = entity.setDirection
	return entity
}

// Render draws this animated entity onto the target
func (v *AnimatedEntity) Render(target d2render.Surface) {
	target.PushTranslation(
		int(v.offsetX)+int((v.subcellX-v.subcellY)*16),
		int(v.offsetY)+int(((v.subcellX+v.subcellY)*8)-5),
	)
	defer target.Pop()
	v.animation.Render(target)
}

func (v AnimatedEntity) GetDirection() int {
	return v.direction
}

// SetTarget sets target coordinates and changes animation based on proximity and direction
func (v *AnimatedEntity) setDirection(angle float64) {
	v.direction = angleToDirection(angle, v.animation.GetDirectionCount())

	var layerDirection int
	switch v.animation.GetDirectionCount() {
	case 4:
		layerDirection = d2dcc.CofToDir4[v.direction]
	case 8:
		layerDirection = d2dcc.CofToDir8[v.direction]
	case 16:
		layerDirection = d2dcc.CofToDir16[v.direction]
	case 32:
		layerDirection = d2dcc.CofToDir32[v.direction]
	}

	v.animation.SetDirection(layerDirection)
}

func (v *AnimatedEntity) Advance(elapsed float64) {
	v.animation.Advance(elapsed)
}
