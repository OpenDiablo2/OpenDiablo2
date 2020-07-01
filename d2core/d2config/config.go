// Package d2config contains configuration objects and functions
package d2config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"runtime"
)

const defaultSfxVolume = 1.0
const defaultBgmVolume = 0.3

func getDefaultConfig() *Configuration {
	config := &Configuration{
		Language:        "ENG",
		FullScreen:      false,
		TicksPerSecond:  -1,
		RunInBackground: true,
		VsyncEnabled:    true,
		SfxVolume:       defaultSfxVolume,
		BgmVolume:       defaultBgmVolume,
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
	case "windows":
		if runtime.GOARCH == "386" {
			config.MpqPath = "C:/Program Files/Diablo II"
		}
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
		if usr, err := user.Current(); err == nil {
			config.MpqPath = path.Join(usr.HomeDir, ".wine/drive_c/Program Files (x86)/Diablo II")
		}
	}

	return config
}

func getDefaultConfigPath() string {
	if configDir, err := os.UserConfigDir(); err == nil {
		return path.Join(configDir, "OpenDiablo2", "config.json")
	}

	return getLocalConfigPath()
}

func getLocalConfigPath() string {
	return path.Join(path.Dir(os.Args[0]), "config.json")
}

func load(configPath string) error {
	configFile, err := os.Open(configPath) //nolint:gosec will fix the security error later

	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(configFile)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, &singleton); err != nil {
		return err
	}

	err = configFile.Close()

	if err != nil {
		return err
	}

	return nil
}

func save(configPath string) error {
	configDir := path.Dir(configPath)

	if err := os.MkdirAll(configDir, 0750); err != nil {
		return err
	}

	configFile, err := os.Create(configPath)

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(singleton, "", "    ")

	if err != nil {
		return err
	}

	if _, err := configFile.Write(data); err != nil {
		return err
	}

	err = configFile.Close()

	if err != nil {
		return err
	}

	return nil
}
