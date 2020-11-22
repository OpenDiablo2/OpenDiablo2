//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that OriginComponent implements Component
var _ akara.Component = &OriginComponent{}

// static check that OriginMap implements ComponentMap
var _ akara.ComponentMap = &OriginMap{}

// OriginComponent is a component that describes the origin point of an entity.
// The values are normalized to the display width/height.
// For example, origin (0,0) is top-left corner, (0.5, 0.5) is center
type OriginComponent struct {
	*akara.BaseComponent
	X, Y float64
}

// OriginMap is a map of entity ID's to Origin
type OriginMap struct {
	*akara.BaseComponentMap
}

// AddOrigin adds a new OriginComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *OriginComponent instead of an akara.Component
func (cm *OriginMap) AddOrigin(id akara.EID) *OriginComponent {
	return cm.Add(id).(*OriginComponent)
}

// GetOrigin returns the OriginComponent associated with the given entity id
func (cm *OriginMap) GetOrigin(id akara.EID) (*OriginComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*OriginComponent), found
}

// Origin is a convenient reference to be used as a component identifier
var Origin = newOrigin() // nolint:gochecknoglobals // global by design

func newOrigin() akara.Component {
	return &OriginComponent{
		BaseComponent: akara.NewBaseComponent(OriginCID, newOrigin, newOriginMap),
	}
}

func newOriginMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(OriginCID, newOrigin, newOriginMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &OriginMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
