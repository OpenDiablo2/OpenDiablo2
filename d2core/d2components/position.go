//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
)

// static check that PositionComponent implements Component
var _ akara.Component = &PositionComponent{}

// static check that PositionMap implements ComponentMap
var _ akara.ComponentMap = &PositionMap{}

// PositionComponent contains an embedded d2vector.Position
type PositionComponent struct {
	*akara.BaseComponent
	*d2vector.Position
}

// PositionMap is a map of entity ID's to Position
type PositionMap struct {
	*akara.BaseComponentMap
}

// AddPosition adds a new PositionComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *PositionComponent instead of an akara.Component
func (cm *PositionMap) AddPosition(id akara.EID) *PositionComponent {
	pos := cm.Add(id).(*PositionComponent)

	p := d2vector.NewPosition(0, 0)

	pos.Position = &p

	return pos
}

// GetPosition returns the PositionComponent associated with the given entity id
func (cm *PositionMap) GetPosition(id akara.EID) (*PositionComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*PositionComponent), found
}

// Position is a convenient reference to be used as a component identifier
var Position = newPosition() // nolint:gochecknoglobals // global by design

func newPosition() akara.Component {
	return &PositionComponent{
		BaseComponent: akara.NewBaseComponent(PositionCID, newPosition, newPositionMap),
	}
}

func newPositionMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(PositionCID, newPosition, newPositionMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &PositionMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
