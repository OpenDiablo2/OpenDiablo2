//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2checkbox"
)

// static check that Checkbox implements Component
var _ akara.Component = &Checkbox{}

// Checkbox represents a UI checkbox. It contains an embedded *d2checkbox.Checkbox
type Checkbox struct {
	*d2checkbox.Checkbox
}

// New returns a Checkbox component. This contains an embedded *d2checkbox.Checkbox
func (*Checkbox) New() akara.Component {
	return &Checkbox{
		Checkbox: d2checkbox.New(),
	}
}

// CheckboxFactory is a wrapper for the generic component factory that returns Checkbox component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Checkbox.
type CheckboxFactory struct {
	*akara.ComponentFactory
}

// Add adds a Checkbox component to the given entity and returns it
func (m *CheckboxFactory) Add(id akara.EID) *Checkbox {
	return m.ComponentFactory.Add(id).(*Checkbox)
}

// Get returns the Button component for the given entity, and a bool for whether or not it exists
func (m *CheckboxFactory) Get(id akara.EID) (*Checkbox, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Checkbox), found
}
