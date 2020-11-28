//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
)

const (
	defaultCameraWidth  = 800
	defaultCameraHeight = 600
	defaultCameraZoom   = 1.0
)

// static check that Camera implements Component
var _ akara.Component = &Camera{}

// Camera represents a camera that can be rendered to
type Camera struct {
	*d2vector.Vector
	Width  int
	Height int
	Zoom   float64
}

// New returns a new Camera component.
// The camera defaults to position (0,0), 800x600 resolution, and zoom of 1.0
func (*Camera) New() akara.Component {
	return &Camera{
		Vector: d2vector.NewVector(0, 0),
		Width:  defaultCameraWidth,
		Height: defaultCameraHeight,
		Zoom:   defaultCameraZoom,
	}
}

// CameraFactory is a wrapper for the generic component factory that returns Camera component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Camera.
type CameraFactory struct {
	Camera *akara.ComponentFactory
}

// AddCamera adds a Camera component to the given entity and returns it
func (m *CameraFactory) AddCamera(id akara.EID) *Camera {
	return m.Camera.Add(id).(*Camera)
}

// GetCamera returns the Camera component for the given entity, and a bool for whether or not it exists
func (m *CameraFactory) GetCamera(id akara.EID) (*Camera, bool) {
	component, found := m.Camera.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Camera), found
}
