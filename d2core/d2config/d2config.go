package d2config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Configuration defines the configuration for the engine, loaded from config.json
type Configuration struct {
	MpqPath         string
	Backend         string
	path            string
	MpqLoadOrder    []string
	TicksPerSecond  int
	BgmVolume       float64
	FpsCap          int
	SfxVolume       float64
	FullScreen      bool
	RunInBackground bool
	VsyncEnabled    bool
}

// Save saves the configuration object to disk
func (c *Configuration) Save() error {
	configDir := filepath.Dir(c.path)
	if err := os.MkdirAll(configDir, 0750); err != nil {
		return err
	}

	configFile, err := os.Create(c.path)
	if err != nil {
		return err
	}

	buf, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	if _, err := configFile.Write(buf); err != nil {
		return err
	}

	return configFile.Close()
}

// Dir returns the directory component of the path
func (c *Configuration) Dir() string {
	return filepath.Dir(c.path)
}

// Base returns the base component of the path
func (c *Configuration) Base() string {
	return filepath.Base(c.path)
}

// Path returns the config file path
func (c *Configuration) Path() string {
	return c.path
}

// SetPath sets where the config file is saved to (a full path)
func (c *Configuration) SetPath(p string) {
	c.path = p
}
