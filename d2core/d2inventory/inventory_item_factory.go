package d2inventory

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// NewInventoryItemFactory creates a new InventoryItemFactory and initializes it
func NewInventoryItemFactory(asset *d2asset.AssetManager) (*InventoryItemFactory, error) {
	factory := &InventoryItemFactory{asset: asset}

	factory.loadHeroObjects()

	return factory, nil
}

// InventoryItemFactory is responsible for creating inventory items
type InventoryItemFactory struct {
	asset            *d2asset.AssetManager
	DefaultHeroItems HeroObjects
}

// LoadHeroObjects loads the equipment objects of the hero
func (f *InventoryItemFactory) loadHeroObjects() {
	//Mode:  d2enum.AnimationModePlayerNeutral.String(),
	//Base:  "/data/global/chars",
	f.DefaultHeroItems = map[d2enum.Hero]CharacterEquipment{
		d2enum.HeroBarbarian: {
			RightHand: f.GetWeaponItemByCode("hax"),
			Shield:    f.GetArmorItemByCode("buc"),
		},
		d2enum.HeroNecromancer: {
			RightHand: f.GetWeaponItemByCode("wnd"),
		},
		d2enum.HeroPaladin: {
			RightHand: f.GetWeaponItemByCode("ssd"),
			Shield:    f.GetArmorItemByCode("buc"),
		},
		d2enum.HeroAssassin: {
			RightHand: f.GetWeaponItemByCode("ktr"),
			Shield:    f.GetArmorItemByCode("buc"),
		},
		d2enum.HeroSorceress: {
			RightHand: f.GetWeaponItemByCode("sst"),
			LeftHand:  f.GetWeaponItemByCode("sst"),
		},
		d2enum.HeroAmazon: {
			RightHand: f.GetWeaponItemByCode("jav"),
			Shield:    f.GetArmorItemByCode("buc"),
		},
		d2enum.HeroDruid: {
			RightHand: f.GetWeaponItemByCode("clb"),
			Shield:    f.GetArmorItemByCode("buc"),
		},
	}
}

// GetArmorItemByCode returns the armor item for the given code
func (f *InventoryItemFactory) GetArmorItemByCode(code string) *InventoryItemArmor {
	result := f.asset.Records.Item.Armors[code]
	if result == nil {
		log.Fatalf("Could not find armor entry for code '%s'", code)
	}

	return &InventoryItemArmor{
		InventorySizeX: result.InventoryWidth,
		InventorySizeY: result.InventoryHeight,
		ItemName:       result.Name,
		ItemCode:       result.Code,
		ArmorClass:     "lit", // TODO: Where does this come from?
	}
}

// GetMiscItemByCode returns the miscellaneous item for the given code
func (f *InventoryItemFactory) GetMiscItemByCode(code string) *InventoryItemMisc {
	result := f.asset.Records.Item.Misc[code]
	if result == nil {
		log.Fatalf("Could not find misc item entry for code '%s'", code)
	}

	return &InventoryItemMisc{
		InventorySizeX: result.InventoryWidth,
		InventorySizeY: result.InventoryHeight,
		ItemName:       result.Name,
		ItemCode:       result.Code,
	}
}

// GetWeaponItemByCode returns the weapon item for the given code
func (f *InventoryItemFactory) GetWeaponItemByCode(code string) *InventoryItemWeapon {
	// TODO: Non-normal codes will fail here...
	result := f.asset.Records.Item.Weapons[code]
	if result == nil {
		log.Fatalf("Could not find weapon entry for code '%s'", code)
	}

	return &InventoryItemWeapon{
		InventorySizeX:     result.InventoryWidth,
		InventorySizeY:     result.InventoryHeight,
		ItemName:           result.Name,
		ItemCode:           result.Code,
		WeaponClass:        result.WeaponClass,
		WeaponClassOffHand: result.WeaponClass2Hand,
	}
}
