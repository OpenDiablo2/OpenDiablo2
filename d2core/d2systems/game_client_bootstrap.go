package d2systems

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/gravestench/akara"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	logPrefixGameClientBootstrap = "Game Client Bootstrap"
)

// static check that the game config system implements the system interface
var _ akara.System = &GameClientBootstrap{}

// GameClientBootstrap is responsible for setting up other
// sceneSystems that are common to both the game client and the headless game server
type GameClientBootstrap struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
}

// Init injects the common sceneSystems required by both the game client and headless server
func (m *GameClientBootstrap) Init(world *akara.World) {
	m.World = world

	m.setupLogger()

	m.Debug("initializing ...")

	m.injectSystems()

	m.Debug("initialization complete")
}

func (m *GameClientBootstrap) setupLogger() {
	m.Logger = d2util.NewLogger()
	m.SetPrefix(logPrefixGameClientBootstrap)
}

func (m *GameClientBootstrap) injectSystems() {
	m.Info("injecting terminal scene")
	m.AddSystem(NewTerminalScene())

	m.Info("injecting timescale scene")
	m.AddSystem(&TimeScaleSystem{})

	m.Info("injecting game object factory system")
	m.AddSystem(&GameObjectFactory{})

	m.Info("injecting loading scene")
	m.AddSystem(NewLoadingScene())

	m.Info("injecting mouse cursor scene")
	m.AddSystem(NewMouseCursorScene())

	skipSplash := kingpin.Flag(skipSplashArg, skipSplashDesc).Bool()
	kingpin.Parse()

	if !*skipSplash {
		m.Info("injecting ebiten splash scene")
		m.AddSystem(NewEbitenSplashScene())
	}

	m.Info("injecting main menu scene")
	m.AddSystem(NewMainMenuScene())
}

// Update does nothing, but exists to satisfy the `akara.System` interface
func (m *GameClientBootstrap) Update() {
	m.Debug("game client bootstrap complete, deactivating")
	m.RemoveSystem(m)
}
