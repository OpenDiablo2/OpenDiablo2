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
// FileHandleResolver systems because this system subscribes to entities
// with components created by these other systems. Nothing will  break if these
// other systems are not present in the world, but no config files will be loaded by
// this system either...
type GameConfigSystem struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
	filesToCheck *akara.Subscription
	gameConfigs  *akara.Subscription
	d2components.GameConfigFactory
	d2components.FilePathFactory
	d2components.FileTypeFactory
	d2components.FileHandleFactory
	d2components.FileSourceFactory
	d2components.DirtyFactory
	activeConfig *d2components.GameConfig
}

// Init the world with the necessary components related to game config files
func (m *GameConfigSystem) Init(world *akara.World) {
	m.World = world

	m.setupLogger()

	m.Info("initializing ...")

	m.setupFactories()
	m.setupSubscriptions()
}

func (m *GameConfigSystem) setupLogger() {
	m.Logger = d2util.NewLogger()
	m.SetPrefix(loggerPrefixGameConfig)
}

func (m *GameConfigSystem) setupFactories() {
	m.Info("setting up component factories")

	filePathID := m.RegisterComponent(&d2components.FilePath{})
	fileTypeID := m.RegisterComponent(&d2components.FileType{})
	fileHandleID := m.RegisterComponent(&d2components.FileHandle{})
	fileSourceID := m.RegisterComponent(&d2components.FileSource{})
	gameConfigID := m.RegisterComponent(&d2components.GameConfig{})
	dirtyID := m.RegisterComponent(&d2components.Dirty{})

	m.FilePath = m.GetComponentFactory(filePathID)
	m.FileType = m.GetComponentFactory(fileTypeID)
	m.FileHandle = m.GetComponentFactory(fileHandleID)
	m.FileSource = m.GetComponentFactory(fileSourceID)
	m.GameConfig = m.GetComponentFactory(gameConfigID)
	m.Dirty = m.GetComponentFactory(dirtyID)
}

func (m *GameConfigSystem) setupSubscriptions() {
	m.Info("setting up component subscriptions")

	// we are going to check entities that dont yet have loaded asset types
	filesToCheck := m.NewComponentFilter().
		Require(
			&d2components.FilePath{},
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
		fp, found := m.GetFilePath(eid)
		if !found {
			continue
		}

		ft, found := m.GetFileType(eid)
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
	fh, found := m.GetFileHandle(eid)
	if !found {
		return
	}

	gameConfig := m.AddGameConfig(eid)

	if err := json.NewDecoder(fh.Data).Decode(gameConfig); err != nil {
		m.RemoveEntity(eid)
	}

	m.activeConfig = gameConfig
}
