package d2components

import (
	"github.com/gravestench/akara"
)

// static check that MainViewportComponent implements Component
var _ akara.Component = &MainViewportComponent{}

// static check that MainViewportMap implements ComponentMap
var _ akara.ComponentMap = &MainViewportMap{}

// MainViewportComponent is a component that is used to tag the main viewport of a scene
type MainViewportComponent struct {}

// ID returns a unique identifier for the component type
func (*MainViewportComponent) ID() akara.ComponentID {
	return MainViewportCID
}

// NewMap returns a new component map for the component type
func (*MainViewportComponent) NewMap() akara.ComponentMap {
	return NewMainViewportMap()
}

// MainViewport is a convenient reference to be used as a component identifier
var MainViewport = (*MainViewportComponent)(nil) // nolint:gochecknoglobals // global by design

// NewMainViewportMap creates a new map of entity ID's to MainViewport
func NewMainViewportMap() *MainViewportMap {
	cm := &MainViewportMap{
		components: make(map[akara.EID]*MainViewportComponent),
	}

	return cm
}

// MainViewportMap is a map of entity ID's to MainViewport
type MainViewportMap struct {
	world      *akara.World
	components map[akara.EID]*MainViewportComponent
}

// Init initializes the component map with the given world
func (cm *MainViewportMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*MainViewportMap) ID() akara.ComponentID {
	return MainViewportCID
}

// NewMap returns a new component map for the component type
func (*MainViewportMap) NewMap() akara.ComponentMap {
	return NewMainViewportMap()
}

// Add a new MainViewportComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *MainViewportMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &MainViewportComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddMainViewport adds a new MainViewportComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *MainViewportComponent instead of an akara.Component
func (cm *MainViewportMap) AddMainViewport(id akara.EID) *MainViewportComponent {
	return cm.Add(id).(*MainViewportComponent)
}

// Get returns the component associated with the given entity id
func (cm *MainViewportMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetMainViewport returns the MainViewportComponent associated with the given entity id
func (cm *MainViewportMap) GetMainViewport(id akara.EID) (*MainViewportComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *MainViewportMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
