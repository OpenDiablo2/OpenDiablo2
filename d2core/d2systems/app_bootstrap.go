package d2systems

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/gravestench/akara"
	"os"
	"path"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	configDirectoryName = "OpenDiablo2"
	configFileName      = "config.json"
)

const (
	logPrefixAppBootstrap = "App Bootstrap"
)

// static check that the game config system implements the system interface
var _ akara.System = &AppBootstrapSystem{}

// NewAppBootstrapSystem creates a system that initializes the common application facilities, namely
// config file source directories and loading/creating the OpenDiablo2 config file.
func NewAppBootstrapSystem() *AppBootstrapSystem {
	// we are going to check entities that dont yet have loaded asset types
	filesToCheck := akara.NewFilter().
		Require( // files that need to be loaded
			d2components.FileType,
			d2components.FileHandle,
			d2components.FilePath,
			).
		Forbid( // files which have been loaded
			d2components.GameConfig,
			d2components.StringTable,
			d2components.DataDictionary,
			d2components.Palette,
			d2components.PaletteTransform,
			d2components.Cof,
			d2components.Dc6,
			d2components.Dcc,
			d2components.Ds1,
			d2components.Dt1,
			d2components.Wav,
			d2components.AnimData,
			).
		Build()

	// we are interested in actual game config instances, too
	gameConfigs := akara.NewFilter().Require(d2components.GameConfig).Build()

	sys := &AppBootstrapSystem{
		SubscriberSystem: akara.NewSubscriberSystem(filesToCheck, gameConfigs),
		Logger: d2util.NewLogger(),
	}

	sys.SetPrefix(logPrefixAppBootstrap)
	sys.Debug("Created")

	return sys
}

// AppBootstrapSystem is responsible for the common initialization process between
// the app modes (eg common to the game client as well as the headless server)
type AppBootstrapSystem struct {
	*akara.SubscriberSystem
	*d2util.Logger
	subscribedFiles   *akara.Subscription
	subscribedConfigs *akara.Subscription
	*d2components.GameConfigMap
	*d2components.FilePathMap
	*d2components.FileTypeMap
	*d2components.FileHandleMap
	*d2components.FileSourceMap
}

// Init will inject (or use existing) components related to setting up the config sources and
func (m *AppBootstrapSystem) Init(world *akara.World) {
	m.World = world

	if world == nil {
		m.Error("world is nil, deactivating.")
		m.SetActive(false)
		return
	}

	m.Info("initializing ...")

	for subIdx := range m.Subscriptions {
		m.Subscriptions[subIdx] = m.AddSubscription(m.Subscriptions[subIdx].Filter)
	}

	m.subscribedFiles = m.Subscriptions[0]
	m.subscribedConfigs = m.Subscriptions[1]

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.GameConfigMap = world.InjectMap(d2components.GameConfig).(*d2components.GameConfigMap)
	m.FilePathMap = world.InjectMap(d2components.FilePath).(*d2components.FilePathMap)
	m.FileTypeMap = world.InjectMap(d2components.FileType).(*d2components.FileTypeMap)
	m.FileHandleMap = world.InjectMap(d2components.FileHandle).(*d2components.FileHandleMap)
	m.FileSourceMap = world.InjectMap(d2components.FileSource).(*d2components.FileSourceMap)

	m.injectSystems()
	m.setupConfigSources()
	m.setupConfigFile()
}

func (m *AppBootstrapSystem) injectSystems() {
	m.World.AddSystem(NewFileTypeResolver())
	m.World.AddSystem(NewFileSourceResolver())
	m.World.AddSystem(NewFileHandleResolver())
	m.World.AddSystem(NewGameConfigSystem())
	m.World.AddSystem(NewAssetLoader())
	m.World.AddSystem(NewGameObjectFactory())
}

// we make two entities and assign file paths for the two directories that
// we assume a config file may be inside of. These will be processed in the future by
// the file type resolver system, and then the file source resolver system. At that point,
// there will be sources for these two directories that can possibly resolve a config file.
// A new config file is created if one is not found.
func (m *AppBootstrapSystem) setupConfigSources() {
	// make the two entities, these will be the file sources
	e1, e2 := m.NewEntity(), m.NewEntity()

	// add file path components to these entities
	fp1, fp2 := m.AddFilePath(e1), m.AddFilePath(e2)

	// the first entity gets a filepath for the od2 directory, this one is checked first
	// eg. if OD2 binary is in `~/src/OpenDiablo2/`, then this directory is checked first for a config file
	cfgPath1 := path.Dir(os.Args[0])
	fp1.Path = cfgPath1
	m.Infof("setting up local directory %s for processing", cfgPath1)

	// now add the user config directory
	// this directory on a linux machine would be something like `~/.config/OpenDiablo2/`
	cfgPath2, err := os.UserConfigDir()
	if err == nil {
		fp2.Path = path.Join(cfgPath2, configDirectoryName)
		m.Infof("setting up user config directory %s for processing", fp2.Path)
	} else {
		// we couldn't find the directory, destroy this entity
		m.Error("user config directory not found, skipping")
		m.RemoveEntity(e2)
	}
}


func (m *AppBootstrapSystem) setupConfigFile() {
	// add an entity that will get picked up by the game config system and loaded
	m.AddFilePath(m.NewEntity()).Path = configFileName
	m.Infof("setting up config file `%s` for processing", configFileName)
}

// Update will look for the first entity with a game config component
// and then add the mpq's as file sources
func (m *AppBootstrapSystem) Update() {
	configs := m.subscribedConfigs.GetEntities()
	if len(configs) < 1 {
		return
	}

	m.Infof("found %d new configs to parse", len(configs))

	firstConfigEntityID := configs[0]
	cfg, found := m.GetGameConfig(firstConfigEntityID)
	if !found {
		return
	}

	m.initMpqSources(cfg)

	m.Info("app bootstrap complete, deactivating system")
	m.SetActive(false) // bootstrap is complete!
}

func (m *AppBootstrapSystem) initMpqSources(cfg *d2components.GameConfigComponent) {
	for _, mpqFileName := range cfg.MpqLoadOrder {
		fullMpqFilePath := path.Join(cfg.MpqPath, mpqFileName)

		m.Infof("adding mpq: %s", fullMpqFilePath)

		// make a new entity for the mpq file source
		mpqSource := m.AddFilePath(m.NewEntity())
		mpqSource.Path = fullMpqFilePath
	}
}
