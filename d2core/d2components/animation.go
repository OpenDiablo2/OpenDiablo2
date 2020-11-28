//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// static check that AnimationComponent implements Component
var _ akara.Component = &Animation{}

// Animation is a component that contains a width and height
type Animation struct {
	d2interface.Animation
}

// New returns an animation component. By default, it contains a nil instance of an animation.
func (*Animation) New() akara.Component {
	return &Animation{}
}

// AnimationFactory is a wrapper for the generic component factory that returns Animation component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Animation.
type AnimationFactory struct {
	Animation *akara.ComponentFactory
}

// AddAnimation adds a Animation component to the given entity and returns it
func (m *AnimationFactory) AddAnimation(id akara.EID) *Animation {
	return m.Animation.Add(id).(*Animation)
}

// GetAnimation returns the Animation component for the given entity, and a bool for whether or not it exists
func (m *AnimationFactory) GetAnimation(id akara.EID) (*Animation, bool) {
	component, found := m.Animation.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Animation), found
}
