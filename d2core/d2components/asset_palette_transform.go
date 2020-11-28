//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2pl2"
)

// static check that PaletteTransform implements Component
var _ akara.Component = &PaletteTransform{}

// PaletteTransform is a component that contains an embedded palette transform (pl2) struct
type PaletteTransform struct {
	*d2pl2.PL2
}

// New returns a new PaletteTransform component. By default, it contains a nil instance.
func (*PaletteTransform) New() akara.Component {
	return &PaletteTransform{}
}

// PaletteTransformFactory is a wrapper for the generic component factory that returns PaletteTransform component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a PaletteTransform.
type PaletteTransformFactory struct {
	PaletteTransform *akara.ComponentFactory
}

// AddPaletteTransform adds a PaletteTransform component to the given entity and returns it
func (m *PaletteTransformFactory) AddPaletteTransform(id akara.EID) *PaletteTransform {
	return m.PaletteTransform.Add(id).(*PaletteTransform)
}

// GetPaletteTransform returns the PaletteTransform component for the given entity, and a bool for whether or not it exists
func (m *PaletteTransformFactory) GetPaletteTransform(id akara.EID) (*PaletteTransform, bool) {
	component, found := m.PaletteTransform.Get(id)
	if !found {
		return nil, found
	}

	return component.(*PaletteTransform), found
}
