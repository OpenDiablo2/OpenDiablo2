//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"io"

	"github.com/gravestench/akara"
)

// static check that WavComponent implements Component
var _ akara.Component = &WavComponent{}

// static check that WavMap implements ComponentMap
var _ akara.ComponentMap = &WavMap{}

// WavComponent is a component that contains an embedded io.ReadSeeker for streaming wav audio files
type WavComponent struct {
	*akara.BaseComponent
	Data io.ReadSeeker
}

// WavMap is a map of entity ID's to Wav
type WavMap struct {
	*akara.BaseComponentMap
}

// AddWav adds a new WavComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *WavComponent instead of an akara.Component
func (cm *WavMap) AddWav(id akara.EID) *WavComponent {
	return cm.Add(id).(*WavComponent)
}

// GetWav returns the WavComponent associated with the given entity id
func (cm *WavMap) GetWav(id akara.EID) (*WavComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*WavComponent), found
}

// Wav is a convenient reference to be used as a component identifier
var Wav = newWav() // nolint:gochecknoglobals // global by design

func newWav() akara.Component {
	return &WavComponent{
		BaseComponent: akara.NewBaseComponent(AssetWavCID, newWav, newWavMap),
	}
}

func newWavMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(AssetWavCID, newWav, newWavMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &WavMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
