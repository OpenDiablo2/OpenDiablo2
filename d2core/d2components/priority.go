//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that Priority implements Component
var _ akara.Component = new(Priority)

// Priority is a component that is used to add a priority value.
// This can generally be used for sorting entities when order matters.
type Priority struct {
	Priority int
}

// New returns a new Priority instance. The default is 0.
func (Priority) New() akara.Component {
	return &Priority{}
}

// PriorityFactory is a wrapper for the generic component factory that returns Priority component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Priority.
type PriorityFactory struct {
	*akara.ComponentFactory
}

// Add adds a Priority component to the given entity and returns it
func (m *PriorityFactory) Add(id akara.EID) *Priority {
	return m.ComponentFactory.Add(id).(*Priority)
}

// Get returns the Priority component for the given entity, and a bool for whether or not it exists
func (m *PriorityFactory) Get(id akara.EID) (*Priority, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Priority), found
}
