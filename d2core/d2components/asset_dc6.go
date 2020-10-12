package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
	"github.com/gravestench/akara"
)

// static check that Dc6Component implements Component
var _ akara.Component = &Dc6Component{}

// static check that Dc6Map implements ComponentMap
var _ akara.ComponentMap = &Dc6Map{}

// Dc6Component is a component that contains an embedded dc6 struct
type Dc6Component struct {
	*d2dc6.DC6
}

// ID returns a unique identifier for the component type
func (*Dc6Component) ID() akara.ComponentID {
	return AssetDc6CID
}

// NewMap returns a new component map for the component type
func (*Dc6Component) NewMap() akara.ComponentMap {
	return NewDc6Map()
}

// Dc6 is a convenient reference to be used as a component identifier
var Dc6 = (*Dc6Component)(nil) // nolint:gochecknoglobals // global by design

// NewDc6Map creates a new map of entity ID's to Dc6
func NewDc6Map() *Dc6Map {
	cm := &Dc6Map{
		components: make(map[akara.EID]*Dc6Component),
	}

	return cm
}

// Dc6Map is a map of entity ID's to Dc6
type Dc6Map struct {
	world      *akara.World
	components map[akara.EID]*Dc6Component
}

// Init initializes the component map with the given world
func (cm *Dc6Map) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*Dc6Map) ID() akara.ComponentID {
	return AssetDc6CID
}

// NewMap returns a new component map for the component type
func (*Dc6Map) NewMap() akara.ComponentMap {
	return NewDc6Map()
}

// Add a new Dc6Component for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *Dc6Map) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &Dc6Component{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddDc6 adds a new Dc6Component for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *Dc6Component instead of an akara.Component
func (cm *Dc6Map) AddDc6(id akara.EID) *Dc6Component {
	return cm.Add(id).(*Dc6Component)
}

// Get returns the component associated with the given entity id
func (cm *Dc6Map) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetDc6 returns the Dc6Component associated with the given entity id
func (cm *Dc6Map) GetDc6(id akara.EID) (*Dc6Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *Dc6Map) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
