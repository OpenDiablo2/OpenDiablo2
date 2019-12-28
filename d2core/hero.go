package d2core

import (
	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"
	"github.com/OpenDiablo2/D2Shared/d2common/d2resource"
	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2surface"
)

type Hero struct {
	AnimatedEntity *d2render.AnimatedEntity
	Equipment      CharacterEquipment
	mode           d2enum.AnimationMode
	direction      int
}

func CreateHero(x, y int32, direction int, heroType d2enum.Hero, equipment CharacterEquipment) *Hero {
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

	entity, err := d2render.CreateAnimatedEntity(x, y, object, d2resource.PaletteUnits)
	if err != nil {
		panic(err)
	}

	result := &Hero{AnimatedEntity: entity, Equipment: equipment, mode: d2enum.AnimationModePlayerTownNeutral, direction: direction}
	result.AnimatedEntity.SetMode(result.mode.String(), equipment.RightHand.GetWeaponClass(), direction)
	return result
}

func (v *Hero) Advance(tickTime float64) {
	// TODO: Pathfinding
	if v.AnimatedEntity.LocationX != v.AnimatedEntity.TargetX ||
		v.AnimatedEntity.LocationY != v.AnimatedEntity.TargetY {
		v.AnimatedEntity.Step(tickTime)
	}

	v.AnimatedEntity.Advance(tickTime)
}

func (v *Hero) Render(target *d2surface.Surface) {
	v.AnimatedEntity.Render(target)
}

func (v *Hero) GetPosition() (float64, float64) {
	return v.AnimatedEntity.GetPosition()
}
