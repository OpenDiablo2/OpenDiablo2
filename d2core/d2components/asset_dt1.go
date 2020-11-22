//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"
)

// static check that Dt1Component implements Component
var _ akara.Component = &Dt1Component{}

// static check that Dt1Map implements ComponentMap
var _ akara.ComponentMap = &Dt1Map{}

// Dt1Component is a component that contains an embedded DT1 struct
type Dt1Component struct {
	*akara.BaseComponent
	*d2dt1.DT1
}

// Dt1Map is a map of entity ID's to Dt1
type Dt1Map struct {
	*akara.BaseComponentMap
}

// AddDt1 adds a new Dt1Component for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *Dt1Component instead of an akara.Component
func (cm *Dt1Map) AddDt1(id akara.EID) *Dt1Component {
	return cm.Add(id).(*Dt1Component)
}

// GetDt1 returns the Dt1Component associated with the given entity id
func (cm *Dt1Map) GetDt1(id akara.EID) (*Dt1Component, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*Dt1Component), found
}

// Dt1 is a convenient reference to be used as a component identifier
var Dt1 = newDt1() // nolint:gochecknoglobals // global by design

func newDt1() akara.Component {
	return &Dt1Component{
		BaseComponent: akara.NewBaseComponent(AssetDt1CID, newDt1, newDt1Map),
	}
}

func newDt1Map() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(AssetDt1CID, newDt1, newDt1Map)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &Dt1Map{
		BaseComponentMap: baseMap,
	}

	return cm
}
