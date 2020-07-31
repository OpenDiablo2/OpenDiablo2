package d2mapentity

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2item/diablo2item"
)

const (
	errInvalidItemCodes = "invalid item codes supplied"
)

// Item is a map entity for an item
type Item struct {
	*AnimatedEntity
	Item *diablo2item.Item
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

// Name returns the item name
func (i *Item) Name() string {
	return i.Item.Name()
}
