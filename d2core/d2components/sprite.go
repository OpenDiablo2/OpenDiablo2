//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// static check that AnimationComponent implements Component
var _ akara.Component = &Sprite{}

// Sprite is a component that contains a width and height
type Sprite struct {
	d2interface.Sprite
}

// New returns an animation component. By default, it contains a nil instance of an animation.
func (*Sprite) New() akara.Component {
	return &Sprite{}
}

// SpriteFactory is a wrapper for the generic component factory that returns Sprite component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Sprite.
type SpriteFactory struct {
	Sprite *akara.ComponentFactory
}

// AddSprite adds a Sprite component to the given entity and returns it
func (m *SpriteFactory) AddSprite(id akara.EID) *Sprite {
	return m.Sprite.Add(id).(*Sprite)
}

// GetSprite returns the Sprite component for the given entity, and a bool for whether or not it exists
func (m *SpriteFactory) GetSprite(id akara.EID) (*Sprite, bool) {
	component, found := m.Sprite.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Sprite), found
}