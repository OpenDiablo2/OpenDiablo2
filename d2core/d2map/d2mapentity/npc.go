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
		d2enum.CompositeTypeHead:      selectEquip(result.monstatEx.HDv),
		d2enum.CompositeTypeTorso:     selectEquip(result.monstatEx.TRv),
		d2enum.CompositeTypeLegs:      selectEquip(result.monstatEx.LGv),
		d2enum.CompositeTypeRightArm:  selectEquip(result.monstatEx.Rav),
		d2enum.CompositeTypeLeftArm:   selectEquip(result.monstatEx.Lav),
		d2enum.CompositeTypeRightHand: selectEquip(result.monstatEx.RHv),
		d2enum.CompositeTypeLeftHand:  selectEquip(result.monstatEx.LHv),
		d2enum.CompositeTypeShield:    selectEquip(result.monstatEx.SHv),
		d2enum.CompositeTypeSpecial1:  selectEquip(result.monstatEx.S1v),
		d2enum.CompositeTypeSpecial2:  selectEquip(result.monstatEx.S2v),
		d2enum.CompositeTypeSpecial3:  selectEquip(result.monstatEx.S3v),
		d2enum.CompositeTypeSpecial4:  selectEquip(result.monstatEx.S4v),
		d2enum.CompositeTypeSpecial5:  selectEquip(result.monstatEx.S5v),
		d2enum.CompositeTypeSpecial6:  selectEquip(result.monstatEx.S6v),
		d2enum.CompositeTypeSpecial7:  selectEquip(result.monstatEx.S7v),
		d2enum.CompositeTypeSpecial8:  selectEquip(result.monstatEx.S8v),
	}

	composite, _ := d2asset.LoadComposite(d2enum.ObjectTypeCharacter, monstat.AnimationDirectoryToken,
		d2resource.PaletteUnits)
	result.composite = composite

	composite.SetMode("NU", result.monstatEx.BaseWeaponClass)
	composite.Equip(equipment)

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
		v.composite.SetMode(newAnimationMode.String(), v.composite.GetWeaponClass())
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

	if newMode.String() != v.composite.GetAnimationMode() {
		v.composite.SetMode(newMode.String(), v.composite.GetWeaponClass())
	}

	if v.composite.GetDirection() != direction {
		v.composite.SetDirection(direction)
	}
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
