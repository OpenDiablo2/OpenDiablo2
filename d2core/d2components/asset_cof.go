//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2cof"
)

// static check that Cof implements Component
var _ akara.Component = &Cof{}

// Cof is a component that contains an embedded cof struct
type Cof struct {
	*d2cof.COF
}

// New returns a new Cof component. By default, it contains a nil instance.
func (*Cof) New() akara.Component {
	return &Cof{}
}

// CofFactory is a wrapper for the generic component factory that returns Cof component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Cof.
type CofFactory struct {
	*akara.ComponentFactory
}

// Add adds a Cof component to the given entity and returns it
func (m *CofFactory) Add(id akara.EID) *Cof {
	return m.ComponentFactory.Add(id).(*Cof)
}

// Get returns the Cof component for the given entity, and a bool for whether or not it exists
func (m *CofFactory) Get(id akara.EID) (*Cof, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Cof), found
}
