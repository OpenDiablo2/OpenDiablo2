package ebiten

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/hajimehoshi/ebiten/audio/wav"

	"github.com/hajimehoshi/ebiten/audio"
)

type AudioProvider struct {
	audioContext *audio.Context // The Audio context
	bgmAudio     *audio.Player  // The audio player
	lastBgm      string
	sfxVolume    float64
	bgmVolume    float64
}

func CreateAudio() (*AudioProvider, error) {
	result := &AudioProvider{}
	var err error
	result.audioContext, err = audio.NewContext(44100)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func (eap *AudioProvider) PlayBGM(song string) {
	if eap.lastBgm == song {
		return
	}
	eap.lastBgm = song
	if song == "" && eap.bgmAudio != nil && eap.bgmAudio.IsPlaying() {
		_ = eap.bgmAudio.Pause()
		return
	}
	go func() {
		if eap.bgmAudio != nil {
			err := eap.bgmAudio.Close()
			if err != nil {
				log.Panic(err)
			}
		}
		audioData, err := d2asset.LoadFile(song)
		if err != nil {
			panic(err)
		}
		d, err := wav.Decode(eap.audioContext, audio.BytesReadSeekCloser(audioData))
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
	}()
}

func (eap *AudioProvider) LoadSoundEffect(sfx string) (d2audio.SoundEffect, error) {
	result := CreateSoundEffect(sfx, eap.audioContext, eap.sfxVolume) // TODO: Split
	return result, nil
}

func (eap *AudioProvider) SetVolumes(bgmVolume, sfxVolume float64) {
	eap.sfxVolume = sfxVolume
	eap.bgmVolume = bgmVolume
}
