package Scenes

import (
	"github.com/OpenDiablo2/OpenDiablo2/Common"
	"github.com/OpenDiablo2/OpenDiablo2/Map"
	"github.com/OpenDiablo2/OpenDiablo2/Sound"
	"github.com/OpenDiablo2/OpenDiablo2/UI"
	"github.com/hajimehoshi/ebiten"
)

type MapEngineTest struct {
	uiManager     *UI.Manager
	soundManager  *Sound.Manager
	fileProvider  Common.FileProvider
	sceneProvider SceneProvider
	gameState     *Common.GameState
	mapEngine     *Map.Engine
}

func CreateMapEngineTest(
	fileProvider Common.FileProvider,
	sceneProvider SceneProvider,
	uiManager *UI.Manager,
	soundManager *Sound.Manager) *MapEngineTest {
	result := &MapEngineTest{
		fileProvider:  fileProvider,
		uiManager:     uiManager,
		soundManager:  soundManager,
		sceneProvider: sceneProvider,
	}
	result.gameState = Common.CreateGameState()
	return result
}

func (v *MapEngineTest) Load() []func() {
	// TODO: Game seed comes from the game state object

	v.soundManager.PlayBGM("")
	return []func(){
		func() {
			v.mapEngine = Map.CreateMapEngine(v.gameState, v.soundManager, v.fileProvider)

			v.mapEngine.GenerateMap(Map.RegionAct1Town, 1)
			//v.mapEngine.GenerateMap(Map.RegionAct1Tristram, 300)
			//v.mapEngine.GenerateMap(Map.RegionAct1Cathedral, 257)
			//v.mapEngine.GenerateMap(Map.RegionAct2Town, 301) // Broken rendering
			//v.mapEngine.GenerateMap(Map.RegionAct2Harem, 353)
			//v.mapEngine.GenerateMap(Map.RegionAct3Town, 529)
			//v.mapEngine.GenerateMap(Map.RegionAct3Jungle, 574)
			//v.mapEngine.GenerateMap(Map.RegionAct4Town, 797)
			//v.mapEngine.GenerateMap(Map.RegonAct5Town, 863)
			//v.mapEngine.GenerateMap(Map.RegionAct5IceCaves, 1038)
			//v.mapEngine.GenerateMap(Map.RegionAct5Siege, 879)
			//v.mapEngine.GenerateMap(Map.RegionAct5Lava, 1057) // PALETTE ISSUE
			//v.mapEngine.GenerateMap(Map.RegionAct5Barricade, 880)

		},
	}
}

func (v *MapEngineTest) Unload() {

}

func (v *MapEngineTest) Render(screen *ebiten.Image) {
	v.mapEngine.Render(screen)
}

func (v *MapEngineTest) Update(tickTime float64) {
	if v.uiManager.KeyPressed(ebiten.KeyDown) {
		v.mapEngine.OffsetY -= tickTime * 800
	}
	if v.uiManager.KeyPressed(ebiten.KeyUp) {
		v.mapEngine.OffsetY += tickTime * 800
	}
	if v.uiManager.KeyPressed(ebiten.KeyLeft) {
		v.mapEngine.OffsetX += tickTime * 800
	}
	if v.uiManager.KeyPressed(ebiten.KeyRight) {
		v.mapEngine.OffsetX -= tickTime * 800
	}
}
