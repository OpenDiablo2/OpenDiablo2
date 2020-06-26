package d2mapentity

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type Player struct {
	mapEntity
	composite     *d2asset.Composite
	Equipment     d2inventory.CharacterEquipment
	Stats         d2hero.HeroStatsState
	Class         d2enum.Hero
	Id            string
	direction     int
	name          string
	nameLabel     d2ui.Label
	lastPathSize  int
	isInTown      bool
	animationMode string
	isRunToggled  bool
	isRunning     bool
}

// run speed should be walkspeed * 1.5, since in the original game it is 6 yards walk and 9 yards run.
var baseWalkSpeed = 6.0
var baseRunSpeed = 9.0

func CreatePlayer(id, name string, x, y int, direction int, heroType d2enum.Hero, stats d2hero.HeroStatsState, equipment d2inventory.CharacterEquipment) *Player {
	object := &d2datadict.ObjectLookupRecord{
		Mode:  d2enum.AnimationModePlayerTownNeutral.String(),
		Base:  "/data/global/chars",
		Token: heroType.GetToken(),
		Class: equipment.RightHand.GetWeaponClass(),
		SH:    equipment.Shield.GetItemCode(),
		// TODO: Offhand class?
		HD: equipment.Head.GetArmorClass(),
		TR: equipment.Torso.GetArmorClass(),
		LG: equipment.Legs.GetArmorClass(),
		RA: equipment.RightArm.GetArmorClass(),
		LA: equipment.LeftArm.GetArmorClass(),
		RH: equipment.RightHand.GetItemCode(),
		LH: equipment.LeftHand.GetItemCode(),
	}

	composite, err := d2asset.LoadComposite(object, d2resource.PaletteUnits)
	if err != nil {
		panic(err)
	}

	stats.NextLevelExp = d2datadict.GetExperienceBreakpoint(heroType, stats.Level)
	stats.Stamina = stats.MaxStamina

	result := &Player{
		Id:           id,
		mapEntity:    createMapEntity(x, y),
		composite:    composite,
		Equipment:    equipment,
		Stats:        stats,
		direction:    direction,
		name:         name,
		Class:        heroType,
		nameLabel:    d2ui.CreateLabel(d2resource.FontFormal11, d2resource.PaletteStatic),
		isRunToggled: true,
		isInTown:     true,
		isRunning:    true,
	}
	result.SetSpeed(baseRunSpeed)
	result.mapEntity.directioner = result.rotate
	result.nameLabel.Alignment = d2ui.LabelAlignCenter
	result.nameLabel.SetText(name)
	result.nameLabel.Color = color.White
	err = result.SetMode(d2enum.AnimationModePlayerTownNeutral.String(), equipment.RightHand.GetWeaponClass(), direction)
	if err != nil {
		panic(err)
	}
	return result
}

func (p *Player) SetIsInTown(isInTown bool) {
	p.isInTown = isInTown
}

func (p *Player) ToggleRunWalk() {
	p.isRunToggled = !p.isRunToggled
}

func (p *Player) IsRunToggled() bool {
	return p.isRunToggled
}

func (p *Player) IsRunning() bool {
	return p.isRunning
}

func (p *Player) SetIsRunning(isRunning bool) {
	p.isRunning = isRunning

	if isRunning {
		p.SetSpeed(baseRunSpeed)
	} else {
		p.SetSpeed(baseWalkSpeed)
	}
}

func (p Player) IsInTown() bool {
	return p.isInTown
}

func (v *Player) Advance(tickTime float64) {
	v.Step(tickTime)
	v.composite.Advance(tickTime)
	if v.lastPathSize != len(v.path) {
		v.lastPathSize = len(v.path)
	}

	if v.composite.GetAnimationMode() != v.animationMode {
		v.animationMode = v.composite.GetAnimationMode()
	}
}

func (v *Player) Render(target d2render.Surface) {
	target.PushTranslation(
		v.offsetX+int((v.subcellX-v.subcellY)*16),
		v.offsetY+int(((v.subcellX+v.subcellY)*8)-5),
	)
	defer target.Pop()
	v.composite.Render(target)
	//v.nameLabel.X = v.offsetX
	//v.nameLabel.Y = v.offsetY - 100
	//v.nameLabel.Render(target)
}

func (v *Player) SetMode(animationMode, weaponClass string, direction int) error {
	v.composite.SetMode(animationMode, weaponClass, direction)
	v.direction = direction
	v.weaponClass = weaponClass

	err := v.composite.SetMode(animationMode, weaponClass, direction)
	if err != nil {
		err = v.composite.SetMode(animationMode, "HTH", direction)
		v.weaponClass = "HTH"
	}

	return err
}

func (v *Player) GetAnimationMode() d2enum.PlayerAnimationMode {
	if v.IsRunning() && !v.IsAtTarget() {
		return d2enum.AnimationModePlayerRun
	}

	if v.IsInTown() {
		if !v.IsAtTarget() {
			return d2enum.AnimationModePlayerTownWalk
		}

		return d2enum.AnimationModePlayerTownNeutral
	}

	if !v.IsAtTarget() {
		return d2enum.AnimationModePlayerWalk
	}

	return d2enum.AnimationModePlayerNeutral
}

func (v *Player) SetAnimationMode(animationMode string) error {
	return v.composite.SetMode(animationMode, v.weaponClass, v.direction)
}

// rotate sets direction and changes animation
func (v *Player) rotate(direction int) {
	newAnimationMode := v.GetAnimationMode().String()

	if newAnimationMode != v.composite.GetAnimationMode() || direction != v.direction {
		v.SetMode(newAnimationMode, v.weaponClass, direction)
	}
}

func (v *Player) Name() string {
	return v.name
}
