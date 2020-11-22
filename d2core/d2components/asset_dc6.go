//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
)

// static check that Dc6Component implements Component
var _ akara.Component = &Dc6Component{}

// static check that Dc6Map implements ComponentMap
var _ akara.ComponentMap = &Dc6Map{}

// Dc6Component is a component that contains an embedded DC6 struct
type Dc6Component struct {
	*akara.BaseComponent
	*d2dc6.DC6
}

// Dc6Map is a map of entity ID's to Dc6
type Dc6Map struct {
	*akara.BaseComponentMap
}

// AddDc6 adds a new Dc6Component for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *Dc6Component instead of an akara.Component
func (cm *Dc6Map) AddDc6(id akara.EID) *Dc6Component {
	return cm.Add(id).(*Dc6Component)
}

// GetDc6 returns the Dc6Component associated with the given entity id
func (cm *Dc6Map) GetDc6(id akara.EID) (*Dc6Component, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*Dc6Component), found
}

// Dc6 is a convenient reference to be used as a component identifier
var Dc6 = newDc6() // nolint:gochecknoglobals // global by design

func newDc6() akara.Component {
	return &Dc6Component{
		BaseComponent: akara.NewBaseComponent(AssetDc6CID, newDc6, newDc6Map),
	}
}

func newDc6Map() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(AssetDc6CID, newDc6, newDc6Map)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &Dc6Map{
		BaseComponentMap: baseMap,
	}

	return cm
}
