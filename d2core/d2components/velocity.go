//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
)

// static check that VelocityComponent implements Component
var _ akara.Component = &VelocityComponent{}

// static check that VelocityMap implements ComponentMap
var _ akara.ComponentMap = &VelocityMap{}

// VelocityComponent contains an embedded velocity as a vector
type VelocityComponent struct {
	*akara.BaseComponent
	*d2vector.Vector
}

// VelocityMap is a map of entity ID's to Velocity
type VelocityMap struct {
	*akara.BaseComponentMap
}

// AddVelocity adds a new VelocityComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *VelocityComponent instead of an akara.Component
func (cm *VelocityMap) AddVelocity(id akara.EID) *VelocityComponent {
	c := cm.Add(id).(*VelocityComponent)

	c.Vector = d2vector.NewVector(0, 0)

	return c
}

// GetVelocity returns the VelocityComponent associated with the given entity id
func (cm *VelocityMap) GetVelocity(id akara.EID) (*VelocityComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*VelocityComponent), found
}

// Velocity is a convenient reference to be used as a component identifier
var Velocity = newVelocity() // nolint:gochecknoglobals // global by design

func newVelocity() akara.Component {
	return &VelocityComponent{
		BaseComponent: akara.NewBaseComponent(VelocityCID, newVelocity, newVelocityMap),
	}
}

func newVelocityMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(VelocityCID, newVelocity, newVelocityMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &VelocityMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
