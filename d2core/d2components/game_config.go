//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"os/user"
	"path"
	"runtime"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

// static check that GameConfig implements Component
var _ akara.Component = &GameConfig{}

// GameConfig represents an OpenDiablo2 game configuration
type GameConfig struct {
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

// New creates the default GameConfig
func (*GameConfig) New() akara.Component {
	const (
		defaultSfxVolume = 1.0
		defaultBgmVolume = 0.3
	)

	config := &GameConfig{}

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

// GameConfigFactory is a wrapper for the generic component factory that returns GameConfig component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a GameConfig.
type GameConfigFactory struct {
	*akara.ComponentFactory
}

// Add adds a GameConfig component to the given entity and returns it
func (m *GameConfigFactory) Add(id akara.EID) *GameConfig {
	return m.ComponentFactory.Add(id).(*GameConfig)
}

// Get returns the GameConfig component for the given entity, and a bool for whether or not it exists
func (m *GameConfigFactory) Get(id akara.EID) (*GameConfig, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*GameConfig), found
}
