package d2player

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2item/diablo2item"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type Inventory struct {
	frame *d2ui.Sprite
	panel *d2ui.Sprite
	grid  *ItemGrid

	hoverLabel     *d2ui.Label
	hoverX, hoverY int
	hovering       bool

	originX, originY       int
	lastMouseX, lastMouseY int

	isOpen bool
}

func NewInventory(record *d2datadict.InventoryRecord) *Inventory {

	hoverLabel := d2ui.CreateLabel(d2resource.FontFormal11, d2resource.PaletteStatic)
	hoverLabel.Alignment = d2gui.HorizontalAlignCenter

	return &Inventory{
		grid:       NewItemGrid(record),
		originX:    record.Panel.Left,
		hoverLabel: &hoverLabel,
		// originY: record.Panel.Top,
		originY: 0, // expansion data has these all offset by +60 ...
	}
}

func (g *Inventory) IsOpen() bool {
	return g.isOpen
}

func (g *Inventory) Toggle() {
	g.isOpen = !g.isOpen
}

func (g *Inventory) Open() {
	g.isOpen = true
}

func (g *Inventory) Close() {
	g.isOpen = false
}

func (g *Inventory) Load() {
	animation, _ := d2asset.LoadAnimation(d2resource.Frame, d2resource.PaletteSky)
	g.frame, _ = d2ui.LoadSprite(animation)

	animation, _ = d2asset.LoadAnimation(d2resource.InventoryCharacterPanel, d2resource.PaletteSky)
	g.panel, _ = d2ui.LoadSprite(animation)
	items := []InventoryItem{
		diablo2item.NewItem("kit", "Crimson", "of the Bat", "of Frost").Identify(),
		diablo2item.NewItem("rin", "Steel", "of Shock").Identify(),
		diablo2item.NewItem("jav").Identify(),
		diablo2item.NewItem("buc").Identify(),
		//diablo2item.NewItem("Arctic Furs", "qui"),
		// TODO: Load the player's actual items
	}
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotLeftArm, diablo2item.NewItem("wnd"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotRightArm, diablo2item.NewItem("buc"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotHead, diablo2item.NewItem("crn"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotTorso, diablo2item.NewItem("plt"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotLegs, diablo2item.NewItem("vbt"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotBelt, diablo2item.NewItem("vbl"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotGloves, diablo2item.NewItem("lgl"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotLeftHand, diablo2item.NewItem("rin"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotRightHand, diablo2item.NewItem("rin"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotNeck, diablo2item.NewItem("amu"))
	// TODO: Load the player's actual items
	g.grid.Add(items...)
}

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

	w, h = g.panel.GetCurrentFrameSize()

	g.panel.SetPosition(x, y+h)

	if err := g.panel.Render(target); err != nil {
		return err
	}

	y += h

	// Bottom right
	if err := g.panel.SetCurrentFrame(7); err != nil {
		return err
	}

	w, h = g.panel.GetCurrentFrameSize()
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
	centerX, centerY := g.hoverX, iy - halfH

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
