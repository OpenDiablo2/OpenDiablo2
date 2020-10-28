package d2components

import (
	"github.com/gravestench/akara"
)

// static check that SizeComponent implements Component
var _ akara.Component = &SizeComponent{}

// static check that SizeMap implements ComponentMap
var _ akara.ComponentMap = &SizeMap{}

// SizeComponent is a component that contains a width and height
type SizeComponent struct {
	Width, Height uint
}

// ID returns a unique identifier for the component type
func (*SizeComponent) ID() akara.ComponentID {
	return SizeCID
}

// NewMap returns a new component map for the component type
func (*SizeComponent) NewMap() akara.ComponentMap {
	return NewSizeMap()
}

// Size is a convenient reference to be used as a component identifier
var Size = (*SizeComponent)(nil) // nolint:gochecknoglobals // global by design

// NewSizeMap creates a new map of entity ID's to Size
func NewSizeMap() *SizeMap {
	cm := &SizeMap{
		components: make(map[akara.EID]*SizeComponent),
	}

	return cm
}

// SizeMap is a map of entity ID's to Size
type SizeMap struct {
	world      *akara.World
	components map[akara.EID]*SizeComponent
}

// Init initializes the component map with the given world
func (cm *SizeMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*SizeMap) ID() akara.ComponentID {
	return SizeCID
}

// NewMap returns a new component map for the component type
func (*SizeMap) NewMap() akara.ComponentMap {
	return NewSizeMap()
}

// Add a new SizeComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *SizeMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &SizeComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddSize adds a new SizeComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *SizeComponent instead of an akara.Component
func (cm *SizeMap) AddSize(id akara.EID) *SizeComponent {
	return cm.Add(id).(*SizeComponent)
}

// Get returns the component associated with the given entity id
func (cm *SizeMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetSize returns the SizeComponent associated with the given entity id
func (cm *SizeMap) GetSize(id akara.EID) (*SizeComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *SizeMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
