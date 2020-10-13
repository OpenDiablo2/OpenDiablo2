package d2components

import (
	"github.com/gravestench/akara"
)

// static check that FontTableComponent implements Component
var _ akara.Component = &FontTableComponent{}

// static check that FontTableMap implements ComponentMap
var _ akara.ComponentMap = &FontTableMap{}

// FontTableComponent is a component that contains a file Type
type FontTableComponent struct {
	Data []byte
}

// ID returns a unique identifier for the component type
func (*FontTableComponent) ID() akara.ComponentID {
	return AssetFontTableCID
}

// NewMap returns a new component map for the component type
func (*FontTableComponent) NewMap() akara.ComponentMap {
	return NewFontTableMap()
}

// FontTable is a convenient reference to be used as a component identifier
var FontTable = (*FontTableComponent)(nil) // nolint:gochecknoglobals // global by design

// NewFontTableMap creates a new map of entity ID's to FontTable
func NewFontTableMap() *FontTableMap {
	cm := &FontTableMap{
		components: make(map[akara.EID]*FontTableComponent),
	}

	return cm
}

// FontTableMap is a map of entity ID's to FontTable
type FontTableMap struct {
	world      *akara.World
	components map[akara.EID]*FontTableComponent
}

// Init initializes the component map with the given world
func (cm *FontTableMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*FontTableMap) ID() akara.ComponentID {
	return AssetFontTableCID
}

// NewMap returns a new component map for the component type
func (*FontTableMap) NewMap() akara.ComponentMap {
	return NewFontTableMap()
}

// Add a new FontTableComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *FontTableMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &FontTableComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddFontTable adds a new FontTableComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *FontTableComponent instead of an akara.Component
func (cm *FontTableMap) AddFontTable(id akara.EID) *FontTableComponent {
	return cm.Add(id).(*FontTableComponent)
}

// Get returns the component associated with the given entity id
func (cm *FontTableMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetFontTable returns the FontTableComponent associated with the given entity id
func (cm *FontTableMap) GetFontTable(id akara.EID) (*FontTableComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *FontTableMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
