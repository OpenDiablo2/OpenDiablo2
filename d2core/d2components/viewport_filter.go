//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that ViewportFilterComponent implements Component
var _ akara.Component = &ViewportFilterComponent{}

// static check that ViewportFilterMap implements ComponentMap
var _ akara.ComponentMap = &ViewportFilterMap{}

// ViewportFilterComponent is a component that contains a bitset that denotes which viewport
// the entity will be rendered.
type ViewportFilterComponent struct {
	*akara.BaseComponent
	*akara.BitSet
}

// ViewportFilterMap is a map of entity ID's to ViewportFilter
type ViewportFilterMap struct {
	*akara.BaseComponentMap
}

// AddViewportFilter adds a new ViewportFilterComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *ViewportFilterComponent instead of an akara.Component
func (cm *ViewportFilterMap) AddViewportFilter(id akara.EID) *ViewportFilterComponent {
	c := cm.Add(id).(*ViewportFilterComponent)

	c.BitSet = akara.NewBitSet(0)

	return c
}

// GetViewportFilter returns the ViewportFilterComponent associated with the given entity id
func (cm *ViewportFilterMap) GetViewportFilter(id akara.EID) (*ViewportFilterComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*ViewportFilterComponent), found
}

// ViewportFilter is a convenient reference to be used as a component identifier
var ViewportFilter = newViewportFilter() // nolint:gochecknoglobals // global by design

func newViewportFilter() akara.Component {
	return &ViewportFilterComponent{
		BaseComponent: akara.NewBaseComponent(ViewportFilterCID, newViewportFilter, newViewportFilterMap),
	}
}

func newViewportFilterMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(ViewportFilterCID, newViewportFilter, newViewportFilterMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &ViewportFilterMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
