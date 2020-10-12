package d2systems

import (
	"encoding/json"
	"os"
	"path"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	configDirectoryName = "OpenDiablo2"
	configFileName      = "config.json"
)

// static check that the game config system implements the system interface
var _ akara.System = &GameConfigSystem{}

func NewGameConfigSystem() *GameConfigSystem {
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

	//the fabled singleton component. the game config will be added once.
	gameConfigs := akara.NewFilter().Require(d2components.GameConfig).Build()

	gcs := &GameConfigSystem{
		SubscriberSystem: akara.NewSubscriberSystem(thingsToCheck, gameConfigs),
		maps: struct {
			gameConfigs *d2components.GameConfigMap
			filePaths   *d2components.FilePathMap
			fileTypes   *d2components.FileTypeMap
			fileHandles *d2components.FileHandleMap
			fileSources *d2components.FileSourceMap
			dirty       *d2components.DirtyMap
		}{},
	}

	return gcs
}

// GameConfigSystem is responsible for game config bootstrap procedure, as well as
// clearing the `Dirty` component of game configs. In the `bootstrap` method of this system
// you can see that this system will add entities for the directories it expects config files
// to be found in, and it also adds an entity for the initial config file to be loaded.
//
// This system is dependant on the FileTypeResolver, FileSourceResolver, and
// FileHandleResolver systems because this system subscribes to entities
// with components created by these other systems. Nothing will  break if these
// other systems are not present in the world, but no config files will be loaded by
// this system either...
type GameConfigSystem struct {
	*akara.SubscriberSystem
	filesToCheck *akara.Subscription
	gameConfigs  *akara.Subscription
	maps         struct {
		gameConfigs *d2components.GameConfigMap
		filePaths   *d2components.FilePathMap
		fileTypes   *d2components.FileTypeMap
		fileHandles *d2components.FileHandleMap
		fileSources *d2components.FileSourceMap
		dirty       *d2components.DirtyMap
	}
}

func (m *GameConfigSystem) Init(world *akara.World) {
	m.World = world

	if world == nil {
		m.SetActive(false)
		return
	}

	for subIdx := range m.Subscriptions {
		m.AddSubscription(m.Subscriptions[subIdx])
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
	m.maps.dirty = world.InjectMap(d2components.Dirty).(*d2components.DirtyMap)

	m.bootstrap()
}

// bootstrap sets up the config directories, which will get turned into file sources,
// as well as the config file it is looking for, which will eventually get added as
// an entity in the subscription for this system. After it loads
func (m *GameConfigSystem) bootstrap() {
	// we make two entities and assign file paths for the two directories that
	// we assume a config file may be inside of.
	e1, e2 := m.NewEntity(), m.NewEntity()
	fp1, fp2 := m.maps.filePaths.AddFilePath(e1), m.maps.filePaths.AddFilePath(e2)

	// we'll add a filepath for user config dir, like `~/.config/OpenDiablo2/`
	configDir, err := os.UserConfigDir()
	if err == nil {
		fp1.Path = path.Join(configDir, configDirectoryName)
	} else {
		// we can safely remove the entity and it's components
		// if we cant find the user config dir, no biggie
		m.RemoveEntity(e1)
	}

	// our second directory is the dir where od2 is located
	fp2.Path = path.Dir(os.Args[0])

	// Now, we add another entity which will be for loading our config file.
	// Assuming that the FileTypeResolver and FileHandleResolver systems are active,
	// this entity should eventually get the components required by our subscription
	// for files to check. Once that happens, we will process the file into a GameConfig.
	e3 := m.NewEntity()
	fp3 := m.maps.filePaths.AddFilePath(e3)
	fp3.Path = configFileName
}

func (m *GameConfigSystem) Process() {
	m.clearDirty(m.gameConfigs.GetEntities())
	m.checkForNewConfig(m.filesToCheck.GetEntities())
}

func (m *GameConfigSystem) clearDirty(entities []akara.EID) {
	for _, eid := range entities {
		dc, found := m.maps.dirty.GetDirty(eid)
		if !found {
			m.maps.dirty.AddDirty(eid) // adds it, but it's false
			continue
		}

		dc.IsDirty = false
	}
}

func (m *GameConfigSystem) checkForNewConfig(entities []akara.EID) {
	for _, eid := range entities {
		fp, found := m.maps.filePaths.GetFilePath(eid)
		if !found {
			continue
		}

		ft, found := m.maps.fileTypes.GetFileType(eid)
		if !found {
			continue
		}

		if fp.Path != configFileName || ft.Type != d2enum.FileTypeJSON {
			continue
		}

		m.loadConfig(eid)
	}
}

func (m *GameConfigSystem) loadConfig(eid akara.EID) {
	fh, found := m.maps.fileHandles.GetFileHandle(eid)
	if !found {
		return
	}

	gameConfig := m.maps.gameConfigs.AddGameConfig(eid)

	if err := json.NewDecoder(fh.Data).Decode(gameConfig); err != nil {
		m.maps.gameConfigs.Remove(eid)
		return
	}

	for _, mpqFileName := range gameConfig.MpqLoadOrder {
		fullMpqFilePath := path.Join(gameConfig.MpqPath, mpqFileName)

		// make a new entity for the mpq file source
		mpqSource := m.maps.filePaths.AddFilePath(m.NewEntity())
		mpqSource.Path = fullMpqFilePath
	}
}
