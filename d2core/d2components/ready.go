//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that Ready implements Component
var _ akara.Component = &Ready{}

// Ready is used to signify when a UI component is ready to be used.
// (files are loaded, surfaces rendered)
type Ready struct {}

// New returns a Ready component. This component is an empty tag component.
func (*Ready) New() akara.Component {
	return &Ready{}
}

// ReadyFactory is a wrapper for the generic component factory that returns Ready component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Ready.
type ReadyFactory struct {
	*akara.ComponentFactory
}

// Add adds a Ready component to the given entity and returns it
func (m *ReadyFactory) Add(id akara.EID) *Ready {
	return m.ComponentFactory.Add(id).(*Ready)
}

// Get returns the Ready component for the given entity, and a bool for whether or not it exists
func (m *ReadyFactory) Get(id akara.EID) (*Ready, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Ready), found
}
