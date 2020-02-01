package d2audio

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var singleton d2interface.AudioProvider

var (
	ErrHasInit error = errors.New("audio system is already initialized")
	ErrNotInit error = errors.New("audio system has not been initialized")
)

// CreateManager creates a sound provider
func Initialize(audioProvider d2interface.AudioProvider) error {
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

func LoadSoundEffect(sfx string) (d2interface.SoundEffect, error) {
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
