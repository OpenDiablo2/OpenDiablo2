//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that FontTableComponent implements Component
var _ akara.Component = &FontTableComponent{}

// static check that FontTableMap implements ComponentMap
var _ akara.ComponentMap = &FontTableMap{}

// FontTableComponent is a component that contains font table data as a byte slice
type FontTableComponent struct {
	*akara.BaseComponent
	Data []byte
}

// FontTableMap is a map of entity ID's to FontTable
type FontTableMap struct {
	*akara.BaseComponentMap
}

// AddFontTable adds a new FontTableComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *FontTableComponent instead of an akara.Component
func (cm *FontTableMap) AddFontTable(id akara.EID) *FontTableComponent {
	return cm.Add(id).(*FontTableComponent)
}

// GetFontTable returns the FontTableComponent associated with the given entity id
func (cm *FontTableMap) GetFontTable(id akara.EID) (*FontTableComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*FontTableComponent), found
}

// FontTable is a convenient reference to be used as a component identifier
var FontTable = newFontTable() // nolint:gochecknoglobals // global by design

func newFontTable() akara.Component {
	return &FontTableComponent{
		BaseComponent: akara.NewBaseComponent(AssetFontTableCID, newFontTable, newFontTableMap),
	}
}

func newFontTableMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(AssetFontTableCID, newFontTable, newFontTableMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &FontTableMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
