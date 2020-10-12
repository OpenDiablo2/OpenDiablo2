package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2animdata"
	"github.com/gravestench/akara"
)

// static check that AnimDataComponent implements Component
var _ akara.Component = &AnimDataComponent{}

// static check that AnimDataMap implements ComponentMap
var _ akara.ComponentMap = &AnimDataMap{}

// AnimDataComponent is a component that contains an embedded animdata struct
type AnimDataComponent struct {
	*d2animdata.AnimationData
}

// ID returns a unique identifier for the component type
func (*AnimDataComponent) ID() akara.ComponentID {
	return AssetD2AnimDataCID
}

// NewMap returns a new component map for the component type
func (*AnimDataComponent) NewMap() akara.ComponentMap {
	return NewAnimDataMap()
}

// AnimData is a convenient reference to be used as a component identifier
var AnimData = (*AnimDataComponent)(nil) // nolint:gochecknoglobals // global by design

// NewAnimDataMap creates a new map of entity ID's to AnimData
func NewAnimDataMap() *AnimDataMap {
	cm := &AnimDataMap{
		components: make(map[akara.EID]*AnimDataComponent),
	}

	return cm
}

// AnimDataMap is a map of entity ID's to AnimData
type AnimDataMap struct {
	world      *akara.World
	components map[akara.EID]*AnimDataComponent
}

// Init initializes the component map with the given world
func (cm *AnimDataMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*AnimDataMap) ID() akara.ComponentID {
	return AssetD2AnimDataCID
}

// NewMap returns a new component map for the component type
func (*AnimDataMap) NewMap() akara.ComponentMap {
	return NewAnimDataMap()
}

// Add a new AnimDataComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *AnimDataMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &AnimDataComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddAnimData adds a new AnimDataComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *AnimDataComponent instead of an akara.Component
func (cm *AnimDataMap) AddAnimData(id akara.EID) *AnimDataComponent {
	return cm.Add(id).(*AnimDataComponent)
}

// Get returns the component associated with the given entity id
func (cm *AnimDataMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetAnimData returns the AnimDataComponent associated with the given entity id
func (cm *AnimDataMap) GetAnimData(id akara.EID) (*AnimDataComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *AnimDataMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
