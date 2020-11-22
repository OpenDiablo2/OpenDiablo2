package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/gravestench/akara"
)

// static check that AnimationComponent implements Component
var _ akara.Component = &AnimationComponent{}

// static check that AnimationMap implements ComponentMap
var _ akara.ComponentMap = &AnimationMap{}

// AnimationComponent is a component that contains a width and height
type AnimationComponent struct {
	d2interface.Animation
}

// ID returns a unique identifier for the component type
func (*AnimationComponent) ID() akara.ComponentID {
	return AnimationCID
}

// NewMap returns a new component map for the component type
func (*AnimationComponent) NewMap() akara.ComponentMap {
	return NewAnimationMap()
}

// Animation is a convenient reference to be used as a component identifier
var Animation = (*AnimationComponent)(nil) // nolint:gochecknoglobals // global by design

// NewAnimationMap creates a new map of entity ID's to Animation
func NewAnimationMap() *AnimationMap {
	cm := &AnimationMap{
		components: make(map[akara.EID]*AnimationComponent),
	}

	return cm
}

// AnimationMap is a map of entity ID's to Animation
type AnimationMap struct {
	world      *akara.World
	components map[akara.EID]*AnimationComponent
}

// Init initializes the component map with the given world
func (cm *AnimationMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*AnimationMap) ID() akara.ComponentID {
	return AnimationCID
}

// NewMap returns a new component map for the component type
func (*AnimationMap) NewMap() akara.ComponentMap {
	return NewAnimationMap()
}

// Add a new AnimationComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *AnimationMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &AnimationComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddAnimation adds a new AnimationComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *AnimationComponent instead of an akara.Component
func (cm *AnimationMap) AddAnimation(id akara.EID) *AnimationComponent {
	return cm.Add(id).(*AnimationComponent)
}

// Get returns the component associated with the given entity id
func (cm *AnimationMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetAnimation returns the AnimationComponent associated with the given entity id
func (cm *AnimationMap) GetAnimation(id akara.EID) (*AnimationComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *AnimationMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
