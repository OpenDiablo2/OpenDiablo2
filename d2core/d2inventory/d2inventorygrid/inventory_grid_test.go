package d2inventorygrid

import (
	"testing"
)

func TestInventoryGrid_CellCount(t *testing.T) {
	x, y, w, h := 0, 0, 1, 1
	grid := NewInventoryGrid(w, h)
	want := w * h
	got := len(grid.GetCells(x, y, w, h))

	if !(want == got) {
		t.Errorf("Incorrect number of cells in grid; wanted %d: got %d", want, got)
	}

	// width/height is minimum of 1
	w, h = -1, -1
	grid = NewInventoryGrid(w, h)
	want = w * h
	got = len(grid.GetCells(x, y, w, h))

	if !(want == got) {
		t.Errorf("Incorrect number of cells in grid; wanted %d: got %d", want, got)
	}

	// now checking for common grid size
	w, h = 10, 4
	grid = NewInventoryGrid(w, h)
	want = w * h
	got = len(grid.GetCells(x, y, w, h))

	if !(want == got) {
		t.Errorf("Incorrect number of cells in grid; wanted %d: got %d", want, got)
	}
}

func TestInventoryGrid_Size(t *testing.T) {
	w, h := 1, 1
	grid := NewInventoryGrid(w, h)
	gotWidth, gotHeight := grid.Size()

	if !(w == gotWidth) || !(h == gotHeight) {
		str := "Incorrect number of cells in grid; wanted (%dx%d): got (%dx%d)"
		t.Errorf(str, w, h, gotWidth, gotHeight)
	}
}

func TestInventoryGrid_Place(t *testing.T) {
	var x, y, w, h int
	var grid *InventoryGrid
	var item *GridItem
	w, h = 3, 4
	grid = NewInventoryGrid(w, h)

	// Item is too big, should return an error
	item = NewGridItem(w+1, h)
	if err := grid.Place(x, y, item); err == nil {
		str := "Item with dimensions (%dx%d) SHOULD NOT fit in grid with dimensions (%dx%d)"
		w1, h1 := item.Size()
		w2, h2 := grid.Size()
		t.Errorf(str, w1, h1, w2, h2)
	}

	// item is just big enough, should be placed
	item = NewGridItem(w,h)
	if err :=grid.Place(x, y, item); err != nil {
		str := "Item with dimensions (%dx%d) SHOULD fit in grid with dimensions (%dx%d)"
		w1, h1 := item.Size()
		w2, h2 := grid.Size()
		t.Errorf(str, w1, h1, w2, h2)
	}

	// since the item was placed, the cells should show they are occupied by the item
	if grid.GetCell(x, y).occupier != item {
		t.Errorf("occupier not correctly set on cells at (%d,%d)", 0, 0)
	}

	// now try placing out of bounds
	oob := -1
	grid = NewInventoryGrid(w, h)
	item = NewGridItem(w, h)

	if err := grid.Place(oob, oob, item); err == nil {
		errStr := "placed %dx%d item out of bounds at (%d, %d)"
		t.Errorf(errStr, w, h, oob, oob)
	}
}

func TestInventoryGrid_AutoPlace(t *testing.T) {
	var gridWidth, gridHeight, itemWidth, itemHeight int
	gridWidth, gridHeight = 3, 6
	itemWidth, itemHeight = 1, 3
	grid := NewInventoryGrid(gridWidth, gridHeight)
	items := [6]*GridItem{
		NewGridItem(itemWidth, itemHeight),
		NewGridItem(itemWidth, itemHeight),
		NewGridItem(itemWidth, itemHeight),
		NewGridItem(itemWidth, itemHeight),
		NewGridItem(itemWidth, itemHeight),
		NewGridItem(itemWidth, itemHeight),
	}

	shouldGetPlacedStr := "item %d of %d SHOULD have been auto-placed"
	shouldNotGetPlacedStr := "item SHOULD NOT have been auto-placed"

	// try auto-placing all items
	for idx:= range items {
		if err := grid.AutoPlace(items[idx]); err != nil {
			t.Errorf(shouldGetPlacedStr, idx+1, len(items))
		}
	}

	// all cells should be filled, should return error
	extraItem := NewGridItem(1, 1)
	if err := grid.AutoPlace(extraItem); err == nil {
		t.Errorf(shouldNotGetPlacedStr)
	}

	// now we test with items of different dimensions
	gridWidth, gridHeight = 10, 4
	grid = NewInventoryGrid(gridWidth, gridHeight)
	items2 := [13]*GridItem{
		NewGridItem(2, 4), // A
		NewGridItem(2, 2), // B
		NewGridItem(2, 3), // C
		NewGridItem(2, 2), // D
		NewGridItem(2, 1), // E
		NewGridItem(1, 1), // F
		NewGridItem(2, 2), // G
		NewGridItem(1, 3), // H
		NewGridItem(2, 1), // I
		NewGridItem(1, 2), // J
		NewGridItem(2, 1), // K
		NewGridItem(1, 1), // L
		NewGridItem(1, 1), // M
	}

	// try auto-placing all items
	for idx:= range items2 {
		if err := grid.AutoPlace(items2[idx]); err != nil {
			t.Errorf(shouldGetPlacedStr, idx+1, len(items2))
		}
	}

	// after auto-place, items should be placed like this:
	//	 K K I I E E D D A A
	//	 M H G G C C D D A A
	//	 J H G G C C B B A A
	//	 J H L F C C B B A A
	
	// test each cell for each item and make sure the occupier is set to that item
	for itemIdx := range items2 {
		item := items2[itemIdx]
		cells := grid.GetCells(item.x, item.y, item.width, item.height)

		for cellIdx := range cells {
			cell := cells[cellIdx]
			got := cell.occupier

			if got != item {
				errStr := "item #%d: cell (%d,%d) reported incorrect occupier"
				t.Errorf(errStr, itemIdx, cell.x, cell.y)
			}
		}
	}

	// all cells should be filled, should return error
	if err := grid.AutoPlace(extraItem); err == nil {
		t.Errorf(shouldNotGetPlacedStr)
	}
}

func TestInventoryGrid_Swap(t *testing.T) {
	grid := NewInventoryGrid(2,2)
	item1 := NewGridItem(1, 1)
	item2 := NewGridItem(2, 2)
	item3 := NewGridItem(1, 1)
	item4 := NewGridItem(1, 1)

	// grid should look like:
	// |_|_|
	// |_|A|
	if grid.AutoPlace(item1) != nil {
		t.Error("could not auto-place during swap test")
	}

	// |B|B|
	// |B|B|
	swapped, err := grid.Swap(0, 0, item2)
	if err != nil {
		t.Error(err)
	}

	if swapped != item1 {
		t.Error("did not receive expected swap item")
	}

	if cell := grid.GetCell(0, 1); cell.occupier != item2 {
		t.Error("grid cell not occupied by expected item")
	}


	// attempting to swap a 2x2 at any coordinate in this grid should fail
	// |_|C|
	// |D|_|
	if _, err = grid.Swap(1, 0, item3); err != nil {
		t.Error(err)
	}

	if err = grid.Place(0, 1, item4); err != nil {
		t.Error(err)
	}

	for x := 0; x < grid.width; x++ {
		for y := 0; y < grid.height; y++ {
			swapped, err = grid.Swap(x, y, item2)

			if swapped != nil {
				t.Error("swap should not have worked")
			}

			if err == nil {
				t.Error("swap should have returned an error")
			}
		}
	}
}
