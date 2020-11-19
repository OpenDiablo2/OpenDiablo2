package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/gravestench/akara"
)

// static check that ScaleComponent implements Component
var _ akara.Component = &ScaleComponent{}

// static check that ScaleMap implements ComponentMap
var _ akara.ComponentMap = &ScaleMap{}

// ScaleComponent is a component that contains scale for x and y axis
type ScaleComponent struct {
	*d2vector.Vector
}

// ID returns a unique identifier for the component type
func (*ScaleComponent) ID() akara.ComponentID {
	return ScaleCID
}

// NewMap returns a new component map for the component type
func (*ScaleComponent) NewMap() akara.ComponentMap {
	return NewScaleMap()
}

// Scale is a convenient reference to be used as a component identifier
var Scale = (*ScaleComponent)(nil) // nolint:gochecknoglobals // global by design

// NewScaleMap creates a new map of entity ID's to Scale
func NewScaleMap() *ScaleMap {
	cm := &ScaleMap{
		components: make(map[akara.EID]*ScaleComponent),
	}

	return cm
}

// ScaleMap is a map of entity ID's to Scale
type ScaleMap struct {
	world      *akara.World
	components map[akara.EID]*ScaleComponent
}

// Init initializes the component map with the given world
func (cm *ScaleMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*ScaleMap) ID() akara.ComponentID {
	return ScaleCID
}

// NewMap returns a new component map for the component type
func (*ScaleMap) NewMap() akara.ComponentMap {
	return NewScaleMap()
}

// Add a new ScaleComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *ScaleMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &ScaleComponent{
		Vector: d2vector.NewVector(1, 1),
	}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddScale adds a new ScaleComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *ScaleComponent instead of an akara.Component
func (cm *ScaleMap) AddScale(id akara.EID) *ScaleComponent {
	return cm.Add(id).(*ScaleComponent)
}

// Get returns the component associated with the given entity id
func (cm *ScaleMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetScale returns the ScaleComponent associated with the given entity id
func (cm *ScaleMap) GetScale(id akara.EID) (*ScaleComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *ScaleMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
