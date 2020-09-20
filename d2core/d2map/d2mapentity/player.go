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
	lastPathSize      int
	isInTown          bool
	isRunToggled      bool
	isRunning         bool
	isCasting         bool
	onFinishedCasting func()
}

// run speed should be walkspeed * 1.5, since in the original game it is 6 yards walk and 9 yards run.
const baseWalkSpeed = 6.0
const baseRunSpeed = 9.0

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

// Advance is called once per frame and processes a
// single game tick.
func (p *Player) Advance(tickTime float64) {
	p.Step(tickTime)

	if p.IsCasting() && p.composite.GetPlayedCount() >= 1 {
		p.isCasting = false
		if p.onFinishedCasting != nil {
			p.onFinishedCasting()
			p.onFinishedCasting = nil
		}

		if err := p.SetAnimationMode(p.GetAnimationMode()); err != nil {
			fmt.Printf("failed to set animationMode to: %d, err: %v\n", p.GetAnimationMode(), err)
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
}

// Render renders the animated composite for this entity.
func (p *Player) Render(target d2interface.Surface) {
	renderOffset := p.Position.RenderOffset()
	target.PushTranslation(
		int((renderOffset.X()-renderOffset.Y())*16),
		int(((renderOffset.X()+renderOffset.Y())*8)-5),
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
func (p *Player) StartCasting(onFinishedCasting func()) {
	p.isCasting = true
	p.onFinishedCasting = onFinishedCasting
	if err := p.SetAnimationMode(d2enum.PlayerAnimationModeCast); err != nil {
		fmtStr := "failed to set animationMode of player: %s to: %d, err: %v\n"
		fmt.Printf(fmtStr, p.ID(), d2enum.PlayerAnimationModeCast, err)
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
	// hack: we need to get full size of composite animations, currently only gets legs
	height = (height * 2) - (height / 2)

	return width, height
}
