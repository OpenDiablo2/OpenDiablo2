package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"
	"github.com/gravestench/akara"
)

// static check that Dt1Component implements Component
var _ akara.Component = &Dt1Component{}

// static check that Dt1Map implements ComponentMap
var _ akara.ComponentMap = &Dt1Map{}

// Dt1Component is a component that contains an embedded dt1 struct
type Dt1Component struct {
	*d2dt1.DT1
}

// ID returns a unique identifier for the component type
func (*Dt1Component) ID() akara.ComponentID {
	return AssetDt1CID
}

// NewMap returns a new component map for the component type
func (*Dt1Component) NewMap() akara.ComponentMap {
	return NewDt1Map()
}

// Dt1 is a convenient reference to be used as a component identifier
var Dt1 = (*Dt1Component)(nil) // nolint:gochecknoglobals // global by design

// NewDt1Map creates a new map of entity ID's to Dt1
func NewDt1Map() *Dt1Map {
	cm := &Dt1Map{
		components: make(map[akara.EID]*Dt1Component),
	}

	return cm
}

// Dt1Map is a map of entity ID's to Dt1
type Dt1Map struct {
	world      *akara.World
	components map[akara.EID]*Dt1Component
}

// Init initializes the component map with the given world
func (cm *Dt1Map) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*Dt1Map) ID() akara.ComponentID {
	return AssetDt1CID
}

// NewMap returns a new component map for the component type
func (*Dt1Map) NewMap() akara.ComponentMap {
	return NewDt1Map()
}

// Add a new Dt1Component for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *Dt1Map) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &Dt1Component{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddDt1 adds a new Dt1Component for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *Dt1Component instead of an akara.Component
func (cm *Dt1Map) AddDt1(id akara.EID) *Dt1Component {
	return cm.Add(id).(*Dt1Component)
}

// Get returns the component associated with the given entity id
func (cm *Dt1Map) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetDt1 returns the Dt1Component associated with the given entity id
func (cm *Dt1Map) GetDt1(id akara.EID) (*Dt1Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *Dt1Map) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
