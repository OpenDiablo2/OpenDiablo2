package d2audio

import (
	"errors"
)

var singleton AudioProvider

var (
	ErrWasInit = errors.New("audio system is already initialized")
	ErrNotInit = errors.New("audio system has not been initialized")
)

type SoundEffect interface {
	Play()
	Stop()
}

type AudioProvider interface {
	PlayBGM(song string)
	LoadSoundEffect(sfx string) (SoundEffect, error)
	SetVolumes(bgmVolume, sfxVolume float64)
}

// CreateManager creates a sound provider
func Initialize(audioProvider AudioProvider) error {
	verifyNotInit()
	singleton = audioProvider
	return nil
}

// PlayBGM plays an infinitely looping background track
func PlayBGM(song string) error {
	verifyWasInit()
	go func() {
		singleton.PlayBGM(song)
	}()
	return nil
}

func LoadSoundEffect(sfx string) (SoundEffect, error) {
	verifyWasInit()
	return singleton.LoadSoundEffect(sfx)
}

func SetVolumes(bgmVolume, sfxVolume float64) {
	verifyWasInit()
	singleton.SetVolumes(bgmVolume, sfxVolume)
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
