package d2common

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/mitchellh/go-homedir"
)

// Configuration defines the configuration for the engine
type Configuration struct {
	Language          string
	FullScreen        bool
	Scale             float64
	PauseInBackground bool
	TicksPerSecond    int
	FpsCap            int
	VsyncDisabled     bool
	MpqPath           string
	MpqLoadOrder      []string
	SfxVolume         float64
	BgmVolume         float64
	MuteSound         bool
}

const ConfigBasePath = "config.json"

var defaultLoadOrder = []string{
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
}

var darwinLoadOrder = []string{
	"Diablo II Patch",            // "patch_d2.mpq"
	"Diablo II Expansion Data",   // "d2exp.mpq"
	"Diablo II Expansion Music",  // "d2xmusic.mpq"
	"Diablo II Expansion Speech", // "d2xtalk.mpq"
	"Diablo II Expansion Movies", // "d2xvideo.mpq"
	"Diablo II Game Data",        // "d2data.mpq"
	// "d2char.mpq",                 // ?
	"Diablo II Music",            // "d2music.mpq"
	"Diablo II Sounds",           // "d2sfx.mpq"
	"Diablo II Movies",           // "d2video.mpq"
	"Diablo II Speech",           // "d2speech.mpq"
}

func LoadConfiguration() (*Configuration, error) {
	var config Configuration

	f, err := os.Open(ConfigBasePath)
	if err == nil {
		defer f.Close()

		if err := json.NewDecoder(f).Decode(&config); err != nil {
			return nil, fmt.Errorf("failed to parse %s: %w", ConfigBasePath, err)
		}
	}

	if config.Language == "" {
		config.Language = "ENG"
	}

	if config.Scale <= 0 {
		config.Scale = 1.0
	}

	if config.TicksPerSecond <= 0 {
		config.TicksPerSecond = -1
	}

	if config.MuteSound {
		config.SfxVolume = 0.0
		config.BgmVolume = 0.0
	} else {
		if config.SfxVolume <= 0 {
			config.SfxVolume = 1.0
		}

		if config.BgmVolume <= 0 {
			config.BgmVolume = 0.3
		}
	}

	// Try to infer mpq path if not set explicitly
	if config.MpqPath == "" {
		mpqPath, err := inferMpqPath()
		if err != nil {
			return nil, err
		}

		if _, err := os.Stat(mpqPath); err != nil {
			return nil, fmt.Errorf("could not find Diablo 2 assets, Diablo II and Diablo II Lord of Destruction must be installed")
		}

		config.MpqPath = mpqPath
	}

	if len(config.MpqLoadOrder) == 0 {
		if runtime.GOOS == "darwin" {
			config.MpqLoadOrder = darwinLoadOrder
		} else {
			config.MpqLoadOrder = defaultLoadOrder
		}
	}

	return &config, nil
}

func inferMpqPath() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return "C:/Program Files (x86)/Diablo II", nil
	case "linux":
		homeDir, err := homedir.Dir()
		if err != nil {
			return "", fmt.Errorf("failed to get home directory: %w", err)
		}
		return filepath.Join(homeDir, ".wine", "drive_c", "Program Files (x86)/Diablo II"), nil
	case "darwin":
		return "/Applications/Diablo II", nil
	default:
		return "", fmt.Errorf("could not infer mpq path; must be provided in %s", ConfigBasePath)
	}
}
