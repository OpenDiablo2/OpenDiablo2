package d2config

import (
	"os"
	"path"
)

const (
	od2ConfigDirName = "OpenDiablo2"
)

const (
	od2ConfigFileName = "config.json"
)

// DefaultConfigPath returns the absolute path for the default config file location
func DefaultConfigPath() string {
	if configDir, err := os.UserConfigDir(); err == nil {
		return path.Join(configDir, od2ConfigDirName, od2ConfigFileName)
	}

	return LocalConfigPath()
}

// LocalConfigPath returns the absolute path to the directory of the OpenDiablo2 executable
func LocalConfigPath() string {
	return path.Join(path.Dir(os.Args[0]), od2ConfigFileName)
}
