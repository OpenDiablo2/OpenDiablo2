package d2components

import (
	"github.com/gravestench/ecs"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
)

// static check that PositionComponent implements Component
var _ ecs.Component = &PositionComponent{}

// static check that PositionMap implements ComponentMap
var _ ecs.ComponentMap = &PositionMap{}

// PositionComponent stores an x,y position
type PositionComponent struct {
	*d2vector.Position
}

// ID returns a unique identifier for the component type
func (*PositionComponent) ID() ecs.ComponentID {
	return PositionCID
}

// NewMap returns a new component map the component type
func (*PositionComponent) NewMap() ecs.ComponentMap {
	return NewPositionMap()
}

// Position is a convenient reference to be used as a component identifier
var Position = (*PositionComponent)(nil) // nolint:gochecknoglobals // global by design

// NewPositionMap creates a new map of entity ID's to position components
func NewPositionMap() *PositionMap {
	cm := &PositionMap{
		components: make(map[ecs.EID]*PositionComponent),
	}

	return cm
}

// PositionMap is a map of entity ID's to position components
type PositionMap struct {
	world      *ecs.World
	components map[ecs.EID]*PositionComponent
}

// Init initializes the component map with the given world
func (cm *PositionMap) Init(world *ecs.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*PositionMap) ID() ecs.ComponentID {
	return PositionCID
}

// NewMap returns a new component map the component type
func (*PositionMap) NewMap() ecs.ComponentMap {
	return NewPositionMap()
}

// Add a new PositionComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *PositionMap) Add(id ecs.EID) ecs.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	position := d2vector.NewPosition(0, 0)
	cm.components[id] = &PositionComponent{Position: &position}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddPosition adds a new PositionComponent for the given entity id and returns it.
// If the entity already has a position component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *PositionComponent instead of an ecs.Component
func (cm *PositionMap) AddPosition(id ecs.EID) *PositionComponent {
	return cm.Add(id).(*PositionComponent)
}

// Get returns the component associated with the given entity id
func (cm *PositionMap) Get(id ecs.EID) (ecs.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetPosition returns the position component associated with the given entity id
func (cm *PositionMap) GetPosition(id ecs.EID) (*PositionComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *PositionMap) Remove(id ecs.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
