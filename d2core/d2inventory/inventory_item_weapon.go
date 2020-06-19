package d2inventory

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

type InventoryItemWeapon struct {
	InventorySizeX     int    `json:"inventorySizeX"`
	InventorySizeY     int    `json:"inventorySizeY"`
	InventorySlotX     int    `json:"inventorySlotX"`
	InventorySlotY     int    `json:"inventorySlotY"`
	ItemName           string `json:"itemName"`
	ItemCode           string `json:"itemCode"`
	WeaponClass        string `json:"weaponClass"`
	WeaponClassOffHand string `json:"weaponClassOffHand"`
}

func GetWeaponItemByCode(code string) *InventoryItemWeapon {
	// TODO: Non-normal codes will fail here...
	result := d2datadict.Weapons[code]
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

func (v *InventoryItemWeapon) GetWeaponClass() string {
	if v == nil || v.ItemCode == "" {
		return "hth"
	}
	return v.WeaponClass
}

func (v *InventoryItemWeapon) GetWeaponClassOffHand() string {
	if v == nil || v.ItemCode == "" {
		return ""
	}
	return v.WeaponClassOffHand
}

func (v *InventoryItemWeapon) InventoryItemName() string {
	if v == nil {
		return ""
	}
	return v.ItemName
}

func (v *InventoryItemWeapon) InventoryItemType() d2enum.InventoryItemType {
	return d2enum.InventoryItemTypeWeapon
}

func (v *InventoryItemWeapon) InventoryGridSize() (int, int) {
	return v.InventorySizeX, v.InventorySizeY
}

func (v *InventoryItemWeapon) InventoryGridSlot() (int, int) {
	return v.InventorySlotX, v.InventorySlotY
}

func (v *InventoryItemWeapon) SetInventoryGridSlot(x int, y int) {
	v.InventorySlotX, v.InventorySlotY = x, y
}

func (v *InventoryItemWeapon) Serialize() []byte {
	return []byte{}
}

func (v *InventoryItemWeapon) GetItemCode() string {
	if v == nil {
		return ""
	}
	return v.ItemCode
}
