// Package ebiten contains ebiten's implementation of the audio interface
package ebiten

import (
	"errors"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/hajimehoshi/ebiten/audio/wav"

	"github.com/hajimehoshi/ebiten/audio"
)

const sampleRate = 44100

// AudioProvider represents a provider capable of playing audio
type AudioProvider struct {
	app          d2interface.App
	audioContext *audio.Context // The Audio context
	bgmAudio     *audio.Player  // The audio player
	lastBgm      string
	sfxVolume    float64
	bgmVolume    float64
}

func (eap *AudioProvider) Advance(_, _ float64) error {
	return nil // AudioProvider doesn't need to advance
}

func (eap *AudioProvider) Render(_ d2interface.Surface) error {
	return nil // AudioProvider doesn't need to render
}

func (eap *AudioProvider) Initialize() error {
	panic("implement me")
}

func (eap *AudioProvider) BindApp(app d2interface.App) error {
	if eap.app != nil {
		return errors.New("audio provider already bound to an app")
	}

	eap.app = app

	return nil
}

func (eap *AudioProvider) UnbindApp(app d2interface.App) error {
	if eap.app == nil {
		return errors.New("audio provider not bound to an app")
	}

	eap.app = nil

	return nil
}

// CreateAudio creates an instance of ebiten's audio provider
func CreateAudio() (d2interface.AudioProvider, error) {
	result := &AudioProvider{}

	var err error
	result.audioContext, err = audio.NewContext(sampleRate)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, nil
}

// PlayBGM loads an audio stream and plays it in the background
func (eap *AudioProvider) PlayBGM(song string) {
	if eap.lastBgm == song {
		return
	}

	eap.lastBgm = song

	if song == "" && eap.bgmAudio != nil && eap.bgmAudio.IsPlaying() {
		_ = eap.bgmAudio.Pause()

		return
	}

	if eap.bgmAudio != nil {
		err := eap.bgmAudio.Close()

		if err != nil {
			log.Panic(err)
		}
	}

	audioStream, err := d2asset.LoadFileStream(song)

	if err != nil {
		panic(err)
	}

	d, err := wav.Decode(eap.audioContext, audioStream)

	if err != nil {
		log.Fatal(err)
	}

	s := audio.NewInfiniteLoop(d, d.Length())
	eap.bgmAudio, err = audio.NewPlayer(eap.audioContext, s)

	if err != nil {
		log.Fatal(err)
	}

	eap.bgmAudio.SetVolume(eap.bgmVolume)

	// Play the infinite-length stream. This never ends.
	err = eap.bgmAudio.Rewind()

	if err != nil {
		panic(err)
	}

	err = eap.bgmAudio.Play()

	if err != nil {
		panic(err)
	}
}

// LoadSoundEffect loads a sound affect so that it canb e played
func (eap *AudioProvider) LoadSoundEffect(sfx string) (d2interface.SoundEffect, error) {
	result := CreateSoundEffect(sfx, eap.audioContext, eap.sfxVolume) // TODO: Split

	return result, nil
}

// SetVolumes sets the volumes of the audio provider
func (eap *AudioProvider) SetVolumes(bgmVolume, sfxVolume float64) {
	eap.sfxVolume = sfxVolume
	eap.bgmVolume = bgmVolume
}
