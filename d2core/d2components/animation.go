//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// static check that AnimationComponent implements Component
var _ akara.Component = &AnimationComponent{}

// static check that AnimationMap implements ComponentMap
var _ akara.ComponentMap = &AnimationMap{}

// AnimationComponent is a component that contains a width and height
type AnimationComponent struct {
	*akara.BaseComponent
	d2interface.Animation
}

// AnimationMap is a map of entity ID's to Animation
type AnimationMap struct {
	*akara.BaseComponentMap
}

// AddAnimation adds a new AnimationComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *AnimationComponent instead of an akara.Component
func (cm *AnimationMap) AddAnimation(id akara.EID) *AnimationComponent {
	return cm.Add(id).(*AnimationComponent)
}

// GetAnimation returns the AnimationComponent associated with the given entity id
func (cm *AnimationMap) GetAnimation(id akara.EID) (*AnimationComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*AnimationComponent), found
}

// Animation is a convenient reference to be used as a component identifier
var Animation = newAnimation() // nolint:gochecknoglobals // global by design

func newAnimation() akara.Component {
	return &AnimationComponent{
		BaseComponent: akara.NewBaseComponent(AnimationCID, newAnimation, newAnimationMap),
	}
}

func newAnimationMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(AnimationCID, newAnimation, newAnimationMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &AnimationMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
