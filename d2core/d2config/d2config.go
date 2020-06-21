package d2config

import (
	"log"
)

// Configuration defines the configuration for the engine, loaded from config.json
type Configuration struct {
	Language        string
	FullScreen      bool
	RunInBackground bool
	TicksPerSecond  int
	FpsCap          int
	VsyncEnabled    bool
	MpqPath         string
	MpqLoadOrder    []string
	SfxVolume       float64
	BgmVolume       float64
}

var singleton *Configuration = getDefaultConfig()

func Load() error {
	configPaths := []string{
		getLocalConfigPath(),
		getDefaultConfigPath(),
	}

	var loaded bool
	for _, configPath := range configPaths {
		log.Printf("loading configuration file from %s...", configPath)
		if err := load(configPath); err == nil {
			loaded = true
			break
		}
	}

	if !loaded {
		log.Println("failed to load configuration file, saving default configuration...")
		if err := Save(); err != nil {
			return err
		}
	}

	return nil
}

func Save() error {
	configPath := getDefaultConfigPath()
	log.Printf("saving configuration file to %s...", configPath)

	var err error
	if err = save(configPath); err != nil {
		log.Printf("failed to write configuration file (%s)", err)
	}

	return err
}

func Set(config Configuration) {
	temp := config
	singleton = &temp
}

func Get() Configuration {
	if singleton == nil {
		panic("configuration is not initialized")
	}

	return *singleton
}
