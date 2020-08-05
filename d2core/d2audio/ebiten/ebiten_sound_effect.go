package ebiten

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

type panStream struct {
	audio.ReadSeekCloser
	pan float64 // -1: left; 0: center; 1: right
}

const (
	bitsPerByte = 8
)

func newPanStreamFromReader(src audio.ReadSeekCloser) *panStream {
	return &panStream{
		ReadSeekCloser: src,
		pan:            0,
	}
}

func (s *panStream) Read(p []byte) (n int, err error) {
	n, err = s.ReadSeekCloser.Read(p)
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

// CreateSoundEffect creates a new instance of ebiten's sound effect implementation.
func CreateSoundEffect(sfx string, context *audio.Context, loop bool) *SoundEffect {
	result := &SoundEffect{}

	soundFile := "/data/global/sfx/"

	if _, exists := d2datadict.Sounds[sfx]; exists {
		soundEntry := d2datadict.Sounds[sfx]
		soundFile += soundEntry.FileName
	} else {
		soundFile += sfx
	}

	audioData, err := d2asset.LoadFileStream(soundFile)

	if err != nil {
		audioData, err = d2asset.LoadFileStream("/data/global/music/" + sfx)
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

	err = v.player.Play()

	if err != nil {
		panic(err)
	}
}

// Stop stops the sound effect
func (v *SoundEffect) Stop() {
	err := v.player.Pause()

	if err != nil {
		panic(err)
	}
}
