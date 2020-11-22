//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2pl2"
)

// static check that PaletteTransformComponent implements Component
var _ akara.Component = &PaletteTransformComponent{}

// static check that PaletteTransformMap implements ComponentMap
var _ akara.ComponentMap = &PaletteTransformMap{}

// PaletteTransformComponent is a component that contains an embedded palette transform (pl2) struct
type PaletteTransformComponent struct {
	*akara.BaseComponent
	Transform *d2pl2.PL2
}

// PaletteTransformMap is a map of entity ID's to PaletteTransform
type PaletteTransformMap struct {
	*akara.BaseComponentMap
}

// AddPaletteTransform adds a new PaletteTransformComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *PaletteTransformComponent instead of an akara.Component
func (cm *PaletteTransformMap) AddPaletteTransform(id akara.EID) *PaletteTransformComponent {
	return cm.Add(id).(*PaletteTransformComponent)
}

// GetPaletteTransform returns the PaletteTransformComponent associated with the given entity id
func (cm *PaletteTransformMap) GetPaletteTransform(id akara.EID) (*PaletteTransformComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*PaletteTransformComponent), found
}

// PaletteTransform is a convenient reference to be used as a component identifier
var PaletteTransform = newPaletteTransform() // nolint:gochecknoglobals // global by design

func newPaletteTransform() akara.Component {
	return &PaletteTransformComponent{
		BaseComponent: akara.NewBaseComponent(AssetPaletteTransformCID, newPaletteTransform, newPaletteTransformMap),
	}
}

func newPaletteTransformMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(AssetPaletteTransformCID, newPaletteTransform, newPaletteTransformMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &PaletteTransformMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
