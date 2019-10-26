package Sound

import (
	"log"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

// Manager provides sound
type Manager struct {
	fileProvider Common.FileProvider
	audioContext *audio.Context // The Audio context
	bgmAudio     *audio.Player  // The audio player
	lastBgm      string
}

// CreateManager creates a sound provider
func CreateManager(fileProvider Common.FileProvider) *Manager {
	result := &Manager{
		fileProvider: fileProvider,
	}
	audioContext, err := audio.NewContext(22050)
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
	go func() {
		if v.bgmAudio != nil {
			v.bgmAudio.Close()
		}
		audioData := v.fileProvider.LoadFile(song)
		d, err := wav.Decode(v.audioContext, audio.BytesReadSeekCloser(audioData))
		if err != nil {
			log.Fatal(err)
		}
		s := audio.NewInfiniteLoop(d, int64(len(audioData)))

		v.bgmAudio, err = audio.NewPlayer(v.audioContext, s)
		if err != nil {
			log.Fatal(err)
		}
		// Play the infinite-length stream. This never ends.
		v.bgmAudio.Rewind()
		v.bgmAudio.Play()
	}()
}
