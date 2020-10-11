package d2components

import (
	"github.com/gravestench/ecs"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// static check that FileSourceComponent implements Component
var _ ecs.Component = &FileSourceComponent{}

// static check that FileSourceMap implements ComponentMap
var _ ecs.ComponentMap = &FileSourceMap{}

// AbstractSource is the abstract representation of what a file source is
type AbstractSource interface {
	Open(path *FilePathComponent) (d2interface.DataStream, error)
}

// FileSourceComponent is a component that contains a FileSourceComponent instance
type FileSourceComponent struct {
	AbstractSource
}

// ID returns a unique identifier for the component type
func (*FileSourceComponent) ID() ecs.ComponentID {
	return FileSourceCID
}

// NewMap returns a new component map the component type
func (*FileSourceComponent) NewMap() ecs.ComponentMap {
	return NewFileSourceMap()
}

// FileSource is a convenient reference to be used as a component identifier
var FileSource = (*FileSourceComponent)(nil) // nolint:gochecknoglobals // global by design

// NewFileSourceMap creates a new map of entity ID's to FileSourceComponent components
func NewFileSourceMap() *FileSourceMap {
	cm := &FileSourceMap{
		components: make(map[ecs.EID]*FileSourceComponent),
	}

	return cm
}

// FileSourceMap is a map of entity ID's to FileSourceComponent type components
type FileSourceMap struct {
	world      *ecs.World
	components map[ecs.EID]*FileSourceComponent
}

// Init initializes the component map with the given world
func (cm *FileSourceMap) Init(world *ecs.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*FileSourceMap) ID() ecs.ComponentID {
	return FileSourceCID
}

// NewMap returns a new component map the component type
func (*FileSourceMap) NewMap() ecs.ComponentMap {
	return NewFileSourceMap()
}

// Add a new FileSourceComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *FileSourceMap) Add(id ecs.EID) ecs.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &FileSourceComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddFileSource adds a new FileSourceComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *FileSourceComponent instead of an ecs.Component
func (cm *FileSourceMap) AddFileSource(id ecs.EID) *FileSourceComponent {
	return cm.Add(id).(*FileSourceComponent)
}

// Get returns the component associated with the given entity id
func (cm *FileSourceMap) Get(id ecs.EID) (ecs.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetFileSource returns the FileSourceComponent type component associated with the given entity id
func (cm *FileSourceMap) GetFileSource(id ecs.EID) (*FileSourceComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetFileSources returns all FileSourceComponent components
func (cm *FileSourceMap) GetFileSources() []*FileSourceComponent {
	result := make([]*FileSourceComponent, 0)

	for _, src := range cm.components {
		result = append(result, src)
	}

	return result
}

// Remove a component for the given entity id, return the component.
func (cm *FileSourceMap) Remove(id ecs.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
