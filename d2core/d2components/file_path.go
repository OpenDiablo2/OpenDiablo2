package d2components

import (
	"github.com/gravestench/akara"
)

// static check that FilePathComponent implements Component
var _ akara.Component = &FilePathComponent{}

// static check that FilePathMap implements ComponentMap
var _ akara.ComponentMap = &FilePathMap{}

// FilePathComponent is a component that contains a file Path string
type FilePathComponent struct {
	Path string
}

// ID returns a unique identifier for the component type
func (*FilePathComponent) ID() akara.ComponentID {
	return FilePathCID
}

// NewMap returns a new component map for the component type
func (*FilePathComponent) NewMap() akara.ComponentMap {
	return NewFilePathMap()
}

// FilePath is a convenient reference to be used as a component identifier
var FilePath = (*FilePathComponent)(nil) // nolint:gochecknoglobals // global by design

// NewFilePathMap creates a new map of entity ID's to FilePath
func NewFilePathMap() *FilePathMap {
	cm := &FilePathMap{
		components: make(map[akara.EID]*FilePathComponent),
	}

	return cm
}

// FilePathMap is a map of entity ID's to FilePath
type FilePathMap struct {
	world      *akara.World
	components map[akara.EID]*FilePathComponent
}

// Init initializes the component map with the given world
func (cm *FilePathMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*FilePathMap) ID() akara.ComponentID {
	return FilePathCID
}

// NewMap returns a new component map for the component type
func (*FilePathMap) NewMap() akara.ComponentMap {
	return NewFilePathMap()
}

// Add a new FilePathComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *FilePathMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &FilePathComponent{Path: ""}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddFilePath adds a new FilePathComponent for the given entity id and returns it.
// If the entity already has a FilePathComponent, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *FilePathComponent instead of an akara.Component
func (cm *FilePathMap) AddFilePath(id akara.EID) *FilePathComponent {
	return cm.Add(id).(*FilePathComponent)
}

// Get returns the component associated with the given entity id
func (cm *FilePathMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetFilePath returns the FilePathComponent associated with the given entity id
func (cm *FilePathMap) GetFilePath(id akara.EID) (*FilePathComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *FilePathMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
