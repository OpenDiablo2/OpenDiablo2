package d2player

import (
	"errors"
	"fmt"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

type InventoryItem interface {
	InventoryGridSize() (width int, height int)
	ItemCode() string
	InventoryGridSlot() (x int, y int)
	SetInventoryGridSlot(x int, y int)
}

var ErrorInventoryFull = errors.New("inventory full")

// Reusable grid for use with player and merchant inventory.
// Handles layout and rendering item icons based on code.
type ItemGrid struct {
	items    []InventoryItem
	width    int
	height   int
	originX  int
	originY  int
	sprites  map[string]*d2ui.Sprite
	slotSize int
}

func NewItemGrid(width int, height int, originX int, originY int) *ItemGrid {
	return &ItemGrid{
		width:    width,
		height:   height,
		originX:  originX,
		originY:  originY,
		slotSize: 29,
		sprites:  make(map[string]*d2ui.Sprite),
	}
}

func (g *ItemGrid) SlotToScreen(slotX int, slotY int) (screenX int, screenY int) {
	screenX = g.originX + slotX*g.slotSize
	screenY = g.originY + slotY*g.slotSize
	return screenX, screenY
}

func (g *ItemGrid) ScreenToSlot(screenX int, screenY int) (slotX int, slotY int) {
	slotX = (screenX - g.originX) / g.slotSize
	slotY = (screenY - g.originY) / g.slotSize
	return slotX, slotY
}

func (g *ItemGrid) GetSlot(x int, y int) InventoryItem {
	for _, item := range g.items {
		slotX, slotY := item.InventoryGridSlot()
		width, height := item.InventoryGridSize()

		if x >= slotX && x < slotX+width && y >= slotY && y < slotY+height {
			return item
		}
	}

	return nil
}

// Add places a given set of items into the first available slots.
// Returns a count of the number of items which could be inserted.
func (g *ItemGrid) Add(items ...InventoryItem) (int, error) {
	added := 0
	var err error

	for _, item := range items {
		if g.add(item) {
			added++
		} else {
			err = ErrorInventoryFull
			break
		}
	}

	g.Load(items...)

	return added, err
}

// Load reads the inventory sprites for items into local cache for rendering.
func (g *ItemGrid) Load(items ...InventoryItem) {
	var itemSprite *d2ui.Sprite

	for _, item := range items {
		if _, exists := g.sprites[item.ItemCode()]; exists {
			// Already loaded, don't reload.
			continue
		}

		// TODO: Put the pattern into D2Shared
		animation, err := d2asset.LoadAnimation(
			fmt.Sprintf("/data/global/items/inv%s.dc6", item.ItemCode()),
			d2resource.PaletteSky,
		)
		if err != nil {
			log.Printf("failed to load sprite for item (%s): %v", item.ItemCode(), err)
			continue
		}
		itemSprite, err = d2ui.LoadSprite(animation)

		g.sprites[item.ItemCode()] = itemSprite
	}

}

// Walk from top left to bottom right until a position large enough to hold the item is found.
// This is inefficient but simplifies the storage.  At most a hundred or so cells will be looped, so impact is minimal.
func (g *ItemGrid) add(item InventoryItem) bool {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			if !g.canFit(x, y, item) {
				continue
			}

			g.set(x, y, item)
			return true
		}
	}

	return false
}

// canFit loops over all items to determine if any other items would overlap the given position.
func (g *ItemGrid) canFit(x int, y int, item InventoryItem) bool {
	insertWidth, insertHeight := item.InventoryGridSize()
	if x+insertWidth > g.width || y+insertHeight > g.height {
		return false
	}

	for _, compItem := range g.items {
		slotX, slotY := compItem.InventoryGridSlot()
		compWidth, compHeight := compItem.InventoryGridSize()

		if x+insertWidth >= slotX &&
			x < slotX+compWidth &&
			y+insertHeight >= slotY &&
			y < slotY+compHeight {
			return false
		}
	}

	return true
}

func (g *ItemGrid) Set(x int, y int, item InventoryItem) error {
	if !g.canFit(x, y, item) {
		return fmt.Errorf("can not set item (%s) to position (%v, %v)", item.ItemCode(), x, y)
	}
	g.set(x, y, item)
	return nil
}

func (g *ItemGrid) set(x int, y int, item InventoryItem) {
	item.SetInventoryGridSlot(x, y)
	g.items = append(g.items, item)

	g.Load(item)
}

// Remove does an in place filter to remove the element from the slice of items.
func (g *ItemGrid) Remove(item InventoryItem) {
	n := 0
	for _, compItem := range g.items {
		if compItem == item {
			continue
		}
		g.items[n] = compItem
		n++
	}

	g.items = g.items[:n]
}

func (g *ItemGrid) Render(target d2render.Surface) {
	for _, item := range g.items {
		if item == nil {
			continue
		}

		itemSprite := g.sprites[item.ItemCode()]
		if itemSprite == nil {
			// In case it failed to load.
			// TODO: fallback to something
			continue
		}

		slotX, slotY := g.SlotToScreen(item.InventoryGridSlot())
		_, h := itemSprite.GetCurrentFrameSize()
		itemSprite.SetPosition(slotX, slotY+h)
		_ = itemSprite.Render(target)

	}
}
