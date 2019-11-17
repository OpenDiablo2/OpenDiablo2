package d2audio

import (
	"log"

	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

// Manager provides sound
type Manager struct {
	fileProvider d2interface.FileProvider
	audioContext *audio.Context // The Audio context
	bgmAudio     *audio.Player  // The audio player
	lastBgm      string
	sfxVolume    float64
	bgmVolume    float64
}

// CreateManager creates a sound provider
func CreateManager(fileProvider d2interface.FileProvider) *Manager {
	result := &Manager{
		fileProvider: fileProvider,
	}
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
		audioData := v.fileProvider.LoadFile(song)
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
	result := CreateSoundEffect(sfx, v.fileProvider, v.audioContext, v.sfxVolume)
	return result
}

func (v *Manager) SetVolumes(bgmVolume, sfxVolume float64) {
	v.sfxVolume = sfxVolume
	v.bgmVolume = bgmVolume
}
