package d2components

import (
	"github.com/gravestench/akara"
)

// static check that ViewportFilterComponent implements Component
var _ akara.Component = &ViewportFilterComponent{}

// static check that ViewportFilterMap implements ComponentMap
var _ akara.ComponentMap = &ViewportFilterMap{}

// ViewportFilterComponent is a component that contains a bitset that denotes which viewport
// the entity will be rendered.
type ViewportFilterComponent struct {
	*akara.BitSet
}

// ID returns a unique identifier for the component type
func (*ViewportFilterComponent) ID() akara.ComponentID {
	return ViewportFilterCID
}

// NewMap returns a new component map for the component type
func (*ViewportFilterComponent) NewMap() akara.ComponentMap {
	return NewViewportFilterMap()
}

// ViewportFilter is a convenient reference to be used as a component identifier
var ViewportFilter = (*ViewportFilterComponent)(nil) // nolint:gochecknoglobals // global by design

// NewViewportFilterMap creates a new map of entity ID's to ViewportFilter
func NewViewportFilterMap() *ViewportFilterMap {
	cm := &ViewportFilterMap{
		components: make(map[akara.EID]*ViewportFilterComponent),
	}

	return cm
}

// ViewportFilterMap is a map of entity ID's to ViewportFilter
type ViewportFilterMap struct {
	world      *akara.World
	components map[akara.EID]*ViewportFilterComponent
}

// Init initializes the component map with the given world
func (cm *ViewportFilterMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*ViewportFilterMap) ID() akara.ComponentID {
	return ViewportFilterCID
}

// NewMap returns a new component map for the component type
func (*ViewportFilterMap) NewMap() akara.ComponentMap {
	return NewViewportFilterMap()
}

// Add a new ViewportFilterComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *ViewportFilterMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &ViewportFilterComponent{
		BitSet: akara.NewBitSet(),
	}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddViewportFilter adds a new ViewportFilterComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *ViewportFilterComponent instead of an akara.Component
func (cm *ViewportFilterMap) AddViewportFilter(id akara.EID) *ViewportFilterComponent {
	return cm.Add(id).(*ViewportFilterComponent)
}

// Get returns the component associated with the given entity id
func (cm *ViewportFilterMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetViewportFilter returns the ViewportFilterComponent associated with the given entity id
func (cm *ViewportFilterMap) GetViewportFilter(id akara.EID) (*ViewportFilterComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *ViewportFilterMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}

