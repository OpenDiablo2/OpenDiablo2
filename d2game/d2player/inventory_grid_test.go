package d2player

type TestItem struct {
	width          int
	height         int
	inventorySlotX int
	inventorySlotY int
}

func (t *TestItem) InventoryGridSize() (int, int) {
	return t.width, t.height
}

func (t *TestItem) GetItemCode() string {
	return ""
}

func (t *TestItem) InventoryGridSlot() (int, int) {
	return t.inventorySlotX, t.inventorySlotY
}

func (t *TestItem) SetInventoryGridSlot(x int, y int) {
	t.inventorySlotX, t.inventorySlotY = x, y
}

func NewTestItem(width int, height int) *TestItem {
	return &TestItem{width: width, height: height}
}

//func TestItemGrid_Add_Basic(t *testing.T) {
//	grid := NewItemGrid(2, 2, 0, 0)
//
//	tl := NewTestItem(1, 1)
//	tr := NewTestItem(1, 1)
//	bl := NewTestItem(1, 1)
//	br := NewTestItem(1, 1)
//
//	added, err := grid.Add(tl, tr, bl, br)
//
//	assert.Equal(t, 4, added)
//	assert.NoError(t, err)
//	assert.Equal(t, tl, grid.GetSlot(0, 0))
//	assert.Equal(t, tr, grid.GetSlot(1, 0))
//	assert.Equal(t, bl, grid.GetSlot(0, 1))
//	assert.Equal(t, br, grid.GetSlot(1, 1))
//
//}

//func TestItemGrid_Add_OverflowBasic(t *testing.T) {
//	grid := NewItemGrid(2, 2, 0, 0)
//
//	tl := NewTestItem(1, 1)
//	tr := NewTestItem(1, 1)
//	bl := NewTestItem(1, 1)
//	br := NewTestItem(1, 1)
//	o := NewTestItem(1, 1)
//
//	added, err := grid.Add(tl, tr, bl, br, o)
//
//	assert.Equal(t, 4, added)
//	assert.Error(t, err)
//	assert.Equal(t, tl, grid.GetSlot(0, 0))
//	assert.Equal(t, tr, grid.GetSlot(1, 0))
//	assert.Equal(t, bl, grid.GetSlot(0, 1))
//	assert.Equal(t, br, grid.GetSlot(1, 1))
//
//}

//func TestItemGrid_Add_LargeItem(t *testing.T) {
//	grid := NewItemGrid(3, 3, 0, 0)
//
//	tl := NewTestItem(1, 1)
//	o := NewTestItem(2, 2)
//
//	added, err := grid.Add(tl, o)
//
//	assert.Equal(t, 2, added)
//	assert.NoError(t, err)
//	assert.Equal(t, tl, grid.GetSlot(0, 0))
//	assert.Equal(t, o, grid.GetSlot(1, 0))
//	assert.Equal(t, o, grid.GetSlot(2, 0))
//	assert.Nil(t, grid.GetSlot(0, 1))
//	assert.Equal(t, o, grid.GetSlot(1, 1))
//	assert.Equal(t, o, grid.GetSlot(2, 1))
//	assert.Nil(t, grid.GetSlot(0, 2))
//	assert.Nil(t, grid.GetSlot(1, 2))
//	assert.Nil(t, grid.GetSlot(2, 2))
//
//}

//func TestItemGrid_Add_OverflowLargeItem(t *testing.T) {
//	grid := NewItemGrid(2, 2, 0, 0)
//
//	tl := NewTestItem(1, 1)
//	o := NewTestItem(2, 2)
//
//	added, err := grid.Add(tl, o)
//
//	assert.Equal(t, 1, added)
//	assert.Error(t, err)
//	assert.Equal(t, tl, grid.GetSlot(0, 0))
//	assert.Nil(t, grid.GetSlot(1, 0))
//	assert.Nil(t, grid.GetSlot(0, 1))
//	assert.Nil(t, grid.GetSlot(1, 1))
//
//}
