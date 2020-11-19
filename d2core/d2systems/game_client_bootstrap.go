package d2systems

import (
	"errors"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	logPrefixGameClientBootstrap = "Game Client Bootstrap"
)

// static check that the game config system implements the system interface
var _ akara.System = &GameClientBootstrapSystem{}

func NewGameClientBootstrapSystem() *GameClientBootstrapSystem {
	// we are interested in actual game config instances, too
	gameConfigs := akara.NewFilter().Require(d2components.GameConfig).Build()

	sys := &GameClientBootstrapSystem{
		BaseSubscriberSystem: akara.NewBaseSubscriberSystem(gameConfigs),
		Logger: d2util.NewLogger(),
	}

	sys.SetPrefix(logPrefixGameClientBootstrap)

	return sys
}

// GameClientBootstrapSystem is responsible for setting up the regular diablo2 game launch
type GameClientBootstrapSystem struct {
	*akara.BaseSubscriberSystem
	*d2util.Logger
	*RenderSystem
}

func (m *GameClientBootstrapSystem) Init(world *akara.World) {
	m.Info("initializing ...")

	m.injectSystems()

	m.Info("game client bootstrap complete, deactivating")
	m.SetActive(false)

	if err := m.RenderSystem.Loop(); err != nil {
		panic(err)
	}
}

func (m *GameClientBootstrapSystem) injectSystems() {
	m.RenderSystem = NewRenderSystem()

	m.World.AddSystem(NewUpdateCounterSystem())
	m.World.AddSystem(NewMainMenuScene())
	m.World.AddSystem(m.RenderSystem)

	loadAttempts := 10

	for loadAttempts > 0 {
		m.World.Update(0)

		loadAttempts--

		if m.RenderSystem.renderer != nil {
			return // we've loaded the config, everything is cool
		}
	}

	panic(errors.New("could not initialize renderer"))
}

func (m *GameClientBootstrapSystem) Update() {}
