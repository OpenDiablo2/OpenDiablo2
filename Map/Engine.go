package Map

import (
	"image"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Sound"
	"github.com/hajimehoshi/ebiten"
)

type EngineRegion struct {
	Rect   image.Rectangle
	Region Region
}

type Engine struct {
	soundManager *Sound.Manager
	gameState    *Common.GameState
	fileProvider Common.FileProvider
	regions      []*EngineRegion
}

func CreateMapEngine(gameState *Common.GameState, soundManager *Sound.Manager, fileProvider Common.FileProvider) *Engine {
	result := &Engine{
		gameState:    gameState,
		soundManager: soundManager,
		fileProvider: fileProvider,
		regions:      make([]*EngineRegion, 0),
	}
	return result
}

func (v *Engine) GenerateMap(regionType RegionIdType, levelPreset int) {
	//randomSource := rand.NewSource(v.gameState.Seed)
	//region := LoadRegion(randomSource, regionType, levelPreset, v.fileProvider)
	v.soundManager.PlayBGM("/data/global/music/Act1/tristram.wav")
	v.ReloadMapCache()
}

func (v *Engine) ReloadMapCache() {

}

func (v *Engine) Render(target *ebiten.Image) {
	//v.region.RenderTile(300, 300, 0, 0, Map.RegionLayerTypeFloors, 0, screen)
}
