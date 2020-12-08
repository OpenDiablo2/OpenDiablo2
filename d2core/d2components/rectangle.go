//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom/rectangle"
	"github.com/gravestench/akara"
)

// static check that Rectangle implements Component
var _ akara.Component = &Rectangle{}

// Rectangle represents an entities x,y axis scale as a vector
type Rectangle struct {
	rectangle.Rectangle
}

// New creates a new Rectangle instance. By default, the scale is (1,1)
func (*Rectangle) New() akara.Component {
	return &Rectangle{
		Rectangle: rectangle.Rectangle{},
	}
}

// RectangleFactory is a wrapper for the generic component factory that returns Rectangle component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Rectangle.
type RectangleFactory struct {
	*akara.ComponentFactory
}

// Add adds a Rectangle component to the given entity and returns it
func (m *RectangleFactory) Add(id akara.EID) *Rectangle {
	return m.ComponentFactory.Add(id).(*Rectangle)
}

// Get returns the Rectangle component for the given entity, and a bool for whether or not it exists
func (m *RectangleFactory) Get(id akara.EID) (*Rectangle, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Rectangle), found
}
