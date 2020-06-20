package d2inventory

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// InventoryItemMisc stores the info of an miscellaneous item in the inventory
type InventoryItemMisc struct {
	InventorySizeX int    `json:"inventorySizeX"`
	InventorySizeY int    `json:"inventorySizeY"`
	InventorySlotX int    `json:"inventorySlotX"`
	InventorySlotY int    `json:"inventorySlotY"`
	ItemName       string `json:"itemName"`
	ItemCode       string `json:"itemCode"`
}

// GetMiscItemByCode returns the miscellaneous item for the given code
func GetMiscItemByCode(code string) *InventoryItemMisc {
	result := d2datadict.MiscItems[code]
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

// InventoryItemName returns the name of the miscellaneous item
func (v *InventoryItemMisc) InventoryItemName() string {
	if v == nil {
		return ""
	}

	return v.ItemName
}

// InventoryItemType returns the item type of the miscellaneous item
func (v *InventoryItemMisc) InventoryItemType() d2enum.InventoryItemType {
	return d2enum.InventoryItemTypeItem
}

// InventoryGridSize returns the grid size of the miscellaneous item
func (v *InventoryItemMisc) InventoryGridSize() (sizeX, sizeY int) {
	return v.InventorySizeX, v.InventorySizeY
}

// InventoryGridSlot returns the grid slot coordinates of the miscellaneous item
func (v *InventoryItemMisc) InventoryGridSlot() (slotX, slotY int) {
	return v.InventorySlotX, v.InventorySlotY
}

// SetInventoryGridSlot sets the InventorySlotX and InventorySlotY of the miscellaneous item with the given x and y values
func (v *InventoryItemMisc) SetInventoryGridSlot(x, y int) {
	v.InventorySlotX, v.InventorySlotY = x, y
}

// Serialize returns the miscellaneous item object as a byte array
func (v *InventoryItemMisc) Serialize() []byte {
	return []byte{}
}

// GetItemCode returns the item code of the miscellaneous item
func (v *InventoryItemMisc) GetItemCode() string {
	if v == nil {
		return ""
	}

	return v.ItemCode
}
