package d2corecommon

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
	"path"
	"runtime"
)

// Configuration defines the configuration for the engine, loaded from config.json
type Configuration struct {
	Language        string
	FullScreen      bool
	Scale           float64
	RunInBackground bool
	TicksPerSecond  int
	FpsCap          int
	VsyncEnabled    bool
	MpqPath         string
	MpqLoadOrder    []string
	SfxVolume       float64
	BgmVolume       float64
}

func LoadConfiguration() *Configuration {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return getDefaultConfiguration()
	}

	configDir = path.Join(configDir, "OpenDiablo2")
	configPath := path.Join(configDir, "config.json")
	log.Printf("loading configuration file from %s...", configPath)
	configFile, err := os.Open(configPath)
	defer configFile.Close()

	if err == nil {
		var config Configuration
		decoder := json.NewDecoder(configFile)
		if err := decoder.Decode(&config); err == nil {
			return &config
		}
	}

	return getDefaultConfiguration()
}

func (c *Configuration) Save() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	configDir = path.Join(configDir, "OpenDiablo2")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	configPath := path.Join(configDir, "config.json")
	log.Printf("saving configuration file to %s...", configPath)
	configFile, err := os.Create(configPath)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(configFile)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(c); err != nil {
		return err
	}

	return nil
}

func getConfigurationPath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "config.json"
	}

	return path.Join(configDir, "OpenDiablo2/config.json")
}

func getDefaultConfiguration() *Configuration {
	config := &Configuration{
		Language:        "ENG",
		FullScreen:      true,
		Scale:           1,
		TicksPerSecond:  -1,
		RunInBackground: true,
		VsyncEnabled:    true,
		SfxVolume:       1.0,
		BgmVolume:       0.3,
		MpqPath:         "C:/Program Files (x86)/Diablo II",
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
	}

	switch runtime.GOOS {
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
