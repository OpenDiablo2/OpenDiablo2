package ebiten

import (
	"io"
	"math"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

type panStream struct {
	io.ReadSeeker
	pan float64 // -1: left; 0: center; 1: right
}

const (
	bitsPerByte = 8
)

func newPanStreamFromReader(src io.ReadSeeker) *panStream {
	return &panStream{
		ReadSeeker: src,
		pan:        0,
	}
}

func (s *panStream) Read(p []byte) (n int, err error) {
	n, err = s.ReadSeeker.Read(p)
	if err != nil {
		return
	}

	ls := math.Min(s.pan*-1+1, 1)
	rs := math.Min(s.pan+1, 1)

	for i := 0; i < len(p); i += 4 {
		lc := int16(float64(int16(p[i])|int16(p[i+1])<<bitsPerByte) * ls)
		rc := int16(float64(int16(p[i+2])|int16(p[i+3])<<bitsPerByte) * rs)

		p[i] = byte(lc)
		p[i+1] = byte(lc >> bitsPerByte)
		p[i+2] = byte(rc)
		p[i+3] = byte(rc >> bitsPerByte)
	}

	return
}

// SoundEffect represents an ebiten implementation of a sound effect
type SoundEffect struct {
	player      *audio.Player
	volumeScale float64
	panStream   *panStream
}

// SetPan sets the audio pan, left is -1.0, center is 0.0, right is 1.0
func (v *SoundEffect) SetPan(pan float64) {
	v.panStream.pan = pan
}

// SetVolume ets the volume
func (v *SoundEffect) SetVolume(volume float64) {
	v.player.SetVolume(volume * v.volumeScale)
}

// IsPlaying returns a bool for whether or not the sound is currently playing
func (v *SoundEffect) IsPlaying() bool {
	return v.player.IsPlaying()
}

// Play plays the sound effect
func (v *SoundEffect) Play() {
	err := v.player.Rewind()

	if err != nil {
		panic(err)
	}

	v.player.Play()
}

// Stop stops the sound effect
func (v *SoundEffect) Stop() {
	v.player.Pause()
}
