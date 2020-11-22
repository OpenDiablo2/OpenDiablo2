//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
)

// static check that ScaleComponent implements Component
var _ akara.Component = &ScaleComponent{}

// static check that ScaleMap implements ComponentMap
var _ akara.ComponentMap = &ScaleMap{}

// ScaleComponent represents an entities x,y axis scale as a vector
type ScaleComponent struct {
	*akara.BaseComponent
	*d2vector.Vector
}

// ScaleMap is a map of entity ID's to Scale
type ScaleMap struct {
	*akara.BaseComponentMap
}

// AddScale adds a new ScaleComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *ScaleComponent instead of an akara.Component
func (cm *ScaleMap) AddScale(id akara.EID) *ScaleComponent {
	c := cm.Add(id).(*ScaleComponent)

	c.Vector = d2vector.NewVector(1, 1)

	return c
}

// GetScale returns the ScaleComponent associated with the given entity id
func (cm *ScaleMap) GetScale(id akara.EID) (*ScaleComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*ScaleComponent), found
}

// Scale is a convenient reference to be used as a component identifier
var Scale = newScale() // nolint:gochecknoglobals // global by design

func newScale() akara.Component {
	return &ScaleComponent{
		BaseComponent: akara.NewBaseComponent(ScaleCID, newScale, newScaleMap),
	}
}

func newScaleMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(ScaleCID, newScale, newScaleMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &ScaleMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
