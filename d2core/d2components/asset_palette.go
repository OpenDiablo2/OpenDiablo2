package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/gravestench/akara"
)

// static check that PaletteComponent implements Component
var _ akara.Component = &PaletteComponent{}

// static check that PaletteMap implements ComponentMap
var _ akara.ComponentMap = &PaletteMap{}

// PaletteComponent is a component that contains a n embedded DATPalette struct
type PaletteComponent struct {
	d2interface.Palette
}

// ID returns a unique identifier for the component type
func (*PaletteComponent) ID() akara.ComponentID {
	return AssetPaletteCID
}

// NewMap returns a new component map for the component type
func (*PaletteComponent) NewMap() akara.ComponentMap {
	return NewPaletteMap()
}

// Palette is a convenient reference to be used as a component identifier
var Palette = (*PaletteComponent)(nil) // nolint:gochecknoglobals // global by design

// NewPaletteMap creates a new map of entity ID's to Palette
func NewPaletteMap() *PaletteMap {
	cm := &PaletteMap{
		components: make(map[akara.EID]*PaletteComponent),
	}

	return cm
}

// PaletteMap is a map of entity ID's to Palette
type PaletteMap struct {
	world      *akara.World
	components map[akara.EID]*PaletteComponent
}

// Init initializes the component map with the given world
func (cm *PaletteMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*PaletteMap) ID() akara.ComponentID {
	return AssetPaletteCID
}

// NewMap returns a new component map for the component type
func (*PaletteMap) NewMap() akara.ComponentMap {
	return NewPaletteMap()
}

// Add a new PaletteComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *PaletteMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &PaletteComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddPalette adds a new PaletteComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *PaletteComponent instead of an akara.Component
func (cm *PaletteMap) AddPalette(id akara.EID) *PaletteComponent {
	return cm.Add(id).(*PaletteComponent)
}

// Get returns the component associated with the given entity id
func (cm *PaletteMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetPalette returns the PaletteComponent associated with the given entity id
func (cm *PaletteMap) GetPalette(id akara.EID) (*PaletteComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *PaletteMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
