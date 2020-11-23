//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// static check that RenderableComponent implements Component
var _ akara.Component = &RenderableComponent{}

// static check that RenderableMap implements ComponentMap
var _ akara.ComponentMap = &RenderableMap{}

// RenderableComponent is a component that contains an embedded surface interface, which is used for rendering
type RenderableComponent struct {
	*akara.BaseComponent
	d2interface.Surface
}

// RenderableMap is a map of entity ID's to Renderable
type RenderableMap struct {
	*akara.BaseComponentMap
}

// AddRenderable adds a new RenderableComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *RenderableComponent instead of an akara.Component
func (cm *RenderableMap) AddRenderable(id akara.EID) *RenderableComponent {
	return cm.Add(id).(*RenderableComponent)
}

// GetRenderable returns the RenderableComponent associated with the given entity id
func (cm *RenderableMap) GetRenderable(id akara.EID) (*RenderableComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*RenderableComponent), found
}

// Renderable is a convenient reference to be used as a component identifier
var Renderable = newRenderable() // nolint:gochecknoglobals // global by design

func newRenderable() akara.Component {
	return &RenderableComponent{
		BaseComponent: akara.NewBaseComponent(RenderableCID, newRenderable, newRenderableMap),
	}
}

func newRenderableMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(RenderableCID, newRenderable, newRenderableMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &RenderableMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
