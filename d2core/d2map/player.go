package d2map

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type Player struct {
	*AnimatedComposite
	Equipment     d2inventory.CharacterEquipment
	Id            string
	direction     int
	Name          string
	nameLabel     d2ui.Label
	lastPathSize  int
	isInTown      bool
	animationMode string
}

func CreatePlayer(id, name string, x, y int, direction int, heroType d2enum.Hero, equipment d2inventory.CharacterEquipment) *Player {
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

	entity, err := CreateAnimatedComposite(x, y, object, d2resource.PaletteUnits)
	if err != nil {
		panic(err)
	}

	result := &Player{
		Id:                id,
		AnimatedComposite: entity,
		Equipment:         equipment,
		direction:         direction,
		Name:              name,
		nameLabel:         d2ui.CreateLabel(d2resource.FontFormal11, d2resource.PaletteStatic),
	}
	result.nameLabel.Alignment = d2ui.LabelAlignCenter
	result.nameLabel.SetText(name)
	result.nameLabel.Color = color.White
	result.SetPlayer(result)
	err = result.SetMode(d2enum.AnimationModePlayerTownNeutral.String(), equipment.RightHand.GetWeaponClass(), direction)
	if err != nil {
		panic(err)
	}
	return result
}

func (p *Player) SetIsInTown(isInTown bool) {
	p.isInTown = isInTown
}

func (p Player) IsInTown() bool {
	return p.isInTown
}

func (v *Player) Advance(tickTime float64) {
	v.Step(tickTime)
	v.AnimatedComposite.Advance(tickTime)
	if v.lastPathSize != len(v.path) {
		v.lastPathSize = len(v.path)
	}

	if v.AnimatedComposite.composite.GetAnimationMode() != v.animationMode {
		v.animationMode = v.AnimatedComposite.composite.GetAnimationMode()
	}
}

func (v *Player) Render(target d2render.Surface) {
	v.AnimatedComposite.Render(target)
	offX := v.AnimatedComposite.offsetX + int((v.AnimatedComposite.subcellX-v.AnimatedComposite.subcellY)*16)
	offY := v.AnimatedComposite.offsetY + int(((v.AnimatedComposite.subcellX+v.AnimatedComposite.subcellY)*8)-5)
	v.nameLabel.X = offX
	v.nameLabel.Y = offY - 100
	v.nameLabel.Render(target)
}

func (v *Player) GetPosition() (float64, float64) {
	return v.AnimatedComposite.GetPosition()
}
