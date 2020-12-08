//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// static check that Palette implements Component
var _ akara.Component = &Palette{}

// Palette is a component that contains an embedded palette interface
type Palette struct {
	d2interface.Palette
}

// New returns a new Palette component. By default, it contains a nil instance.
func (*Palette) New() akara.Component {
	return &Palette{}
}

// PaletteFactory is a wrapper for the generic component factory that returns Palette component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Palette.
type PaletteFactory struct {
	*akara.ComponentFactory
}

// Add adds a Palette component to the given entity and returns it
func (m *PaletteFactory) Add(id akara.EID) *Palette {
	return m.ComponentFactory.Add(id).(*Palette)
}

// Get returns the Palette component for the given entity, and a bool for whether or not it exists
func (m *PaletteFactory) Get(id akara.EID) (*Palette, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Palette), found
}
