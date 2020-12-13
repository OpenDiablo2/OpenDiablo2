package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2bitmapfont"
)

// static check that BitmapFont implements Component
var _ akara.Component = &BitmapFont{}

// BitmapFont represent a font made from a font table, a sprite, and a palette (d2 files)
type BitmapFont struct {
	*d2bitmapfont.BitmapFont
}

// New creates a new BitmapFont.
func (*BitmapFont) New() akara.Component {
	return &BitmapFont{}
}

// BitmapFontFactory is a wrapper for the generic component factory that returns BitmapFont component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a BitmapFont.
type BitmapFontFactory struct {
	*akara.ComponentFactory
}

// Add adds a BitmapFont component to the given entity and returns it
func (m *BitmapFontFactory) Add(id akara.EID) *BitmapFont {
	return m.ComponentFactory.Add(id).(*BitmapFont)
}

// Get returns the BitmapFont component for the given entity, and a bool for whether or not it exists
func (m *BitmapFontFactory) Get(id akara.EID) (*BitmapFont, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*BitmapFont), found
}

