package d2audio

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2data/datadict"

	"github.com/hajimehoshi/ebiten/audio/wav"

	"github.com/hajimehoshi/ebiten/audio"
)

type SoundEffect struct {
	player *audio.Player
}

func CreateSoundEffect(sfx string, fileProvider d2interface.FileProvider, context *audio.Context, volume float64) *SoundEffect {
	result := &SoundEffect{}
	var soundFile string
	if _, exists := datadict.Sounds[sfx]; exists {
		soundEntry := datadict.Sounds[sfx]
		soundFile = soundEntry.FileName
	} else {
		soundFile = sfx
	}
	audioData := fileProvider.LoadFile(soundFile)
	d, err := wav.Decode(context, audio.BytesReadSeekCloser(audioData))
	if err != nil {
		log.Fatal(err)
	}

	player, err := audio.NewPlayer(context, d)
	if err != nil {
		log.Fatal(err)
	}
	player.SetVolume(volume)
	result.player = player
	return result
}

func (v *SoundEffect) Play() {
	v.player.Rewind()
	v.player.Play()
}

func (v *SoundEffect) Stop() {
	v.player.Pause()
}
