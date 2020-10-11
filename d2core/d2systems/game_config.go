package d2systems

import (
	"github.com/gravestench/ecs"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

// static check that the game config system implements the system interface
var _ ecs.System = &GameConfigSystem{}

func NewGameConfigSystem() *GameConfigSystem {
	gameConfigs := ecs.NewFilter().
		Require(d2components.GameConfig).
		Build()

	gcs := &GameConfigSystem{
		SubscriberSystem: ecs.NewSubscriberSystem(gameConfigs),
	}

	return gcs
}

type GameConfigSystem struct {
	*ecs.SubscriberSystem
	configs     *d2components.GameConfigMap
	filePaths   *d2components.FilePathMap
	fileTypes   *d2components.FileTypeMap
	fileHandles *d2components.FileHandleMap
	fileSources *d2components.FileSourceMap
}

func (m *GameConfigSystem) Init(world *ecs.World) {
	m.World = world

	if world == nil {
		m.SetActive(false)
		return
	}

	for subIdx := range m.Subscriptions {
		m.AddSubscription(m.Subscriptions[subIdx])
	}

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.filePaths = world.InjectMap(d2components.FilePath).(*d2components.FilePathMap)
	m.fileTypes = world.InjectMap(d2components.FileType).(*d2components.FileTypeMap)
	m.fileHandles = world.InjectMap(d2components.FileHandle).(*d2components.FileHandleMap)
	m.fileSources = world.InjectMap(d2components.FileSource).(*d2components.FileSourceMap)
}

func (m *GameConfigSystem) Process() {
	for subIdx := range m.Subscriptions {
		for _, EID := range m.Subscriptions[subIdx].GetEntities() {
			_ = EID
		}
	}
}
