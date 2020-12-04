//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
	"github.com/gravestench/akara"
)

// static check that Scale implements Component
var _ akara.Component = &Scale{}

// Scale represents an entities x,y axis scale as a vector
type Scale struct {
	*d2math.Vector3
}

// New creates a new Scale instance. By default, the scale is (1,1)
func (*Scale) New() akara.Component {
	return &Scale{
		Vector3: d2math.NewVector3(1, 1, 1),
	}
}

// ScaleFactory is a wrapper for the generic component factory that returns Scale component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Scale.
type ScaleFactory struct {
	Scale *akara.ComponentFactory
}

// AddScale adds a Scale component to the given entity and returns it
func (m *ScaleFactory) AddScale(id akara.EID) *Scale {
	return m.Scale.Add(id).(*Scale)
}

// GetScale returns the Scale component for the given entity, and a bool for whether or not it exists
func (m *ScaleFactory) GetScale(id akara.EID) (*Scale, bool) {
	component, found := m.Scale.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Scale), found
}
