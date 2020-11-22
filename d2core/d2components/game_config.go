//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"os/user"
	"path"
	"runtime"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/gravestench/akara"
)

// static check that GameConfigComponent implements Component
var _ akara.Component = &GameConfigComponent{}

// static check that GameConfigMap implements ComponentMap
var _ akara.ComponentMap = &GameConfigMap{}

// GameConfigComponent represents an OpenDiablo2 game configuration
type GameConfigComponent struct {
	*akara.BaseComponent
	MpqLoadOrder    []string
	MpqPath         string
	TicksPerSecond  int
	FpsCap          int
	SfxVolume       float64
	BgmVolume       float64
	FullScreen      bool
	RunInBackground bool
	VsyncEnabled    bool
	Backend         string
	LogLevel        d2util.LogLevel
}

// GameConfigMap is a map of entity ID's to GameConfig
type GameConfigMap struct {
	*akara.BaseComponentMap
}

// AddGameConfig adds a new GameConfigComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *GameConfigComponent instead of an akara.Component
func (cm *GameConfigMap) AddGameConfig(id akara.EID) *GameConfigComponent {
	return defaultConfig(cm.Add(id).(*GameConfigComponent))
}

// GetGameConfig returns the GameConfigComponent associated with the given entity id
func (cm *GameConfigMap) GetGameConfig(id akara.EID) (*GameConfigComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*GameConfigComponent), found
}

// GameConfig is a convenient reference to be used as a component identifier
var GameConfig = newGameConfig() // nolint:gochecknoglobals // global by design

func newGameConfig() akara.Component {
	return &GameConfigComponent{
		BaseComponent: akara.NewBaseComponent(GameConfigCID, newGameConfig, newGameConfigMap),
	}
}

func newGameConfigMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(GameConfigCID, newGameConfig, newGameConfigMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &GameConfigMap{
		BaseComponentMap: baseMap,
	}

	return cm
}

func defaultConfig(config *GameConfigComponent) *GameConfigComponent {
	const (
		defaultSfxVolume = 1.0
		defaultBgmVolume = 0.3
	)

	config.FullScreen = false
	config.TicksPerSecond = -1
	config.RunInBackground = true
	config.VsyncEnabled = true
	config.SfxVolume = defaultSfxVolume
	config.BgmVolume = defaultBgmVolume
	config.MpqPath = "C:/Program Files (x86)/Diablo II"
	config.Backend = "Ebiten"
	config.MpqLoadOrder = []string{
		"Patch_D2.mpq",
		"d2exp.mpq",
		"d2xmusic.mpq",
		"d2xtalk.mpq",
		"d2xvideo.mpq",
		"d2data.mpq",
		"d2char.mpq",
		"d2music.mpq",
		"d2sfx.mpq",
		"d2video.mpq",
		"d2speech.mpq",
	}
	config.LogLevel = d2util.LogLevelDefault

	switch runtime.GOOS {
	case "windows":
		if runtime.GOARCH == "386" {
			config.MpqPath = "C:/Program Files/Diablo II"
		}
	case "darwin":
		config.MpqPath = "/Applications/Diablo II/"
		config.MpqLoadOrder = []string{
			"Diablo II Patch",
			"Diablo II Expansion Data",
			"Diablo II Expansion Movies",
			"Diablo II Expansion Music",
			"Diablo II Expansion Speech",
			"Diablo II Game Data",
			"Diablo II Graphics",
			"Diablo II Movies",
			"Diablo II Music",
			"Diablo II Sounds",
			"Diablo II Speech",
		}
	case "linux":
		if usr, err := user.Current(); err == nil {
			config.MpqPath = path.Join(usr.HomeDir, ".wine/drive_c/Program Files (x86)/Diablo II")
		}
	}

	return config
}
