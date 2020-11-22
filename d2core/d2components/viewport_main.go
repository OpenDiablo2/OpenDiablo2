package d2components

import (
	"github.com/gravestench/akara"
)

// static check that MainViewportComponent implements Component
var _ akara.Component = &MainViewportComponent{}

// static check that MainViewportMap implements ComponentMap
var _ akara.ComponentMap = &MainViewportMap{}

// MainViewportComponent is used to flag viewports as the main viewport of a scene
type MainViewportComponent struct {
	*akara.BaseComponent
}

// MainViewportMap is a map of entity ID's to MainViewport
type MainViewportMap struct {
	*akara.BaseComponentMap
}

// AddMainViewport adds a new MainViewportComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *MainViewportComponent instead of an akara.Component
func (cm *MainViewportMap) AddMainViewport(id akara.EID) *MainViewportComponent {
	return cm.Add(id).(*MainViewportComponent)
}

// GetMainViewport returns the MainViewportComponent associated with the given entity id
func (cm *MainViewportMap) GetMainViewport(id akara.EID) (*MainViewportComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*MainViewportComponent), found
}

// MainViewport is a convenient reference to be used as a component identifier
var MainViewport = newMainViewport() // nolint:gochecknoglobals // global by design

func newMainViewport() akara.Component {
	return &MainViewportComponent{
		BaseComponent: akara.NewBaseComponent(MainViewportCID, newMainViewport, newMainViewportMap),
	}
}

func newMainViewportMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(MainViewportCID, newMainViewport, newMainViewportMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &MainViewportMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
