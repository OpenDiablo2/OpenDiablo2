package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// static check that DataDictionaryComponent implements Component
var _ akara.Component = &DataDictionaryComponent{}

// static check that DataDictionaryMap implements ComponentMap
var _ akara.ComponentMap = &DataDictionaryMap{}

// DataDictionaryComponent is a component that contains a file Type
type DataDictionaryComponent struct {
	*d2txt.DataDictionary
}

// ID returns a unique identifier for the component type
func (*DataDictionaryComponent) ID() akara.ComponentID {
	return AssetDataDictionaryCID
}

// NewMap returns a new component map for the component type
func (*DataDictionaryComponent) NewMap() akara.ComponentMap {
	return NewDataDictionaryMap()
}

// DataDictionary is a convenient reference to be used as a component identifier
var DataDictionary = (*DataDictionaryComponent)(nil) // nolint:gochecknoglobals // global by design

// NewDataDictionaryMap creates a new map of entity ID's to DataDictionary
func NewDataDictionaryMap() *DataDictionaryMap {
	cm := &DataDictionaryMap{
		components: make(map[akara.EID]*DataDictionaryComponent),
	}

	return cm
}

// DataDictionaryMap is a map of entity ID's to DataDictionary
type DataDictionaryMap struct {
	world      *akara.World
	components map[akara.EID]*DataDictionaryComponent
}

// Init initializes the component map with the given world
func (cm *DataDictionaryMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*DataDictionaryMap) ID() akara.ComponentID {
	return AssetDataDictionaryCID
}

// NewMap returns a new component map for the component type
func (*DataDictionaryMap) NewMap() akara.ComponentMap {
	return NewDataDictionaryMap()
}

// Add a new DataDictionaryComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *DataDictionaryMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &DataDictionaryComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddDataDictionary adds a new DataDictionaryComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *DataDictionaryComponent instead of an akara.Component
func (cm *DataDictionaryMap) AddDataDictionary(id akara.EID) *DataDictionaryComponent {
	return cm.Add(id).(*DataDictionaryComponent)
}

// Get returns the component associated with the given entity id
func (cm *DataDictionaryMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetDataDictionary returns the DataDictionaryComponent associated with the given entity id
func (cm *DataDictionaryMap) GetDataDictionary(id akara.EID) (*DataDictionaryComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *DataDictionaryMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
