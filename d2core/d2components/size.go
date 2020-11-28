//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
)

// static check that Size implements Component
var _ akara.Component = &Size{}

// Size represents an entities width and height as a vector
type Size struct {
	*d2vector.Vector
}

// New creates a new Size. By default, size is (0,0).
func (*Size) New() akara.Component {
	return &Size{
		Vector: d2vector.NewVector(0, 0),
	}
}

// SizeFactory is a wrapper for the generic component factory that returns Size component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Size.
type SizeFactory struct {
	Size *akara.ComponentFactory
}

// AddSize adds a Size component to the given entity and returns it
func (m *SizeFactory) AddSize(id akara.EID) *Size {
	return m.Size.Add(id).(*Size)
}

// GetSize returns the Size component for the given entity, and a bool for whether or not it exists
func (m *SizeFactory) GetSize(id akara.EID) (*Size, bool) {
	component, found := m.Size.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Size), found
}
