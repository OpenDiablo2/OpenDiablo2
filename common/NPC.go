package common

import (
	"github.com/OpenDiablo2/OpenDiablo2/palettedefs"
	"github.com/hajimehoshi/ebiten"
)

type NPC struct {
	AnimatedEntity *AnimatedEntity
	Paths          []Path
}

func CreateNPC(object Object, fileProvider FileProvider) *NPC {
	result := &NPC{
		AnimatedEntity: CreateAnimatedEntity(object, fileProvider, palettedefs.Units),
		Paths:          object.Paths,
	}
	result.AnimatedEntity.SetMode(object.Lookup.Mode, object.Lookup.Class, 0, fileProvider)

	return result
}

func (v *NPC) Render(target *ebiten.Image, offsetX, offsetY int) {
	v.AnimatedEntity.Render(target, offsetX, offsetY)
}
