package d2core

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2data"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/hajimehoshi/ebiten"
)

type NPC struct {
	AnimatedEntity d2render.AnimatedEntity
	Paths          []d2common.Path
}

func CreateNPC(object d2data.Object, fileProvider d2interface.FileProvider) *NPC {
	result := &NPC{
		AnimatedEntity: d2render.CreateAnimatedEntity(object, fileProvider, d2enum.Units),
		Paths:          object.Paths,
	}
	result.AnimatedEntity.SetMode(object.Lookup.Mode, object.Lookup.Class, 1, fileProvider)
	return result
}

func (v *NPC) Render(target *ebiten.Image, offsetX, offsetY int) {
	v.AnimatedEntity.Render(target, offsetX, offsetY)
}
