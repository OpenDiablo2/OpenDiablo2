package d2player

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2item/diablo2item"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const (
	frameInventoryTopLeft     = 4
	frameInventoryTopRight    = 5
	frameInventoryBottomLeft  = 6
	frameInventoryBottomRight = 7
)

const (
	invCloseButtonX, invCloseButtonY = 419, 449
	invGoldButtonX, invGoldButtonY   = 485, 455
	invGoldLabelX, invGoldLabelY     = 510, 455
)

// NewInventory creates an inventory instance and returns a pointer to it
func NewInventory(asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	l d2util.LogLevel,
	gold int,
	record *d2records.InventoryRecord) *Inventory {
	itemTooltip := ui.NewTooltip(d2resource.FontFormal11, d2resource.PaletteStatic, d2ui.TooltipXCenter, d2ui.TooltipYBottom)

	// https://github.com/OpenDiablo2/OpenDiablo2/issues/797
	itemFactory, _ := diablo2item.NewItemFactory(asset)

	inventory := &Inventory{
		asset:       asset,
		uiManager:   ui,
		item:        itemFactory,
		grid:        NewItemGrid(asset, ui, l, record),
		originX:     record.Panel.Left,
		itemTooltip: itemTooltip,
		// originY: record.Panel.Top,
		originY: 0, // expansion data has these all offset by +60 ...
		gold:    gold,
	}

	inventory.Logger = d2util.NewLogger()
	inventory.Logger.SetLevel(l)
	inventory.Logger.SetPrefix(logPrefix)

	return inventory
}

// Inventory represents the inventory
type Inventory struct {
	asset         *d2asset.AssetManager
	item          *diablo2item.ItemFactory
	uiManager     *d2ui.UIManager
	frame         *d2ui.UIFrame
	panel         *d2ui.Sprite
	grid          *ItemGrid
	itemTooltip   *d2ui.Tooltip
	closeButton   *d2ui.Button
	goldButton    *d2ui.Button
	goldLabel     *d2ui.Label
	panelGroup    *d2ui.WidgetGroup
	panelMoveGold *d2ui.WidgetGroup
	hoverX        int
	hoverY        int
	originX       int
	originY       int
	lastMouseX    int
	lastMouseY    int
	hovering      bool
	isOpen        bool
	onCloseCb     func()
	gold          int

	*d2util.Logger
}

// Toggle negates the open state of the inventory
func (g *Inventory) Toggle() {
	if g.isOpen {
		g.Close()
	} else {
		g.Open()
	}
}

// Load the resources required by the inventory
func (g *Inventory) Load() {
	var err error

	g.panelGroup = g.uiManager.NewWidgetGroup(d2ui.RenderPriorityHeroStatsPanel)
	g.panelMoveGold = g.uiManager.NewWidgetGroup(d2ui.RenderPriorityHeroStatsPanel)

	g.frame = d2ui.NewUIFrame(g.asset, g.uiManager, d2ui.FrameRight)
	g.panelGroup.AddWidget(g.frame)

	g.closeButton = g.uiManager.NewButton(d2ui.ButtonTypeSquareClose, "")
	g.closeButton.SetVisible(false)
	g.closeButton.SetPosition(invCloseButtonX, invCloseButtonY)
	g.closeButton.OnActivated(func() { g.Close() })
	g.panelGroup.AddWidget(g.closeButton)

	g.goldButton = g.uiManager.NewButton(d2ui.ButtonTypeGoldCoin, "")
	g.goldButton.SetVisible(false)
	g.goldButton.SetPosition(invGoldButtonX, invGoldButtonY)
	g.goldButton.OnActivated(func() { g.onGoldClicked() })
	g.panelGroup.AddWidget(g.goldButton)

	g.goldLabel = g.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteStatic)
	g.goldLabel.Alignment = d2ui.HorizontalAlignLeft
	g.goldLabel.SetText(fmt.Sprintln(g.gold))
	g.goldLabel.SetPosition(invGoldLabelX, invGoldLabelY)
	g.panelGroup.AddWidget(g.goldLabel)

	g.panel, err = g.uiManager.NewSprite(d2resource.InventoryCharacterPanel, d2resource.PaletteSky)
	if err != nil {
		g.Error(err.Error())
	}

	g.panelGroup.SetVisible(false)

	// https://github.com/OpenDiablo2/OpenDiablo2/issues/795
	testInventoryCodes := [][]string{
		{"kit", "Crimson", "of the Bat", "of Frost"},
		{"rin", "Steel", "of Shock"},
		{"jav"},
		{"buc"},
	}

	inventoryItems := make([]InventoryItem, 0)

	for idx := range testInventoryCodes {
		item, itemErr := g.item.NewItem(testInventoryCodes[idx]...)
		if itemErr != nil {
			continue
		}

		item.Identify()
		inventoryItems = append(inventoryItems, item)
	}

	// https://github.com/OpenDiablo2/OpenDiablo2/issues/795
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
		item, itemErr := g.item.NewItem(testEquippedItemCodes[slot]...)
		if itemErr != nil {
			continue
		}

		g.grid.ChangeEquippedSlot(slot, item)
	}

	_, err = g.grid.Add(inventoryItems...)
	if err != nil {
		g.Errorf("could not add items to the inventory, err: %v", err.Error())
	}
}

// Open opens the inventory
func (g *Inventory) Open() {
	g.isOpen = true
	g.panelGroup.SetVisible(true)
}

// Close closes the inventory
func (g *Inventory) Close() {
	g.isOpen = false
	g.panelGroup.SetVisible(false)
	g.onCloseCb()
}

// SetOnCloseCb the callback run on closing the inventory
func (g *Inventory) SetOnCloseCb(cb func()) {
	g.onCloseCb = cb
}

func (g *Inventory) onGoldClicked() {
	g.Info("Gold action clicked")
}

// IsOpen returns true if the inventory is open
func (g *Inventory) IsOpen() bool {
	return g.isOpen
}

// Render draws the inventory onto the given surface
func (g *Inventory) Render(target d2interface.Surface) {
	if !g.isOpen {
		return
	}

	g.goldLabel.Render(target)

	err := g.renderFrame(target)
	if err != nil {
		g.Error(err.Error())
	}

	g.grid.Render(target)
	g.renderItemHover(target)
}

func (g *Inventory) renderFrame(target d2interface.Surface) error {
	g.frame.Render(target)

	frames := []int{
		frameInventoryTopLeft,
		frameInventoryTopRight,
		frameInventoryBottomRight,
		frameInventoryBottomLeft,
	}

	x, y := g.originX+1, g.originY
	y += 64

	for _, frame := range frames {
		if err := g.panel.SetCurrentFrame(frame); err != nil {
			return err
		}

		w, h := g.panel.GetCurrentFrameSize()

		g.panel.SetPosition(x, y+h)
		g.panel.Render(target)

		switch frame {
		case frameInventoryTopLeft:
			x += w
		case frameInventoryTopRight:
			y += h
		case frameInventoryBottomRight:
			x = g.originX + 1
		}
	}

	return nil
}

func (g *Inventory) renderItemHover(target d2interface.Surface) {
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
}

func (g *Inventory) renderItemDescription(target d2interface.Surface, i InventoryItem) {
	lines := i.GetItemDescription()
	g.itemTooltip.SetTextLines(lines)
	_, y := g.grid.SlotToScreen(i.InventoryGridSlot())

	g.itemTooltip.SetPosition(g.hoverX, y)
	g.itemTooltip.Render(target)
}
