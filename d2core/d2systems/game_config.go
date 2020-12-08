package d2systems

import (
	"encoding/json"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

// static check that the game config system implements the system interface
var _ akara.System = &GameConfigSystem{}

const (
	loggerPrefixGameConfig = "Game Config"
)

// GameConfigSystem is responsible for game config configFileBootstrap procedure, as well as
// clearing the `Dirty` component of game configs. In the `configFileBootstrap` method of this system
// you can see that this system will add entities for the directories it expects config files
// to be found in, and it also adds an entity for the initial config file to be loaded.
//
// This system is dependant on the FileTypeResolver, FileSourceResolver, and
// FileHandleResolver sceneSystems because this system subscribes to entities
// with components created by these other sceneSystems. Nothing will  break if these
// other sceneSystems are not present in the world, but no config files will be loaded by
// this system either...
type GameConfigSystem struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
	filesToCheck *akara.Subscription
	gameConfigs  *akara.Subscription
	Components struct {
		GameConfig d2components.GameConfigFactory
		File d2components.FileFactory
		FileType d2components.FileTypeFactory
		FileHandle d2components.FileHandleFactory
		FileSource d2components.FileSourceFactory
		Dirty d2components.DirtyFactory
	}
	activeConfig *d2components.GameConfig
}

// Init the world with the necessary components related to game config files
func (m *GameConfigSystem) Init(world *akara.World) {
	m.World = world

	m.setupLogger()

	m.Debug("initializing ...")

	m.setupFactories()
	m.setupSubscriptions()
}

func (m *GameConfigSystem) setupLogger() {
	m.Logger = d2util.NewLogger()
	m.SetPrefix(loggerPrefixGameConfig)
}

func (m *GameConfigSystem) setupFactories() {
	m.Debug("setting up component factories")

	m.InjectComponent(&d2components.File{}, &m.Components.File.ComponentFactory)
	m.InjectComponent(&d2components.FileType{}, &m.Components.FileType.ComponentFactory)
	m.InjectComponent(&d2components.FileHandle{}, &m.Components.FileHandle.ComponentFactory)
	m.InjectComponent(&d2components.FileSource{}, &m.Components.FileSource.ComponentFactory)
	m.InjectComponent(&d2components.GameConfig{}, &m.Components.GameConfig.ComponentFactory)
	m.InjectComponent(&d2components.Dirty{}, &m.Components.Dirty.ComponentFactory)
}

func (m *GameConfigSystem) setupSubscriptions() {
	m.Debug("setting up component subscriptions")

	// we are going to check entities that dont yet have loaded asset types
	filesToCheck := m.NewComponentFilter().
		Require(
			&d2components.File{},
			&d2components.FileType{},
			&d2components.FileHandle{},
		).
		Forbid(
			&d2components.GameConfig{},
			&d2components.StringTable{},
			&d2components.DataDictionary{},
			&d2components.Palette{},
			&d2components.PaletteTransform{},
			&d2components.Cof{},
			&d2components.Dc6{},
			&d2components.Dcc{},
			&d2components.Ds1{},
			&d2components.Dt1{},
			&d2components.Wav{},
			&d2components.AnimationData{},
		).
		Build()

	// we are interested in actual game config instances, too
	gameConfigs := m.NewComponentFilter().
		Require(&d2components.GameConfig{}).
		Build()

	m.filesToCheck = m.AddSubscription(filesToCheck)
	m.gameConfigs = m.AddSubscription(gameConfigs)
}

// Update checks for new config files
func (m *GameConfigSystem) Update() {
	m.checkForNewConfig(m.filesToCheck.GetEntities())
}

func (m *GameConfigSystem) checkForNewConfig(entities []akara.EID) {
	for _, eid := range entities {
		fp, found := m.Components.File.Get(eid)
		if !found {
			continue
		}

		ft, found := m.Components.FileType.Get(eid)
		if !found {
			continue
		}

		if fp.Path != configFileName || ft.Type != d2enum.FileTypeJSON {
			continue
		}

		m.Info("loading config file ...")
		m.loadConfig(eid)
	}
}

func (m *GameConfigSystem) loadConfig(eid akara.EID) {
	fh, found := m.Components.FileHandle.Get(eid)
	if !found {
		return
	}

	gameConfig := m.Components.GameConfig.Add(eid)

	if err := json.NewDecoder(fh.Data).Decode(gameConfig); err != nil {
		m.RemoveEntity(eid)
	}

	m.activeConfig = gameConfig
}
