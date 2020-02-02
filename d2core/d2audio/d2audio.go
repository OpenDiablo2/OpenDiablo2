package d2audio

import (
	"errors"
)

var singleton AudioProvider

var (
	ErrHasInit = errors.New("audio system is already initialized")
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
	if singleton != nil {
		return ErrHasInit
	}
	singleton = audioProvider
	return nil
}

// PlayBGM plays an infinitely looping background track
func PlayBGM(song string) error {
	if singleton == nil {
		return ErrNotInit
	}
	singleton.PlayBGM(song)
	return nil
}

func LoadSoundEffect(sfx string) (SoundEffect, error) {
	if singleton == nil {
		return nil, ErrNotInit
	}
	return singleton.LoadSoundEffect(sfx)
}

func SetVolumes(bgmVolume, sfxVolume float64) error {
	if singleton == nil {
		return ErrNotInit
	}
	singleton.SetVolumes(bgmVolume, sfxVolume)
	return nil
}
