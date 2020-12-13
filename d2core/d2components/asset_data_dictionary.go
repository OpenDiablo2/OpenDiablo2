//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// static check that DataDictionaryComponent implements Component
var _ akara.Component = &DataDictionary{}

// DataDictionary is a component that contains an embedded txt data dictionary struct
type DataDictionary struct {
	*d2txt.DataDictionary
}

// New returns a DataDictionary component. By default, it contains a nil instance.
func (*DataDictionary) New() akara.Component {
	return &AnimationData{}
}

// DataDictionaryFactory is a wrapper for the generic component factory that returns DataDictionary component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a DataDictionary.
type DataDictionaryFactory struct {
	*akara.ComponentFactory
}

// Add adds a DataDictionary component to the given entity and returns it
func (m *DataDictionaryFactory) Add(id akara.EID) *DataDictionary {
	return m.ComponentFactory.Add(id).(*DataDictionary)
}

// Get returns the DataDictionary component for the given entity, and a bool for whether or not it exists
func (m *DataDictionaryFactory) Get(id akara.EID) (*DataDictionary, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*DataDictionary), found
}
