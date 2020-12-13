package d2systems

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/pkg/profile"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path"
	"strings"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	configDirectoryName = "OpenDiablo2"
	configFileName      = "config.json"
)

const (
	logPrefixAppBootstrap = "App Bootstrap"
)

const (
	sceneTestArg  = "testscene"
	sceneTestDesc = "test a scene"

	serverArg  = "server"
	serverDesc = "run dedicated server"

	counterArg  = "counter"
	counterDesc = "print updates/sec"

	skipSplashArg  = "nosplash"
	skipSplashDesc = "skip the ebiten splash screen"

	logLevelArg  = "loglevel"
	logLevelShort  = 'l'
	logLevelDesc = "sets the logging level for all loggers at startup"

	profilerArg  = "profile"
	profilerDesc = "Profiles the program, one of (cpu, mem, block, goroutine, trace, thread, mutex)"
)

// static check that the game config system implements the system interface
var _ akara.System = &AppBootstrap{}

// AppBootstrap is responsible for the common initialization process between
// the app modes (eg common to the game client as well as the headless server)
type AppBootstrap struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
	subscribedFiles   *akara.Subscription
	subscribedConfigs *akara.Subscription
	Components        struct {
		GameConfig d2components.GameConfigFactory
		File       d2components.FileFactory
		FileType   d2components.FileTypeFactory
		FileHandle d2components.FileHandleFactory
		FileSource d2components.FileSourceFactory
	}
	logLevel d2util.LogLevel
}

// Init will inject (or use existing) components related to setting up the config sources
func (m *AppBootstrap) Init(world *akara.World) {
	m.World = world

	m.setupLogger()

	m.Debug("initializing ...")

	m.setupSubscriptions()
	m.setupFactories()
	m.injectSystems()
	m.setupConfigSources()
	m.setupConfigFile()
	m.setupLocaleFile()
	m.parseCommandLineArgs()

	m.Debug("... initialization complete!")

	if err := m.World.Update(0); err != nil {
		m.Error(err.Error())
	}
}

func (m *AppBootstrap) setupLogger() {
	m.Logger = d2util.NewLogger()
	m.SetPrefix(logPrefixAppBootstrap)
}

func (m *AppBootstrap) setupSubscriptions() {
	m.Debug("setting up component subscriptions")

	// we are going to check entities that dont yet have loaded asset types
	filesToCheck := m.NewComponentFilter().
		Require( // files that need to be loaded
			&d2components.FileType{},
			&d2components.FileHandle{},
			&d2components.File{},
		).
		Forbid( // files which have been loaded
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
	gameConfigs := m.NewComponentFilter().Require(&d2components.GameConfig{}).Build()

	m.subscribedFiles = m.World.AddSubscription(filesToCheck)
	m.subscribedConfigs = m.World.AddSubscription(gameConfigs)
}

func (m *AppBootstrap) setupFactories() {
	m.Debug("setting up component factories")

	m.InjectComponent(&d2components.GameConfig{}, &m.Components.GameConfig.ComponentFactory)
	m.InjectComponent(&d2components.File{}, &m.Components.File.ComponentFactory)
	m.InjectComponent(&d2components.FileType{}, &m.Components.FileType.ComponentFactory)
	m.InjectComponent(&d2components.FileHandle{}, &m.Components.FileHandle.ComponentFactory)
	m.InjectComponent(&d2components.FileSource{}, &m.Components.FileSource.ComponentFactory)
}

func (m *AppBootstrap) injectSystems() {
	m.Info("injecting file type resolution system")
	m.AddSystem(&FileTypeResolver{})

	m.Info("injecting file source resolution system")
	m.AddSystem(&FileSourceResolver{})

	m.Info("injecting file handle resolution system")
	m.AddSystem(&FileHandleResolver{})

	m.Info("injecting game configuration system")
	m.AddSystem(&GameConfigSystem{})

	m.Info("injecting asset loader system")
	m.AddSystem(&AssetLoaderSystem{})
}

// we make two entities and assign file paths for the two directories that
// we assume a config file may be inside of. These will be processed in the future by
// the file type resolver system, and then the file source resolver system. At that point,
// there will be sources for these two directories that can possibly resolve a config file.
// A new config file is created if one is not found.
func (m *AppBootstrap) setupConfigSources() {
	// make the two entities, these will be the file sources
	e1, e2 := m.NewEntity(), m.NewEntity()

	// add file path components to these entities
	fp1, fp2 := m.Components.File.Add(e1), m.Components.File.Add(e2)

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

func (m *AppBootstrap) setupConfigFile() {
	// add an entity that will get picked up by the game config system and loaded
	m.Components.File.Add(m.NewEntity()).Path = configFileName
	m.Infof("setting up config file `%s` for processing", configFileName)
}

func (m *AppBootstrap) setupLocaleFile() {
	// add an entity that will get picked up by the game config system and loaded
	m.Components.File.Add(m.NewEntity()).Path = d2resource.LocalLanguage
	m.Infof("setting up locale file `%s` for processing", d2resource.LocalLanguage)
}

// Update will look for the first entity with a game config component
// and then add the mpq's as file sources
func (m *AppBootstrap) Update() {
	configs := m.subscribedConfigs.GetEntities()
	if len(configs) < 1 {
		return
	}

	m.Debugf("found %d new configs to parse", len(configs))

	firstConfigEntityID := configs[0]

	cfg, found := m.Components.GameConfig.Get(firstConfigEntityID)
	if !found {
		return
	}

	cfg.LogLevel = m.logLevel

	m.Info("adding mpq's from config file")
	m.initMpqSources(cfg)

	m.Info("app bootstrap complete, deactivating system")
	m.SetActive(false)
}

func (m *AppBootstrap) initMpqSources(cfg *d2components.GameConfig) {
	for _, mpqFileName := range cfg.MpqLoadOrder {
		fullMpqFilePath := path.Join(cfg.MpqPath, mpqFileName)

		m.Infof("adding mpq: %s", fullMpqFilePath)

		// make a new entity for the mpq file source
		mpqSource := m.Components.File.Add(m.NewEntity())
		mpqSource.Path = fullMpqFilePath
	}
}

func (m *AppBootstrap) parseCommandLineArgs() {
	profilerOptions := kingpin.Flag(profilerArg, profilerDesc).String()
	sceneTest := kingpin.Flag(sceneTestArg, sceneTestDesc).String()
	server := kingpin.Flag(serverArg, serverDesc).Bool()
	enableCounter := kingpin.Flag(counterArg, counterDesc).Bool()
	logLevel := kingpin.Flag(logLevelArg, logLevelDesc).Short(logLevelShort).Int()
	_ = kingpin.Flag(skipSplashArg, skipSplashDesc).Bool() // see game client bootstrap

	kingpin.Parse()

	m.parseProfilerOptions(*profilerOptions)

	if *enableCounter {
		m.World.AddSystem(&UpdateCounter{})
	}

	if *server {
		fmt.Println("not yet implemented")
		os.Exit(0)
	}

	m.logLevel = *logLevel

	m.SetLevel(m.logLevel)

	m.World.AddSystem(&RenderSystem{})
	m.World.AddSystem(&InputSystem{})
	m.World.AddSystem(&GameObjectFactory{})

	switch *sceneTest {
	case "splash":
		m.Info("running ebiten splash scene")
		m.World.AddSystem(NewEbitenSplashScene())
	case "load":
		m.Info("running loading scene")
		m.World.AddSystem(NewLoadingScene())
	case "mouse":
		m.Info("running mouse cursor scene")
		m.World.AddSystem(NewMouseCursorScene())
	case "main":
		m.Info("running main menu scene")
		m.World.AddSystem(NewMainMenuScene())
	case "terminal":
		m.Info("running terminal scene")
		m.World.AddSystem(NewTerminalScene())
	case "labels":
		m.Info("running label test scene")
		m.World.AddSystem(NewLabelTestScene())
	case "buttons":
		m.Info("running button test scene")
		m.World.AddSystem(NewButtonTestScene())
	default:
		m.World.AddSystem(&GameClientBootstrap{})
	}
}

func (m *AppBootstrap) parseProfilerOptions(profileOption string) interface{ Stop() } {
	var options []func(*profile.Profile)

	switch strings.ToLower(strings.Trim(profileOption, " ")) {
	case "cpu":
		m.Debug("CPU profiling is enabled.")

		options = append(options, profile.CPUProfile)
	case "mem":
		m.Debug("Memory profiling is enabled.")

		options = append(options, profile.MemProfile)
	case "block":
		m.Debug("Block profiling is enabled.")

		options = append(options, profile.BlockProfile)
	case "goroutine":
		m.Debug("Goroutine profiling is enabled.")

		options = append(options, profile.GoroutineProfile)
	case "trace":
		m.Debug("Trace profiling is enabled.")

		options = append(options, profile.TraceProfile)
	case "thread":
		m.Debug("Thread creation profiling is enabled.")

		options = append(options, profile.ThreadcreationProfile)
	case "mutex":
		m.Debug("Mutex profiling is enabled.")

		options = append(options, profile.MutexProfile)
	}

	options = append(options, profile.ProfilePath("./pprof/"))

	if len(options) > 1 {
		return profile.Start(options...)
	}

	return nil
}
