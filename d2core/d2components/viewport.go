//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom"
)

// static check that ViewportComponent implements Component
var _ akara.Component = &ViewportComponent{}

// static check that ViewportMap implements ComponentMap
var _ akara.ComponentMap = &ViewportMap{}

// ViewportComponent represents the size and position of a scene viewport. This is used
// to control where on screen a viewport is rendered.
type ViewportComponent struct {
	*akara.BaseComponent
	*d2geom.Rectangle
}

// ViewportMap is a map of entity ID's to Viewport
type ViewportMap struct {
	*akara.BaseComponentMap
}

// AddViewport adds a new ViewportComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *ViewportComponent instead of an akara.Component
func (cm *ViewportMap) AddViewport(id akara.EID) *ViewportComponent {
	c := cm.Add(id).(*ViewportComponent)

	const defaultWidth, defaultHeight = 800, 600

	c.Rectangle = &d2geom.Rectangle{
		Left:   0,
		Top:    0,
		Width:  defaultWidth,
		Height: defaultHeight,
	}

	return c
}

// GetViewport returns the ViewportComponent associated with the given entity id
func (cm *ViewportMap) GetViewport(id akara.EID) (*ViewportComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*ViewportComponent), found
}

// Viewport is a convenient reference to be used as a component identifier
var Viewport = newViewport() // nolint:gochecknoglobals // global by design

func newViewport() akara.Component {
	return &ViewportComponent{
		BaseComponent: akara.NewBaseComponent(ViewportCID, newViewport, newViewportMap),
	}
}

func newViewportMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(ViewportCID, newViewport, newViewportMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &ViewportMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
