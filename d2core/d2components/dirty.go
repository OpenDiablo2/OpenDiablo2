//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that Dirty implements Component
var _ akara.Component = &Dirty{}

// Dirty is a flag component that is used to denote a "dirty" state
type Dirty struct {
	IsDirty bool
}

// New creates a new Dirty. By default, IsDirty is false.
func (*Dirty) New() akara.Component {
	return &Dirty{}
}

// DirtyFactory is a wrapper for the generic component factory that returns Dirty component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Dirty.
type DirtyFactory struct {
	Dirty *akara.ComponentFactory
}

// AddDirty adds a Dirty component to the given entity and returns it
func (m *DirtyFactory) AddDirty(id akara.EID) *Dirty {
	return m.Dirty.Add(id).(*Dirty)
}

// GetDirty returns the Dirty component for the given entity, and a bool for whether or not it exists
func (m *DirtyFactory) GetDirty(id akara.EID) (*Dirty, bool) {
	component, found := m.Dirty.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Dirty), found
}
