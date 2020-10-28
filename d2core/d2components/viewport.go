package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/gravestench/akara"
)

// static check that ViewPortComponent implements Component
var _ akara.Component = &ViewPortComponent{}

// static check that ViewPortMap implements ComponentMap
var _ akara.ComponentMap = &ViewPortMap{}

// ViewPortComponent is a component that contains a file Type
type ViewPortComponent struct {
	d2interface.Surface
	*d2geom.Rectangle
}

// ID returns a unique identifier for the component type
func (*ViewPortComponent) ID() akara.ComponentID {
	return ViewportCID
}

// NewMap returns a new component map for the component type
func (*ViewPortComponent) NewMap() akara.ComponentMap {
	return NewViewPortMap()
}

// ViewPort is a convenient reference to be used as a component identifier
var ViewPort = (*ViewPortComponent)(nil) // nolint:gochecknoglobals // global by design

// NewViewPortMap creates a new map of entity ID's to ViewPort
func NewViewPortMap() *ViewPortMap {
	cm := &ViewPortMap{
		components: make(map[akara.EID]*ViewPortComponent),
	}

	return cm
}

// ViewPortMap is a map of entity ID's to ViewPort
type ViewPortMap struct {
	world      *akara.World
	components map[akara.EID]*ViewPortComponent
}

// Init initializes the component map with the given world
func (cm *ViewPortMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*ViewPortMap) ID() akara.ComponentID {
	return ViewportCID
}

// NewMap returns a new component map for the component type
func (*ViewPortMap) NewMap() akara.ComponentMap {
	return NewViewPortMap()
}

// Add a new ViewPortComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *ViewPortMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &ViewPortComponent{Rectangle: &d2geom.Rectangle{}}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddViewPort adds a new ViewPortComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *ViewPortComponent instead of an akara.Component
func (cm *ViewPortMap) AddViewPort(id akara.EID) *ViewPortComponent {
	return cm.Add(id).(*ViewPortComponent)
}

// Get returns the component associated with the given entity id
func (cm *ViewPortMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetViewPort returns the ViewPortComponent associated with the given entity id
func (cm *ViewPortMap) GetViewPort(id akara.EID) (*ViewPortComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *ViewPortMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
