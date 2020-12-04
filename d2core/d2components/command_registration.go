//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that CommandRegistration implements Component
var _ akara.Component = &CommandRegistration{}

// CommandRegistration is a flag component that is used to denote a "dirty" state
type CommandRegistration struct {
	Enabled     bool
	Name        string
	Description string
	Callback    interface{}
}

// New creates a new CommandRegistration. By default, IsCommandRegistration is false.
func (*CommandRegistration) New() akara.Component {
	return &CommandRegistration{
		Enabled: true,
	}
}

// CommandRegistrationFactory is a wrapper for the generic component factory that returns CommandRegistration component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a CommandRegistration.
type CommandRegistrationFactory struct {
	CommandRegistration *akara.ComponentFactory
}

// AddCommandRegistration adds a CommandRegistration component to the given entity and returns it
func (m *CommandRegistrationFactory) AddCommandRegistration(id akara.EID) *CommandRegistration {
	return m.CommandRegistration.Add(id).(*CommandRegistration)
}

// GetCommandRegistration returns the CommandRegistration component for the given entity, and a bool for whether or not it exists
func (m *CommandRegistrationFactory) GetCommandRegistration(id akara.EID) (*CommandRegistration, bool) {
	component, found := m.CommandRegistration.Get(id)
	if !found {
		return nil, found
	}

	return component.(*CommandRegistration), found
}
