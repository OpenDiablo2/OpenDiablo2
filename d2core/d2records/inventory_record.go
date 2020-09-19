package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// Inventory holds all of the inventory records from inventory.txt
type Inventory map[string]*InventoryRecord //nolint:gochecknoglobals // Currently global by design

// InventoryRecord represents a single row from inventory.txt, it describes the grid
// layout and positioning of various inventory-related ui panels.
type InventoryRecord struct {
	Name  string
	Panel *box
	Grid  *grid
	Slots map[d2enum.EquippedSlot]*box
}

type box struct {
	Left   int
	Right  int
	Top    int
	Bottom int
	Width  int
	Height int
}

type grid struct {
	Box        *box
	Rows       int
	Columns    int
	CellWidth  int
	CellHeight int
}
