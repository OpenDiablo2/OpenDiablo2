package d2item

type Equipper interface {
	EquippedItems() []Item
	CarriedItems() []Item
}
