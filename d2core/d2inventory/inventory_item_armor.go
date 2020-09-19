package d2inventory

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// InventoryItemArmor stores the info of an armor item in the inventory
type InventoryItemArmor struct {
	InventorySizeX int    `json:"inventorySizeX"`
	InventorySizeY int    `json:"inventorySizeY"`
	InventorySlotX int    `json:"inventorySlotX"`
	InventorySlotY int    `json:"inventorySlotY"`
	ItemName       string `json:"itemName"`
	ItemCode       string `json:"itemCode"`
	ArmorClass     string `json:"armorClass"`
}

// GetArmorClass returns the class of the armor
func (v *InventoryItemArmor) GetArmorClass() string {
	if v == nil || v.ItemCode == "" {
		return "lit"
	}

	return v.ArmorClass
}

// InventoryItemName returns the name of the armor
func (v *InventoryItemArmor) InventoryItemName() string {
	if v == nil {
		return ""
	}

	return v.ItemName
}

// InventoryItemType returns the item type of the armor
func (v *InventoryItemArmor) InventoryItemType() d2enum.InventoryItemType {
	return d2enum.InventoryItemTypeArmor
}

// InventoryGridSize returns the grid size of the armor
func (v *InventoryItemArmor) InventoryGridSize() (sizeX, sizeY int) {
	return v.InventorySizeX, v.InventorySizeY
}

// InventoryGridSlot returns the grid slot coordinates of the armor
func (v *InventoryItemArmor) InventoryGridSlot() (slotX, slotY int) {
	return v.InventorySlotX, v.InventorySlotY
}

// SetInventoryGridSlot sets the InventorySlotX and InventorySlotY of the armor with the given x and y values
func (v *InventoryItemArmor) SetInventoryGridSlot(x, y int) {
	v.InventorySlotX, v.InventorySlotY = x, y
}

// Serialize returns the armor object as a byte array
func (v *InventoryItemArmor) Serialize() []byte {
	return []byte{}
}

// GetItemCode returns the item code of the armor
func (v *InventoryItemArmor) GetItemCode() string {
	if v == nil {
		return ""
	}

	return v.ItemCode
}
