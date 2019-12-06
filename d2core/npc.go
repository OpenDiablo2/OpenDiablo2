package d2core

import (
	"github.com/OpenDiablo2/D2Shared/d2common"
	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"
	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"
	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/hajimehoshi/ebiten"
)

type NPC struct {
	AnimatedEntity d2render.AnimatedEntity
	HasPaths       bool
	Paths          []d2common.Path
	path           int
}

func CreateNPC(x, y int32, object *d2datadict.ObjectLookupRecord, fileProvider d2interface.FileProvider, direction int) *NPC {
	result := &NPC{
		AnimatedEntity: d2render.CreateAnimatedEntity(x, y, object, fileProvider, d2enum.Units),
		HasPaths:       false,
	}
	result.AnimatedEntity.SetMode(object.Mode, object.Class, direction)
	return result
}

func (v *NPC) Path() d2common.Path {
	return v.Paths[v.path]
}

func (v *NPC) NextPath() d2common.Path {
	v.path++
	if v.path == len(v.Paths) {
		v.path = 0
	}

	return v.Paths[v.path]
}

func (v *NPC) SetPaths(paths []d2common.Path) {
	v.Paths = paths
	v.HasPaths = len(paths) > 0
}

func (v *NPC) Render(target *ebiten.Image, offsetX, offsetY int) {
	v.AnimatedEntity.Render(target, offsetX, offsetY)
}
