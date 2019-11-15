package d2core

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2datadict"
)

type InventoryItemWeapon struct {
	inventorySizeX     int
	inventorySizeY     int
	itemName           string
	itemCode           string
	weaponClass        string
	weaponClassOffHand string
}

func GetWeaponItemByCode(code string) InventoryItemWeapon {
	// TODO: Non-normal codes will fail here...
	result := d2datadict.Weapons[code]
	if result == nil {
		log.Fatalf("Could not find weapon entry for code '%s'", code)
	}
	return InventoryItemWeapon{
		inventorySizeX:     result.InventoryWidth,
		inventorySizeY:     result.InventoryHeight,
		itemName:           result.Name,
		itemCode:           result.Code,
		weaponClass:        result.WeaponClass,
		weaponClassOffHand: result.WeaponClass2Hand,
	}
}

func (v InventoryItemWeapon) GetWeaponClass() string {
	if v.itemCode == "" {
		return "hth"
	}
	return v.weaponClass
}

func (v InventoryItemWeapon) GetWeaponClassOffHand() string {
	if v.itemCode == "" {
		return ""
	}
	return v.weaponClassOffHand
}

func (v InventoryItemWeapon) GetInventoryItemName() string {
	return v.itemName
}

func (v InventoryItemWeapon) GetInventoryItemType() d2enum.InventoryItemType {
	return d2enum.InventoryItemTypeWeapon
}

func (v InventoryItemWeapon) GetInventoryGridSize() (int, int) {
	return v.inventorySizeX, v.inventorySizeY
}

func (v InventoryItemWeapon) Serialize() []byte {
	return []byte{}
}

func (v InventoryItemWeapon) GetItemCode() string {
	return v.itemCode
}
