package d2player

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type Inventory struct {
	frame   *d2ui.Sprite
	panel   *d2ui.Sprite
	grid    *ItemGrid
	originX int
	originY int
	isOpen  bool
}

func NewInventory(record *d2datadict.InventoryRecord) *Inventory {
	return &Inventory{
		grid:    NewItemGrid(record),
		originX: record.Panel.Left,
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
		d2inventory.GetWeaponItemByCode("wnd"),
		d2inventory.GetWeaponItemByCode("sst"),
		d2inventory.GetWeaponItemByCode("jav"),
		d2inventory.GetArmorItemByCode("buc"),
		d2inventory.GetWeaponItemByCode("clb"),
		// TODO: Load the player's actual items
	}
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotLeftArm, d2inventory.GetWeaponItemByCode("wnd"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotRightArm, d2inventory.GetArmorItemByCode("buc"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotHead, d2inventory.GetArmorItemByCode("crn"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotTorso, d2inventory.GetArmorItemByCode("plt"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotLegs, d2inventory.GetArmorItemByCode("vbt"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotBelt, d2inventory.GetArmorItemByCode("vbl"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotGloves, d2inventory.GetArmorItemByCode("lgl"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotLeftHand, d2inventory.GetMiscItemByCode("rin"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotRightHand, d2inventory.GetMiscItemByCode("rin"))
	g.grid.ChangeEquippedSlot(d2enum.EquippedSlotNeck, d2inventory.GetMiscItemByCode("amu"))
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

	x, y = g.originX, g.originY
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

	return nil
}
