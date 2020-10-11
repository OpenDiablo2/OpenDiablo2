package d2components

import (
	"github.com/gravestench/ecs"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// static check that FileHandleComponent implements Component
var _ ecs.Component = &FileHandleComponent{}

// static check that FileHandleMap implements ComponentMap
var _ ecs.ComponentMap = &FileHandleMap{}

// FileHandleComponent is a component that contains a data stream
type FileHandleComponent struct {
	Data d2interface.DataStream
}

// ID returns a unique identifier for the component type
func (*FileHandleComponent) ID() ecs.ComponentID {
	return FileHandleCID
}

// NewMap returns a new component map the component type
func (*FileHandleComponent) NewMap() ecs.ComponentMap {
	return NewFileHandleMap()
}

// FileHandle is a convenient reference to be used as a component identifier
var FileHandle = (*FileHandleComponent)(nil) // nolint:gochecknoglobals // global by design

// NewFileHandleMap creates a new map of entity ID's to FileHandleComponent components
func NewFileHandleMap() *FileHandleMap {
	cm := &FileHandleMap{
		components: make(map[ecs.EID]*FileHandleComponent),
	}

	return cm
}

// FileHandleMap is a map of entity ID's to FileHandleComponent components
type FileHandleMap struct {
	world      *ecs.World
	components map[ecs.EID]*FileHandleComponent
}

// Init initializes the component map with the given world
func (cm *FileHandleMap) Init(world *ecs.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*FileHandleMap) ID() ecs.ComponentID {
	return FileHandleCID
}

// NewMap returns a new component map the component type
func (*FileHandleMap) NewMap() ecs.ComponentMap {
	return NewFileHandleMap()
}

// Add a new FileHandleComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *FileHandleMap) Add(id ecs.EID) ecs.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &FileHandleComponent{Data: nil}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddFileHandle adds a new FileHandleComponent for the given entity id and returns it.
// If the entity already has a FileHandleComponent, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *FileHandleComponent instead of an ecs.Component
func (cm *FileHandleMap) AddFileHandle(id ecs.EID) *FileHandleComponent {
	return cm.Add(id).(*FileHandleComponent)
}

// Get returns the component associated with the given entity id
func (cm *FileHandleMap) Get(id ecs.EID) (ecs.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetFileHandle returns the FileHandleComponent component associated with the given entity id
func (cm *FileHandleMap) GetFileHandle(id ecs.EID) (*FileHandleComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *FileHandleMap) Remove(id ecs.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
