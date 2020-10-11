package d2components

import (
	"github.com/gravestench/ecs"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
)

// static check that VelocityComponent implements Component
var _ ecs.Component = &VelocityComponent{}

// static check that VelocityMap implements ComponentMap
var _ ecs.ComponentMap = &VelocityMap{}

// VelocityComponent stores the velocity as a vec2
type VelocityComponent struct {
	*d2vector.Vector
}

// ID returns a unique identifier for the component type
func (*VelocityComponent) ID() ecs.ComponentID {
	return VelocityCID
}

// NewMap returns a new component map the component type
func (*VelocityComponent) NewMap() ecs.ComponentMap {
	return NewVelocityMap()
}

// Velocity is a convenient reference to be used as a component identifier
var Velocity = (*VelocityComponent)(nil) // nolint:gochecknoglobals // global by design

// NewVelocityMap creates a new map of entity ID's to velocity components
func NewVelocityMap() *VelocityMap {
	return &VelocityMap{
		components: make(map[ecs.EID]*VelocityComponent),
	}
}

// VelocityMap is a map of entity ID's to velocity components
type VelocityMap struct {
	world      *ecs.World
	components map[ecs.EID]*VelocityComponent
}

// Init initializes the component map with the given world
func (cm *VelocityMap) Init(world *ecs.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*VelocityMap) ID() ecs.ComponentID {
	return VelocityCID
}

// NewMap returns a new component map the component type
func (*VelocityMap) NewMap() ecs.ComponentMap {
	return NewVelocityMap()
}

// Add a new VelocityComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *VelocityMap) Add(id ecs.EID) ecs.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	v := d2vector.NewVector(0, 0)
	com := &VelocityComponent{Vector: v}
	cm.components[id] = com

	cm.world.UpdateEntity(id)

	return com
}

// AddVelocity adds a new VelocityComponent for the given entity id and returns it.
// If the entity already has a component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *VelocityComponent instead of an ecs.Component
func (cm *VelocityMap) AddVelocity(id ecs.EID) *VelocityComponent {
	return cm.Add(id).(*VelocityComponent)
}

// Get returns the component associated with the given entity id
func (cm *VelocityMap) Get(id ecs.EID) (ecs.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetVelocity returns the velocity component associated with the given entity id.
// This is used to return a *VelocityComponent, as opposed to an ecs.Component
func (cm *VelocityMap) GetVelocity(id ecs.EID) (*VelocityComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *VelocityMap) Remove(id ecs.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
