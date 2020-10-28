package d2systems

import (
	"os"
	"path"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	configDirectoryName = "OpenDiablo2"
	configFileName      = "config.json"
)

// static check that the game config system implements the system interface
var _ akara.System = &GameBootstrapSystem{}

func NewGameBootstrapSystem() *GameBootstrapSystem {
	// we are going to check entities that dont yet have loaded asset types
	thingsToCheck := akara.NewFilter().
		Require(d2components.FilePath).
		Require(d2components.FileType).
		Require(d2components.FileHandle).
		Forbid(d2components.GameConfig).
		Forbid(d2components.StringTable).
		Forbid(d2components.DataDictionary).
		Forbid(d2components.Palette).
		Forbid(d2components.PaletteTransform).
		Forbid(d2components.Cof).
		Forbid(d2components.Dc6).
		Forbid(d2components.Dcc).
		Forbid(d2components.Ds1).
		Forbid(d2components.Dt1).
		Forbid(d2components.Wav).
		Forbid(d2components.AnimData).
		Build()

	// we are interested in actual game config instances, too
	gameConfigs := akara.NewFilter().Require(d2components.GameConfig).Build()

	gcs := &GameBootstrapSystem{
		SubscriberSystem: akara.NewSubscriberSystem(thingsToCheck, gameConfigs),
		maps: struct {
			gameConfigs *d2components.GameConfigMap
			filePaths   *d2components.FilePathMap
			fileTypes   *d2components.FileTypeMap
			fileHandles *d2components.FileHandleMap
			fileSources *d2components.FileSourceMap
		}{},
	}

	return gcs
}

// GameBootstrapSystem is responsible for setting up the regular diablo2 game launch
type GameBootstrapSystem struct {
	*akara.SubscriberSystem
	filesToCheck *akara.Subscription
	gameConfigs  *akara.Subscription
	maps         struct {
		gameConfigs *d2components.GameConfigMap
		filePaths   *d2components.FilePathMap
		fileTypes   *d2components.FileTypeMap
		fileHandles *d2components.FileHandleMap
		fileSources *d2components.FileSourceMap
	}
}

func (m *GameBootstrapSystem) Init(world *akara.World) {
	m.World = world

	if world == nil {
		m.SetActive(false)
		return
	}

	for subIdx := range m.Subscriptions {
		m.Subscriptions[subIdx] = m.AddSubscription(m.Subscriptions[subIdx].Filter)
	}

	m.filesToCheck = m.Subscriptions[0]
	m.gameConfigs = m.Subscriptions[1]

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.maps.filePaths = world.InjectMap(d2components.FilePath).(*d2components.FilePathMap)
	m.maps.fileTypes = world.InjectMap(d2components.FileType).(*d2components.FileTypeMap)
	m.maps.fileHandles = world.InjectMap(d2components.FileHandle).(*d2components.FileHandleMap)
	m.maps.fileSources = world.InjectMap(d2components.FileSource).(*d2components.FileSourceMap)
	m.maps.gameConfigs = world.InjectMap(d2components.GameConfig).(*d2components.GameConfigMap)

	m.bootstrap()
}

// bootstrap sets up the config directories and config file for processing by other systems.
// when the config is loaded, it sets up the mpq files as sources.
func (m *GameBootstrapSystem) bootstrap() {
	// we make two entities and assign file paths for the two directories that
	// we assume a config file may be inside of. These will be processed in the future by
	// the file type resolver system, and then the file source resolver system. At that point,
	// there will be sources for these two directories that can resolve the config file.
	e1, e2 := m.NewEntity(), m.NewEntity()
	fp1, fp2 := m.maps.filePaths.AddFilePath(e1), m.maps.filePaths.AddFilePath(e2)

	// the od2 directory has the highest priority
	fp1.Path = path.Dir(os.Args[0])

	// the user config directory is second highest
	configDir, err := os.UserConfigDir()
	if err == nil {
		fp2.Path = path.Join(configDir, configDirectoryName)
	} else {
		// we couldn't find the directory
		m.RemoveEntity(e2)
	}

	// now we set up the config file to be loaded. this happens after the directories
	// above are recognized as file sources.
	e3 := m.NewEntity()
	fp3 := m.maps.filePaths.AddFilePath(e3)
	fp3.Path = configFileName
}

func (m *GameBootstrapSystem) Process() {
	configs := m.gameConfigs.GetEntities()
	if len(configs) < 1 {
		return
	}

	cfg, found := m.maps.gameConfigs.GetGameConfig(configs[0])
	if !found {
		return
	}

	m.initMpqSources(cfg)
	m.SetActive(false) // bootstrap is complete!
}

func (m *GameBootstrapSystem) initMpqSources(cfg *d2components.GameConfigComponent) {
	for _, mpqFileName := range cfg.MpqLoadOrder {
		fullMpqFilePath := path.Join(cfg.MpqPath, mpqFileName)

		// make a new entity for the mpq file source
		mpqSource := m.maps.filePaths.AddFilePath(m.NewEntity())
		mpqSource.Path = fullMpqFilePath
	}
}
