package d2map

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type Player struct {
	*AnimatedComposite
	Equipment d2inventory.CharacterEquipment
	Id        string
	mode      d2enum.AnimationMode
	direction int
}

func CreatePlayer(id string, x, y int, direction int, heroType d2enum.Hero, equipment d2inventory.CharacterEquipment) *Player {
	object := &d2datadict.ObjectLookupRecord{
		Mode:  d2enum.AnimationModePlayerNeutral.String(),
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
		mode:              d2enum.AnimationModePlayerTownNeutral,
		direction:         direction,
	}
	result.SetMode(result.mode.String(), equipment.RightHand.GetWeaponClass(), direction)
	return result
}

func (v *Player) Advance(tickTime float64) {
	v.Step(tickTime)
	v.AnimatedComposite.Advance(tickTime)
}

func (v *Player) Render(target d2render.Surface) {
	v.AnimatedComposite.Render(target)
}

func (v *Player) GetPosition() (float64, float64) {
	return v.AnimatedComposite.GetPosition()
}
