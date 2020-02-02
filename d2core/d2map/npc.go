package d2map

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type NPC struct {
	AnimatedEntity *AnimatedEntity
	HasPaths       bool
	Paths          []d2common.Path
	path           int
}

func CreateNPC(x, y int32, object *d2datadict.ObjectLookupRecord, direction int) *NPC {
	entity, err := CreateAnimatedEntity(x, y, object, d2resource.PaletteUnits)
	if err != nil {
		panic(err)
	}

	result := &NPC{AnimatedEntity: entity, HasPaths: false}
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

func (v *NPC) Render(target d2render.Surface) {
	v.AnimatedEntity.Render(target)
}

func (v *NPC) GetPosition() (float64, float64) {
	return v.AnimatedEntity.GetPosition()
}

func (v *NPC) Advance(tickTime float64) {
	if v.HasPaths &&
		v.AnimatedEntity.LocationX == v.AnimatedEntity.TargetX &&
		v.AnimatedEntity.LocationY == v.AnimatedEntity.TargetY &&
		v.AnimatedEntity.Wait() {
		// If at the target, set target to the next path.
		path := v.NextPath()
		v.AnimatedEntity.SetTarget(
			float64(path.X),
			float64(path.Y),
			path.Action,
		)
	}

	if v.AnimatedEntity.LocationX != v.AnimatedEntity.TargetX ||
		v.AnimatedEntity.LocationY != v.AnimatedEntity.TargetY {
		v.AnimatedEntity.Step(tickTime)
	}

	v.AnimatedEntity.Advance(tickTime)
}
