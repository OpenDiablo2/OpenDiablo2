package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// static check that FileHandleComponent implements Component
var _ akara.Component = &FileHandleComponent{}

// static check that FileHandleMap implements ComponentMap
var _ akara.ComponentMap = &FileHandleMap{}

// FileHandleComponent is a component that contains a data stream
type FileHandleComponent struct {
	Data d2interface.DataStream
}

// ID returns a unique identifier for the component type
func (*FileHandleComponent) ID() akara.ComponentID {
	return FileHandleCID
}

// NewMap returns a new component map for the component type
func (*FileHandleComponent) NewMap() akara.ComponentMap {
	return NewFileHandleMap()
}

// FileHandle is a convenient reference to be used as a component identifier
var FileHandle = (*FileHandleComponent)(nil) // nolint:gochecknoglobals // global by design

// NewFileHandleMap creates a new map of entity ID's to FileHandleComponent components
func NewFileHandleMap() *FileHandleMap {
	cm := &FileHandleMap{
		components: make(map[akara.EID]*FileHandleComponent),
	}

	return cm
}

// FileHandleMap is a map of entity ID's to FileHandleComponent components
type FileHandleMap struct {
	world      *akara.World
	components map[akara.EID]*FileHandleComponent
}

// Init initializes the component map with the given world
func (cm *FileHandleMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*FileHandleMap) ID() akara.ComponentID {
	return FileHandleCID
}

// NewMap returns a new component map for the component type
func (*FileHandleMap) NewMap() akara.ComponentMap {
	return NewFileHandleMap()
}

// Add a new FileHandleComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *FileHandleMap) Add(id akara.EID) akara.Component {
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
// *FileHandleComponent instead of an akara.Component
func (cm *FileHandleMap) AddFileHandle(id akara.EID) *FileHandleComponent {
	return cm.Add(id).(*FileHandleComponent)
}

// Get returns the component associated with the given entity id
func (cm *FileHandleMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetFileHandle returns the FileHandleComponent component associated with the given entity id
func (cm *FileHandleMap) GetFileHandle(id akara.EID) (*FileHandleComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *FileHandleMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
