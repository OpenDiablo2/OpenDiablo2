package d2scene

import (
	"fmt"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2helper"

	"github.com/OpenDiablo2/OpenDiablo2/d2core"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	_map "github.com/OpenDiablo2/OpenDiablo2/d2render/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2ui"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"os"
)

type MapEngineTest struct {
	uiManager     *d2ui.Manager
	soundManager  *d2audio.Manager
	fileProvider  d2interface.FileProvider
	sceneProvider d2interface.SceneProvider
	gameState     *d2core.GameState
	mapEngine     *_map.Engine
	currentRegion int
	levelPreset int
	fileIndex int
	keyLocked     bool
}

func CreateMapEngineTest(
	fileProvider d2interface.FileProvider,
	sceneProvider d2interface.SceneProvider,
	uiManager *d2ui.Manager,
	soundManager *d2audio.Manager,
	currentRegion int, levelPreset int) *MapEngineTest {
	result := &MapEngineTest{
		fileProvider:  fileProvider,
		uiManager:     uiManager,
		soundManager:  soundManager,
		sceneProvider: sceneProvider,
		currentRegion: currentRegion,
		levelPreset: levelPreset,
		fileIndex: -1,
		keyLocked:     false,
	}
	result.gameState = d2core.CreateTestGameState()
	return result
}

func (v *MapEngineTest) LoadRegionByIndex(n int, levelPreset, fileIndex int) {
	if n == 0 {
		v.mapEngine.GenerateAct1Overworld()
		return
	}
	// region := regions[n-1]

	v.mapEngine = _map.CreateMapEngine(v.gameState, v.soundManager, v.fileProvider) // necessary for map name update
	v.mapEngine.OffsetY = 0
	v.mapEngine.OffsetX = 0
	// v.mapEngine.GenerateMap(region.regionType, region.levelPreset)
	v.mapEngine.GenerateMap(d2enum.RegionIdType(n), levelPreset, fileIndex)
}

func (v *MapEngineTest) Load() []func() {
	// TODO: Game seed comes from the game state object

	v.soundManager.PlayBGM("")
	return []func(){
		func() {
			v.mapEngine = _map.CreateMapEngine(v.gameState, v.soundManager, v.fileProvider)

			v.LoadRegionByIndex(v.currentRegion, v.levelPreset, v.fileIndex)
		},
	}
}

func (v *MapEngineTest) Unload() {

}

func (v *MapEngineTest) Render(screen *ebiten.Image) {
	v.mapEngine.Render(screen)
	actualX := float64(v.uiManager.CursorX) - v.mapEngine.OffsetX
	actualY := float64(v.uiManager.CursorY) - v.mapEngine.OffsetY
	tileX, tileY := d2helper.ScreenToIso(actualX, actualY)
	subtileX := int(math.Ceil(math.Mod((tileX*10), 10))) / 2
	subtileY := int(math.Ceil(math.Mod((tileY*10), 10))) / 2
	curRegion := v.mapEngine.GetRegionAt(int(tileX), int(tileY))
	if curRegion == nil {
		return
	}
	line := fmt.Sprintf("%d, %d (Tile %d.%d, %d.%d)",
		int(math.Ceil(actualX)),
		int(math.Ceil(actualY)),
		int(math.Ceil(tileX))-curRegion.Rect.Left,
		subtileX,
		int(math.Ceil(tileY))-curRegion.Rect.Top,
		subtileY,
	)

	levelFilesToPick := make([]string, 0)
	for _, fileRecord := range curRegion.Region.LevelPreset.Files {
		if len(fileRecord) == 0 || fileRecord == "" || fileRecord == "0" {
			continue
		}
		levelFilesToPick = append(levelFilesToPick, fileRecord)
	}
	ebitenutil.DebugPrintAt(screen, line, 5, 5)
	ebitenutil.DebugPrintAt(screen, "Map: "+curRegion.Region.LevelType.Name, 5, 17)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%v: %v/%v [%v, %v]", curRegion.Region.RegionPath, v.fileIndex, len(levelFilesToPick), v.currentRegion, v.levelPreset), 5, 29)
	ebitenutil.DebugPrintAt(screen, "N - next region, P - previous region", 5, 41)
	ebitenutil.DebugPrintAt(screen, "Shift+N - next preset, Shift+P - previous preset", 5, 53)
	ebitenutil.DebugPrintAt(screen, "Ctrl+N - next file, Ctrl+P - previous file", 5, 65)
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
	if v.uiManager.KeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if v.uiManager.KeyPressed(ebiten.KeyN) && !v.keyLocked && v.uiManager.KeyPressed(ebiten.KeyControl) {
		v.fileIndex++
		if v.fileIndex == 1091 {
			v.fileIndex = 0
		}
		v.keyLocked = true
		v.sceneProvider.SetNextScene(v)
		return
	}

	if v.uiManager.KeyPressed(ebiten.KeyP) && !v.keyLocked && v.uiManager.KeyPressed(ebiten.KeyControl) {
		v.fileIndex--
		if v.fileIndex == 0 {
			v.fileIndex = 1090
		}
		v.keyLocked = true
		v.sceneProvider.SetNextScene(v)
		return
	}

	if v.uiManager.KeyPressed(ebiten.KeyN) && !v.keyLocked && v.uiManager.KeyPressed(ebiten.KeyShift) {
		v.levelPreset++
		if v.levelPreset == 1091 {
			v.levelPreset = 0
		}
		v.keyLocked = true
		v.sceneProvider.SetNextScene(v)
		return
	}

	if v.uiManager.KeyPressed(ebiten.KeyP) && !v.keyLocked && v.uiManager.KeyPressed(ebiten.KeyShift) {
		v.levelPreset--
		if v.levelPreset == 0 {
			v.levelPreset = 1090
		}
		v.keyLocked = true
		v.sceneProvider.SetNextScene(v)
		return
	}

	if v.uiManager.KeyPressed(ebiten.KeyN) && !v.keyLocked {
		v.currentRegion++
		v.levelPreset++
		if v.currentRegion == 36 {
			v.currentRegion = 0
		}
		v.keyLocked = true
		fmt.Println("---")
		v.sceneProvider.SetNextScene(v)
		return
	}

	if v.uiManager.KeyPressed(ebiten.KeyP) && !v.keyLocked {
		v.currentRegion--
		v.levelPreset--
		if v.currentRegion == -1 {
			v.currentRegion = 35
		}
		v.keyLocked = true
		fmt.Println("---")
		v.sceneProvider.SetNextScene(v)
		return
	}

	//FIXME: do it better
	if !v.uiManager.KeyPressed(ebiten.KeyP) && !v.uiManager.KeyPressed(ebiten.KeyN) {
		v.keyLocked = false
	}
}
