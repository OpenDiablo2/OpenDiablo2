package ebiten

import (
	"log"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// SoundEffect represents an ebiten implementation of a sound effect
type SoundEffect struct {
	player *audio.Player
}

// CreateSoundEffect creates a new instance of ebiten's sound effect implementation.
func CreateSoundEffect(sfx string, context *audio.Context, volume float64, loop bool) *SoundEffect {
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
		player, err = audio.NewPlayer(context, s)
	} else {
		player, err = audio.NewPlayer(context, d)
	}

	if err != nil {
		log.Fatal(err)
	}

	player.SetVolume(volume)

	result.player = player

	return result
}

func (v *SoundEffect) SetVolume(volume float64) {
	v.player.SetVolume(volume)
}

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
