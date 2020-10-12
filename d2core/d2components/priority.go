package d2components

import (
	"github.com/gravestench/akara"
)

// static check that PriorityComponent implements Component
var _ akara.Component = &PriorityComponent{}

// static check that PriorityMap implements ComponentMap
var _ akara.ComponentMap = &PriorityMap{}

// PriorityComponent is a component that is used to add a priority value.
// This can generally be used for sorting entities when order matters.
type PriorityComponent struct {
	Priority int
}

// ID returns a unique identifier for the component type
func (*PriorityComponent) ID() akara.ComponentID {
	return PriorityCID
}

// NewMap returns a new component map for the component type
func (*PriorityComponent) NewMap() akara.ComponentMap {
	return NewPriorityMap()
}

// Priority is a convenient reference to be used as a component identifier
var Priority = (*PriorityComponent)(nil) // nolint:gochecknoglobals // global by design

// NewPriorityMap creates a new map of entity ID's to Priority
func NewPriorityMap() *PriorityMap {
	cm := &PriorityMap{
		components: make(map[akara.EID]*PriorityComponent),
	}

	return cm
}

// PriorityMap is a map of entity ID's to Priority
type PriorityMap struct {
	world      *akara.World
	components map[akara.EID]*PriorityComponent
}

// Init initializes the component map with the given world
func (cm *PriorityMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*PriorityMap) ID() akara.ComponentID {
	return PriorityCID
}

// NewMap returns a new component map for the component type
func (*PriorityMap) NewMap() akara.ComponentMap {
	return NewPriorityMap()
}

// Add a new PriorityComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *PriorityMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &PriorityComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddPriority adds a new PriorityComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *PriorityComponent instead of an akara.Component
func (cm *PriorityMap) AddPriority(id akara.EID) *PriorityComponent {
	return cm.Add(id).(*PriorityComponent)
}

// Get returns the component associated with the given entity id
func (cm *PriorityMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetPriority returns the PriorityComponent associated with the given entity id
func (cm *PriorityMap) GetPriority(id akara.EID) (*PriorityComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *PriorityMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
