package d2map

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type Hero struct {
	AnimatedEntity *AnimatedEntity
	Equipment      d2inventory.CharacterEquipment
	mode           d2enum.AnimationMode
	direction      int
}

func CreateHero(x, y int32, direction int, heroType d2enum.Hero, equipment d2inventory.CharacterEquipment) *Hero {
	object := &d2datadict.ObjectLookupRecord{
		Mode:  d2enum.AnimationModePlayerNeutral.String(),
		Base:  "/data/global/chars",
		Token: heroType.GetToken(),
		Class: equipment.RightHand.WeaponClass(),
		SH:    equipment.Shield.ItemCode(),
		// TODO: Offhand class?
		HD: equipment.Head.ArmorClass(),
		TR: equipment.Torso.ArmorClass(),
		LG: equipment.Legs.ArmorClass(),
		RA: equipment.RightArm.ArmorClass(),
		LA: equipment.LeftArm.ArmorClass(),
		RH: equipment.RightHand.ItemCode(),
		LH: equipment.LeftHand.ItemCode(),
	}

	entity, err := CreateAnimatedEntity(x, y, object, d2resource.PaletteUnits)
	if err != nil {
		panic(err)
	}

	result := &Hero{AnimatedEntity: entity, Equipment: equipment, mode: d2enum.AnimationModePlayerTownNeutral, direction: direction}
	result.AnimatedEntity.SetMode(result.mode.String(), equipment.RightHand.WeaponClass(), direction)
	return result
}

func (v *Hero) Advance(tickTime float64) {
	if v.AnimatedEntity.LocationX != v.AnimatedEntity.TargetX ||
		v.AnimatedEntity.LocationY != v.AnimatedEntity.TargetY ||
		v.AnimatedEntity.HasPathFinding(){
		v.AnimatedEntity.Step(tickTime)
	}

	v.AnimatedEntity.Advance(tickTime)
}

func (v *Hero) Render(target d2render.Surface) {
	v.AnimatedEntity.Render(target)
}

func (v *Hero) GetPosition() (float64, float64) {
	return v.AnimatedEntity.GetPosition()
}
