//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

// static check that Origin implements Component
var _ akara.Component = &Origin{}

// Origin is a component that describes the origin point of an entity as a vector.
// The values should be interpreted as normalized to the width/height of the entity (depends on other components...).
// For example, origin (0,0) should be top-left corner, (0.5, 0.5) should be center.
type Origin struct {
	*d2math.Vector3
}

// New creates a new Origin. By default, the origin is the top-left corner (0,0)
func (*Origin) New() akara.Component {
	return &Origin{
		Vector3: d2math.NewVector3(0, 0, 0),
	}
}

// OriginFactory is a wrapper for the generic component factory that returns Origin component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Origin.
type OriginFactory struct {
	*akara.ComponentFactory
}

// Add adds a Origin component to the given entity and returns it
func (m *OriginFactory) Add(id akara.EID) *Origin {
	return m.ComponentFactory.Add(id).(*Origin)
}

// Get returns the Origin component for the given entity, and a bool for whether or not it exists
func (m *OriginFactory) Get(id akara.EID) (*Origin, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Origin), found
}
