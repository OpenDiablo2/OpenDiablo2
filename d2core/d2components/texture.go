//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// static check that Texture implements Component
var _ akara.Component = &Texture{}

// Texture is a component that contains an embedded surface interface, which is used for rendering
type Texture struct {
	Texture d2interface.Surface
}

// New returns a Texture component. By default, it contains a nil instance.
func (*Texture) New() akara.Component {
	return &Texture{}
}

// TextureFactory is a wrapper for the generic component factory that returns Texture component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Texture.
type TextureFactory struct {
	*akara.ComponentFactory
}

// Add adds a Texture component to the given entity and returns it
func (m *TextureFactory) Add(id akara.EID) *Texture {
	return m.ComponentFactory.Add(id).(*Texture)
}

// Get returns the Texture component for the given entity, and a bool for whether or not it exists
func (m *TextureFactory) Get(id akara.EID) (*Texture, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Texture), found
}
