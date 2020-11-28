//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// static check that Renderable implements Component
var _ akara.Component = &Renderable{}

// Renderable is a component that contains an embedded surface interface, which is used for rendering
type Renderable struct {
	d2interface.Surface
}

// New returns a Renderable component. By default, it contains a nil instance.
func (*Renderable) New() akara.Component {
	return &Renderable{}
}

// RenderableFactory is a wrapper for the generic component factory that returns Renderable component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Renderable.
type RenderableFactory struct {
	Renderable *akara.ComponentFactory
}

// AddRenderable adds a Renderable component to the given entity and returns it
func (m *RenderableFactory) AddRenderable(id akara.EID) *Renderable {
	return m.Renderable.Add(id).(*Renderable)
}

// GetRenderable returns the Renderable component for the given entity, and a bool for whether or not it exists
func (m *RenderableFactory) GetRenderable(id akara.EID) (*Renderable, bool) {
	component, found := m.Renderable.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Renderable), found
}
