package ebiten

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

type EbitenSoundEffect struct {
	player *audio.Player
}

func CreateSoundEffect(sfx string, context *audio.Context, volume float64) *EbitenSoundEffect {
	result := &EbitenSoundEffect{}
	var soundFile string
	if _, exists := d2datadict.Sounds[sfx]; exists {
		soundEntry := d2datadict.Sounds[sfx]
		soundFile = soundEntry.FileName
	} else {
		soundFile = sfx
	}

	audioData, err := d2asset.LoadFile(soundFile)
	if err != nil {
		panic(err)
	}

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

func (v *EbitenSoundEffect) Play() {
	v.player.Rewind()
	v.player.Play()
}

func (v *EbitenSoundEffect) Stop() {
	v.player.Pause()
}
