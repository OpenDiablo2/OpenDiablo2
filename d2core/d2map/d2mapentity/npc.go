package d2mapentity

import (
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type NPC struct {
	mapEntity
	composite    *d2asset.Composite
	action       int
	HasPaths     bool
	Paths        []d2common.Path
	path         int
	isDone       bool
	repetitions  int
	direction    int
	objectLookup *d2datadict.ObjectLookupRecord
	name string
}

func CreateNPC(x, y int, object *d2datadict.ObjectLookupRecord, direction int) *NPC {
	composite, err := d2asset.LoadComposite(object, d2resource.PaletteUnits)
	if err != nil {
		panic(err)
	}

	result := &NPC{
		mapEntity:    createMapEntity(x, y),
		composite:    composite,
		objectLookup: object,
		HasPaths:     false,
	}
	result.SetMode(object.Mode, object.Class, direction)
	result.mapEntity.directioner = result.rotate

	switch object.Act {
	case 1:
		switch object.Id {
		case 0:
			result.name = "Gheed"
		case 1:
			result.name = "Cain"
		case 2:
			result.name = "Akara"
		case 5:
			result.name = "Kashya"
		case 7:
			result.name = "Warriv"
		case 8:
			result.name = "Charsi"
		case 9:
			result.name = "Andariel"
		}
	}

	return result
}

func (v *NPC) Render(target d2render.Surface) {
	target.PushTranslation(
		v.offsetX+int((v.subcellX-v.subcellY)*16),
		v.offsetY+int(((v.subcellX+v.subcellY)*8)-5),
	)
	defer target.Pop()
	v.composite.Render(target)
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
	v.composite.Advance(tickTime)

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
	newAnimationMode := d2enum.AnimationModeMonsterNeutral
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

// rotate sets direction and changes animation
func (v *NPC) rotate(direction int) {
	var newMode d2enum.MonsterAnimationMode
	if !v.IsAtTarget() {
		newMode = d2enum.AnimationModeMonsterWalk
	} else {
		newMode = d2enum.AnimationModeMonsterNeutral
	}
	if newMode.String() != v.composite.GetAnimationMode() || direction != v.direction {
		v.SetMode(newMode.String(), v.weaponClass, direction)
	}
}

// SetMode changes the graphical mode of this animated entity
func (v *NPC) SetMode(animationMode, weaponClass string, direction int) error {
	v.direction = direction
	v.weaponClass = weaponClass

	err := v.composite.SetMode(animationMode, weaponClass, direction)
	if err != nil {
		err = v.composite.SetMode(animationMode, "HTH", direction)
		v.weaponClass = "HTH"
	}

	return err
}

func (m *NPC) Name() string {
	return m.name
}
