//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
)

// static check that Dcc implements Component
var _ akara.Component = &Dcc{}

// Dcc is a component that contains an embedded DCC struct
type Dcc struct {
	*d2dcc.DCC
}

// New returns a Dcc component. By default, it contains a nil instance.
func (*Dcc) New() akara.Component {
	return &Dcc{}
}

// DccFactory is a wrapper for the generic component factory that returns Dcc component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Dcc.
type DccFactory struct {
	*akara.ComponentFactory
}

// Add adds a Dcc component to the given entity and returns it
func (m *DccFactory) Add(id akara.EID) *Dcc {
	return m.ComponentFactory.Add(id).(*Dcc)
}

// Get returns the Dcc component for the given entity, and a bool for whether or not it exists
func (m *DccFactory) Get(id akara.EID) (*Dcc, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Dcc), found
}
