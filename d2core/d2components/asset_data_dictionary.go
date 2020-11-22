//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// static check that DataDictionaryComponent implements Component
var _ akara.Component = &DataDictionaryComponent{}

// static check that DataDictionaryMap implements ComponentMap
var _ akara.ComponentMap = &DataDictionaryMap{}

// DataDictionaryComponent is a component that contains an embedded txt data dictionary struct
type DataDictionaryComponent struct {
	*akara.BaseComponent
	*d2txt.DataDictionary
}

// DataDictionaryMap is a map of entity ID's to DataDictionary
type DataDictionaryMap struct {
	*akara.BaseComponentMap
}

// AddDataDictionary adds a new DataDictionaryComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *DataDictionaryComponent instead of an akara.Component
func (cm *DataDictionaryMap) AddDataDictionary(id akara.EID) *DataDictionaryComponent {
	return cm.Add(id).(*DataDictionaryComponent)
}

// GetDataDictionary returns the DataDictionaryComponent associated with the given entity id
func (cm *DataDictionaryMap) GetDataDictionary(id akara.EID) (*DataDictionaryComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*DataDictionaryComponent), found
}

// DataDictionary is a convenient reference to be used as a component identifier
var DataDictionary = newDataDictionary() // nolint:gochecknoglobals // global by design

func newDataDictionary() akara.Component {
	return &DataDictionaryComponent{
		BaseComponent: akara.NewBaseComponent(AssetDataDictionaryCID, newDataDictionary, newDataDictionaryMap),
	}
}

func newDataDictionaryMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(AssetDataDictionaryCID, newDataDictionary, newDataDictionaryMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &DataDictionaryMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
