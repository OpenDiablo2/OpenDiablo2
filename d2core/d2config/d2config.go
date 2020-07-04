package d2config

import (
	"log"
)

// Configuration defines the configuration for the engine, loaded from config.json
type Configuration struct {
	mpqLoadOrder    []string
	language        string
	mpqPath         string
	ticksPerSecond  int
	fpsCap          int
	sfxVolume       float64
	bgmVolume       float64
	fullScreen      bool
	runInBackground bool
	vsyncEnabled    bool
	backend         string
}

func (c *Configuration) MpqLoadOrder() []string {
	return c.mpqLoadOrder
}

func (c *Configuration) Language() string {
	return c.language
}

func (c *Configuration) MpqPath() string {
	return c.mpqPath
}

func (c *Configuration) TicksPerSecond() int {
	return c.ticksPerSecond
}

func (c *Configuration) FpsCap() int {
	return c.fpsCap
}

func (c *Configuration) SfxVolume() float64 {
	return c.sfxVolume
}

func (c *Configuration) BgmVolume() float64 {
	return c.bgmVolume
}

func (c *Configuration) FullScreen() bool {
	return c.fullScreen
}

func (c *Configuration) RunInBackground() bool {
	return c.runInBackground
}

func (c *Configuration) VsyncEnabled() bool {
	return c.vsyncEnabled
}

func (c *Configuration) Backend() string {
	return c.backend
}

// Load loads a configuration object from disk
func (c *Configuration) Load() error {
	configPaths := []string{
		getDefaultConfigPath(),
		getLocalConfigPath(),
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

// Save saves the configuration object to disk
func (c *Configuration) Save() error {
	configPath := getDefaultConfigPath()
	log.Printf("saving configuration file to %s...", configPath)

	var err error
	if err = save(configPath); err != nil {
		log.Printf("failed to write configuration file (%s)", err)
	}

	return err
}
