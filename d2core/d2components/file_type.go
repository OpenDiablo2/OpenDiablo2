package d2components

import (
	"github.com/gravestench/ecs"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// static check that FileTypeComponent implements Component
var _ ecs.Component = &FileTypeComponent{}

// static check that FileTypeMap implements ComponentMap
var _ ecs.ComponentMap = &FileTypeMap{}

// FileTypeComponent is a component that contains a file Type
type FileTypeComponent struct {
	Type d2enum.FileType
}

// ID returns a unique identifier for the component type
func (*FileTypeComponent) ID() ecs.ComponentID {
	return FileTypeCID
}

// NewMap returns a new component map the component type
func (*FileTypeComponent) NewMap() ecs.ComponentMap {
	return NewFileTypeMap()
}

// FileType is a convenient reference to be used as a component identifier
var FileType = (*FileTypeComponent)(nil) // nolint:gochecknoglobals // global by design

// NewFileTypeMap creates a new map of entity ID's to FileType
func NewFileTypeMap() *FileTypeMap {
	cm := &FileTypeMap{
		components: make(map[ecs.EID]*FileTypeComponent),
	}

	return cm
}

// FileTypeMap is a map of entity ID's to FileType
type FileTypeMap struct {
	world      *ecs.World
	components map[ecs.EID]*FileTypeComponent
}

// Init initializes the component map with the given world
func (cm *FileTypeMap) Init(world *ecs.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*FileTypeMap) ID() ecs.ComponentID {
	return FileTypeCID
}

// NewMap returns a new component map the component type
func (*FileTypeMap) NewMap() ecs.ComponentMap {
	return NewFileTypeMap()
}

// Add a new FileTypeComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *FileTypeMap) Add(id ecs.EID) ecs.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &FileTypeComponent{Type: d2enum.FileTypeUnknown}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddFileType adds a new FileTypeComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *FileTypeComponent instead of an ecs.Component
func (cm *FileTypeMap) AddFileType(id ecs.EID) *FileTypeComponent {
	return cm.Add(id).(*FileTypeComponent)
}

// Get returns the component associated with the given entity id
func (cm *FileTypeMap) Get(id ecs.EID) (ecs.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetFileType returns the FileTypeComponent associated with the given entity id
func (cm *FileTypeMap) GetFileType(id ecs.EID) (*FileTypeComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *FileTypeMap) Remove(id ecs.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
