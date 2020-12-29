package d2mapentity

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
)

// Player is the player character entity.
type Player struct {
	mapEntity
	name              string
	animationMode     string
	composite         *d2asset.Composite
	Equipment         *d2inventory.CharacterEquipment
	Stats             *d2hero.HeroStatsState
	Skills            map[int]*d2hero.HeroSkill
	LeftSkill         *d2hero.HeroSkill
	RightSkill        *d2hero.HeroSkill
	Class             d2enum.Hero
	Gold              int
	lastPathSize      int
	isInTown          bool
	isRunToggled      bool
	isRunning         bool
	isCasting         bool
	onFinishedCasting func()
	Act               int
}

// run speed should be walkspeed * 1.5, since in the original game it is 6 yards walk and 9 yards run.
const (
	baseWalkSpeed = 9.0
	baseRunSpeed  = 13.0
)

// ID returns the Player uuid
func (p *Player) ID() string {
	return p.mapEntity.uuid
}

// SetIsInTown sets a flag indicating that the player is in town.
func (p *Player) SetIsInTown(isInTown bool) {
	p.isInTown = isInTown
}

// ToggleRunWalk sets a flag indicating whether the player is running.
func (p *Player) ToggleRunWalk() {
	p.isRunToggled = !p.isRunToggled
}

// IsRunToggled returns true if the UI button to toggle running is,
// toggled i.e. not in it's default state.
func (p *Player) IsRunToggled() bool {
	return p.isRunToggled
}

// IsRunning returns true if the player is currently
func (p *Player) IsRunning() bool {
	return p.isRunning
}

// SetIsRunning alters the player speed and sets a flag indicating
// that the player is running.
func (p *Player) SetIsRunning(isRunning bool) {
	p.isRunning = isRunning

	if isRunning {
		p.SetSpeed(baseRunSpeed)
	} else {
		p.SetSpeed(baseWalkSpeed)
	}
}

// IsInTown returns true if the player is currently in town.
func (p *Player) IsInTown() bool {
	return p.isInTown
}

const (
	half = 0.5
)

// Advance is called once per frame and processes a
// single game tick.
func (p *Player) Advance(tickTime float64) {
	p.Step(tickTime)

	if err := p.SetAnimationMode(p.GetAnimationMode()); err != nil {
		fmt.Printf("failed to set animationMode to: %d, err: %v\n", p.GetAnimationMode(), err)
	}

	if p.IsCasting() {
		if p.composite.GetPlayedCount() >= 1 {
			p.isCasting = false
		}

		// skills are casted after the first half of the casting animation is played
		percentDone := float64(p.composite.GetCurrentFrame()) / float64(p.composite.GetFrameCount())
		isHalfDoneCasting := percentDone >= half

		if isHalfDoneCasting && p.onFinishedCasting != nil {
			p.onFinishedCasting()
			p.onFinishedCasting = nil
		}
	}

	if err := p.composite.Advance(tickTime); err != nil {
		fmt.Printf("failed to advance composite animation of player: %s, err: %v\n", p.ID(), err)
	}

	if p.lastPathSize != len(p.path) {
		p.lastPathSize = len(p.path)
	}

	if p.composite.GetAnimationMode() != p.animationMode {
		p.animationMode = p.composite.GetAnimationMode()
	}

	charstats := p.composite.AssetManager.Records.Character.Stats[p.Class]
	staminaDrain := float64(charstats.StaminaRunDrain)

	// This number has been determined by trying it out and checking if the stamina drain is
	// the same as in d2 with the drain value from the assets.
	// (We stopped the time for a lvl 1 babarian to loose all stamina which is around 25 seconds
	// if i Remember correctly)
	const magicStaminaDrainDivisor = 5

	// Drain and regenerate Stamina
	if p.IsRunning() && !p.atTarget() && !p.IsInTown() {
		p.Stats.Stamina -= staminaDrain * tickTime / magicStaminaDrainDivisor
		if p.Stats.Stamina <= 0 {
			p.SetSpeed(baseWalkSpeed)
			p.Stats.Stamina = 0
		}
	} else if p.Stats.Stamina < float64(p.Stats.MaxStamina) {
		p.Stats.Stamina += staminaDrain * tickTime / magicStaminaDrainDivisor
		if p.IsRunning() {
			p.SetSpeed(baseRunSpeed)
		}
	}
}

// Render renders the animated composite for this entity.
func (p *Player) Render(target d2interface.Surface) {
	renderOffset := p.Position.RenderOffset()
	target.PushTranslation(
		int((renderOffset.X()-renderOffset.Y())*subtileWidth),
		int(((renderOffset.X()+renderOffset.Y())*subtileHeight)+subtileOffsetY),
	)

	defer target.Pop()

	if err := p.composite.Render(target); err != nil {
		fmt.Printf("failed to render the composite of player: %s, err: %v\n", p.ID(), err)
	}
}

// GetAnimationMode returns the current animation mode based on what the player is doing and where they are.
func (p *Player) GetAnimationMode() d2enum.PlayerAnimationMode {
	if p.IsRunning() && !p.atTarget() {
		return d2enum.PlayerAnimationModeRun
	}

	if p.IsInTown() {
		if !p.atTarget() {
			return d2enum.PlayerAnimationModeTownWalk
		}

		return d2enum.PlayerAnimationModeTownNeutral
	}

	if !p.atTarget() {
		return d2enum.PlayerAnimationModeWalk
	}

	if p.IsCasting() {
		return d2enum.PlayerAnimationModeCast
	}

	return d2enum.PlayerAnimationModeNeutral
}

// SetAnimationMode sets the Composite's animation mode weapon class and direction.
func (p *Player) SetAnimationMode(animationMode d2enum.PlayerAnimationMode) error {
	return p.composite.SetMode(animationMode, p.composite.GetWeaponClass())
}

// rotate sets direction and changes animation
func (p *Player) rotate(direction int) {
	newAnimationMode := p.GetAnimationMode()

	if newAnimationMode.String() != p.composite.GetAnimationMode() {
		if err := p.composite.SetMode(newAnimationMode, p.composite.GetWeaponClass()); err != nil {
			fmt.Printf("failed to update animationMode of %s, err: %v\n", p.composite.GetWeaponClass(), err)
		}
	}

	if direction != p.composite.GetDirection() {
		p.composite.SetDirection(direction)
	}
}

// SetDirection will rotate the player and change the animation
func (p *Player) SetDirection(direction int) {
	p.rotate(direction)
}

// Name returns the player name.
func (p *Player) Name() string {
	return p.name
}

// IsCasting returns true if
func (p *Player) IsCasting() bool {
	return p.isCasting
}

// StartCasting sets a flag indicating the player is casting a skill and
// sets the animation mode to the casting animation.
// This handles all types of skills - melee, ranged, kick, summon, etc.
// NB: onFinishedCasting is called when the casting animation is >50% complete
func (p *Player) StartCasting(animMode d2enum.PlayerAnimationMode, onFinishedCasting func()) {
	// passive skills, auras, etc.
	if animMode == d2enum.PlayerAnimationModeNone {
		return
	}

	p.isCasting = true
	p.onFinishedCasting = onFinishedCasting

	if err := p.SetAnimationMode(animMode); err != nil {
		fmtStr := "failed to set animationMode of player: %s to: %d, err: %v\n"
		fmt.Printf(fmtStr, p.ID(), animMode, err)
	}
}

// Selectable returns true if the player is in town.
func (p *Player) Selectable() bool {
	// Players are selectable when in town
	return p.IsInTown()
}

// GetPosition returns the entity's position
func (p *Player) GetPosition() d2vector.Position {
	return p.mapEntity.Position
}

// GetVelocity returns the entity's velocity vector
func (p *Player) GetVelocity() d2vector.Vector {
	return p.mapEntity.velocity
}

// GetSize returns the current frame size
func (p *Player) GetSize() (width, height int) {
	width, height = p.composite.GetSize()
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/820
	height = (height * 2) - (height / 2)

	return width, height
}
