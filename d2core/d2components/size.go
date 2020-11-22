//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
)

// static check that SizeComponent implements Component
var _ akara.Component = &SizeComponent{}

// static check that SizeMap implements ComponentMap
var _ akara.ComponentMap = &SizeMap{}

// SizeComponent represents an entities width and height as a vector
type SizeComponent struct {
	*akara.BaseComponent
	*d2vector.Vector
}

// SizeMap is a map of entity ID's to Size
type SizeMap struct {
	*akara.BaseComponentMap
}

// AddSize adds a new SizeComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *SizeComponent instead of an akara.Component
func (cm *SizeMap) AddSize(id akara.EID) *SizeComponent {
	c := cm.Add(id).(*SizeComponent)

	c.Vector = d2vector.NewVector(1, 1)

	return c
}

// GetSize returns the SizeComponent associated with the given entity id
func (cm *SizeMap) GetSize(id akara.EID) (*SizeComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*SizeComponent), found
}

// Size is a convenient reference to be used as a component identifier
var Size = newSize() // nolint:gochecknoglobals // global by design

func newSize() akara.Component {
	return &SizeComponent{
		BaseComponent: akara.NewBaseComponent(SizeCID, newSize, newSizeMap),
	}
}

func newSizeMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(SizeCID, newSize, newSizeMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &SizeMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
