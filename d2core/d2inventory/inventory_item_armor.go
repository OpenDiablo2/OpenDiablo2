package d2inventory

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

type InventoryItemArmor struct {
	inventorySizeX int
	inventorySizeY int
	inventorySlotX int
	inventorySlotY int
	itemName       string
	itemCode       string
	armorClass     string
}

func GetArmorItemByCode(code string) *InventoryItemArmor {
	result := d2datadict.Armors[code]
	if result == nil {
		log.Fatalf("Could not find armor entry for code '%s'", code)
	}
	return &InventoryItemArmor{
		inventorySizeX: result.InventoryWidth,
		inventorySizeY: result.InventoryHeight,
		itemName:       result.Name,
		itemCode:       result.Code,
		armorClass:     "lit", // TODO: Where does this come from?
	}
}

func (v *InventoryItemArmor) ArmorClass() string {
	if v == nil || v.itemCode == "" {
		return "lit"
	}
	return v.armorClass
}

func (v *InventoryItemArmor) InventoryItemName() string {
	if v == nil {
		return ""
	}

	return v.itemName
}

func (v *InventoryItemArmor) InventoryItemType() d2enum.InventoryItemType {
	return d2enum.InventoryItemTypeArmor
}

func (v *InventoryItemArmor) InventoryGridSize() (int, int) {
	return v.inventorySizeX, v.inventorySizeY
}

func (v *InventoryItemArmor) InventoryGridSlot() (int, int) {
	return v.inventorySlotX, v.inventorySlotY
}

func (v *InventoryItemArmor) SetInventoryGridSlot(x int, y int) {
	v.inventorySlotX, v.inventorySlotY = x, y
}

func (v *InventoryItemArmor) Serialize() []byte {
	return []byte{}
}

func (v *InventoryItemArmor) ItemCode() string {
	if v == nil {
		return ""
	}

	return v.itemCode
}
