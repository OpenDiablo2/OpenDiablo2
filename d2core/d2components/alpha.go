//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that AlphaComponent implements Component
var _ akara.Component = &AlphaComponent{}

// static check that AlphaMap implements ComponentMap
var _ akara.ComponentMap = &AlphaMap{}

// AlphaComponent is a component that contains an embedded cof struct
type AlphaComponent struct {
	*akara.BaseComponent
	Alpha float64
}

// AlphaMap is a map of entity ID's to Alpha
type AlphaMap struct {
	*akara.BaseComponentMap
}

// AddAlpha adds a new AlphaComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *AlphaComponent instead of an akara.Component
func (cm *AlphaMap) AddAlpha(id akara.EID) *AlphaComponent {
	c := cm.Add(id).(*AlphaComponent)

	c.Alpha = 1

	return c
}

// GetAlpha returns the AlphaComponent associated with the given entity id
func (cm *AlphaMap) GetAlpha(id akara.EID) (*AlphaComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*AlphaComponent), found
}

// Alpha is a convenient reference to be used as a component identifier
var Alpha = newAlpha() // nolint:gochecknoglobals // global by design

func newAlpha() akara.Component {
	return &AlphaComponent{
		BaseComponent: akara.NewBaseComponent(AlphaCID, newAlpha, newAlphaMap),
	}
}

func newAlphaMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(AlphaCID, newAlpha, newAlphaMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &AlphaMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
