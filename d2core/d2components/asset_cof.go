package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2cof"
	"github.com/gravestench/akara"
)

// static check that CofComponent implements Component
var _ akara.Component = &CofComponent{}

// static check that CofMap implements ComponentMap
var _ akara.ComponentMap = &CofMap{}

// CofComponent is a component that contains an embedded cof struct
type CofComponent struct {
	*d2cof.COF
}

// ID returns a unique identifier for the component type
func (*CofComponent) ID() akara.ComponentID {
	return AssetCofCID
}

// NewMap returns a new component map for the component type
func (*CofComponent) NewMap() akara.ComponentMap {
	return NewCofMap()
}

// Cof is a convenient reference to be used as a component identifier
var Cof = (*CofComponent)(nil) // nolint:gochecknoglobals // global by design

// NewCofMap creates a new map of entity ID's to Cof
func NewCofMap() *CofMap {
	cm := &CofMap{
		components: make(map[akara.EID]*CofComponent),
	}

	return cm
}

// CofMap is a map of entity ID's to Cof
type CofMap struct {
	world      *akara.World
	components map[akara.EID]*CofComponent
}

// Init initializes the component map with the given world
func (cm *CofMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*CofMap) ID() akara.ComponentID {
	return AssetCofCID
}

// NewMap returns a new component map for the component type
func (*CofMap) NewMap() akara.ComponentMap {
	return NewCofMap()
}

// Add a new CofComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *CofMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &CofComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddCof adds a new CofComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *CofComponent instead of an akara.Component
func (cm *CofMap) AddCof(id akara.EID) *CofComponent {
	return cm.Add(id).(*CofComponent)
}

// Get returns the component associated with the given entity id
func (cm *CofMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetCof returns the CofComponent associated with the given entity id
func (cm *CofMap) GetCof(id akara.EID) (*CofComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *CofMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
