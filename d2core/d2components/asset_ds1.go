package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
	"github.com/gravestench/akara"
)

// static check that Ds1Component implements Component
var _ akara.Component = &Ds1Component{}

// static check that Ds1Map implements ComponentMap
var _ akara.ComponentMap = &Ds1Map{}

// Ds1Component is a component that contains an embedded ds1 struct
type Ds1Component struct {
	*d2ds1.DS1
}

// ID returns a unique identifier for the component type
func (*Ds1Component) ID() akara.ComponentID {
	return AssetDs1CID
}

// NewMap returns a new component map for the component type
func (*Ds1Component) NewMap() akara.ComponentMap {
	return NewDs1Map()
}

// Ds1 is a convenient reference to be used as a component identifier
var Ds1 = (*Ds1Component)(nil) // nolint:gochecknoglobals // global by design

// NewDs1Map creates a new map of entity ID's to Ds1Component
func NewDs1Map() *Ds1Map {
	cm := &Ds1Map{
		components: make(map[akara.EID]*Ds1Component),
	}

	return cm
}

// Ds1Map is a map of entity ID's to Ds1Component
type Ds1Map struct {
	world      *akara.World
	components map[akara.EID]*Ds1Component
}

// Init initializes the component map with the given world
func (cm *Ds1Map) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*Ds1Map) ID() akara.ComponentID {
	return AssetDs1CID
}

// NewMap returns a new component map for the component type
func (*Ds1Map) NewMap() akara.ComponentMap {
	return NewDs1Map()
}

// Add a new Ds1Component for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *Ds1Map) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &Ds1Component{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddDs1 adds a new Ds1Component for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *Ds1Component instead of an akara.Component
func (cm *Ds1Map) AddDs1(id akara.EID) *Ds1Component {
	return cm.Add(id).(*Ds1Component)
}

// Get returns the component associated with the given entity id
func (cm *Ds1Map) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetDs1Component returns the Ds1Component associated with the given entity id
func (cm *Ds1Map) GetDs1Component(id akara.EID) (*Ds1Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *Ds1Map) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
