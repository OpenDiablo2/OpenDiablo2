package d2core

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/hajimehoshi/ebiten"
)

type NPC struct {
	AnimatedEntity d2render.AnimatedEntity
	Paths          []d2common.Path
}

func CreateNPC(x, y int32, object *d2datadict.ObjectLookupRecord, fileProvider d2interface.FileProvider, direction int) *NPC {
	result := &NPC{
		AnimatedEntity: d2render.CreateAnimatedEntity(x, y, object, fileProvider, d2enum.Units),
	}
	result.AnimatedEntity.SetMode(object.Mode, object.Class, direction, fileProvider)
	return result
}

func (v *NPC) SetPaths(paths []d2common.Path) {
	v.Paths = paths
}

func (v *NPC) Render(target *ebiten.Image, offsetX, offsetY int) {
	v.AnimatedEntity.Render(target, offsetX, offsetY)
}
