//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2tbl"
)

// static check that StringTableComponent implements Component
var _ akara.Component = &StringTableComponent{}

// static check that StringTableMap implements ComponentMap
var _ akara.ComponentMap = &StringTableMap{}

// StringTableComponent is a component that contains an embedded text table struct
type StringTableComponent struct {
	*akara.BaseComponent
	*d2tbl.TextDictionary
}

// StringTableMap is a map of entity ID's to StringTable
type StringTableMap struct {
	*akara.BaseComponentMap
}

// AddStringTable adds a new StringTableComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *StringTableComponent instead of an akara.Component
func (cm *StringTableMap) AddStringTable(id akara.EID) *StringTableComponent {
	return cm.Add(id).(*StringTableComponent)
}

// GetStringTable returns the StringTableComponent associated with the given entity id
func (cm *StringTableMap) GetStringTable(id akara.EID) (*StringTableComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*StringTableComponent), found
}

// StringTable is a convenient reference to be used as a component identifier
var StringTable = newStringTable() // nolint:gochecknoglobals // global by design

func newStringTable() akara.Component {
	return &StringTableComponent{
		BaseComponent: akara.NewBaseComponent(AssetStringTableCID, newStringTable, newStringTableMap),
	}
}

func newStringTableMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(AssetStringTableCID, newStringTable, newStringTableMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &StringTableMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
