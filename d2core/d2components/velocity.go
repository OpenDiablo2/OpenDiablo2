package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
)

// static check that VelocityComponent implements Component
var _ akara.Component = &VelocityComponent{}

// static check that VelocityMap implements ComponentMap
var _ akara.ComponentMap = &VelocityMap{}

// VelocityComponent stores the velocity as a vec2
type VelocityComponent struct {
	*d2vector.Vector
}

// ID returns a unique identifier for the component type
func (*VelocityComponent) ID() akara.ComponentID {
	return VelocityCID
}

// NewMap returns a new component map for the component type
func (*VelocityComponent) NewMap() akara.ComponentMap {
	return NewVelocityMap()
}

// Velocity is a convenient reference to be used as a component identifier
var Velocity = (*VelocityComponent)(nil) // nolint:gochecknoglobals // global by design

// NewVelocityMap creates a new map of entity ID's to velocity components
func NewVelocityMap() *VelocityMap {
	return &VelocityMap{
		components: make(map[akara.EID]*VelocityComponent),
	}
}

// VelocityMap is a map of entity ID's to velocity components
type VelocityMap struct {
	world      *akara.World
	components map[akara.EID]*VelocityComponent
}

// Init initializes the component map with the given world
func (cm *VelocityMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*VelocityMap) ID() akara.ComponentID {
	return VelocityCID
}

// NewMap returns a new component map for the component type
func (*VelocityMap) NewMap() akara.ComponentMap {
	return NewVelocityMap()
}

// Add a new VelocityComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *VelocityMap) Add(id akara.EID) akara.Component {
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
// *VelocityComponent instead of an akara.Component
func (cm *VelocityMap) AddVelocity(id akara.EID) *VelocityComponent {
	return cm.Add(id).(*VelocityComponent)
}

// Get returns the component associated with the given entity id
func (cm *VelocityMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetVelocity returns the velocity component associated with the given entity id.
// This is used to return a *VelocityComponent, as opposed to an akara.Component
func (cm *VelocityMap) GetVelocity(id akara.EID) (*VelocityComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *VelocityMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
