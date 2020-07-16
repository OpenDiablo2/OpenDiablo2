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

// NPC is a passive complex entity with which the player can interact.
// For example, Deckard Cain.
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

// CreateNPC creates a new NPC and returns a pointer to it.
func CreateNPC(x, y int, monstat *d2datadict.MonStatsRecord, direction int) *NPC {
	result := &NPC{
		mapEntity:     newMapEntity(x, y),
		HasPaths:      false,
		monstatRecord: monstat,
		monstatEx:     d2datadict.MonStats2[monstat.ExtraDataKey],
	}

	var equipment [16]string

	for compType, opts := range result.monstatEx.EquipmentOptions {
		equipment[compType] = selectEquip(opts)
	}

	composite, _ := d2asset.LoadComposite(d2enum.ObjectTypeCharacter, monstat.AnimationDirectoryToken,
		d2resource.PaletteUnits)
	result.composite = composite

	composite.SetMode(d2enum.MonsterAnimationModeNeutral, result.monstatEx.BaseWeaponClass)
	composite.Equip(&equipment)

	result.SetSpeed(float64(monstat.SpeedBase))
	result.mapEntity.directioner = result.rotate

	result.composite.SetDirection(direction)

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

// Render renders this entity's animated composite.
func (v *NPC) Render(target d2interface.Surface) {
	renderOffset := v.Position.RenderOffset()
	target.PushTranslation(
		int((renderOffset.X()-renderOffset.Y())*16),
		int(((renderOffset.X()+renderOffset.Y())*8)-5),
	)

	defer target.Pop()
	v.composite.Render(target)
}

// Path returns the current part of the entity's path.
func (v *NPC) Path() d2common.Path {
	return v.Paths[v.path]
}

// NextPath returns the next part of the entity's path.
func (v *NPC) NextPath() d2common.Path {
	v.path++
	if v.path == len(v.Paths) {
		v.path = 0
	}

	return v.Paths[v.path]
}

// SetPaths sets the entity's paths to the given slice. It also sets flags
// on the entity indicating that it has paths and has completed the
// previous none.
func (v *NPC) SetPaths(paths []d2common.Path) {
	v.Paths = paths
	v.HasPaths = len(paths) > 0
	v.isDone = true
}

// Advance is called once per frame and processes a
// single game tick.
func (v *NPC) Advance(tickTime float64) {
	v.Step(tickTime)
	v.composite.Advance(tickTime)

	if v.HasPaths && v.wait() {
		// If at the target, set target to the next path.
		v.isDone = false
		path := v.NextPath()
		v.setTarget(
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
	newAnimationMode := d2enum.MonsterAnimationModeNeutral
	// TODO: Figure out what 1-3 are for, 4 is correct.
	switch v.action {
	case 1:
		newAnimationMode = d2enum.MonsterAnimationModeNeutral
	case 2:
		newAnimationMode = d2enum.MonsterAnimationModeNeutral
	case 3:
		newAnimationMode = d2enum.MonsterAnimationModeNeutral
	case 4:
		newAnimationMode = d2enum.MonsterAnimationModeSkill1
		v.repetitions = 0
	default:
		v.repetitions = 0
	}

	if v.composite.GetAnimationMode() != newAnimationMode.String() {
		v.composite.SetMode(newAnimationMode, v.composite.GetWeaponClass())
	}
}

// rotate sets direction and changes animation
func (v *NPC) rotate(direction int) {
	var newMode d2enum.MonsterAnimationMode
	if !v.atTarget() {
		newMode = d2enum.MonsterAnimationModeWalk
	} else {
		newMode = d2enum.MonsterAnimationModeNeutral
	}

	if newMode.String() != v.composite.GetAnimationMode() {
		v.composite.SetMode(newMode, v.composite.GetWeaponClass())
	}

	if v.composite.GetDirection() != direction {
		v.composite.SetDirection(direction)
	}
}

// Selectable returns true if the object can be highlighted/selected.
func (m *NPC) Selectable() bool {
	// is there something handy that determines selectable npc's?
	if m.name != "" {
		return true
	}

	return false
}

// Name returns the NPC's in-game name (e.g. "Deckard Cain") or an empty string if it does not have a name.
func (m *NPC) Name() string {
	return m.name
}
