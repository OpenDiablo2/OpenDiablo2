package d2core

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2datadict"
)

type InventoryItemArmor struct {
	inventorySizeX int
	inventorySizeY int
	itemName       string
	itemCode       string
	armorClass     string
}

func GetArmorItemByCode(code string) InventoryItemArmor {
	result := d2datadict.Armors[code]
	if result == nil {
		log.Fatalf("Could not find armor entry for code '%s'", code)
	}
	return InventoryItemArmor{
		inventorySizeX: result.InventoryWidth,
		inventorySizeY: result.InventoryHeight,
		itemName:       result.Name,
		itemCode:       result.Code,
		armorClass:     "lit", // TODO: Where does this come from?
	}
}

func (v InventoryItemArmor) GetArmorClass() string {
	if v.itemCode == "" {
		return "lit"
	}
	return v.armorClass
}

func (v InventoryItemArmor) GetInventoryItemName() string {
	return v.itemName
}

func (v InventoryItemArmor) GetInventoryItemType() d2enum.InventoryItemType {
	return d2enum.InventoryItemTypeArmor
}

func (v InventoryItemArmor) GetInventoryGridSize() (int, int) {
	return v.inventorySizeX, v.inventorySizeY
}

func (v InventoryItemArmor) Serialize() []byte {
	return []byte{}
}

func (v InventoryItemArmor) GetItemCode() string {
	return v.itemCode
}
