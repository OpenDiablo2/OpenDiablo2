package d2inventory

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

type InventoryItemArmor struct {
	InventorySizeX int    `json:"inventorySizeX"`
	InventorySizeY int    `json:"inventorySizeY"`
	InventorySlotX int    `json:"inventorySlotX"`
	InventorySlotY int    `json:"inventorySlotY"`
	ItemName       string `json:"itemName"`
	ItemCode       string `json:"itemCode"`
	ArmorClass     string `json:"armorClass"`
}

func GetArmorItemByCode(code string) *InventoryItemArmor {
	result := d2datadict.Armors[code]
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

func (v *InventoryItemArmor) GetArmorClass() string {
	if v == nil || v.ItemCode == "" {
		return "lit"
	}
	return v.ArmorClass
}

func (v *InventoryItemArmor) InventoryItemName() string {
	if v == nil {
		return ""
	}

	return v.ItemName
}

func (v *InventoryItemArmor) InventoryItemType() d2enum.InventoryItemType {
	return d2enum.InventoryItemTypeArmor
}

func (v *InventoryItemArmor) InventoryGridSize() (int, int) {
	return v.InventorySizeX, v.InventorySizeY
}

func (v *InventoryItemArmor) InventoryGridSlot() (int, int) {
	return v.InventorySlotX, v.InventorySlotY
}

func (v *InventoryItemArmor) SetInventoryGridSlot(x int, y int) {
	v.InventorySlotX, v.InventorySlotY = x, y
}

func (v *InventoryItemArmor) Serialize() []byte {
	return []byte{}
}

func (v *InventoryItemArmor) GetItemCode() string {
	if v == nil {
		return ""
	}

	return v.ItemCode
}
