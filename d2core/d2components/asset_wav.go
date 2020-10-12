package d2components

import (
	"io"

	"github.com/gravestench/akara"
)

// static check that WavComponent implements Component
var _ akara.Component = &WavComponent{}

// static check that WavMap implements ComponentMap
var _ akara.ComponentMap = &WavMap{}

// WavComponent is a component that contains an embedded wav.Stream
type WavComponent struct {
	Data io.ReadSeeker
}

// ID returns a unique identifier for the component type
func (*WavComponent) ID() akara.ComponentID {
	return AssetWavCID
}

// NewMap returns a new component map for the component type
func (*WavComponent) NewMap() akara.ComponentMap {
	return NewWavMap()
}

// Wav is a convenient reference to be used as a component identifier
var Wav = (*WavComponent)(nil) // nolint:gochecknoglobals // global by design

// NewWavMap creates a new map of entity ID's to Wav
func NewWavMap() *WavMap {
	cm := &WavMap{
		components: make(map[akara.EID]*WavComponent),
	}

	return cm
}

// WavMap is a map of entity ID's to Wav
type WavMap struct {
	world      *akara.World
	components map[akara.EID]*WavComponent
}

// Init initializes the component map with the given world
func (cm *WavMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*WavMap) ID() akara.ComponentID {
	return AssetWavCID
}

// NewMap returns a new component map for the component type
func (*WavMap) NewMap() akara.ComponentMap {
	return NewWavMap()
}

// Add a new WavComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *WavMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = &WavComponent{}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddWav adds a new WavComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *WavComponent instead of an akara.Component
func (cm *WavMap) AddWav(id akara.EID) *WavComponent {
	return cm.Add(id).(*WavComponent)
}

// Get returns the component associated with the given entity id
func (cm *WavMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetWav returns the WavComponent associated with the given entity id
func (cm *WavMap) GetWav(id akara.EID) (*WavComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *WavMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
