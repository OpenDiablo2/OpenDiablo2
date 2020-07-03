package d2mapentity

import (
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

type NPC struct {
	mapEntity
	composite     *d2asset.Composite
	action        int
	HasPaths      bool
	Paths         []d2common.Path
	path          int
	isDone        bool
	repetitions   int
	direction     int
	monstatRecord *d2datadict.MonStatsRecord
	monstatEx     *d2datadict.MonStats2Record
	name          string
}

func CreateNPC(x, y int, monstat *d2datadict.MonStatsRecord, direction int) *NPC {
	result := &NPC{
		mapEntity:     createMapEntity(x, y),
		HasPaths:      false,
		monstatRecord: monstat,
		monstatEx:     d2datadict.MonStats2[monstat.ExtraDataKey],
	}

	equipment := &[d2enum.CompositeTypeMax]string{
		selectEquip(result.monstatEx.HDv),
		selectEquip(result.monstatEx.TRv),
		selectEquip(result.monstatEx.LGv),
		selectEquip(result.monstatEx.Rav),
		selectEquip(result.monstatEx.Lav),
		selectEquip(result.monstatEx.RHv),
		selectEquip(result.monstatEx.LHv),
		selectEquip(result.monstatEx.SHv),
		selectEquip(result.monstatEx.S1v),
		selectEquip(result.monstatEx.S2v),
		selectEquip(result.monstatEx.S3v),
		selectEquip(result.monstatEx.S4v),
		selectEquip(result.monstatEx.S5v),
		selectEquip(result.monstatEx.S6v),
		selectEquip(result.monstatEx.S7v),
		selectEquip(result.monstatEx.S8v),
	}

	composite, err := d2asset.LoadComposite("/Data/Global/Monsters", monstat.AnimationDirectoryToken,
		d2resource.PaletteUnits, equipment)
	result.composite = composite

	if err != nil {
		panic(err)
	}

	result.setMode("NU", result.monstatEx.BaseWeaponClass, 0)

	result.mapEntity.directioner = result.rotate

	if result.monstatRecord != nil && result.monstatRecord.IsInteractable {
		result.name = d2common.TranslateString(result.monstatRecord.NameStringTableKey)
	}

	return result
}

func selectEquip(slice []string) string {
	if len(slice) != 0 {
		return slice[rand.Intn(len(slice))]
	}

	return ""
}

func (v *NPC) Render(target d2interface.Surface) {
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
		v.setMode(newAnimationMode.String(), v.weaponClass, v.direction)
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
		v.setMode(newMode.String(), v.weaponClass, direction)
	}
}

// setMode changes the graphical mode of this animated entity
func (v *NPC) setMode(animationMode, weaponClass string, direction int) error {
	v.direction = direction
	v.weaponClass = weaponClass

	err := v.composite.SetMode(animationMode, weaponClass)
	if err != nil {
		err = v.composite.SetMode(animationMode, "HTH")
		v.weaponClass = "HTH"
	}

	v.composite.SetDirection(direction)

	return err
}

func (m *NPC) Selectable() bool {
	// is there something handy that determines selectable npc's?
	if m.name != "" {
		return true
	}
	return false
}

func (m *NPC) Name() string {
	return m.name
}
