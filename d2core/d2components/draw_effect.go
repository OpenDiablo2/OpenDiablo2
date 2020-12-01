package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/gravestench/akara"
)

// static check that DrawEffect implements Component
var _ akara.Component = &DrawEffect{}

// DrawEffect is a flag component that is used to denote a "dirty" state
type DrawEffect struct {
	DrawEffect d2enum.DrawEffect
}

// New creates a new DrawEffect. By default, IsDrawEffect is false.
func (*DrawEffect) New() akara.Component {
	return &DrawEffect{}
}

// DrawEffectFactory is a wrapper for the generic component factory that returns DrawEffect component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a DrawEffect.
type DrawEffectFactory struct {
	DrawEffect *akara.ComponentFactory
}

// AddDrawEffect adds a DrawEffect component to the given entity and returns it
func (m *DrawEffectFactory) AddDrawEffect(id akara.EID) *DrawEffect {
	return m.DrawEffect.Add(id).(*DrawEffect)
}

// GetDrawEffect returns the DrawEffect component for the given entity, and a bool for whether or not it exists
func (m *DrawEffectFactory) GetDrawEffect(id akara.EID) (*DrawEffect, bool) {
	component, found := m.DrawEffect.Get(id)
	if !found {
		return nil, found
	}

	return component.(*DrawEffect), found
}
