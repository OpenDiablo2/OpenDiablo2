//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2animdata"
)

// static check that AnimationData implements Component
var _ akara.Component = &AnimationData{}

// AnimationData is a component that contains an embedded AnimationData struct
type AnimationData struct {
	*d2animdata.AnimationData
}

// New returns an AnimationData component. By default, it contains a nil instance.
func (*AnimationData) New() akara.Component {
	return &AnimationData{}
}

// AnimationDataFactory is a wrapper for the generic component factory that returns AnimationData component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a AnimationData.
type AnimationDataFactory struct {
	*akara.ComponentFactory
}

// Add adds a AnimationData component to the given entity and returns it
func (m *AnimationDataFactory) Add(id akara.EID) *AnimationData {
	return m.ComponentFactory.Add(id).(*AnimationData)
}

// Get returns the AnimationData component for the given entity, and a bool for whether or not it exists
func (m *AnimationDataFactory) Get(id akara.EID) (*AnimationData, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*AnimationData), found
}
