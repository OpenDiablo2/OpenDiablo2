//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

// static check that Velocity implements Component
var _ akara.Component = &Velocity{}

// Velocity contains an embedded velocity as a vector
type Velocity struct {
	*d2math.Vector3
}

// New creates a new Velocity. By default, the velocity is (0,0).
func (*Velocity) New() akara.Component {
	return &Velocity{
		Vector3: d2math.NewVector3(0, 0, 0),
	}
}

// VelocityFactory is a wrapper for the generic component factory that returns Velocity component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Velocity.
type VelocityFactory struct {
	*akara.ComponentFactory
}

// Add adds a Velocity component to the given entity and returns it
func (m *VelocityFactory) Add(id akara.EID) *Velocity {
	return m.ComponentFactory.Add(id).(*Velocity)
}

// Get returns the Velocity component for the given entity, and a bool for whether or not it exists
func (m *VelocityFactory) Get(id akara.EID) (*Velocity, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Velocity), found
}
