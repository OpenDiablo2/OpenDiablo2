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

func NewMouseCursorScene() *MouseCursorScene {
	scene := &MouseCursorScene{
		BaseScene: NewBaseScene(sceneKeyMouseCursor),
	}

	return scene
}

// static check that MouseCursorScene implements the scene interface
var _ d2interface.Scene = &MouseCursorScene{}

type MouseCursorScene struct {
	*BaseScene
	booted        bool
	cursor        akara.EID
	lastTimeMoved time.Time
	debug         struct {
		enabled bool
	}
}

func (s *MouseCursorScene) Init(world *akara.World) {
	s.World = world

	s.Info("initializing ...")
}

func (s *MouseCursorScene) boot() {
	if !s.BaseScene.booted {
		s.BaseScene.boot()
		return
	}

	s.registerTerminalCommands()

	s.createMouseCursor()

	s.booted = true
}

func (s *MouseCursorScene) createMouseCursor() {
	s.Info("creating mouse cursor")
	s.cursor = s.Add.Sprite(0, 0, d2resource.CursorDefault, d2resource.PaletteUnits)
}

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

		switch s.debug.enabled {
		case true:
			s.Infof("position: (%d, %d)", int(position.X), int(position.Y))
		default:
			s.Debugf("position: (%d, %d)", int(position.X), int(position.Y))
		}
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

	if alpha.Alpha > 1e-1 && alpha.Alpha < 1 {
		switch s.debug.enabled {
		case true:
			s.Infof("fading %.2f", alpha.Alpha)
		default:
			s.Debugf("fading %.2f", alpha.Alpha)
		}
	}
}

func (s *MouseCursorScene) registerTerminalCommands() {
	s.registerDebugCommand()
}

func (s *MouseCursorScene) registerDebugCommand() {
	s.Info("registering debug command")

	const (
		command     = "debug_mouse"
		description = "show debug information about the mouse"
	)

	s.RegisterTerminalCommand(command, description, func(val bool) {
		s.setDebug(val)
	})
}

func (s *MouseCursorScene) setDebug(val bool) {
	s.debug.enabled = val
}
