package d2components

import (
	"github.com/gravestench/akara"
)

// static check that MainViewport implements Component
var _ akara.Component = &MainViewport{}

// MainViewport is used to flag viewports as the main viewport of a scene
type MainViewport struct{}

// New returns a new MainViewport instance. It is always a nil instance.
func (*MainViewport) New() akara.Component {
	return (*MainViewport)(nil)
}

// MainViewportFactory is a wrapper for the generic component factory that returns MainViewport component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a MainViewport.
type MainViewportFactory struct {
	*akara.ComponentFactory
}

// Add adds a MainViewport component to the given entity and returns it
func (m *MainViewportFactory) Add(id akara.EID) *MainViewport {
	return m.ComponentFactory.Add(id).(*MainViewport)
}

// Get returns the MainViewport component for the given entity, and a bool for whether or not it exists
func (m *MainViewportFactory) Get(id akara.EID) (*MainViewport, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*MainViewport), found
}
