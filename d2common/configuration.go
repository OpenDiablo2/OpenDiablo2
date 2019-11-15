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
}

const configPath = "config.json"

var defaultLoadOrder = []string{
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

var darwinLoadOrder = []string{
	"Diablo II Patch",
	"Diablo II Expansion Data",
	"Diablo II Expansion Music",
	"Diablo II Expansion Speech",
	"Diablo II Expansion Movies",
	"Diablo II Game Data",
	"Diablo II Speech",
	"Diablo II Music",
	"Diablo II Sounds",
	"Diablo II Movies",
	"Diablo II Speech",
}

func LoadConfiguration() (*Configuration, error) {
	var config Configuration

	f, err := os.Open(configPath)
	if err == nil {
		defer f.Close()

		if err := json.NewDecoder(f).Decode(&config); err != nil {
			return nil, fmt.Errorf("failed to parse %s: %w", configPath, err)
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

	if config.SfxVolume <= 0 {
		config.SfxVolume = 1.0
	}

	if config.BgmVolume <= 0 {
		config.BgmVolume = 0.3
	}

	// Try to infer mpq path if not set explicitly
	if config.MpqPath == "" {
		mpqPath, err := inferMpqPath()
		if err != nil {
			return nil, err
		}

		if _, err := os.Stat(mpqPath); err != nil {
			return nil, fmt.Errorf("could not find Diablo 2 assets. Install Diablo 2 or provide asset path in %s", configPath)
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
		return "", fmt.Errorf("could not infer mpq path; must be provided in %s", configPath)
	}
}
