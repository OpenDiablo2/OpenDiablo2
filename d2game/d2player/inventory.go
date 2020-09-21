package d2player

import (
	"fmt"
	"image/color"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2item/diablo2item"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

// Inventory represents the inventory
type Inventory struct {
	asset      *d2asset.AssetManager
	item       *diablo2item.ItemFactory
	uiManager  *d2ui.UIManager
	frame      *d2ui.Sprite
	panel      *d2ui.Sprite
	grid       *ItemGrid
	hoverLabel *d2ui.Label
	hoverX     int
	hoverY     int
	originX    int
	originY    int
	lastMouseX int
	lastMouseY int
	hovering   bool
	isOpen     bool
}

// NewInventory creates an inventory instance and returns a pointer to it
func NewInventory(asset *d2asset.AssetManager, ui *d2ui.UIManager,
	record *d2records.InventoryRecord) *Inventory {
	hoverLabel := ui.NewLabel(d2resource.FontFormal11, d2resource.PaletteStatic)
	hoverLabel.Alignment = d2gui.HorizontalAlignCenter

	itemFactory, _ := diablo2item.NewItemFactory(asset) // TODO handle errors

	return &Inventory{
		asset:      asset,
		uiManager:  ui,
		item:       itemFactory,
		grid:       NewItemGrid(asset, ui, record),
		originX:    record.Panel.Left,
		hoverLabel: hoverLabel,
		// originY: record.Panel.Top,
		originY: 0, // expansion data has these all offset by +60 ...
	}
}

// IsOpen returns true if the inventory is open
func (g *Inventory) IsOpen() bool {
	return g.isOpen
}

// Toggle negates the open state of the inventory
func (g *Inventory) Toggle() {
	g.isOpen = !g.isOpen
}

// Open opens the inventory
func (g *Inventory) Open() {
	g.isOpen = true
}

// Close closes the inventory
func (g *Inventory) Close() {
	g.isOpen = false
}

// Load the resources required by the inventory
func (g *Inventory) Load() {
	sprite, err := g.uiManager.NewSprite(d2resource.Frame, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	g.frame = sprite

	g.panel, _ = g.uiManager.NewSprite(d2resource.InventoryCharacterPanel, d2resource.PaletteSky)

	// TODO: remove this item test code
	testInventoryCodes := [][]string{
		{"kit", "Crimson", "of the Bat", "of Frost"},
		{"rin", "Steel", "of Shock"},
		{"jav"},
		{"buc"},
	}

	inventoryItems := make([]InventoryItem, 0)

	for idx := range testInventoryCodes {
		item, err := g.item.NewItem(testInventoryCodes[idx]...)
		if err != nil {
			continue
		}

		item.Identify()
		inventoryItems = append(inventoryItems, item)
	}

	testEquippedItemCodes := map[d2enum.EquippedSlot][]string{
		d2enum.EquippedSlotLeftArm:   {"wnd"},
		d2enum.EquippedSlotRightArm:  {"buc"},
		d2enum.EquippedSlotHead:      {"crn"},
		d2enum.EquippedSlotTorso:     {"plt"},
		d2enum.EquippedSlotLegs:      {"vbt"},
		d2enum.EquippedSlotBelt:      {"vbl"},
		d2enum.EquippedSlotGloves:    {"lgl"},
		d2enum.EquippedSlotLeftHand:  {"rin"},
		d2enum.EquippedSlotRightHand: {"rin"},
		d2enum.EquippedSlotNeck:      {"amu"},
	}

	for slot := range testEquippedItemCodes {
		item, err := g.item.NewItem(testEquippedItemCodes[slot]...)
		if err != nil {
			continue
		}

		g.grid.ChangeEquippedSlot(slot, item)
	}

	// TODO: Load the player's actual items
	_, err = g.grid.Add(inventoryItems...)
	if err != nil {
		fmt.Printf("could not add items to the inventory, err: %v\n", err)
	}
}

// Render draws the inventory onto the given surface
func (g *Inventory) Render(target d2interface.Surface) error {
	if !g.isOpen {
		return nil
	}

	x, y := g.originX, g.originY

	// Frame
	// Top left
	if err := g.frame.SetCurrentFrame(5); err != nil {
		return err
	}

	w, h := g.frame.GetCurrentFrameSize()

	g.frame.SetPosition(x, y+h)

	if err := g.frame.Render(target); err != nil {
		return err
	}

	x += w

	// Top right
	if err := g.frame.SetCurrentFrame(6); err != nil {
		return err
	}

	w, h = g.frame.GetCurrentFrameSize()

	g.frame.SetPosition(x, y+h)

	if err := g.frame.Render(target); err != nil {
		return err
	}

	x += w
	y += h

	// Right
	if err := g.frame.SetCurrentFrame(7); err != nil {
		return err
	}

	w, h = g.frame.GetCurrentFrameSize()

	g.frame.SetPosition(x-w, y+h)

	if err := g.frame.Render(target); err != nil {
		return err
	}

	y += h

	// Bottom right
	if err := g.frame.SetCurrentFrame(8); err != nil {
		return err
	}

	w, h = g.frame.GetCurrentFrameSize()

	g.frame.SetPosition(x-w, y+h)

	if err := g.frame.Render(target); err != nil {
		return err
	}

	x -= w

	// Bottom left
	if err := g.frame.SetCurrentFrame(9); err != nil {
		return err
	}

	w, h = g.frame.GetCurrentFrameSize()

	g.frame.SetPosition(x-w, y+h)

	if err := g.frame.Render(target); err != nil {
		return err
	}

	x, y = g.originX+1, g.originY
	y += 64

	// Panel
	// Top left
	if err := g.panel.SetCurrentFrame(4); err != nil {
		return err
	}

	w, h = g.panel.GetCurrentFrameSize()

	g.panel.SetPosition(x, y+h)

	if err := g.panel.Render(target); err != nil {
		return err
	}

	x += w

	// Top right
	if err := g.panel.SetCurrentFrame(5); err != nil {
		return err
	}

	_, h = g.panel.GetCurrentFrameSize()

	g.panel.SetPosition(x, y+h)

	if err := g.panel.Render(target); err != nil {
		return err
	}

	y += h

	// Bottom right
	if err := g.panel.SetCurrentFrame(7); err != nil {
		return err
	}

	_, h = g.panel.GetCurrentFrameSize()
	g.panel.SetPosition(x, y+h)

	if err := g.panel.Render(target); err != nil {
		return err
	}

	// Bottom left
	if err := g.panel.SetCurrentFrame(6); err != nil {
		return err
	}

	w, h = g.panel.GetCurrentFrameSize()

	g.panel.SetPosition(x-w, y+h)

	if err := g.panel.Render(target); err != nil {
		return err
	}

	g.grid.Render(target)

	hovering := false

	for idx := range g.grid.items {
		item := g.grid.items[idx]
		ix, iy := g.grid.SlotToScreen(item.InventoryGridSlot())
		iw, ih := g.grid.sprites[item.GetItemCode()].GetCurrentFrameSize()
		mx, my := g.lastMouseX, g.lastMouseY
		hovering = hovering || ((mx > ix) && (mx < ix+iw) && (my > iy) && (my < iy+ih))

		if hovering {
			if !g.hovering {
				// set the initial hover coordinates
				// this is so that moving mouse doesnt move the description
				g.hoverX, g.hoverY = mx, my
			}

			g.renderItemDescription(target, item)

			break
		}
	}

	g.hovering = hovering

	return nil
}

func (g *Inventory) renderItemDescription(target d2interface.Surface, i InventoryItem) {
	lines := i.GetItemDescription()

	maxW, maxH := 0, 0
	_, iy := g.grid.SlotToScreen(i.InventoryGridSlot())

	for idx := range lines {
		w, h := g.hoverLabel.GetTextMetrics(lines[idx])

		if maxW < w {
			maxW = w
		}

		maxH += h
	}

	halfW, halfH := maxW/2, maxH/2
	centerX, centerY := g.hoverX, iy-halfH

	if (centerX + halfW) > 800 {
		centerX = 800 - halfW
	}

	if (centerY + halfH) > 600 {
		centerY = 600 - halfH
	}

	target.PushTranslation(centerX, centerY)
	target.PushTranslation(-halfW, -halfH)
	target.DrawRect(maxW, maxH, color.RGBA{0, 0, 0, uint8(200)})
	target.PushTranslation(halfW, 0)

	for idx := range lines {
		g.hoverLabel.SetText(lines[idx])
		_, h := g.hoverLabel.GetTextMetrics(lines[idx])
		g.hoverLabel.Render(target)
		target.PushTranslation(0, h)
	}

	target.PopN(len(lines))
	target.PopN(3)
}
