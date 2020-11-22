//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that DirtyComponent implements Component
var _ akara.Component = &DirtyComponent{}

// static check that DirtyMap implements ComponentMap
var _ akara.ComponentMap = &DirtyMap{}

// DirtyComponent is a flag component that is used to denote a "dirty" state
type DirtyComponent struct {
	*akara.BaseComponent
	IsDirty bool
}

// DirtyMap is a map of entity ID's to Dirty
type DirtyMap struct {
	*akara.BaseComponentMap
}

// AddDirty adds a new DirtyComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *DirtyComponent instead of an akara.Component
func (cm *DirtyMap) AddDirty(id akara.EID) *DirtyComponent {
	return cm.Add(id).(*DirtyComponent)
}

// GetDirty returns the DirtyComponent associated with the given entity id
func (cm *DirtyMap) GetDirty(id akara.EID) (*DirtyComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*DirtyComponent), found
}

// Dirty is a convenient reference to be used as a component identifier
var Dirty = newDirty() // nolint:gochecknoglobals // global by design

func newDirty() akara.Component {
	return &DirtyComponent{
		BaseComponent: akara.NewBaseComponent(DirtyCID, newDirty, newDirtyMap),
	}
}

func newDirtyMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(DirtyCID, newDirty, newDirtyMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &DirtyMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
