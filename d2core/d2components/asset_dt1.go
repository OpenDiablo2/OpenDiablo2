//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"
)

// static check that Dt1 implements Component
var _ akara.Component = &Dt1{}

// Dt1 is a component that contains an embedded DT1 struct
type Dt1 struct {
	*d2dt1.DT1
}

// New returns a Dt1 component. By default, it contains a nil instance.
func (*Dt1) New() akara.Component {
	return &Dt1{}
}

// Dt1Factory is a wrapper for the generic component factory that returns Dt1 component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Dt1.
type Dt1Factory struct {
	*akara.ComponentFactory
}

// Add adds a Dt1 component to the given entity and returns it
func (m *Dt1Factory) Add(id akara.EID) *Dt1 {
	return m.ComponentFactory.Add(id).(*Dt1)
}

// Get returns the Dt1 component for the given entity, and a bool for whether or not it exists
func (m *Dt1Factory) Get(id akara.EID) (*Dt1, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Dt1), found
}
