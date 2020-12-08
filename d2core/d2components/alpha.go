//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that Alpha implements Component
var _ akara.Component = &Alpha{}

// Alpha is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Alpha struct {
	Alpha float64
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Alpha) New() akara.Component {
	const defaultAlpha = 1.0

	return &Alpha{
		Alpha: defaultAlpha,
	}
}

// AlphaFactory is a wrapper for the generic component factory that returns Alpha component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Alpha.
type AlphaFactory struct {
	*akara.ComponentFactory
}

// Add adds a Alpha component to the given entity and returns it
func (m *AlphaFactory) Add(id akara.EID) *Alpha {
	return m.ComponentFactory.Add(id).(*Alpha)
}

// Get returns the Alpha component for the given entity, and a bool for whether or not it exists
func (m *AlphaFactory) Get(id akara.EID) (*Alpha, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Alpha), found
}
