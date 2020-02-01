package d2inventory

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

type InventoryItemWeapon struct {
	inventorySizeX     int
	inventorySizeY     int
	inventorySlotX     int
	inventorySlotY     int
	itemName           string
	itemCode           string
	weaponClass        string
	weaponClassOffHand string
}

func GetWeaponItemByCode(code string) *InventoryItemWeapon {
	// TODO: Non-normal codes will fail here...
	result := d2datadict.Weapons[code]
	if result == nil {
		log.Fatalf("Could not find weapon entry for code '%s'", code)
	}
	return &InventoryItemWeapon{
		inventorySizeX:     result.InventoryWidth,
		inventorySizeY:     result.InventoryHeight,
		itemName:           result.Name,
		itemCode:           result.Code,
		weaponClass:        result.WeaponClass,
		weaponClassOffHand: result.WeaponClass2Hand,
	}
}

func (v *InventoryItemWeapon) WeaponClass() string {
	if v == nil || v.itemCode == "" {
		return "hth"
	}
	return v.weaponClass
}

func (v *InventoryItemWeapon) WeaponClassOffHand() string {
	if v == nil || v.itemCode == "" {
		return ""
	}
	return v.weaponClassOffHand
}

func (v *InventoryItemWeapon) InventoryItemName() string {
	if v == nil {
		return ""
	}
	return v.itemName
}

func (v *InventoryItemWeapon) InventoryItemType() d2enum.InventoryItemType {
	return d2enum.InventoryItemTypeWeapon
}

func (v *InventoryItemWeapon) InventoryGridSize() (int, int) {
	return v.inventorySizeX, v.inventorySizeY
}

func (v *InventoryItemWeapon) InventoryGridSlot() (int, int) {
	return v.inventorySlotX, v.inventorySlotY
}

func (v *InventoryItemWeapon) SetInventoryGridSlot(x int, y int) {
	v.inventorySlotX, v.inventorySlotY = x, y
}

func (v *InventoryItemWeapon) Serialize() []byte {
	return []byte{}
}

func (v *InventoryItemWeapon) ItemCode() string {
	if v == nil {
		return ""
	}
	return v.itemCode
}
