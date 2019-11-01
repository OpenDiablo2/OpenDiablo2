package Map

import (
	"image"
	"math/rand"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Sound"
	"github.com/hajimehoshi/ebiten"
)

type EngineRegion struct {
	Rect   image.Rectangle
	Region *Region
}

type Engine struct {
	soundManager *Sound.Manager
	gameState    *Common.GameState
	fileProvider Common.FileProvider
	regions      []EngineRegion
	OffsetX      float64
	OffsetY      float64
}

func CreateMapEngine(gameState *Common.GameState, soundManager *Sound.Manager, fileProvider Common.FileProvider) *Engine {
	result := &Engine{
		gameState:    gameState,
		soundManager: soundManager,
		fileProvider: fileProvider,
		regions:      make([]EngineRegion, 0),
	}
	return result
}

func (v *Engine) GenerateMap(regionType RegionIdType, levelPreset int) {
	randomSource := rand.NewSource(v.gameState.Seed)
	region := LoadRegion(randomSource, regionType, levelPreset, v.fileProvider)
	v.regions = append(v.regions, EngineRegion{
		Rect:   image.Rectangle{image.Point{0, 0}, image.Point{int(region.TileWidth), int(region.TileHeight)}},
		Region: region,
	})
	v.soundManager.PlayBGM("/data/global/music/Act1/tristram.wav") // TODO: Temp stuff here
}

func (v *Engine) Render(target *ebiten.Image) {
	// TODO: Temporary hack for testing
	for y := 0; y < int(v.regions[0].Region.TileHeight); y++ {
		offX := -(y * 80)
		offY := y * 40
		for x := 0; x < int(v.regions[0].Region.TileWidth); x++ {
			v.regions[0].Region.RenderTile(400+offX+int(v.OffsetX), offY+int(v.OffsetY), x, y, RegionLayerTypeFloors, 0, target)
			offX += 80
			offY += 40
		}
	}

}
