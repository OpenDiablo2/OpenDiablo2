package d2mapentity

import (
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2path"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// NPC is a passive complex entity with which the player can interact.
// For example, Deckard Cain.
type NPC struct {
	mapEntity
	Paths         []d2path.Path
	name          string
	composite     *d2asset.Composite
	action        int
	path          int
	repetitions   int
	monstatRecord *d2records.MonStatsRecord
	monstatEx     *d2records.MonStats2Record
	HasPaths      bool
	isDone        bool
}

const (
	magicOffsetX            = 5
	magicOffsetScalarX      = 8
	magicOffsetScalarY      = 16
	minAnimationRepetitions = 3
	maxAnimationRepetitions = 5
)

func selectEquip(slice []string) string {
	if len(slice) != 0 {
		// nolint:gosec // not concerned with crypto-strong randomness
		return slice[rand.Intn(len(slice))]
	}

	return ""
}

// ID returns the NPC uuid
func (v *NPC) ID() string {
	return v.mapEntity.uuid
}

// Render renders this entity's animated composite.
func (v *NPC) Render(target d2interface.Surface) {
	renderOffset := v.Position.RenderOffset()
	target.PushTranslation(
		int((renderOffset.X()-renderOffset.Y())*magicOffsetScalarY),
		int(((renderOffset.X()+renderOffset.Y())*magicOffsetScalarX)-magicOffsetX),
	)

	defer target.Pop()

	if v.composite.Render(target) != nil {
		return
	}
}

// Path returns the current part of the entity's path.
func (v *NPC) Path() d2path.Path {
	return v.Paths[v.path]
}

// NextPath returns the next part of the entity's path.
func (v *NPC) NextPath() d2path.Path {
	v.path++
	if v.path == len(v.Paths) {
		v.path = 0
	}

	return v.Paths[v.path]
}

// SetPaths sets the entity's paths to the given slice. It also sets flags
// on the entity indicating that it has paths and has completed the
// previous none.
func (v *NPC) SetPaths(paths []d2path.Path) {
	v.Paths = paths
	v.HasPaths = len(paths) > 0
	v.isDone = true
}

// Advance is called once per frame and processes a
// single game tick.
func (v *NPC) Advance(tickTime float64) {
	v.Step(tickTime)

	if err := v.composite.Advance(tickTime); err != nil {
		return
	}

	if v.HasPaths && v.wait() {
		// If at the target, set target to the next path.
		v.isDone = false
		path := v.NextPath()
		v.setTarget(
			path.Position,
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
	var newAnimationMode d2enum.MonsterAnimationMode

	v.isDone = true

	// nolint:gosec // not concerned with crypto-strong randomness
	v.repetitions = minAnimationRepetitions + rand.Intn(maxAnimationRepetitions)

	switch d2enum.NPCActionType(v.action) {
	case d2enum.NPCActionSkill1:
		newAnimationMode = d2enum.MonsterAnimationModeSkill1
		v.repetitions = 0
	case d2enum.NPCActionInvalid, d2enum.NPCAction1, d2enum.NPCAction2, d2enum.NPCAction3:
		newAnimationMode = d2enum.MonsterAnimationModeNeutral
		v.repetitions = 0
	default:
		newAnimationMode = d2enum.MonsterAnimationModeNeutral
		v.repetitions = 0
	}

	if v.composite.GetAnimationMode() != newAnimationMode.String() {
		if err := v.composite.SetMode(newAnimationMode, v.composite.GetWeaponClass()); err != nil {
			return
		}
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
		if err := v.composite.SetMode(newMode, v.composite.GetWeaponClass()); err != nil {
			return
		}
	}

	if v.composite.GetDirection() != direction {
		v.composite.SetDirection(direction)
	}
}

// Selectable returns true if the object can be highlighted/selected.
func (v *NPC) Selectable() bool {
	// is there something handy that determines selectable npc's?
	return v.name != ""
}

// Label returns the NPC's in-game name (e.g. "Deckard Cain") or an empty string if it does not have a name.
func (v *NPC) Label() string {
	return v.name
}

// GetPosition returns the NPC's position
func (v *NPC) GetPosition() d2vector.Position {
	return v.mapEntity.Position
}

// GetVelocity returns the NPC's velocity vector
func (v *NPC) GetVelocity() d2vector.Vector {
	return v.mapEntity.velocity
}

// GetSize returns the current frame size
func (v *NPC) GetSize() (width, height int) {
	return v.composite.GetSize()
}
