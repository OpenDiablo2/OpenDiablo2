package d2audio

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2asset"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

// Manager provides sound
type Manager struct {
	audioContext *audio.Context // The Audio context
	bgmAudio     *audio.Player  // The audio player
	lastBgm      string
	sfxVolume    float64
	bgmVolume    float64
}

// CreateManager creates a sound provider
func CreateManager() *Manager {
	result := &Manager{}
	audioContext, err := audio.NewContext(44100)
	if err != nil {
		log.Fatal(err)
	}
	result.audioContext = audioContext
	return result
}

// PlayBGM plays an infinitely looping background track
func (v *Manager) PlayBGM(song string) {
	if v.lastBgm == song {
		return
	}
	v.lastBgm = song
	if song == "" && v.bgmAudio != nil && v.bgmAudio.IsPlaying() {
		_ = v.bgmAudio.Pause()
		return
	}
	go func() {
		if v.bgmAudio != nil {
			err := v.bgmAudio.Close()
			if err != nil {
				log.Panic(err)
			}
		}
		audioData, err := d2asset.LoadFile(song)
		if err != nil {
			panic(err)
		}
		d, err := wav.Decode(v.audioContext, audio.BytesReadSeekCloser(audioData))
		if err != nil {
			log.Fatal(err)
		}
		s := audio.NewInfiniteLoop(d, d.Length())
		v.bgmAudio, err = audio.NewPlayer(v.audioContext, s)
		if err != nil {
			log.Fatal(err)
		}
		v.bgmAudio.SetVolume(v.bgmVolume)
		// Play the infinite-length stream. This never ends.
		err = v.bgmAudio.Rewind()
		if err != nil {
			panic(err)
		}
		err = v.bgmAudio.Play()
		if err != nil {
			panic(err)
		}
	}()
}

func (v *Manager) LoadSoundEffect(sfx string) *SoundEffect {
	result := CreateSoundEffect(sfx, v.audioContext, v.sfxVolume)
	return result
}

func (v *Manager) SetVolumes(bgmVolume, sfxVolume float64) {
	v.sfxVolume = sfxVolume
	v.bgmVolume = bgmVolume
}
