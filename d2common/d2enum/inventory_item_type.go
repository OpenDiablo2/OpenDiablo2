package d2enum

// InventoryItemType represents a inventory item type
type InventoryItemType int

// Inventry item types
const (
	InventoryItemTypeItem InventoryItemType = iota
	InventoryItemTypeWeapon
	InventoryItemTypeArmor
)
