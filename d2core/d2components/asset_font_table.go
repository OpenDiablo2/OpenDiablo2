//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that FontTable implements Component
var _ akara.Component = &FontTable{}

// FontTable is a component that contains font table data as a byte slice
type FontTable struct {
	Data []byte
}

// New returns a FontTable component. By default, Data is a nil instance.
func (*FontTable) New() akara.Component {
	return &FontTable{}
}

// FontTableFactory is a wrapper for the generic component factory that returns FontTable component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a FontTable.
type FontTableFactory struct {
	*akara.ComponentFactory
}

// Add adds a FontTable component to the given entity and returns it
func (m *FontTableFactory) Add(id akara.EID) *FontTable {
	return m.ComponentFactory.Add(id).(*FontTable)
}

// Get returns the FontTable component for the given entity, and a bool for whether or not it exists
func (m *FontTableFactory) Get(id akara.EID) (*FontTable, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*FontTable), found
}
