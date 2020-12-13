package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2input"
)

// static check that Interactive implements Component
var _ akara.Component = &Interactive{}

func noop() bool {
	return false
}

// Interactive is used to flag file entities with a file type
type Interactive struct {
	Enabled bool
	*d2input.InputVector
	Callback func() (preventPropagation bool)
}

// New returns a Interactive component. By default, it contains a nil instance.
func (*Interactive) New() akara.Component {
	return &Interactive{
		Enabled:     true,
		InputVector: d2input.NewInputVector(),
		Callback:    noop,
	}
}

// InteractiveFactory is a wrapper for the generic component factory that returns Interactive component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Interactive.
type InteractiveFactory struct {
	*akara.ComponentFactory
}

// Add adds a Interactive component to the given entity and returns it
func (m *InteractiveFactory) Add(id akara.EID) *Interactive {
	return m.ComponentFactory.Add(id).(*Interactive)
}

// Get returns the Interactive component for the given entity, and a bool for whether or not it exists
func (m *InteractiveFactory) Get(id akara.EID) (*Interactive, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Interactive), found
}
