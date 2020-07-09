package d2inventorygrid

type InventoryCellState int

const (
	InventoryCellVacant InventoryCellState = iota
	InventoryCellOccupied
)

type GridCell struct{
	x, y int
	grid *InventoryGrid
	state InventoryCellState
	occupier *GridItem
}

func (gc *GridCell) IsVacant() bool {
	return gc.state == InventoryCellVacant
}

func (gc *GridCell) IsOccupied() bool {
	return gc.state == InventoryCellOccupied
}

func (gc *GridCell) Item() *GridItem {
	return gc.occupier
}

