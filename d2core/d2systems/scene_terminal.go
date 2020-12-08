package d2systems

import (
	"image/color"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2term"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

const (
	sceneKeyTerminal = "Terminal"
)

const (
	viewportTerminal = iota + 1
)

// static check that TerminalScene implements the scene interface
var _ d2interface.Scene = &TerminalScene{}

// NewTerminalScene creates a new main menu scene. This is the first screen that the user
// will see when launching the game.
func NewTerminalScene() *TerminalScene {
	scene := &TerminalScene{
		BaseScene: NewBaseScene(sceneKeyTerminal),
	}

	scene.backgroundColor = color.Transparent

	return scene
}

// TerminalScene represents the in-game terminal for typing commands
type TerminalScene struct {
	*BaseScene
	d2interface.Terminal
	d2interface.InputManager
	commandsToRegister *akara.Subscription
	booted             bool
}

// Init the terminal
func (s *TerminalScene) Init(world *akara.World) {
	s.World = world

	s.Debug("initializing ...")

	s.setupSubscriptions()
}

func (s *TerminalScene) setupSubscriptions() {
	s.Debug("setting up component subscriptions")

	commandsToRegister := s.NewComponentFilter().
		Require(
			&d2components.CommandRegistration{},
			&d2components.Dirty{},
		).
		Build()

	s.commandsToRegister = s.World.AddSubscription(commandsToRegister)
}

func (s *TerminalScene) boot() {
	if !s.BaseScene.booted {
		s.BaseScene.boot()
		return
	}

	s.createTerminal()

	s.booted = true
}

// Update and render the terminal to the terminal viewport
func (s *TerminalScene) Update() {
	for _, id := range s.Viewports {
		s.Components.Priority.Add(id).Priority = scenePriorityTerminal
	}

	if s.Paused() {
		return
	}

	if !s.booted {
		s.boot()
	}

	if !s.booted {
		return
	}

	s.processCommandRegistrations()
	s.updateTerminal()

	s.BaseScene.Update()
}

func (s *TerminalScene) processCommandRegistrations() {
	for _, eid := range s.commandsToRegister.GetEntities() {
		s.processCommand(eid)
	}
}

func (s *TerminalScene) processCommand(eid akara.EID) {
	reg, found := s.Components.CommandRegistration.Get(eid)
	if !found {
		return
	}

	s.Infof("Registering command `%s` - %s", reg.Name, reg.Description)

	err := s.Terminal.BindAction(reg.Name, reg.Description, reg.Callback)
	if err != nil {
		s.Error(err.Error())
	}

	s.Components.Dirty.Remove(eid)
}

func (s *TerminalScene) createTerminal() {
	inputManager := d2input.NewInputManager()

	term, err := d2term.New(inputManager)
	if err != nil {
		panic(err)
	}

	s.InputManager = inputManager
	s.Terminal = term

	termVP := s.Add.Viewport(viewportTerminal, 800, 600)

	texture, _ := s.Components.Texture.Get(termVP)
	texture.Texture.Clear(color.Transparent)

	alpha := s.Components.Alpha.Add(termVP)
	alpha.Alpha = 0.5

	vpFilter := s.Components.ViewportFilter.Add(termVP)

	vpFilter.Set(viewportTerminal, true)
}

func (s *TerminalScene) updateTerminal() {
	// this is a hack to use the old-style input manager and terminal
	// stuff inside of this ECS implementation
	_ = s.InputManager.Advance(s.TimeDelta.Seconds(), float64(time.Now().Second()))
	_ = s.Terminal.Advance(s.TimeDelta.Seconds())

	termVP := s.Viewports[viewportTerminal]

	texture, _ := s.Components.Texture.Get(termVP)
	texture.Texture.Clear(color.Transparent)

	_ = s.Terminal.Render(texture.Texture)
}
