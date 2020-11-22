package d2systems

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	logPrefixGameClientBootstrap = "Game Client Bootstrap"
)

// static check that the game config system implements the system interface
var _ akara.System = &GameClientBootstrapSystem{}

// NewGameClientBootstrapSystem makes a new client bootstrap system
func NewGameClientBootstrapSystem() *GameClientBootstrapSystem {
	// we are interested in actual game config instances, too
	gameConfigs := akara.NewFilter().Require(d2components.GameConfig).Build()

	sys := &GameClientBootstrapSystem{
		BaseSubscriberSystem: akara.NewBaseSubscriberSystem(gameConfigs),
		Logger:               d2util.NewLogger(),
	}

	sys.SetPrefix(logPrefixGameClientBootstrap)

	return sys
}

// GameClientBootstrapSystem is responsible for setting up other
// systems that are common to both the game client and the headless game server
type GameClientBootstrapSystem struct {
	*akara.BaseSubscriberSystem
	*d2util.Logger
	*RenderSystem
}

// Init injects the common systems required by both the game client and headless server
func (m *GameClientBootstrapSystem) Init(_ *akara.World) {
	m.Info("initializing ...")

	m.injectSystems()

	m.Info("game client bootstrap complete, deactivating")
	m.SetActive(false)

	if err := m.World.Update(0); err != nil {
		m.Error(err.Error())
	}
}

func (m *GameClientBootstrapSystem) injectSystems() {
	m.RenderSystem = NewRenderSystem()

	m.World.AddSystem(m.RenderSystem)
	m.World.AddSystem(NewUpdateCounterSystem())
	m.World.AddSystem(NewMainMenuScene())
}

// Update does nothing, but exists to satisfy the `akara.System` interface
func (m *GameClientBootstrapSystem) Update() {
	// nothing to do after init ...
}
