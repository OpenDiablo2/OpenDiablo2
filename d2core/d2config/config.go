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
		language:        "ENG",
		fullScreen:      false,
		ticksPerSecond:  -1,
		runInBackground: true,
		vsyncEnabled:    true,
		sfxVolume:       defaultSfxVolume,
		bgmVolume:       defaultBgmVolume,
		mpqPath:         "C:/Program Files (x86)/Diablo II",
		backend:         "Ebiten",
		mpqLoadOrder: []string{
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
			config.mpqPath = "C:/Program Files/Diablo II"
		}
	case "darwin":
		config.mpqPath = "/Applications/Diablo II/"
		config.mpqLoadOrder = []string{
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
			config.mpqPath = path.Join(usr.HomeDir, ".wine/drive_c/Program Files (x86)/Diablo II")
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

	if err := unmarshalIntoInterface(data); err != nil {
		return err
	}

	err = configFile.Close()

	if err != nil {
		return err
	}

	return nil
}

func unmarshalIntoInterface(d []byte) error {
	tmp := &hack{} // an empty concrete implementation
	if err := json.Unmarshal(d, tmp); err != nil {
		return err
	}

	tmp2cfg(tmp, singleton) // transfer tmp values to singleton

	return nil
}

// TODO figure out a way to unmarshal into an interface
type hack struct{
	MpqLoadOrder []string
	Language string
	MpqPath string
	TicksPerSecond int
	FpsCap int
	SfxVolume float64
	BgmVolume float64
	FullScreen bool
	RunInBackground bool
	VsyncEnabled bool
	Backend string
}

func cfg2tmp (a *Configuration, b *hack) {
	b.MpqLoadOrder = a.mpqLoadOrder
	b.Language = a.language
	b.MpqPath = a.mpqPath
	b.TicksPerSecond = a.ticksPerSecond
	b.FpsCap = a.fpsCap
	b.SfxVolume = a.sfxVolume
	b.BgmVolume = a.bgmVolume
	b.FullScreen = a.fullScreen
	b.RunInBackground = a.runInBackground
	b.VsyncEnabled = a.vsyncEnabled
	b.Backend = a.backend
}

func tmp2cfg (b *hack, a *Configuration) {
		a.mpqLoadOrder = b.MpqLoadOrder
		a.language = b.Language
		a.mpqPath = b.MpqPath
		a.ticksPerSecond = b.TicksPerSecond
		a.fpsCap = b.FpsCap
		a.sfxVolume = b.SfxVolume
		a.bgmVolume = b.BgmVolume
		a.fullScreen = b.FullScreen
		a.runInBackground = b.RunInBackground
		a.vsyncEnabled = b.VsyncEnabled
		a.backend = b.Backend
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

	tmp := &hack{}
	cfg2tmp(singleton, tmp)
	data, err := json.MarshalIndent(tmp, "", "    ")

	if err != nil {
		return err
	}

	if _, writeErr := configFile.Write(data); writeErr != nil {
		return writeErr
	}

	err = configFile.Close()

	if err != nil {
		return err
	}

	return nil
}
