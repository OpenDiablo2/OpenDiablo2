package d2components

import (
	"github.com/gravestench/akara"
)

// static check that RenderCullComponent implements Component
var _ akara.Component = &RenderCullComponent{}

// static check that RenderCullMap implements ComponentMap
var _ akara.ComponentMap = &RenderCullMap{}

// RenderCullComponent is an empty component used for
// tagging entities that should not be rendered
type RenderCullComponent struct{}

// ID returns a unique identifier for the component type
func (*RenderCullComponent) ID() akara.ComponentID {
	return RenderCullCID
}

// NewMap returns a new component map for the component type
func (*RenderCullComponent) NewMap() akara.ComponentMap {
	return NewRenderCullMap()
}

// RenderCull is a convenient reference to be used as a component identifier
var RenderCull = (*RenderCullComponent)(nil) // nolint:gochecknoglobals // global by design

// NewRenderCullMap creates a new map of entity ID's to RenderCull
func NewRenderCullMap() *RenderCullMap {
	cm := &RenderCullMap{
		components: make(map[akara.EID]*RenderCullComponent),
	}

	return cm
}

// RenderCullMap is a map of entity ID's to RenderCull
type RenderCullMap struct {
	world      *akara.World
	components map[akara.EID]*RenderCullComponent
}

// Init initializes the component map with the given world
func (cm *RenderCullMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*RenderCullMap) ID() akara.ComponentID {
	return RenderCullCID
}

// NewMap returns a new component map for the component type
func (*RenderCullMap) NewMap() akara.ComponentMap {
	return NewRenderCullMap()
}

// Add a new RenderCullComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *RenderCullMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &RenderCullComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddRenderCull adds a new RenderCullComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *RenderCullComponent instead of an akara.Component
func (cm *RenderCullMap) AddRenderCull(id akara.EID) *RenderCullComponent {
	return cm.Add(id).(*RenderCullComponent)
}

// Get returns the component associated with the given entity id
func (cm *RenderCullMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetRenderCull returns the RenderCullComponent associated with the given entity id
func (cm *RenderCullMap) GetRenderCull(id akara.EID) (*RenderCullComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *RenderCullMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
