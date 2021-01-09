package d2config

import (
	"os/user"
	"path"
	"runtime"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

// DefaultConfig creates and returns a default configuration
func DefaultConfig() *Configuration {
	const (
		defaultSfxVolume = 1.0
		defaultBgmVolume = 0.3
	)

	config := &Configuration{
		FullScreen:      false,
		TicksPerSecond:  -1,
		RunInBackground: true,
		VsyncEnabled:    true,
		SfxVolume:       defaultSfxVolume,
		BgmVolume:       defaultBgmVolume,
		MpqPath:         "C:/Program Files (x86)/Diablo II",
		Backend:         "Ebiten",
		MpqLoadOrder: []string{
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
		},
		LogLevel: d2util.LogLevelDefault,
		path:     DefaultConfigPath(),
	}

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
