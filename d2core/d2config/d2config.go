package d2config

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

// Config holds the configuration from config.json
var Config *Configuration //nolint:gochecknoglobals // Currently global by design

// Configuration defines the configuration for the engine, loaded from config.json
type Configuration struct {
	MpqLoadOrder    []string
	Language        string
	MpqPath         string
	TicksPerSecond  int
	FpsCap          int
	SfxVolume       float64
	BgmVolume       float64
	FullScreen      bool
	RunInBackground bool
	VsyncEnabled    bool
	Backend         string
}

// Load loads a configuration object from disk
func Load() error {
	Config = new(Configuration)
	return Config.Load()
}

// Load loads a configuration object from disk
func (c *Configuration) Load() error {
	configPaths := []string{
		defaultConfigPath(),
		localConfigPath(),
	}

	for _, configPath := range configPaths {
		log.Printf("loading configuration file from %s...", configPath)

		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			continue
		}

		configFile, err := os.Open(path.Clean(configPath))
		if err != nil {
			return err
		}

		if err := json.NewDecoder(configFile).Decode(&Config); err != nil {
			return err
		}

		if err := configFile.Close(); err != nil {
			return err
		}

		return nil
	}

	log.Println("failed to load configuration file, saving default configuration...")

	Config = defaultConfig()

	return Config.Save()
}

// Save saves the configuration object to disk
func (c *Configuration) Save() error {
	configPath := defaultConfigPath()
	log.Printf("saving configuration file to %s...", configPath)

	configDir := path.Dir(configPath)
	if err := os.MkdirAll(configDir, 0750); err != nil {
		return err
	}

	configFile, err := os.Create(configPath)
	if err != nil {
		return err
	}

	buf, err := json.MarshalIndent(Config, "", "  ")
	if err != nil {
		return err
	}

	if _, err := configFile.Write(buf); err != nil {
		return err
	}

	return configFile.Close()
}

func defaultConfigPath() string {
	if configDir, err := os.UserConfigDir(); err == nil {
		return path.Join(configDir, "OpenDiablo2", "config.json")
	}

	return localConfigPath()
}

func localConfigPath() string {
	return path.Join(path.Dir(os.Args[0]), "config.json")
}
