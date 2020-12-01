//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

// static check that Position implements Component
var _ akara.Component = &Position{}

// Position contains an embedded d2vector.Position, which is a vector with
// helper methods for translating between screen, isometric, tile, and sub-tile space.
type Position struct {
	*d2math.Vector3
}

// New creates a new Position. By default, the position is (0,0)
func (*Position) New() akara.Component {
	return &Position{
		Vector3: d2math.NewVector3(0, 0, 0),
	}
}

// PositionFactory is a wrapper for the generic component factory that returns Position component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Position.
type PositionFactory struct {
	Position *akara.ComponentFactory
}

// AddPosition adds a Position component to the given entity and returns it
func (m *PositionFactory) AddPosition(id akara.EID) *Position {
	return m.Position.Add(id).(*Position)
}

// GetPosition returns the Position component for the given entity, and a bool for whether or not it exists
func (m *PositionFactory) GetPosition(id akara.EID) (*Position, bool) {
	component, found := m.Position.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Position), found
}
