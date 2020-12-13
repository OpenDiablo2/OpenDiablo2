//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
)

// static check that Dc6 implements Component
var _ akara.Component = &Dc6{}

// Dc6 is a component that contains an embedded DC6 struct
type Dc6 struct {
	*d2dc6.DC6
}

// New returns a Dc6 component. By default, it contains a nil instance.
func (*Dc6) New() akara.Component {
	return &Dc6{}
}

// Dc6Factory is a wrapper for the generic component factory that returns Dc6 component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Dc6.
type Dc6Factory struct {
	*akara.ComponentFactory
}

// Add adds a Dc6 component to the given entity and returns it
func (m *Dc6Factory) Add(id akara.EID) *Dc6 {
	return m.ComponentFactory.Add(id).(*Dc6)
}

// Get returns the Dc6 component for the given entity, and a bool for whether or not it exists
func (m *Dc6Factory) Get(id akara.EID) (*Dc6, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Dc6), found
}
