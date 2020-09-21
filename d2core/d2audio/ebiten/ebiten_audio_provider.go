// Package ebiten contains ebiten's implementation of the audio interface
package ebiten

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

const sampleRate = 44100

var _ d2interface.AudioProvider = &AudioProvider{} // Static check to confirm struct conforms to interface

// CreateAudio creates an instance of ebiten's audio provider
func CreateAudio(am *d2asset.AssetManager) (*AudioProvider, error) {
	result := &AudioProvider{
		asset: am,
	}

	var err error
	result.audioContext, err = audio.NewContext(sampleRate)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, nil
}

// AudioProvider represents a provider capable of playing audio
type AudioProvider struct {
	asset        *d2asset.AssetManager
	audioContext *audio.Context // The Audio context
	bgmAudio     *audio.Player  // The audio player
	lastBgm      string
	sfxVolume    float64
	bgmVolume    float64
}

// PlayBGM loads an audio stream and plays it in the background
func (eap *AudioProvider) PlayBGM(song string) {
	if eap.lastBgm == song {
		return
	}

	eap.lastBgm = song

	if song == "" && eap.bgmAudio != nil && eap.bgmAudio.IsPlaying() {
		err := eap.bgmAudio.Pause()
		if err != nil {
			log.Print(err)
		}

		return
	}

	if eap.bgmAudio != nil {
		err := eap.bgmAudio.Close()

		if err != nil {
			log.Panic(err)
		}
	}

	audioStream, err := eap.asset.LoadFileStream(song)

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

// LoadSound loads a sound affect so that it canb e played
func (eap *AudioProvider) LoadSound(sfx string, loop, bgm bool) (d2interface.SoundEffect, error) {
	volume := eap.sfxVolume
	if bgm {
		volume = eap.bgmVolume
	}

	result := eap.createSoundEffect(sfx, eap.audioContext, loop)

	result.volumeScale = volume
	result.SetVolume(volume)

	return result, nil
}

// SetVolumes sets the volumes of the audio provider
func (eap *AudioProvider) SetVolumes(bgmVolume, sfxVolume float64) {
	eap.sfxVolume = sfxVolume
	eap.bgmVolume = bgmVolume
}

// createSoundEffect creates a new instance of ebiten's sound effect implementation.
func (eap *AudioProvider) createSoundEffect(sfx string, context *audio.Context,
	loop bool) *SoundEffect {
	result := &SoundEffect{}

	soundFile := "/data/global/sfx/"

	if _, exists := eap.asset.Records.Sound.Details[sfx]; exists {
		soundEntry := eap.asset.Records.Sound.Details[sfx]
		soundFile += soundEntry.FileName
	} else {
		soundFile += sfx
	}

	audioData, err := eap.asset.LoadFileStream(soundFile)

	if err != nil {
		audioData, err = eap.asset.LoadFileStream("/data/global/music/" + sfx)
	}

	if err != nil {
		panic(err)
	}

	d, err := wav.Decode(context, audioData)

	if err != nil {
		log.Fatal(err)
	}

	var player *audio.Player

	if loop {
		s := audio.NewInfiniteLoop(d, d.Length())
		result.panStream = newPanStreamFromReader(s)
		player, err = audio.NewPlayer(context, result.panStream)
	} else {
		result.panStream = newPanStreamFromReader(d)
		player, err = audio.NewPlayer(context, result.panStream)
	}

	if err != nil {
		log.Fatal(err)
	}

	result.player = player

	return result
}
