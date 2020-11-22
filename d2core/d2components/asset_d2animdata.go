//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2animdata"
)

// static check that AnimDataComponent implements Component
var _ akara.Component = &AnimDataComponent{}

// static check that AnimDataMap implements ComponentMap
var _ akara.ComponentMap = &AnimDataMap{}

// AnimDataComponent is a component that contains an embedded AnimationData struct
type AnimDataComponent struct {
	*akara.BaseComponent
	*d2animdata.AnimationData
}

// AnimDataMap is a map of entity ID's to AnimData
type AnimDataMap struct {
	*akara.BaseComponentMap
}

// AddAnimData adds a new AnimDataComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *AnimDataComponent instead of an akara.Component
func (cm *AnimDataMap) AddAnimData(id akara.EID) *AnimDataComponent {
	return cm.Add(id).(*AnimDataComponent)
}

// GetAnimData returns the AnimDataComponent associated with the given entity id
func (cm *AnimDataMap) GetAnimData(id akara.EID) (*AnimDataComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*AnimDataComponent), found
}

// AnimData is a convenient reference to be used as a component identifier
var AnimData = newAnimData() // nolint:gochecknoglobals // global by design

func newAnimData() akara.Component {
	return &AnimDataComponent{
		BaseComponent: akara.NewBaseComponent(AssetD2AnimDataCID, newAnimData, newAnimDataMap),
	}
}

func newAnimDataMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(AssetD2AnimDataCID, newAnimData, newAnimDataMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &AnimDataMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
