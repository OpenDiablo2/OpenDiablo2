//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that ViewportFilter implements Component
var _ akara.Component = &ViewportFilter{}

// ViewportFilter is a component that contains a bitset that denotes which viewport
// the entity will be rendered.
type ViewportFilter struct {
	*akara.BitSet
}

// New creates a new ViewportFilter.
// By default, the filter is set to only allow the main scene viewport.
func (*ViewportFilter) New() akara.Component {
	const mainViewport = 0

	return &ViewportFilter{
		BitSet: akara.NewBitSet(mainViewport),
	}
}

// ViewportFilterFactory is a wrapper for the generic component factory that returns ViewportFilter component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a ViewportFilter.
type ViewportFilterFactory struct {
	*akara.ComponentFactory
}

// Add adds a ViewportFilter component to the given entity and returns it
func (m *ViewportFilterFactory) Add(id akara.EID) *ViewportFilter {
	return m.ComponentFactory.Add(id).(*ViewportFilter)
}

// Get returns the ViewportFilter component for the given entity, and a bool for whether or not it exists
func (m *ViewportFilterFactory) Get(id akara.EID) (*ViewportFilter, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*ViewportFilter), found
}
