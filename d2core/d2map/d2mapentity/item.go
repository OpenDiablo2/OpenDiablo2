package d2mapentity

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2item/diablo2item"
)

// static check that item implements map entity interface
var _ d2interface.MapEntity = &Item{}

// Item is a map entity for an item
type Item struct {
	*AnimatedEntity
	Item *diablo2item.Item
}

// ID returns the item uuid
func (i *Item) ID() string {
	return i.AnimatedEntity.uuid
}

// GetPosition returns the item position vector
func (i *Item) GetPosition() d2vector.Position {
	return i.AnimatedEntity.Position
}

// GetVelocity returns the item velocity vector
func (i *Item) GetVelocity() d2vector.Vector {
	return i.AnimatedEntity.velocity
}

// Selectable always returns true for items
func (i *Item) Selectable() bool {
	return true
}

// Highlight sets the highlight flag for a single render tick
func (i *Item) Highlight() {
	i.AnimatedEntity.highlight = true
}

// Label returns the item label
func (i *Item) Label() string {
	return i.Item.Label()
}

// GetSize returns the current frame size
func (i *Item) GetSize() (width, height int) {
	w, h := i.animation.GetCurrentFrameSize()

	if w < minHitboxSize {
		w = minHitboxSize
	}

	if h < minHitboxSize {
		h = minHitboxSize
	}

	return w, h
}
