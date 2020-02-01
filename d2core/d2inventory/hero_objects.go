package d2inventory

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

var HeroObjects map[d2enum.Hero]CharacterEquipment

func LoadHeroObjects() {
	//Mode:  d2enum.AnimationModePlayerNeutral.String(),
	//Base:  "/data/global/chars",
	HeroObjects = map[d2enum.Hero]CharacterEquipment{
		d2enum.HeroBarbarian: {
			RightHand: GetWeaponItemByCode("hax"),
			Shield:    GetArmorItemByCode("buc"),
		},
		d2enum.HeroNecromancer: {
			RightHand: GetWeaponItemByCode("wnd"),
		},
		d2enum.HeroPaladin: {
			RightHand: GetWeaponItemByCode("ssd"),
			Shield:    GetArmorItemByCode("buc"),
		},
		d2enum.HeroAssassin: {
			RightHand: GetWeaponItemByCode("ktr"),
			Shield:    GetArmorItemByCode("buc"),
		},
		d2enum.HeroSorceress: {
			RightHand: GetWeaponItemByCode("sst"),
			LeftHand:  GetWeaponItemByCode("sst"),
		},
		d2enum.HeroAmazon: {
			RightHand: GetWeaponItemByCode("jav"),
			Shield:    GetArmorItemByCode("buc"),
		},
		d2enum.HeroDruid: {
			RightHand: GetWeaponItemByCode("clb"),
			Shield:    GetArmorItemByCode("buc"),
		},
	}
}
