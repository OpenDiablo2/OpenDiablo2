//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that Dirty implements Component
var _ akara.Component = &Dirty{}

// Dirty is a flag component that is used to denote a "dirty" state
type Dirty struct {}

// New creates a new Dirty. By default, IsDirty is false.
func (*Dirty) New() akara.Component {
	return &Dirty{}
}

// DirtyFactory is a wrapper for the generic component factory that returns Dirty component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Dirty.
type DirtyFactory struct {
	*akara.ComponentFactory
}

// Add adds a Dirty component to the given entity and returns it
func (m *DirtyFactory) Add(id akara.EID) *Dirty {
	return m.ComponentFactory.Add(id).(*Dirty)
}

// Get returns the Dirty component for the given entity, and a bool for whether or not it exists
func (m *DirtyFactory) Get(id akara.EID) (*Dirty, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Dirty), found
}
