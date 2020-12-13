//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
	"github.com/gravestench/akara"
)

// static check that Size implements Component
var _ akara.Component = &Size{}

// Size represents an entities width and height as a vector
type Size struct {
	*d2math.Vector2
}

// New creates a new Size. By default, size is (0,0).
func (*Size) New() akara.Component {
	return &Size{
		Vector2: d2math.NewVector2(0, 0),
	}
}

// SizeFactory is a wrapper for the generic component factory that returns Size component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Size.
type SizeFactory struct {
	*akara.ComponentFactory
}

// Add adds a Size component to the given entity and returns it
func (m *SizeFactory) Add(id akara.EID) *Size {
	return m.ComponentFactory.Add(id).(*Size)
}

// Get returns the Size component for the given entity, and a bool for whether or not it exists
func (m *SizeFactory) Get(id akara.EID) (*Size, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Size), found
}
