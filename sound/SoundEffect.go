package sound

import (
	"log"

	"github.com/hajimehoshi/ebiten/audio/wav"

	"github.com/OpenDiablo2/OpenDiablo2/common"

	"github.com/hajimehoshi/ebiten/audio"
)

type SoundEffect struct {
	player *audio.Player
}

func CreateSoundEffect(sfx string, fileProvider common.FileProvider, context *audio.Context, volume float64) *SoundEffect {
	result := &SoundEffect{}
	var soundFile string
	if _, exists := common.Sounds[sfx]; exists {
		soundEntry := common.Sounds[sfx]
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
