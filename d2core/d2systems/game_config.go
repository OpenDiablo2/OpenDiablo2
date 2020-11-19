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

func NewGameConfigSystem() *GameConfigSystem {
	// we are going to check entities that dont yet have loaded asset types
	filesToCheck := akara.NewFilter().
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
	gameConfigs := akara.NewFilter().
		Require(d2components.GameConfig).
		Build()

	gcs := &GameConfigSystem{
		BaseSubscriberSystem: akara.NewBaseSubscriberSystem(filesToCheck, gameConfigs),
		Logger: d2util.NewLogger(),
	}

	gcs.SetPrefix(loggerPrefixGameConfig)

	return gcs
}

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
	*akara.BaseSubscriberSystem
	*d2util.Logger
	filesToCheck *akara.Subscription
	gameConfigs  *akara.Subscription
	*d2components.GameConfigMap
	*d2components.FilePathMap
	*d2components.FileTypeMap
	*d2components.FileHandleMap
	*d2components.FileSourceMap
	*d2components.DirtyMap
	ActiveConfig *d2components.GameConfigComponent
}

func (m *GameConfigSystem) Init(world *akara.World) {
	m.Info("initializing ...")

	m.filesToCheck = m.Subscriptions[0]
	m.gameConfigs = m.Subscriptions[1]

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.FilePathMap = world.InjectMap(d2components.FilePath).(*d2components.FilePathMap)
	m.FileTypeMap = world.InjectMap(d2components.FileType).(*d2components.FileTypeMap)
	m.FileHandleMap = world.InjectMap(d2components.FileHandle).(*d2components.FileHandleMap)
	m.FileSourceMap = world.InjectMap(d2components.FileSource).(*d2components.FileSourceMap)
	m.GameConfigMap = world.InjectMap(d2components.GameConfig).(*d2components.GameConfigMap)
	m.DirtyMap = world.InjectMap(d2components.Dirty).(*d2components.DirtyMap)
}

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
		m.GameConfigMap.Remove(eid)
	}

	m.ActiveConfig = gameConfig
}
