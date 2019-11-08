package tests

import (
	"testing"

	"github.com/hajimehoshi/ebiten"

	_map "github.com/OpenDiablo2/OpenDiablo2/map"

	"github.com/OpenDiablo2/OpenDiablo2/common"
	"github.com/OpenDiablo2/OpenDiablo2/core"
	"github.com/OpenDiablo2/OpenDiablo2/mpq"
)

func TestMapGenerationPerformance(t *testing.T) {
	mpq.InitializeCryptoBuffer()
	common.ConfigBasePath = "../"

	engine := core.CreateEngine()
	gameState := common.CreateGameState()
	mapEngine := _map.CreateMapEngine(gameState, engine.SoundManager, engine)
	mapEngine.GenerateAct1Overworld()
	surface, _ := ebiten.NewImage(800, 600, ebiten.FilterNearest)
	for y := 0; y < 1000; y++ {
		mapEngine.Render(surface)
		mapEngine.OffsetY = float64(-y)
	}

}
