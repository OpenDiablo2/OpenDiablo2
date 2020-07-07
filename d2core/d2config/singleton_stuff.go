package d2config

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

// TODO remove this shit

var singleton = getDefaultConfig()

// Load loads a configuration object from disk
func Load() error {
	return singleton.Load()
}

// Save saves the configuration object to disk
func Save() error {
	return singleton.Save()
}

// Get returns a configuration object
func Get() d2interface.Configuration {
	return singleton
}
