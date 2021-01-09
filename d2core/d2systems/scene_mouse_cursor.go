package d2systems

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"math"
	"strconv"
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
)

// NewMouseCursorScene creates a mouse cursor scene
func NewMouseCursorScene() *MouseCursorScene {
	scene := &MouseCursorScene{
		BaseScene: NewBaseScene(sceneKeyMouseCursor),
	}

	return scene
}

// static check that MouseCursorScene implements the scene interface
var _ d2interface.Scene = &MouseCursorScene{}

// MouseCursorScene is a scene that renders a mouse cursor in the window
type MouseCursorScene struct {
	lastTimeMoved time.Time
	*BaseScene
	cursor akara.EID
	state  d2enum.SceneState
	debug  struct {
		enabled bool
	}
}

// Init does basic scene initialization
func (s *MouseCursorScene) Init(world *akara.World) {
	s.World = world

	s.Debug("initializing ...")
}

func (s *MouseCursorScene) boot() {
	if !s.BaseScene.Booted() {
		s.BaseScene.boot()
		return
	}

	s.registerTerminalCommands()

	s.createMouseCursor()

	s.state = d2enum.SceneStateBooted
}

func (s *MouseCursorScene) createMouseCursor() {
	s.Debug("creating mouse cursor")
	s.cursor = s.Add.Sprite(0, 0, d2resource.CursorDefault, d2resource.PaletteUnits)
}

// Update updates the state of the scene
func (s *MouseCursorScene) Update() {
	for _, id := range s.Viewports {
		s.Components.Priority.Add(id).Priority = scenePriorityMouseCursor
	}

	if s.Paused() {
		return
	}

	if s.state == d2enum.SceneStateUninitialized {
		s.boot()
	}

	if s.state != d2enum.SceneStateBooted {
		return
	}

	s.updateCursorTransform()
	s.handleCursorFade()

	s.BaseScene.Update()
}

func (s *MouseCursorScene) updateCursorTransform() {
	transform, found := s.Components.Transform.Get(s.cursor)
	if !found {
		return
	}

	node, _ := s.Components.SceneGraphNode.Get(s.cursor)
	_ = node

	cx, cy := s.Input.CursorPosition()
	tx, ty := transform.Translation.XY()

	if int(tx) != cx || int(ty) != cy {
		s.lastTimeMoved = time.Now()
	}

	transform.Translation.X, transform.Translation.Y = float64(cx), float64(cy)
}

func (s *MouseCursorScene) handleCursorFade() {
	alpha, found := s.Components.Alpha.Get(s.cursor)
	if !found {
		return
	}

	shouldFadeOut := time.Since(s.lastTimeMoved) > fadeTimeout

	if shouldFadeOut {
		alpha.Alpha = math.Max(alpha.Alpha*0.825, 0) // nolint:gomnd // arbitrary example number for test scene
	} else {
		alpha.Alpha = math.Min(alpha.Alpha+0.125, 1) // nolint:gomnd // arbitrary example number for test scene
	}

	if alpha.Alpha > 1e-1 && alpha.Alpha < 1 { // nolint:gomnd // arbitrary example number for test scene
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
	s.Debug("registering debug command")

	const (
		command     = "debug_mouse"
		description = "show debug information about the mouse"
	)

	s.RegisterTerminalCommand(command, description, []string{"val"}, func(args []string) error {
		val, err := strconv.ParseBool(args[0])
		if err != nil {
			return fmt.Errorf("invalid argument")
		}

		s.setDebug(val)

		return nil
	})
}

func (s *MouseCursorScene) setDebug(val bool) {
	s.debug.enabled = val
}
