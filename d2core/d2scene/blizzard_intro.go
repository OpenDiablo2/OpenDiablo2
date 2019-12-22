package d2scene

import (
	"github.com/OpenDiablo2/D2Shared/d2data/d2video"
	"github.com/OpenDiablo2/OpenDiablo2/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2corecommon/d2coreinterface"
	"github.com/hajimehoshi/ebiten"
)

type BlizzardIntro struct {
	sceneProvider d2coreinterface.SceneProvider
	videoDecoder  *d2video.BinkDecoder
}

func CreateBlizzardIntro(sceneProvider d2coreinterface.SceneProvider) *BlizzardIntro {
	result := &BlizzardIntro{sceneProvider: sceneProvider}

	return result
}

func (v *BlizzardIntro) Load() []func() {
	return []func(){
		func() {
			videoBytes := d2asset.MustLoadFile("/data/local/video/BlizNorth640x480.bik")
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
