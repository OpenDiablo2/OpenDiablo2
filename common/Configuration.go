package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/mitchellh/go-homedir"
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

// ConfigBasePath is used for tests to find the base config json file
var ConfigBasePath = "./"

func LoadConfiguration() *Configuration {
	configJSON, err := ioutil.ReadFile(path.Join(ConfigBasePath, "config.json"))
	if err != nil {
		log.Fatal(err)
	}
	var config Configuration
	err = json.Unmarshal(configJSON, &config)
	if err != nil {
		log.Fatal(err)
	}
	// Path fixup for wine-installed diablo 2 in linux
	if config.MpqPath[0] != '/' {
		if _, err := os.Stat(config.MpqPath); os.IsNotExist(err) {
			homeDir, _ := homedir.Dir()
			newPath := strings.ReplaceAll(config.MpqPath, `C:\`, homeDir+"/.wine/drive_c/")
			newPath = strings.ReplaceAll(newPath, "C:/", homeDir+"/.wine/drive_c/")
			newPath = strings.ReplaceAll(newPath, `\`, "/")
			if _, err := os.Stat(newPath); !os.IsNotExist(err) {
				config.MpqPath = newPath
			}
		}
	}
	return &config
}
