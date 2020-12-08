//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2tbl"
)

// static check that StringTable implements Component
var _ akara.Component = &StringTable{}

// StringTable is a component that contains an embedded text table struct
type StringTable struct {
	*d2tbl.TextDictionary
}

// New returns a new StringTable component. By default, it contains a nil instance.
func (*StringTable) New() akara.Component {
	return &StringTable{}
}

// StringTableFactory is a wrapper for the generic component factory that returns StringTable component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a StringTable.
type StringTableFactory struct {
	*akara.ComponentFactory
}

// Add adds a StringTable component to the given entity and returns it
func (m *StringTableFactory) Add(id akara.EID) *StringTable {
	return m.ComponentFactory.Add(id).(*StringTable)
}

// Get returns the StringTable component for the given entity, and a bool for whether or not it exists
func (m *StringTableFactory) Get(id akara.EID) (*StringTable, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*StringTable), found
}
