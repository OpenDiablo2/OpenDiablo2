//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

const (
	defaultCameraWidth  = 800
	defaultCameraHeight = 600
	defaultCameraNear   = -100
	defaultCameraFar    = 100
	defaultCameraZ      = -200
)

// static check that Camera implements Component
var _ akara.Component = &Camera{}

// Camera represents a camera that can be rendered to
type Camera struct {
	PerspectiveMatrix *d2math.Matrix4
	OrthogonalMatrix  *d2math.Matrix4
	Size              *d2math.Vector2
	Clip              *d2math.Vector2
}

// New returns a new Camera component.
// The camera defaults to position (0,0), 800x600 resolution, and zoom of 1.0
func (*Camera) New() akara.Component {
	c := &Camera{
		Size:     d2math.NewVector2(defaultCameraWidth, defaultCameraHeight),
		Clip:     d2math.NewVector2(defaultCameraNear, defaultCameraFar),
	}

	w, h := c.Size.XY()
	n, f := c.Clip.XY()

	c.PerspectiveMatrix = d2math.NewMatrix4(nil).PerspectiveLH(w, h, n, f)

	l, r, t, b := -(w / 2), w/2, -(h / 2), h/2

	c.OrthogonalMatrix = d2math.NewMatrix4(nil).Ortho(l, r, t, b, n, f)

	return c
}

// CameraFactory is a wrapper for the generic component factory that returns Camera component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Camera.
type CameraFactory struct {
	*akara.ComponentFactory
}

// Add adds a Camera component to the given entity and returns it
func (m *CameraFactory) Add(id akara.EID) *Camera {
	return m.ComponentFactory.Add(id).(*Camera)
}

// Get returns the Camera component for the given entity, and a bool for whether or not it exists
func (m *CameraFactory) Get(id akara.EID) (*Camera, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Camera), found
}
