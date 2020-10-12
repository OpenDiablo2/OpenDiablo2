package d2components

import (
	"github.com/gravestench/akara"
)

// static check that DirtyComponent implements Component
var _ akara.Component = &DirtyComponent{}

// static check that DirtyMap implements ComponentMap
var _ akara.ComponentMap = &DirtyMap{}

// DirtyComponent is a component that is used to signify when something is "dirty".
// What this means depends on the context, but it is typically used as a condition to
// perform some processing.
type DirtyComponent struct {
	IsDirty bool
}

// ID returns a unique identifier for the component type
func (*DirtyComponent) ID() akara.ComponentID {
	return DirtyCID
}

// NewMap returns a new component map for the component type
func (*DirtyComponent) NewMap() akara.ComponentMap {
	return NewDirtyMap()
}

// Dirty is a convenient reference to be used as a component identifier
var Dirty = (*DirtyComponent)(nil) // nolint:gochecknoglobals // global by design

// NewDirtyMap creates a new map of entity ID's to Dirty
func NewDirtyMap() *DirtyMap {
	cm := &DirtyMap{
		components: make(map[akara.EID]*DirtyComponent),
	}

	return cm
}

// DirtyMap is a map of entity ID's to Dirty
type DirtyMap struct {
	world      *akara.World
	components map[akara.EID]*DirtyComponent
}

// Init initializes the component map with the given world
func (cm *DirtyMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*DirtyMap) ID() akara.ComponentID {
	return DirtyCID
}

// NewMap returns a new component map for the component type
func (*DirtyMap) NewMap() akara.ComponentMap {
	return NewDirtyMap()
}

// Add a new DirtyComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *DirtyMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &DirtyComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddDirty adds a new DirtyComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *DirtyComponent instead of an akara.Component
func (cm *DirtyMap) AddDirty(id akara.EID) *DirtyComponent {
	return cm.Add(id).(*DirtyComponent)
}

// Get returns the component associated with the given entity id
func (cm *DirtyMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetDirty returns the DirtyComponent associated with the given entity id
func (cm *DirtyMap) GetDirty(id akara.EID) (*DirtyComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *DirtyMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
