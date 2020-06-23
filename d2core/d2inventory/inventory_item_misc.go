package d2inventory

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"log"
)

type InventoryItemMisc struct {
	InventorySizeX int    `json:"inventorySizeX"`
	InventorySizeY int    `json:"inventorySizeY"`
	InventorySlotX int    `json:"inventorySlotX"`
	InventorySlotY int    `json:"inventorySlotY"`
	ItemName       string `json:"itemName"`
	ItemCode       string `json:"itemCode"`
}

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

func (v *InventoryItemMisc) InventoryItemName() string {
	if v == nil {
		return ""
	}
	return v.ItemName
}

func (v *InventoryItemMisc) InventoryItemType() d2enum.InventoryItemType {
	return d2enum.InventoryItemTypeItem
}

func (v *InventoryItemMisc) InventoryGridSize() (int, int) {
	return v.InventorySizeX, v.InventorySizeY
}

func (v *InventoryItemMisc) InventoryGridSlot() (int, int) {
	return v.InventorySlotX, v.InventorySlotY
}

func (v *InventoryItemMisc) SetInventoryGridSlot(x int, y int) {
	v.InventorySlotX, v.InventorySlotY = x, y
}

func (v *InventoryItemMisc) Serialize() []byte {
	return []byte{}
}

func (v *InventoryItemMisc) GetItemCode() string {
	if v == nil {
		return ""
	}

	return v.ItemCode
}
