package d2systems

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/gravestench/akara"
)

const (
	logPrefixGameObjectFactory = "Object Factory"
)

// NewGameObjectFactory creates a game object factory
func NewGameObjectFactory() *GameObjectFactory {
	m := &GameObjectFactory{
		BaseSystem: &akara.BaseSystem{},
		Logger: d2util.NewLogger(),
	}

	m.SetPrefix(logPrefixGameObjectFactory)

	return m
}

// static check that GameObjectFactory implements the System interface
var _ akara.System = &GameObjectFactory{}

// GameObjectFactory
type GameObjectFactory struct {
	*akara.BaseSystem
	*d2util.Logger
	SpriteFactory *SpriteFactory
}

// Init will initialize the Game Object Factory by injecting all of the factory subsystems into the world
func (t *GameObjectFactory) Init(world *akara.World) {
	t.Info("initializing ...")
	t.injectSubSystems()
}

func (t *GameObjectFactory) injectSubSystems() {
	t.SpriteFactory = NewSpriteFactorySubsystem(t.BaseSystem, t.Logger)
}

func (t *GameObjectFactory) Update() {
	t.SpriteFactory.Update()
}


