package d2config

import (
	"os/user"
	"path/filepath"
	"runtime"
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
			"patch_d2.mpq",
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
		path: DefaultConfigPath(),
	}

	switch runtime.GOOS {
	case "windows":
		if runtime.GOARCH == "386" {
			config.MpqPath = "C:/Program Files/Diablo II"
		}
	case "darwin":
		config.MpqPath = "/Applications/Diablo II/"
	case "linux":
		if usr, err := user.Current(); err == nil {
			config.MpqPath = filepath.Join(usr.HomeDir, ".wine", "drive_c", "Program Files (x86)", "Diablo II")
		}
	}

	return config
}
