package d2systems

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
	"github.com/gravestench/akara"
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
		SubscriberSystem: akara.NewSubscriberSystem(gameConfigs),
		Logger: d2util.NewLogger(),
	}

	sys.SetPrefix(logPrefixGameClientBootstrap)

	return sys
}

// GameClientBootstrapSystem is responsible for setting up the regular diablo2 game launch
type GameClientBootstrapSystem struct {
	*akara.SubscriberSystem
	*d2util.Logger
}

func (m *GameClientBootstrapSystem) Init(world *akara.World) {
	m.World = world

	if world == nil {
		m.SetActive(false)
		return
	}

	m.Info("initializing ...")

	m.injectSystems()

	m.SetActive(false)
}

func (m *GameClientBootstrapSystem) injectSystems() {
	m.World.AddSystem(NewMainMenuScene())
}

func (m *GameClientBootstrapSystem) Update() {}
