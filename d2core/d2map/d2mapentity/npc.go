package d2mapentity

import (
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

type NPC struct {
	*AnimatedComposite
	action      int
	HasPaths    bool
	Paths       []d2common.Path
	path        int
	isDone      bool
	repetitions int
}

func CreateNPC(x, y int, object *d2datadict.ObjectLookupRecord, direction int) *NPC {
	entity, err := CreateAnimatedComposite(x, y, object, d2resource.PaletteUnits)
	if err != nil {
		panic(err)
	}

	result := &NPC{AnimatedComposite: entity, HasPaths: false}
	result.SetMode(object.Mode, object.Class, direction)
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
	v.isDone = true
}

func (v *NPC) Advance(tickTime float64) {
	v.Step(tickTime)
	v.AnimatedComposite.Advance(tickTime)

	if v.HasPaths && v.wait() {
		// If at the target, set target to the next path.
		v.isDone = false
		path := v.NextPath()
		v.SetTarget(
			float64(path.X),
			float64(path.Y),
			v.next,
		)
		v.action = path.Action
	}
}

// If an npc has a path to pause at each location.
// Waits for animation to end and all repetitions to be exhausted.
func (v *NPC) wait() bool {
	return v.isDone && v.composite.GetPlayedCount() > v.repetitions
}

func (v *NPC) next() {
	v.isDone = true
	v.repetitions = 3 + rand.Intn(5)
	newAnimationMode := d2enum.AnimationModeObjectNeutral
	// TODO: Figure out what 1-3 are for, 4 is correct.
	switch v.action {
	case 1:
		newAnimationMode = d2enum.AnimationModeMonsterNeutral
	case 2:
		newAnimationMode = d2enum.AnimationModeMonsterNeutral
	case 3:
		newAnimationMode = d2enum.AnimationModeMonsterNeutral
	case 4:
		newAnimationMode = d2enum.AnimationModeMonsterSkill1
		v.repetitions = 0
	default:
		v.repetitions = 0
	}

	if v.composite.GetAnimationMode() != newAnimationMode.String() {
		v.SetMode(newAnimationMode.String(), v.weaponClass, v.direction)
	}
}
