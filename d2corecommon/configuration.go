package d2corecommon

import (
	"encoding/json"
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
	configFile, err := os.Open("config.json")
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
	configFile, err := os.Create("config.json")
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
		usr, err := user.Current()
		if err == nil {
			config.MpqPath = path.Join(usr.HomeDir, ".wine/drive_c/Program Files (x86)/Diablo II")
		}
	}

	return config
}
