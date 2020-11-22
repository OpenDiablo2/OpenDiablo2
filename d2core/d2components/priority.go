//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that PriorityComponent implements Component
var _ akara.Component = &PriorityComponent{}

// static check that PriorityMap implements ComponentMap
var _ akara.ComponentMap = &PriorityMap{}

// PriorityComponent is a component that is used to add a priority value.
// This can generally be used for sorting entities when order matters.
type PriorityComponent struct {
	*akara.BaseComponent
	Priority int
}

// PriorityMap is a map of entity ID's to Priority
type PriorityMap struct {
	*akara.BaseComponentMap
}

// AddPriority adds a new PriorityComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *PriorityComponent instead of an akara.Component
func (cm *PriorityMap) AddPriority(id akara.EID) *PriorityComponent {
	return cm.Add(id).(*PriorityComponent)
}

// GetPriority returns the PriorityComponent associated with the given entity id
func (cm *PriorityMap) GetPriority(id akara.EID) (*PriorityComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*PriorityComponent), found
}

// Priority is a convenient reference to be used as a component identifier
var Priority = newPriority() // nolint:gochecknoglobals // global by design

func newPriority() akara.Component {
	return &PriorityComponent{
		BaseComponent: akara.NewBaseComponent(PriorityCID, newPriority, newPriorityMap),
	}
}

func newPriorityMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(PriorityCID, newPriority, newPriorityMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &PriorityMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
