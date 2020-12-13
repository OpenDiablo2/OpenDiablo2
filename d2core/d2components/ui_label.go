//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2label"
)

// static check that Label implements Component
var _ akara.Component = &Label{}

// Label represents a ui label. It contains an embedded *d2label.Label
type Label struct {
	*d2label.Label
}

// New returns a Label component. This component contains an embedded *d2label.Label
func (*Label) New() akara.Component {
	return &Label{
		d2label.New(),
	}
}

// LabelFactory is a wrapper for the generic component factory that returns Label component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Label.
type LabelFactory struct {
	*akara.ComponentFactory
}

// Add adds a Label component to the given entity and returns it
func (m *LabelFactory) Add(id akara.EID) *Label {
	return m.ComponentFactory.Add(id).(*Label)
}

// Get returns the Label component for the given entity, and a bool for whether or not it exists
func (m *LabelFactory) Get(id akara.EID) (*Label, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Label), found
}
