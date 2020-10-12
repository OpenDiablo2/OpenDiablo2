package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/gravestench/akara"
)

// static check that DccComponent implements Component
var _ akara.Component = &DccComponent{}

// static check that DccMap implements ComponentMap
var _ akara.ComponentMap = &DccMap{}

// DccComponent is a component that contains an embedded dcc struct
type DccComponent struct {
	*d2dcc.DCC
}

// ID returns a unique identifier for the component type
func (*DccComponent) ID() akara.ComponentID {
	return AssetDccCID
}

// NewMap returns a new component map for the component type
func (*DccComponent) NewMap() akara.ComponentMap {
	return NewDccMap()
}

// Dcc is a convenient reference to be used as a component identifier
var Dcc = (*DccComponent)(nil) // nolint:gochecknoglobals // global by design

// NewDccMap creates a new map of entity ID's to Dcc
func NewDccMap() *DccMap {
	cm := &DccMap{
		components: make(map[akara.EID]*DccComponent),
	}

	return cm
}

// DccMap is a map of entity ID's to Dcc
type DccMap struct {
	world      *akara.World
	components map[akara.EID]*DccComponent
}

// Init initializes the component map with the given world
func (cm *DccMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*DccMap) ID() akara.ComponentID {
	return AssetDccCID
}

// NewMap returns a new component map for the component type
func (*DccMap) NewMap() akara.ComponentMap {
	return NewDccMap()
}

// Add a new DccComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *DccMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &DccComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddDcc adds a new DccComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *DccComponent instead of an akara.Component
func (cm *DccMap) AddDcc(id akara.EID) *DccComponent {
	return cm.Add(id).(*DccComponent)
}

// Get returns the component associated with the given entity id
func (cm *DccMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetDcc returns the DccComponent associated with the given entity id
func (cm *DccMap) GetDcc(id akara.EID) (*DccComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *DccMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
