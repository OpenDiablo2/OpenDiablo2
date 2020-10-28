package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/gravestench/akara"
)

// static check that StaticPositionComponent implements Component
var _ akara.Component = &StaticPositionComponent{}

// static check that StaticPositionMap implements ComponentMap
var _ akara.ComponentMap = &StaticPositionMap{}

// StaticPositionComponent is a component that contains an embedded position struct.
// StaticPosition is used for positions that do not change with camera position (like ui overlay)
type StaticPositionComponent struct {
	*d2vector.Position
}

// ID returns a unique identifier for the component type
func (*StaticPositionComponent) ID() akara.ComponentID {
	return StaticPositionCID
}

// NewMap returns a new component map for the component type
func (*StaticPositionComponent) NewMap() akara.ComponentMap {
	return NewStaticPositionMap()
}

// StaticPosition is a convenient reference to be used as a component identifier
var StaticPosition = (*StaticPositionComponent)(nil) // nolint:gochecknoglobals // global by design

// NewStaticPositionMap creates a new map of entity ID's to StaticPosition
func NewStaticPositionMap() *StaticPositionMap {
	cm := &StaticPositionMap{
		components: make(map[akara.EID]*StaticPositionComponent),
	}

	return cm
}

// StaticPositionMap is a map of entity ID's to StaticPosition
type StaticPositionMap struct {
	world      *akara.World
	components map[akara.EID]*StaticPositionComponent
}

// Init initializes the component map with the given world
func (cm *StaticPositionMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*StaticPositionMap) ID() akara.ComponentID {
	return StaticPositionCID
}

// NewMap returns a new component map for the component type
func (*StaticPositionMap) NewMap() akara.ComponentMap {
	return NewStaticPositionMap()
}

// Add a new StaticPositionComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *StaticPositionMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	position := d2vector.NewPosition(0, 0)
	cm.components[id] = &StaticPositionComponent{Position: &position}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddStaticPosition adds a new StaticPositionComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *StaticPositionComponent instead of an akara.Component
func (cm *StaticPositionMap) AddStaticPosition(id akara.EID) *StaticPositionComponent {
	return cm.Add(id).(*StaticPositionComponent)
}

// Get returns the component associated with the given entity id
func (cm *StaticPositionMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetStaticPosition returns the StaticPositionComponent associated with the given entity id
func (cm *StaticPositionMap) GetStaticPosition(id akara.EID) (*StaticPositionComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *StaticPositionMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
