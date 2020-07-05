package ebiten

import (
	"github.com/hajimehoshi/ebiten/audio"
)

// SoundEffect represents an ebiten implementation of a sound effect
type SoundEffect struct {
	player *audio.Player
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
