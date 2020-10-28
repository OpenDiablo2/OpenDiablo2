package d2components

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/gravestench/akara"
)

// static check that CameraComponent implements Component
var _ akara.Component = &CameraComponent{}

// static check that CameraMap implements ComponentMap
var _ akara.ComponentMap = &CameraMap{}

// CameraComponent represents a camera that can be rendered to
type CameraComponent struct {
	*d2vector.Position
	Width  uint
	Height uint
	Zoom   float64
}

// ID returns a unique identifier for the component type
func (*CameraComponent) ID() akara.ComponentID {
	return CameraCID
}

// NewMap returns a new component map for the component type
func (*CameraComponent) NewMap() akara.ComponentMap {
	return NewCameraMap()
}

// Camera is a convenient reference to be used as a component identifier
var Camera = (*CameraComponent)(nil) // nolint:gochecknoglobals // global by design

// NewCameraMap creates a new map of entity ID's to Camera
func NewCameraMap() *CameraMap {
	cm := &CameraMap{
		components: make(map[akara.EID]*CameraComponent),
	}

	return cm
}

// CameraMap is a map of entity ID's to Camera
type CameraMap struct {
	world      *akara.World
	components map[akara.EID]*CameraComponent
}

// Init initializes the component map with the given world
func (cm *CameraMap) Init(world *akara.World) {
	cm.world = world
}

// ID returns a unique identifier for the component type
func (*CameraMap) ID() akara.ComponentID {
	return CameraCID
}

// NewMap returns a new component map for the component type
func (*CameraMap) NewMap() akara.ComponentMap {
	return NewCameraMap()
}

// Add a new CameraComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *CameraMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	position := d2vector.NewPosition(0, 0)
	cm.components[id] = &CameraComponent{Position: &position}

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// AddCamera adds a new CameraComponent for the given entity id and returns it.
// If the entity already has a file type component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *CameraComponent instead of an akara.Component
func (cm *CameraMap) AddCamera(id akara.EID) *CameraComponent {
	return cm.Add(id).(*CameraComponent)
}

// Get returns the component associated with the given entity id
func (cm *CameraMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetCamera returns the CameraComponent associated with the given entity id
func (cm *CameraMap) GetCamera(id akara.EID) (*CameraComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *CameraMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}
