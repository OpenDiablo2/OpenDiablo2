package d2config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"
)

var (
	ErrNotInit = errors.New("configuration is not initialized")
	ErrWasInit = errors.New("configuration has already been initialized")
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
	verifyNotInit()

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
	} else {
		log.Printf("configuration file not found, writing default")
		os.MkdirAll(filepath.Dir(configPath), os.ModePerm)
		configFile, err := os.Create(configPath)
		if err == nil {
			encoder := json.NewEncoder(configFile)
			defer configFile.Close()
			encoder.SetIndent("", "    ")
			encoder.Encode(getDefaultConfiguration())
		} else {
			log.Printf("failed to write default configuration (%s)", err)
		}
	}

	singleton = getDefaultConfiguration()
	return nil
}

func Save() error {
	verifyWasInit()

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

func Get() *Configuration {
	verifyWasInit()
	return singleton
}

func verifyWasInit() {
	if singleton == nil {
		panic(ErrNotInit)
	}
}

func verifyNotInit() {
	if singleton != nil {
		panic(ErrWasInit)
	}
}
