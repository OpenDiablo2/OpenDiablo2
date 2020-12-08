package d2systems

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

const (
	logPrefixGameObjectFactory = "Object Factory"
)

// static check that GameObjectFactory implements the System interface
var _ akara.System = &GameObjectFactory{}

// GameObjectFactory is a wrapper system for subordinate sceneSystems that
// do the actual object creation work.
type GameObjectFactory struct {
	akara.BaseSystem
	*d2util.Logger
	*SpriteFactory
	*ShapeSystem
	*UIWidgetFactory
}

// Init will initialize the Game Object Factory by injecting all of the factory subsystems into the world
func (t *GameObjectFactory) Init(world *akara.World) {
	t.World = world

	t.setupLogger()

	t.Debug("initializing ...")

	t.injectSubSystems()
}

func (t *GameObjectFactory) setupLogger() {
	t.Logger = d2util.NewLogger()
	t.SetPrefix(logPrefixGameObjectFactory)
}

func (t *GameObjectFactory) injectSubSystems() {
	t.Debug("creating sprite factory")
	t.SpriteFactory = NewSpriteFactory(t.BaseSystem, t.Logger)
	t.ShapeSystem = NewShapeSystem(t.BaseSystem, t.Logger)
	t.UIWidgetFactory = NewUIWidgetFactory(t.BaseSystem, t.Logger, t.SpriteFactory, t.ShapeSystem)
}

// Update updates all the sub-sceneSystems
func (t *GameObjectFactory) Update() {
	t.SpriteFactory.Update()
	t.ShapeSystem.Update()
	t.UIWidgetFactory.Update()
}
