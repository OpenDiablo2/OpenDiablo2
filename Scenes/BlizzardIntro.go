package Scenes

import (
	"github.com/OpenDiablo2/OpenDiablo2/Common"
	"github.com/OpenDiablo2/OpenDiablo2/Video"
	"github.com/hajimehoshi/ebiten"
)

type BlizzardIntro struct {
	fileProvider  Common.FileProvider
	sceneProvider SceneProvider
	videoDecoder  *Video.BinkDecoder
}

func CreateBlizzardIntro(fileProvider Common.FileProvider, sceneProvider SceneProvider) *BlizzardIntro {
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
			v.videoDecoder = Video.CreateBinkDecoder(videoBytes)
		},
	}
}

func (v *BlizzardIntro) Unload() {

}

func (v *BlizzardIntro) Render(screen *ebiten.Image) {

}

func (v *BlizzardIntro) Update(tickTime float64) {

}
