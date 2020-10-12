package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2pl2"
	"github.com/gravestench/akara"
)

// static check that PaletteTransformComponent implements Component
var _ akara.Component = &PaletteTransformComponent{}

// static check that PaletteTransformMap implements ComponentMap
var _ akara.ComponentMap = &PaletteTransformMap{}

// PaletteTransformComponent is a component that contains an embedded PL2 struct
type PaletteTransformComponent struct {
	Transform *d2pl2.PL2
}

// ID returns a unique identifier for the component type
func (*PaletteTransformComponent) ID() akara.ComponentID {
	return AssetPaletteTransformCID
}

// NewMap returns a new component map for the component type
func (*PaletteTransformComponent) NewMap() akara.ComponentMap {
	return NewPaletteTransformMap()
}

// PaletteTransform is a convenient reference to be used as a component identifier
var PaletteTransform = (*PaletteTransformComponent)(
	nil) // nolint:gochecknoglobals // global by design

// NewPaletteTransformMap creates a new map of entity ID's to PaletteTransform
func NewPaletteTransformMap() *PaletteTransformMap {
	cm := &PaletteTransformMap{
		components: make(map[akara.EID]*PaletteTransformComponent),
	}

	return cm
}

// PaletteTransformMap is a map of entity ID's to PaletteTransform
type PaletteTransformMap struct {
	world      *akara.World
	components map[akara.EID]*PaletteTransformComponent
}

// Init initializes the component map with the given world
func (cm *PaletteTransformMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*PaletteTransformMap) ID() akara.ComponentID {
	return AssetPaletteTransformCID
}

// NewMap returns a new component map for the component type
func (*PaletteTransformMap) NewMap() akara.ComponentMap {
	return NewPaletteTransformMap()
}

// Add a new PaletteTransformComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *PaletteTransformMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &PaletteTransformComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddPaletteTransform adds a new PaletteTransformComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *PaletteTransformComponent instead of an akara.Component
func (cm *PaletteTransformMap) AddPaletteTransform(id akara.EID) *PaletteTransformComponent {
	return cm.Add(id).(*PaletteTransformComponent)
}

// Get returns the component associated with the given entity id
func (cm *PaletteTransformMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetPaletteTransform returns the PaletteTransformComponent associated with the given entity id
func (cm *PaletteTransformMap) GetPaletteTransform(id akara.EID) (*PaletteTransformComponent,
	bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *PaletteTransformMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
