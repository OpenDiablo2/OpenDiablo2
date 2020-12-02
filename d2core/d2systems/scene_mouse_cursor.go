package d2systems

import (
	"math"
	"time"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

const (
	sceneKeyMouseCursor = "Mouse Cursor"
)

const (
	fadeTimeout = time.Second * 4
	fadeTime    = time.Second
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
	booted        bool
	cursor        akara.EID
	lastTimeMoved time.Time
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
	for _, id := range s.Viewports {
		s.AddPriority(id).Priority = scenePriorityMouseCursor
	}

	if s.Paused() {
		return
	}

	if !s.booted {
		s.boot()
	}

	s.updateCursorPosition()
	s.handleCursorFade()

	s.BaseScene.Update()
}

func (s *MouseCursorScene) updateCursorPosition() {
	position, found := s.GetPosition(s.cursor)
	if !found {
		return
	}

	cx, cy := s.CursorPosition()

	if int(position.X) != cx || int(position.Y) != cy {
		s.lastTimeMoved = time.Now()
	}

	position.X, position.Y = float64(cx), float64(cy)
}

func (s *MouseCursorScene) handleCursorFade() {
	alpha, found := s.GetAlpha(s.cursor)
	if !found {
		return
	}

	shouldFadeOut := time.Now().Sub(s.lastTimeMoved) > fadeTimeout

	if shouldFadeOut {
		alpha.Alpha = math.Max(alpha.Alpha*0.825, 0)
	} else {
		alpha.Alpha = math.Min(alpha.Alpha+0.125, 1)
	}
}
