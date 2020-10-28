package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/gravestench/akara"
)

// static check that SurfaceComponent implements Component
var _ akara.Component = &SurfaceComponent{}

// static check that SurfaceMap implements ComponentMap
var _ akara.ComponentMap = &SurfaceMap{}

// SurfaceComponent is a component that contains an embedded surface interface
type SurfaceComponent struct {
	d2interface.Surface
}

// ID returns a unique identifier for the component type
func (*SurfaceComponent) ID() akara.ComponentID {
	return SurfaceCID
}

// NewMap returns a new component map for the component type
func (*SurfaceComponent) NewMap() akara.ComponentMap {
	return NewSurfaceMap()
}

// Surface is a convenient reference to be used as a component identifier
var Surface = (*SurfaceComponent)(nil) // nolint:gochecknoglobals // global by design

// NewSurfaceMap creates a new map of entity ID's to Surface
func NewSurfaceMap() *SurfaceMap {
	cm := &SurfaceMap{
		components: make(map[akara.EID]*SurfaceComponent),
	}

	return cm
}

// SurfaceMap is a map of entity ID's to Surface
type SurfaceMap struct {
	world      *akara.World
	components map[akara.EID]*SurfaceComponent
}

// Init initializes the component map with the given world
func (cm *SurfaceMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*SurfaceMap) ID() akara.ComponentID {
	return SurfaceCID
}

// NewMap returns a new component map for the component type
func (*SurfaceMap) NewMap() akara.ComponentMap {
	return NewSurfaceMap()
}

// Add a new SurfaceComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *SurfaceMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &SurfaceComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddSurface adds a new SurfaceComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *SurfaceComponent instead of an akara.Component
func (cm *SurfaceMap) AddSurface(id akara.EID) *SurfaceComponent {
	return cm.Add(id).(*SurfaceComponent)
}

// Get returns the component associated with the given entity id
func (cm *SurfaceMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetSurface returns the SurfaceComponent associated with the given entity id
func (cm *SurfaceMap) GetSurface(id akara.EID) (*SurfaceComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *SurfaceMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
