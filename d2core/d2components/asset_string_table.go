package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2tbl"
)

// static check that StringTableComponent implements Component
var _ akara.Component = &StringTableComponent{}

// static check that StringTableMap implements ComponentMap
var _ akara.ComponentMap = &StringTableMap{}

// StringTableComponent is a component that contains an embedded string table struct
type StringTableComponent struct {
	*d2tbl.TextDictionary
}

// ID returns a unique identifier for the component type
func (*StringTableComponent) ID() akara.ComponentID {
	return AssetStringTableCID
}

// NewMap returns a new component map for the component type
func (*StringTableComponent) NewMap() akara.ComponentMap {
	return NewStringTableMap()
}

// StringTable is a convenient reference to be used as a component identifier
var StringTable = (*StringTableComponent)(nil) // nolint:gochecknoglobals // global by design

// NewStringTableMap creates a new map of entity ID's to StringTable
func NewStringTableMap() *StringTableMap {
	cm := &StringTableMap{
		components: make(map[akara.EID]*StringTableComponent),
	}

	return cm
}

// StringTableMap is a map of entity ID's to StringTable
type StringTableMap struct {
	world      *akara.World
	components map[akara.EID]*StringTableComponent
}

// Init initializes the component map with the given world
func (cm *StringTableMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*StringTableMap) ID() akara.ComponentID {
	return AssetStringTableCID
}

// NewMap returns a new component map for the component type
func (*StringTableMap) NewMap() akara.ComponentMap {
	return NewStringTableMap()
}

// Add a new StringTableComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *StringTableMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &StringTableComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddStringTable adds a new StringTableComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *StringTableComponent instead of an akara.Component
func (cm *StringTableMap) AddStringTable(id akara.EID) *StringTableComponent {
	return cm.Add(id).(*StringTableComponent)
}

// Get returns the component associated with the given entity id
func (cm *StringTableMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetStringTable returns the StringTableComponent associated with the given entity id
func (cm *StringTableMap) GetStringTable(id akara.EID) (*StringTableComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *StringTableMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
