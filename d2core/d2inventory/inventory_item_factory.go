package d2inventory

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// NewInventoryItemFactory creates a new InventoryItemFactory and initializes it
func NewInventoryItemFactory(asset *d2asset.AssetManager) (*InventoryItemFactory, error) {
	factory := &InventoryItemFactory{asset: asset}

	err := factory.loadHeroObjects()
	if err != nil {
		return nil, err
	}

	return factory, nil
}

// InventoryItemFactory is responsible for creating inventory items
type InventoryItemFactory struct {
	asset            *d2asset.AssetManager
	DefaultHeroItems HeroObjects
}

// LoadHeroObjects loads the equipment objects of the hero
func (f *InventoryItemFactory) loadHeroObjects() error {
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/795
	//Mode:  d2enum.AnimationModePlayerNeutral.String(),
	//Base:  "/data/global/chars",
	shield, err := f.GetArmorItemByCode("buc")
	if err != nil {
		return err
	}

	rhhex, err := f.GetWeaponItemByCode("hax")
	if err != nil {
		return err
	}

	rhwnd, err := f.GetWeaponItemByCode("wnd")
	if err != nil {
		return err
	}

	rhssd, err := f.GetWeaponItemByCode("ssd")
	if err != nil {
		return err
	}

	rhktr, err := f.GetWeaponItemByCode("ktr")
	if err != nil {
		return err
	}

	rhsst, err := f.GetWeaponItemByCode("sst")
	if err != nil {
		return err
	}

	rhjav, err := f.GetWeaponItemByCode("jav")
	if err != nil {
		return err
	}

	rhclb, err := f.GetWeaponItemByCode("clb")
	if err != nil {
		return err
	}

	f.DefaultHeroItems = map[d2enum.Hero]CharacterEquipment{
		d2enum.HeroBarbarian: {
			RightHand: rhhex,
			Shield:    shield,
		},
		d2enum.HeroNecromancer: {
			RightHand: rhwnd,
		},
		d2enum.HeroPaladin: {
			RightHand: rhssd,
			Shield:    shield,
		},
		d2enum.HeroAssassin: {
			RightHand: rhktr,
			Shield:    shield,
		},
		d2enum.HeroSorceress: {
			RightHand: rhsst,
			LeftHand:  rhsst,
		},
		d2enum.HeroAmazon: {
			RightHand: rhjav,
			Shield:    shield,
		},
		d2enum.HeroDruid: {
			RightHand: rhclb,
			Shield:    shield,
		},
	}

	return nil
}

// GetArmorItemByCode returns the armor item for the given code
func (f *InventoryItemFactory) GetArmorItemByCode(code string) (*InventoryItemArmor, error) {
	result := f.asset.Records.Item.Armors[code]
	if result == nil {
		return nil, fmt.Errorf("could not find armor entry for code '%s'", code)
	}

	return &InventoryItemArmor{
		InventorySizeX: result.InventoryWidth,
		InventorySizeY: result.InventoryHeight,
		ItemName:       result.Name,
		ItemCode:       result.Code,
		ArmorClass:     d2enum.ArmorClassLite, // comes from ArmType.txt
	}, nil
}

// GetMiscItemByCode returns the miscellaneous item for the given code
func (f *InventoryItemFactory) GetMiscItemByCode(code string) (*InventoryItemMisc, error) {
	result := f.asset.Records.Item.Misc[code]
	if result == nil {
		return nil, fmt.Errorf("could not find misc item entry for code '%s'", code)
	}

	return &InventoryItemMisc{
		InventorySizeX: result.InventoryWidth,
		InventorySizeY: result.InventoryHeight,
		ItemName:       result.Name,
		ItemCode:       result.Code,
	}, nil
}

// GetWeaponItemByCode returns the weapon item for the given code
func (f *InventoryItemFactory) GetWeaponItemByCode(code string) (*InventoryItemWeapon, error) {
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/796
	result := f.asset.Records.Item.Weapons[code]
	if result == nil {
		return nil, fmt.Errorf("could not find weapon entry for code '%s'", code)
	}

	return &InventoryItemWeapon{
		InventorySizeX:     result.InventoryWidth,
		InventorySizeY:     result.InventoryHeight,
		ItemName:           result.Name,
		ItemCode:           result.Code,
		WeaponClass:        result.WeaponClass,
		WeaponClassOffHand: result.WeaponClass2Hand,
	}, nil
}
