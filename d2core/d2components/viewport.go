//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom"
)

// static check that Viewport implements Component
var _ akara.Component = &Viewport{}

// Viewport represents the size and position of a scene viewport. This is used
// to control where on screen a viewport is rendered.
type Viewport struct {
	*d2geom.Rectangle
}

// New creates a new Viewport. By default, the viewport size is 800x600,
// and is positioned at the top-left of the screen.
func (*Viewport) New() akara.Component {
	c := &Viewport{}

	const defaultWidth, defaultHeight = 800, 600

	c.Rectangle = &d2geom.Rectangle{
		Left:   0,
		Top:    0,
		Width:  defaultWidth,
		Height: defaultHeight,
	}

	return c
}

// ViewportFactory is a wrapper for the generic component factory that returns Viewport component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Viewport.
type ViewportFactory struct {
	*akara.ComponentFactory
}

// Add adds a Viewport component to the given entity and returns it
func (m *ViewportFactory) Add(id akara.EID) *Viewport {
	return m.ComponentFactory.Add(id).(*Viewport)
}

// Get returns the Viewport component for the given entity, and a bool for whether or not it exists
func (m *ViewportFactory) Get(id akara.EID) (*Viewport, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Viewport), found
}
