package d2inventory

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
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
