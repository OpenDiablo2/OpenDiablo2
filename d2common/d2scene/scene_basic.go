package d2scene

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2events"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const MainCamera = 0

func NewBasicScene(world *akara.World) *BasicScene {
	scene := &BasicScene{
		World:      world,
		Events:     d2events.NewEventEmitter(),
		Cameras:    make([]akara.EID, 0),
		Components: NewBasicSceneComponents(world),
	}

	scene.Add = NewSceneObjectFactory(world, scene)
	scene.createMainCamera()

	return scene
}

type BasicScene struct {
	Key        string
	World      *akara.World
	Components *BasicSceneComponents
	Events     *d2events.EventEmitter
	Cameras    []akara.EID
	Camera     *d2components.CameraComponent
	Viewport   *d2components.ViewPortComponent
	Add        *SceneObjectFactory
	Active     bool
	Paused     bool
	booted     bool
}

func (s *BasicScene) Init() {
	if s.World == nil {
		return
	}

	s.Active = true
	s.booted = true
}

func (s *BasicScene) createMainCamera() {
	cameraEID := s.World.NewEntity()

	mainCamera := s.Components.AddCamera(cameraEID)
	mainViewport := s.Components.AddViewPort(cameraEID)

	if len(s.Cameras) < 1 {
		s.Cameras = append(s.Cameras, cameraEID)
	}

	s.Cameras[MainCamera] = cameraEID

	s.Camera = mainCamera
	s.Viewport = mainViewport
}
