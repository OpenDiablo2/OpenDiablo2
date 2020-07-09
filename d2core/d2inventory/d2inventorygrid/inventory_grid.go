package d2inventorygrid

import (
	"errors"
	"fmt"
)

// PlacementConflict represents a type of conflict situation during item placement
type PlacementConflict int

// PlacementConflict types
const (
	PlacementPossible PlacementConflict = iota
	PlacementSwappable
	PlacementImpossible
)

// NewInventoryGrid creates an inventory grid
func NewInventoryGrid(width, height int) *InventoryGrid {
	if width < 1 {
		width = 1
	}

	if height < 1 {
		height = 1
	}

	grid := &InventoryGrid{
		width:  width,
		height: height,
		items:  make([]*GridItem, 0),
	}

	grid.cells = make([]*GridCell, width*height)

	for idx := 0; idx < (width * height); idx++ {
		x := idx % width
		y := idx / width
		grid.cells[idx] = &GridCell{x: x, y: y, grid: grid}
	}

	return grid
}

// InventoryGrid represents an inventory grid, which can contain items
type InventoryGrid struct {
	width, height int
	items         []*GridItem
	cells         []*GridCell
}

// Size returns the width and height of the grid
func (grid *InventoryGrid) Size() (width, height int) {
	return grid.width, grid.height
}

// Items returns a slice of the items within the inventory grid
func (grid *InventoryGrid) Items() []*GridItem {
	return grid.items
}

// Grab attempts to remove an item from the grid, returning the grid item and an error state.
// The x,y parameters do not have to be the "root" cell of the item
func (grid *InventoryGrid) Grab(x, y int) (*GridItem, error) {
	rootCell := grid.GetCell(x, y)

	if !grid.grabPossible(rootCell) {
		return nil, errors.New("cell (%d, %d) is not occupied by an item")
	}

	item := rootCell.occupier

	return item, grid.removeItem(item)
}

// Place attempts to add a grid item to the grid.
// The x,y parameters are where the "root" cell of the item will be.
func (grid *InventoryGrid) Place(x, y int, item *GridItem) error {
	conflict := grid.checkConflicts(x, y, item)

	if conflict == PlacementPossible {
		return grid.addItem(x, y, item)
	}

	errStr := "could not place item with dimensions %dx%d t (%d,%d)"

	return fmt.Errorf(errStr, item.width, item.height, x, y)
}

// Swap will attempt to swap an item with a possible existing item.
// The x,y parameters denote the "root" cell for the item being swapped-in.
// The item being swapped out will be returned, along with an error state.
// If there is no placement conflict (no item to swap out) the item will still
// be added to the grid.
func (grid *InventoryGrid) Swap(x, y int, swapIn *GridItem) (swapped *GridItem, err error) {
	switch grid.checkConflicts(x, y, swapIn) {
	case PlacementPossible:
		return nil, grid.addItem(x, y, swapIn)
	case PlacementSwappable:
		return grid.swapItem(x, y, swapIn)
	}

	return nil, fmt.Errorf("could not perform swap at (%d,%d)", x, y)
}

// AutoPlace will attempt to automatically place an item in the grid
func (grid *InventoryGrid) AutoPlace(item *GridItem) error {
	// first candidate x,y is grid width/height - item width/height
	// eg. for 10x10 grid placing a 2x4 item, first candidate coordinate is 0-indexed (7,5)
	candidateX, candidateY := grid.width-item.width, grid.height-item.height

	// starting from highest candidate coordinates, work our way down to 0,0
	// if we find a group of cells that isn't occupied at all, we place the item
	for x := candidateX; x >= 0; x-- {
		for y := candidateY; y >= 0; y-- {
			candidateCell := grid.GetCell(candidateX, candidateY)

			// happens on out of bounds cell
			if candidateCell == nil {
				continue
			}

			if candidateCell.IsOccupied() {
				continue
			}

			if grid.checkConflicts(x, y, item) == PlacementPossible {
				return grid.Place(x, y, item)
			}
		}
	}

	return fmt.Errorf("could not autoplace item")
}

// GetCells returns a slice of cells for a given "rectangle" in the grid.
// The rectangle parameters are origin x,y (the top-left corner), width, and height.
// Cells which are out of the grid bounds will be nil entries in the slice.
func (grid *InventoryGrid) GetCells(x, y, width, height int) []*GridCell {
	// out of bounds cells will remain nil, we expect it
	result := make([]*GridCell, width*height)
	resultIndex := 0

	for rowIndex := 0; rowIndex < height; rowIndex++ {
		for colIndex := 0; colIndex < width; colIndex++ {
			result[resultIndex] = grid.GetCell(x+colIndex, y+rowIndex)
			resultIndex++
		}
	}

	return result
}

// GetCell returns a cell at the given grid coordinate (nil if out of bounds)
func (grid *InventoryGrid) GetCell(x, y int) *GridCell {
	index := (y * grid.width) + x

	if index < 0 || index > (len(grid.cells)-1) {
		return nil
	}

	return grid.cells[index]
}

func (grid *InventoryGrid) grabPossible(cell *GridCell) bool {
	return cell.occupier != nil
}

func (grid *InventoryGrid) addItem(x, y int, item *GridItem) error {
	cells := grid.GetCells(x, y, item.width, item.height)
	for idx := range cells {
		if cells[idx].occupier != nil {
			return errors.New("cannot add item to cells that are occupied")
		}
	}

	for idx := range cells {
		cells[idx].occupier = item
	}

	item.x, item.y = cells[0].x, cells[0].y
	grid.items = append(grid.items, item)

	return nil
}

func (grid *InventoryGrid) removeItem(item *GridItem) error {
	itemCells := grid.GetCells(item.x, item.y, item.width, item.height)
	for idx := range itemCells {
		itemCells[idx].occupier = nil
	}

	item.x, item.y = OutOfBounds, OutOfBounds

	for idx := 0; idx < len(grid.items); idx++ {
		if grid.items[idx] != item {
			continue
		}

		items := make([]*GridItem, 0)
		items = append(items, grid.items[:idx]...)
		items = append(items, grid.items[idx+1:]...)
		grid.items = items

		return nil
	}

	return errors.New("cannot remove item from grid")
}

func (grid *InventoryGrid) swapItem(x, y int, swapIn *GridItem) (swapOut *GridItem, err error) {
	cells := grid.GetCells(x, y, swapIn.width, swapIn.height)
	for idx := range cells {
		if cells[idx].occupier != nil {
			swapOut = cells[idx].occupier
			break
		}
	}

	if swapOut == nil {
		return swapOut, errors.New("could not perform swap with nil item")
	}

	if swapErr := grid.removeItem(swapOut); swapErr != nil {
		err = swapErr
		return swapOut, err
	}

	if placeErr := grid.Place(x, y, swapIn); placeErr != nil {
		err = placeErr
		return swapOut, err
	}

	return swapOut, err
}

func (grid *InventoryGrid) checkConflicts(x, y int, item *GridItem) PlacementConflict {
	var encountered *GridItem

	w, h := item.Size()
	cells := grid.GetCells(x, y, w, h)

	// conditions for impossible placement:
	// 	* a cell is out of bounds
	// 	* there are 2 distinct occupiers in the range of cells we check
	for idx := range cells {
		if cells[idx] == nil {
			// the cell is out of bounds
			return PlacementImpossible
		}

		occupier := cells[idx].occupier
		if encountered == nil && occupier != nil {
			// we've encountered an occupier
			encountered = occupier
		}

		if occupier != nil && encountered != occupier {
			// we've encountered another occupier
			return PlacementImpossible
		}
	}

	if encountered != nil {
		// we've encountered only one distinct occupier
		return PlacementSwappable
	}

	// we've encountered no occupiers and all cells were within bounds
	return PlacementPossible
}
