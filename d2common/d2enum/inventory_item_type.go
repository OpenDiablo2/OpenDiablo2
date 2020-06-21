package d2enum

type InventoryItemType int

const (
	InventoryItemTypeItem   InventoryItemType = 0 // item
	InventoryItemTypeWeapon InventoryItemType = 1 // Weapon
	InventoryItemTypeArmor  InventoryItemType = 2 // Armor
)
