//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
)

// static check that CameraComponent implements Component
var _ akara.Component = &CameraComponent{}

// static check that CameraMap implements ComponentMap
var _ akara.ComponentMap = &CameraMap{}

// CameraComponent represents a camera that can be rendered to
type CameraComponent struct {
	*akara.BaseComponent
	*d2vector.Vector
	Width  int
	Height int
	Zoom   float64
}

// CameraMap is a map of entity ID's to Camera
type CameraMap struct {
	*akara.BaseComponentMap
}

// AddCamera adds a new CameraComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *CameraComponent instead of an akara.Component
func (cm *CameraMap) AddCamera(id akara.EID) *CameraComponent {
	camera := cm.Add(id).(*CameraComponent)

	camera.Vector = d2vector.NewVector(0, 0)

	return camera
}

// GetCamera returns the CameraComponent associated with the given entity id
func (cm *CameraMap) GetCamera(id akara.EID) (*CameraComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*CameraComponent), found
}

// Camera is a convenient reference to be used as a component identifier
var Camera = newCamera() // nolint:gochecknoglobals // global by design

func newCamera() akara.Component {
	return &CameraComponent{
		BaseComponent: akara.NewBaseComponent(CameraCID, newCamera, newCameraMap),
	}
}

func newCameraMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(CameraCID, newCamera, newCameraMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &CameraMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
