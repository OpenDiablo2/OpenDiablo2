package d2inventory

type InventoryItemMisc struct {
	InventorySizeX int    `json:"inventorySizeX"`
	InventorySizeY int    `json:"inventorySizeY"`
	InventorySlotX int    `json:"inventorySlotX"`
	InventorySlotY int    `json:"inventorySlotY"`
	ItemName       string `json:"itemName"`
	ItemCode       string `json:"itemCode"`
}
