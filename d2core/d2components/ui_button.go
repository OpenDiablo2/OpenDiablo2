//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2button"
)

// static check that Button implements Component
var _ akara.Component = &Button{}

// Button represents a ui label. It contains an embedded *d2button.Button
type Button struct {
	*d2button.Button
	States struct {
		// id's of segmented sprites
		Normal, Pressed, Toggled, PressedToggled, Disabled akara.EID
	}
}

// New returns a Button component. This contains an embedded *d2button.Button.
func (*Button) New() akara.Component {
	return &Button{
		Button: d2button.New(),
	}
}

// ButtonFactory is a wrapper for the generic component factory that returns Button component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Button.
type ButtonFactory struct {
	*akara.ComponentFactory
}

// Add adds a Button component to the given entity and returns it
func (m *ButtonFactory) Add(id akara.EID) *Button {
	return m.ComponentFactory.Add(id).(*Button)
}

// Get returns the Button component for the given entity, and a bool for whether or not it exists
func (m *ButtonFactory) Get(id akara.EID) (*Button, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Button), found
}
