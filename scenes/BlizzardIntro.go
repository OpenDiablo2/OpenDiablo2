package scenes

import (
	"github.com/OpenDiablo2/OpenDiablo2/common"
	"github.com/OpenDiablo2/OpenDiablo2/video"
	"github.com/hajimehoshi/ebiten"
)

type BlizzardIntro struct {
	fileProvider  common.FileProvider
	sceneProvider SceneProvider
	videoDecoder  *video.BinkDecoder
}

func CreateBlizzardIntro(fileProvider common.FileProvider, sceneProvider SceneProvider) *BlizzardIntro {
	result := &BlizzardIntro{
		fileProvider:  fileProvider,
		sceneProvider: sceneProvider,
	}

	return result
}

func (v *BlizzardIntro) Load() []func() {
	return []func(){
		func() {
			videoBytes := v.fileProvider.LoadFile("/data/local/video/BlizNorth640x480.bik")
			v.videoDecoder = video.CreateBinkDecoder(videoBytes)
		},
	}
}

func (v *BlizzardIntro) Unload() {

}

func (v *BlizzardIntro) Render(screen *ebiten.Image) {

}

func (v *BlizzardIntro) Update(tickTime float64) {

}
