package d2systems

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

const (
	sceneKeyMouseCursor = "Mouse Cursor"
)

// NewMouseCursorScene creates a new main menu scene. This is the first screen that the user
// will see when launching the game.
func NewMouseCursorScene() *MouseCursorScene {
	scene := &MouseCursorScene{
		BaseScene: NewBaseScene(sceneKeyMouseCursor),
	}

	return scene
}

// static check that MouseCursorScene implements the scene interface
var _ d2interface.Scene = &MouseCursorScene{}

// MouseCursorScene represents the game's main menu, where users can select single or multi player,
// or start the map engine test.
type MouseCursorScene struct {
	*BaseScene
	booted bool
	cursor akara.EID
}

// Init the main menu scene
func (s *MouseCursorScene) Init(world *akara.World) {
	s.World = world

	s.Info("initializing ...")
}

func (s *MouseCursorScene) boot() {
	if !s.BaseScene.booted {
		s.BaseScene.boot()
		return
	}

	s.createMouseCursor()

	s.booted = true
}

func (s *MouseCursorScene) createMouseCursor() {
	s.Info("creating mouse cursor")
	s.cursor = s.Add.Sprite(0, 0, d2resource.CursorDefault, d2resource.PaletteUnits)
}

// Update the main menu scene
func (s *MouseCursorScene) Update() {
	if s.Paused() {
		return
	}

	if !s.booted {
		s.boot()
		return
	}

	s.updateCursorPosition()

	s.BaseScene.Update()
}

func (s *MouseCursorScene) updateCursorPosition() {
	spritePosition, found := s.GetPosition(s.cursor)
	if !found {
		return
	}

	cx, cy := s.systems.InputSystem.renderer.GetCursorPos()

	spritePosition.Set(float64(cx), float64(cy))
}
