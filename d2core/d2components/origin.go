package d2components

import (
	"github.com/gravestench/akara"
)

// static check that OriginComponent implements Component
var _ akara.Component = &OriginComponent{}

// static check that OriginMap implements ComponentMap
var _ akara.ComponentMap = &OriginMap{}

// OriginComponent is a component that describes the origin point of an entity.
// The values are normalized to the display width/height.
// For example, origin (0,0) is top-left corner, (0.5, 0.5) is center
type OriginComponent struct {
	X, Y float64 // normalized
}

// ID returns a unique identifier for the component type
func (*OriginComponent) ID() akara.ComponentID {
	return OriginCID
}

// NewMap returns a new component map for the component type
func (*OriginComponent) NewMap() akara.ComponentMap {
	return NewOriginMap()
}

// Origin is a convenient reference to be used as a component identifier
var Origin = (*OriginComponent)(nil) // nolint:gochecknoglobals // global by design

// NewOriginMap creates a new map of entity ID's to Origin
func NewOriginMap() *OriginMap {
	cm := &OriginMap{
		components: make(map[akara.EID]*OriginComponent),
	}

	return cm
}

// OriginMap is a map of entity ID's to Origin
type OriginMap struct {
	world      *akara.World
	components map[akara.EID]*OriginComponent
}

// Init initializes the component map with the given world
func (cm *OriginMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*OriginMap) ID() akara.ComponentID {
	return OriginCID
}

// NewMap returns a new component map for the component type
func (*OriginMap) NewMap() akara.ComponentMap {
	return NewOriginMap()
}

// Add a new OriginComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *OriginMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &OriginComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddOrigin adds a new OriginComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *OriginComponent instead of an akara.Component
func (cm *OriginMap) AddOrigin(id akara.EID) *OriginComponent {
	return cm.Add(id).(*OriginComponent)
}

// Get returns the component associated with the given entity id
func (cm *OriginMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetOrigin returns the OriginComponent associated with the given entity id
func (cm *OriginMap) GetOrigin(id akara.EID) (*OriginComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *OriginMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
