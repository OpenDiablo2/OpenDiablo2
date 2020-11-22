//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// static check that PaletteComponent implements Component
var _ akara.Component = &PaletteComponent{}

// static check that PaletteMap implements ComponentMap
var _ akara.ComponentMap = &PaletteMap{}

// PaletteComponent is a component that contains an embedded palette interface
type PaletteComponent struct {
	*akara.BaseComponent
	d2interface.Palette
}

// PaletteMap is a map of entity ID's to Palette
type PaletteMap struct {
	*akara.BaseComponentMap
}

// AddPalette adds a new PaletteComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *PaletteComponent instead of an akara.Component
func (cm *PaletteMap) AddPalette(id akara.EID) *PaletteComponent {
	return cm.Add(id).(*PaletteComponent)
}

// GetPalette returns the PaletteComponent associated with the given entity id
func (cm *PaletteMap) GetPalette(id akara.EID) (*PaletteComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*PaletteComponent), found
}

// Palette is a convenient reference to be used as a component identifier
var Palette = newPalette() // nolint:gochecknoglobals // global by design

func newPalette() akara.Component {
	return &PaletteComponent{
		BaseComponent: akara.NewBaseComponent(AssetPaletteCID, newPalette, newPaletteMap),
	}
}

func newPaletteMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(AssetPaletteCID, newPalette, newPaletteMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &PaletteMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
