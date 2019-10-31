package Scenes

import (
	"math/rand"
	"time"

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
	region        *MapEngine.Region
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
	// TODO: Game seed comes from the game state object
	randomSource := rand.NewSource(time.Now().UnixNano())
	v.soundManager.PlayBGM("")
	return []func(){
		func() {
			v.region = MapEngine.LoadRegion(randomSource, MapEngine.RegionAct1Tristram, 300, v.fileProvider)
		},
	}
}

func (v *MapEngineTest) Unload() {

}

func (v *MapEngineTest) Render(screen *ebiten.Image) {
	v.region.RenderTile(300, 300, 0, 0, MapEngine.RegionLayerTypeFloors, 0, screen)
}

func (v *MapEngineTest) Update(tickTime float64) {

}
