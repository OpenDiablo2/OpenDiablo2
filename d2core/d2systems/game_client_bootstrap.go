package d2systems

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

const (
	logPrefixGameClientBootstrap = "Game Client Bootstrap"
)

// static check that the game config system implements the system interface
var _ akara.System = &GameClientBootstrapSystem{}

// GameClientBootstrapSystem is responsible for setting up other
// baseSystems that are common to both the game client and the headless game server
type GameClientBootstrapSystem struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
}

// Init injects the common baseSystems required by both the game client and headless server
func (m *GameClientBootstrapSystem) Init(world *akara.World) {
	m.World = world

	m.setupLogger()

	m.Info("initializing ...")

	m.injectSystems()

	m.Info("initialization complete")

	if err := m.World.Update(0); err != nil {
		m.Error(err.Error())
	}
}

func (m *GameClientBootstrapSystem) setupLogger() {
	m.Logger = d2util.NewLogger()
	m.SetPrefix(logPrefixGameClientBootstrap)
}

func (m *GameClientBootstrapSystem) injectSystems() {
	m.Info("injecting render system")
	m.AddSystem(&RenderSystem{})

	m.Info("injecting input system")
	m.AddSystem(&InputSystem{})

	m.Info("injecting update counter system")
	m.AddSystem(&UpdateCounter{})

	m.Info("injecting loading scene")
	m.AddSystem(NewLoadingScene())

	m.Info("injecting mouse cursor scene")
	m.AddSystem(NewMouseCursorScene())

	m.Info("injecting main menu scene")
	m.AddSystem(NewMainMenuScene())
}

// Update does nothing, but exists to satisfy the `akara.System` interface
func (m *GameClientBootstrapSystem) Update() {
	m.Info("game client bootstrap complete, deactivating")
	m.RemoveSystem(m)
}
