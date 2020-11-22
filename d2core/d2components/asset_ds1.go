//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
)

// static check that Ds1Component implements Component
var _ akara.Component = &Ds1Component{}

// static check that Ds1Map implements ComponentMap
var _ akara.ComponentMap = &Ds1Map{}

// Ds1Component is a component that contains an embedded DS1 struct
type Ds1Component struct {
	*akara.BaseComponent
	*d2ds1.DS1
}

// Ds1Map is a map of entity ID's to Ds1
type Ds1Map struct {
	*akara.BaseComponentMap
}

// AddDs1 adds a new Ds1Component for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *Ds1Component instead of an akara.Component
func (cm *Ds1Map) AddDs1(id akara.EID) *Ds1Component {
	return cm.Add(id).(*Ds1Component)
}

// GetDs1 returns the Ds1Component associated with the given entity id
func (cm *Ds1Map) GetDs1(id akara.EID) (*Ds1Component, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*Ds1Component), found
}

// Ds1 is a convenient reference to be used as a component identifier
var Ds1 = newDs1() // nolint:gochecknoglobals // global by design

func newDs1() akara.Component {
	return &Ds1Component{
		BaseComponent: akara.NewBaseComponent(AssetDs1CID, newDs1, newDs1Map),
	}
}

func newDs1Map() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(AssetDs1CID, newDs1, newDs1Map)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &Ds1Map{
		BaseComponentMap: baseMap,
	}

	return cm
}
