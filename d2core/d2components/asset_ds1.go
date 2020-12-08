//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
)

// static check that Ds1 implements Component
var _ akara.Component = &Ds1{}

// Ds1 is a component that contains an embedded DS1 struct
type Ds1 struct {
	*d2ds1.DS1
}

// New returns a Ds1 component. By default, it contains a nil instance.
func (*Ds1) New() akara.Component {
	return &Ds1{}
}

// Ds1Factory is a wrapper for the generic component factory that returns Ds1 component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Ds1.
type Ds1Factory struct {
	*akara.ComponentFactory
}

// Add adds a Ds1 component to the given entity and returns it
func (m *Ds1Factory) Add(id akara.EID) *Ds1 {
	return m.ComponentFactory.Add(id).(*Ds1)
}

// Get returns the Ds1 component for the given entity, and a bool for whether or not it exists
func (m *Ds1Factory) Get(id akara.EID) (*Ds1, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Ds1), found
}
