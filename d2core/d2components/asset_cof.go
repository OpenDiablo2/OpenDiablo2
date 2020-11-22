//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2cof"
)

// static check that CofComponent implements Component
var _ akara.Component = &CofComponent{}

// static check that CofMap implements ComponentMap
var _ akara.ComponentMap = &CofMap{}

// CofComponent is a component that contains an embedded cof struct
type CofComponent struct {
	*akara.BaseComponent
	*d2cof.COF
}

// CofMap is a map of entity ID's to Cof
type CofMap struct {
	*akara.BaseComponentMap
}

// AddCof adds a new CofComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *CofComponent instead of an akara.Component
func (cm *CofMap) AddCof(id akara.EID) *CofComponent {
	return cm.Add(id).(*CofComponent)
}

// GetCof returns the CofComponent associated with the given entity id
func (cm *CofMap) GetCof(id akara.EID) (*CofComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*CofComponent), found
}

// Cof is a convenient reference to be used as a component identifier
var Cof = newCof() // nolint:gochecknoglobals // global by design

func newCof() akara.Component {
	return &CofComponent{
		BaseComponent: akara.NewBaseComponent(AssetCofCID, newCof, newCofMap),
	}
}

func newCofMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(AssetCofCID, newCof, newCofMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &CofMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
