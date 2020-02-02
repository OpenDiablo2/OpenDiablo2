package d2config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
)

var (
	ErrNotInit = errors.New("configuration is not initialized")
	ErrHasInit = errors.New("configuration has already been initialized")
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

var singleton *Configuration

func Initialize() error {
	if singleton != nil {
		return ErrHasInit
	}
	configDir, err := os.UserConfigDir()
	if err != nil {
		singleton = getDefaultConfiguration()
		return nil
	}

	configDir = path.Join(configDir, "OpenDiablo2")
	configPath := path.Join(configDir, "config.json")
	log.Printf("loading configuration file from %s...", configPath)
	configFile, err := os.Open(configPath)

	if err == nil {
		var config Configuration
		decoder := json.NewDecoder(configFile)
		defer configFile.Close()
		if err := decoder.Decode(&config); err == nil {
			singleton = &config
			return nil
		}
	}

	singleton = getDefaultConfiguration()
	return nil
}

func Save() error {
	if singleton == nil {
		return ErrNotInit
	}
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
	if err := encoder.Encode(singleton); err != nil {
		return err
	}

	return nil
}

func Get() (*Configuration, error) {
	if singleton == nil {
		return nil, ErrNotInit
	}
	return singleton, nil
}
