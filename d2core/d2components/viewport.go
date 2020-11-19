package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom"
	"github.com/gravestench/akara"
)

// static check that ViewportComponent implements Component
var _ akara.Component = &ViewportComponent{}

// static check that ViewportMap implements ComponentMap
var _ akara.ComponentMap = &ViewportMap{}

// ViewportComponent is a component that contains a file Type
type ViewportComponent struct {
	*d2geom.Rectangle
}

// ID returns a unique identifier for the component type
func (*ViewportComponent) ID() akara.ComponentID {
	return ViewportCID
}

// NewMap returns a new component map for the component type
func (*ViewportComponent) NewMap() akara.ComponentMap {
	return NewViewportMap()
}

// Viewport is a convenient reference to be used as a component identifier
var Viewport = (*ViewportComponent)(nil) // nolint:gochecknoglobals // global by design

// NewViewportMap creates a new map of entity ID's to Viewport
func NewViewportMap() *ViewportMap {
	cm := &ViewportMap{
		components: make(map[akara.EID]*ViewportComponent),
	}

	return cm
}

// ViewportMap is a map of entity ID's to Viewport
type ViewportMap struct {
	world      *akara.World
	components map[akara.EID]*ViewportComponent
}

// Init initializes the component map with the given world
func (cm *ViewportMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*ViewportMap) ID() akara.ComponentID {
	return ViewportCID
}

// NewMap returns a new component map for the component type
func (*ViewportMap) NewMap() akara.ComponentMap {
	return NewViewportMap()
}

// Add a new ViewportComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *ViewportMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &ViewportComponent{
		Rectangle: &d2geom.Rectangle{},
	}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddViewport adds a new ViewportComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *ViewportComponent instead of an akara.Component
func (cm *ViewportMap) AddViewport(id akara.EID) *ViewportComponent {
	return cm.Add(id).(*ViewportComponent)
}

// Get returns the component associated with the given entity id
func (cm *ViewportMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetViewport returns the ViewportComponent associated with the given entity id
func (cm *ViewportMap) GetViewport(id akara.EID) (*ViewportComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *ViewportMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
