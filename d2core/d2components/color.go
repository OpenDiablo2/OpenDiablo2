//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"image/color"

	"github.com/gravestench/akara"
)

// static check that Color implements Component
var _ akara.Component = &Color{}

// Color is a flag component that is used to denote a "dirty" state
type Color struct {
	color.Color
}

// New creates a new Color. By default, IsColor is false.
func (*Color) New() akara.Component {
	return &Color{
		color.Transparent,
	}
}

// ColorFactory is a wrapper for the generic component factory that returns Color component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Color.
type ColorFactory struct {
	Color *akara.ComponentFactory
}

// AddColor adds a Color component to the given entity and returns it
func (m *ColorFactory) AddColor(id akara.EID) *Color {
	return m.Color.Add(id).(*Color)
}

// GetColor returns the Color component for the given entity, and a bool for whether or not it exists
func (m *ColorFactory) GetColor(id akara.EID) (*Color, bool) {
	component, found := m.Color.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Color), found
}
