package d2inventory

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// InventoryItemWeapon stores the info of an weapon item in the inventory
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

// GetWeaponItemByCode returns the weapon item for the given code
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

// GetWeaponClass returns the class of the weapon
func (v *InventoryItemWeapon) GetWeaponClass() string {
	if v == nil || v.ItemCode == "" {
		return "hth"
	}

	return v.WeaponClass
}

// GetWeaponClassOffHand returns the class of the off hand weapon
func (v *InventoryItemWeapon) GetWeaponClassOffHand() string {
	if v == nil || v.ItemCode == "" {
		return ""
	}

	return v.WeaponClassOffHand
}

// InventoryItemName returns the name of the weapon
func (v *InventoryItemWeapon) InventoryItemName() string {
	if v == nil {
		return ""
	}

	return v.ItemName
}

// InventoryItemType returns the item type of the weapon
func (v *InventoryItemWeapon) InventoryItemType() d2enum.InventoryItemType {
	return d2enum.InventoryItemTypeWeapon
}

// InventoryGridSize returns the grid size of the weapon
func (v *InventoryItemWeapon) InventoryGridSize() (sizeX, sizeY int) {
	return v.InventorySizeX, v.InventorySizeY
}

// InventoryGridSlot returns the grid slot coordinates of the weapon
func (v *InventoryItemWeapon) InventoryGridSlot() (slotX, slotY int) {
	return v.InventorySlotX, v.InventorySlotY
}

// SetInventoryGridSlot sets the InventorySlotX and InventorySlotY of the weapon with the given x and y values
func (v *InventoryItemWeapon) SetInventoryGridSlot(x, y int) {
	v.InventorySlotX, v.InventorySlotY = x, y
}

// Serialize returns the weapon object as a byte array
func (v *InventoryItemWeapon) Serialize() []byte {
	return []byte{}
}

// GetItemCode returns the item code of the weapon
func (v *InventoryItemWeapon) GetItemCode() string {
	if v == nil {
		return ""
	}

	return v.ItemCode
}
