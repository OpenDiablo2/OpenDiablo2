package d2gamescene

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2video"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/hajimehoshi/ebiten"
)

type BlizzardIntro struct {
	videoDecoder *d2video.BinkDecoder
}

func CreateBlizzardIntro() *BlizzardIntro {
	result := &BlizzardIntro{}

	return result
}

func (v *BlizzardIntro) Load() []func() {
	return []func(){
		func() {
			videoBytes, err := d2asset.LoadFile("/data/local/video/BlizNorth640x480.bik")
			if err != nil {
				panic(err)
			}
			v.videoDecoder = d2video.CreateBinkDecoder(videoBytes)
		},
	}
}

func (v *BlizzardIntro) Unload() {

}

func (v *BlizzardIntro) Render(screen *ebiten.Image) {

}

func (v *BlizzardIntro) Update(tickTime float64) {

}
