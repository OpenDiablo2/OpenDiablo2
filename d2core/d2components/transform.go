//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

// static check that Transform implements Component
var _ akara.Component = &Transform{}

// Transform contains a vec3 for Translation, Rotation, and Scale
type Transform struct {
	Translation *d2math.Vector3
	Rotation *d2math.Vector3
	Scale *d2math.Vector3
}

func (t *Transform) GetMatrix() *d2math.Matrix4 {
	return d2math.NewMatrix4(nil).
		Translate(t.Translation).
		RotateX(t.Rotation.X).
		RotateY(t.Rotation.Y).
		RotateZ(t.Rotation.Z).
		ScaleXYZ(t.Scale.XYZ())
}

// New creates a new Transform. By default, the position is (0,0)
func (*Transform) New() akara.Component {
	return &Transform{
		Translation: d2math.NewVector3(0, 0, 0),
		Rotation: d2math.NewVector3(0, 0, 0),
		Scale: d2math.NewVector3(1, 1, 1),
	}
}

// TransformFactory is a wrapper for the generic component factory that returns Transform component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Transform.
type TransformFactory struct {
	*akara.ComponentFactory
}

// Add adds a Transform component to the given entity and returns it
func (m *TransformFactory) Add(id akara.EID) *Transform {
	return m.ComponentFactory.Add(id).(*Transform)
}

// Get returns the Transform component for the given entity, and a bool for whether or not it exists
func (m *TransformFactory) Get(id akara.EID) (*Transform, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Transform), found
}
