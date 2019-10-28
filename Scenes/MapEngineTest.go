package Scenes

import (
	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/MapEngine"
	"github.com/essial/OpenDiablo2/Sound"
	"github.com/essial/OpenDiablo2/UI"
	"github.com/hajimehoshi/ebiten"
)

type MapEngineTest struct {
	uiManager     *UI.Manager
	soundManager  *Sound.Manager
	fileProvider  Common.FileProvider
	sceneProvider SceneProvider
}

func CreateMapEngineTest(fileProvider Common.FileProvider, sceneProvider SceneProvider, uiManager *UI.Manager, soundManager *Sound.Manager) *MapEngineTest {
	result := &MapEngineTest{
		fileProvider:  fileProvider,
		uiManager:     uiManager,
		soundManager:  soundManager,
		sceneProvider: sceneProvider,
	}
	return result
}

func (v *MapEngineTest) Load() []func() {
	v.soundManager.PlayBGM("")
	return []func(){
		func() {
			_ = MapEngine.LoadDS1("/data/global/tiles/ACT1/town/townE1.ds1", v.fileProvider)
		},
	}
}

func (v *MapEngineTest) Unload() {

}

func (v *MapEngineTest) Render(screen *ebiten.Image) {

}

func (v *MapEngineTest) Update(tickTime float64) {

}
