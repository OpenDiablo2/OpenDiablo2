//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
)

// static check that DccComponent implements Component
var _ akara.Component = &DccComponent{}

// static check that DccMap implements ComponentMap
var _ akara.ComponentMap = &DccMap{}

// DccComponent is a component that contains an embedded DCC struct
type DccComponent struct {
	*akara.BaseComponent
	*d2dcc.DCC
}

// DccMap is a map of entity ID's to Dcc
type DccMap struct {
	*akara.BaseComponentMap
}

// AddDcc adds a new DccComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *DccComponent instead of an akara.Component
func (cm *DccMap) AddDcc(id akara.EID) *DccComponent {
	return cm.Add(id).(*DccComponent)
}

// GetDcc returns the DccComponent associated with the given entity id
func (cm *DccMap) GetDcc(id akara.EID) (*DccComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*DccComponent), found
}

// Dcc is a convenient reference to be used as a component identifier
var Dcc = newDcc() // nolint:gochecknoglobals // global by design

func newDcc() akara.Component {
	return &DccComponent{
		BaseComponent: akara.NewBaseComponent(AssetDccCID, newDcc, newDccMap),
	}
}

func newDccMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(AssetDccCID, newDcc, newDccMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &DccMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
