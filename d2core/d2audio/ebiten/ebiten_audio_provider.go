// Package ebiten contains ebiten's implementation of the audio interface
package ebiten

import (
	"errors"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

const sampleRate = 44100

var _ d2interface.AudioProvider = &AudioProvider{} // Static check to confirm struct conforms to interface

// AudioProvider represents a provider capable of playing audio
type AudioProvider struct {
	app          d2interface.App
	assetManager d2interface.AssetManager
	audioContext *audio.Context // The Audio context
	bgmAudio     *audio.Player  // The audio player
	lastBgm      string
	sfxVolume    float64
	bgmVolume    float64
}

// BindApp binds to the OpenDiablo2 app
func (eap *AudioProvider) BindApp(app d2interface.App) error {
	if eap.app != nil {
		return errors.New("audio provider already bound to an app instance")
	}

	eap.app = app

	return nil
}

// Initialize the audio provider
func (eap *AudioProvider) Initialize() error {
	assetManager, err := eap.app.Asset()
	if err != nil {
		return err
	}

	eap.assetManager = assetManager

	return nil
}

// Bind to an asset manager
func (eap *AudioProvider) Bind(manager d2interface.AssetManager) error {
	if eap.assetManager != nil {
		return errors.New("font manager already bound to an asset manager")
	}

	eap.assetManager = manager

	return nil
}

// CreateAudio creates an instance of ebiten's audio provider
func CreateAudio() (*AudioProvider, error) {
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

	audioStream, err := eap.assetManager.LoadFileStream(song)

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
	asset, err := eap.app.Asset()
	if err != nil {
		return nil, errors.New("cannot load audio without an asset manager")
	}

	result := &SoundEffect{}

	var soundFile string

	if _, exists := d2datadict.Sounds[sfx]; exists {
		soundEntry := d2datadict.Sounds[sfx]
		soundFile = soundEntry.FileName
	} else {
		soundFile = sfx
	}

	audioData, err := asset.LoadFile(soundFile)

	if err != nil {
		panic(err)
	}

	d, err := wav.Decode(eap.audioContext, audio.BytesReadSeekCloser(audioData))

	if err != nil {
		log.Fatal(err)
	}

	player, err := audio.NewPlayer(eap.audioContext, d)

	if err != nil {
		log.Fatal(err)
	}

	player.SetVolume(eap.sfxVolume)

	result.player = player

	return result, nil
}

// SetVolumes sets the volumes of the audio provider
func (eap *AudioProvider) SetVolumes(bgmVolume, sfxVolume float64) {
	eap.sfxVolume = sfxVolume
	eap.bgmVolume = bgmVolume
}
