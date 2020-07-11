package d2mapentity

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
)

// Player is the player character entity.
type Player struct {
	mapEntity
	composite *d2asset.Composite
	Equipment d2inventory.CharacterEquipment
	Stats     d2hero.HeroStatsState
	Class     d2enum.Hero
	Id        string
	name      string
	// nameLabel     d2ui.Label
	lastPathSize  int
	isInTown      bool
	animationMode string
	isRunToggled  bool
	isRunning     bool
	isCasting     bool
}

// run speed should be walkspeed * 1.5, since in the original game it is 6 yards walk and 9 yards run.
var baseWalkSpeed = 6.0
var baseRunSpeed = 9.0

// CreatePlayer creates a new player entity and returns a pointer to it.
func CreatePlayer(id, name string, x, y int, direction int, heroType d2enum.Hero, stats d2hero.HeroStatsState, equipment d2inventory.CharacterEquipment) *Player {
	layerEquipment := &[d2enum.CompositeTypeMax]string{
		d2enum.CompositeTypeHead:      equipment.Head.GetArmorClass(),
		d2enum.CompositeTypeTorso:     equipment.Torso.GetArmorClass(),
		d2enum.CompositeTypeLegs:      equipment.Legs.GetArmorClass(),
		d2enum.CompositeTypeRightArm:  equipment.RightArm.GetArmorClass(),
		d2enum.CompositeTypeLeftArm:   equipment.LeftArm.GetArmorClass(),
		d2enum.CompositeTypeRightHand: equipment.RightHand.GetItemCode(),
		d2enum.CompositeTypeLeftHand:  equipment.LeftHand.GetItemCode(),
		d2enum.CompositeTypeShield:    equipment.Shield.GetItemCode(),
	}

	composite, err := d2asset.LoadComposite(d2enum.ObjectTypePlayer, heroType.GetToken(),
		d2resource.PaletteUnits)
	if err != nil {
		panic(err)
	}

	stats.NextLevelExp = d2datadict.GetExperienceBreakpoint(heroType, stats.Level)
	stats.Stamina = stats.MaxStamina

	result := &Player{
		Id:        id,
		mapEntity: newMapEntity(x, y),
		composite: composite,
		Equipment: equipment,
		Stats:     stats,
		name:      name,
		Class:     heroType,
		//nameLabel:    d2ui.CreateLabel(d2resource.FontFormal11, d2resource.PaletteStatic),
		isRunToggled: true,
		isInTown:     true,
		isRunning:    true,
	}
	result.SetSpeed(baseRunSpeed)
	result.mapEntity.directioner = result.rotate
	//result.nameLabel.Alignment = d2ui.LabelAlignCenter
	//result.nameLabel.SetText(name)
	//result.nameLabel.Color = color.White
	err = composite.SetMode(d2enum.PlayerAnimationModeTownNeutral, equipment.RightHand.GetWeaponClass())
	if err != nil {
		panic(err)
	}

	composite.SetDirection(direction)
	composite.Equip(layerEquipment)

	return result
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
func (p Player) IsInTown() bool {
	return p.isInTown
}

// Advance is called once per frame and processes a
// single game tick.
func (v *Player) Advance(tickTime float64) {
	v.Step(tickTime)

	if v.IsCasting() && v.composite.GetPlayedCount() >= 1 {
		v.isCasting = false
		v.SetAnimationMode(v.GetAnimationMode())
	}
	v.composite.Advance(tickTime)

	if v.lastPathSize != len(v.path) {
		v.lastPathSize = len(v.path)
	}

	if v.composite.GetAnimationMode() != v.animationMode {
		v.animationMode = v.composite.GetAnimationMode()
	}
}

// Render renders the animated composite for this entity.
func (v *Player) Render(target d2interface.Surface) {
	renderOffset := v.Position.RenderOffset()
	target.PushTranslation(
		int((renderOffset.X()-renderOffset.Y())*16),
		int(((renderOffset.X()+renderOffset.Y())*8)-5),
	)

	defer target.Pop()
	v.composite.Render(target)
	// v.nameLabel.X = v.offsetX
	// v.nameLabel.Y = v.offsetY - 100
	// v.nameLabel.Render(target)
}

// GetAnimationMode returns the current animation mode based on what the player is doing and where they are.
func (v *Player) GetAnimationMode() d2enum.PlayerAnimationMode {
	if v.IsRunning() && !v.atTarget() {
		return d2enum.PlayerAnimationModeRun
	}

	if v.IsInTown() {
		if !v.atTarget() {
			return d2enum.PlayerAnimationModeTownWalk
		}

		return d2enum.PlayerAnimationModeTownNeutral
	}

	if !v.atTarget() {
		return d2enum.PlayerAnimationModeWalk
	}

	if v.IsCasting() {
		return d2enum.PlayerAnimationModeCast
	}

	return d2enum.PlayerAnimationModeNeutral
}

// SetAnimationMode sets the Composite's animation mode weapon class and direction.
func (v *Player) SetAnimationMode(animationMode d2enum.PlayerAnimationMode) error {
	return v.composite.SetMode(animationMode, v.composite.GetWeaponClass())
}

// rotate sets direction and changes animation
func (v *Player) rotate(direction int) {
	newAnimationMode := v.GetAnimationMode()

	if newAnimationMode.String() != v.composite.GetAnimationMode() {
		v.composite.SetMode(newAnimationMode, v.composite.GetWeaponClass())
	}

	if direction != v.composite.GetDirection() {
		v.composite.SetDirection(direction)
	}
}

// Name returns the player name.
func (v *Player) Name() string {
	return v.name
}

// IsCasting returns true if
func (v *Player) IsCasting() bool {
	return v.isCasting
}

// SetCasting sets a flag indicating the player is casting a skill and
// sets the animation mode to the casting animation.
func (v *Player) SetCasting() {
	v.isCasting = true
	v.SetAnimationMode(d2enum.PlayerAnimationModeCast)
}

// Selectable returns true if the player is in town.
func (v *Player) Selectable() bool {
	// Players are selectable when in town
	return v.IsInTown()
}
