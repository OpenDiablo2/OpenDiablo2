package d2scene

import (
	"fmt"
	"math"
	"os"

	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2mapengine"

	"github.com/OpenDiablo2/OpenDiablo2/d2helper"

	"github.com/OpenDiablo2/OpenDiablo2/d2core"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2ui"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type MapEngineTest struct {
	uiManager     *d2ui.Manager
	soundManager  *d2audio.Manager
	fileProvider  d2interface.FileProvider
	sceneProvider d2interface.SceneProvider
	gameState     *d2core.GameState
	mapEngine     *d2mapengine.Engine
	currentRegion int
	keyLocked     bool
}

func CreateMapEngineTest(
	fileProvider d2interface.FileProvider,
	sceneProvider d2interface.SceneProvider,
	uiManager *d2ui.Manager,
	soundManager *d2audio.Manager,
	currentRegion int) *MapEngineTest {
	result := &MapEngineTest{
		fileProvider:  fileProvider,
		uiManager:     uiManager,
		soundManager:  soundManager,
		sceneProvider: sceneProvider,
		currentRegion: currentRegion,
		keyLocked:     false,
	}
	result.gameState = d2core.CreateTestGameState()
	return result
}

type RegionSpec struct {
	regionType  d2enum.RegionIdType
	levelPreset int
}

var regions []RegionSpec = []RegionSpec{
	{d2enum.RegionAct1Tristram, 300},
	{d2enum.RegionAct1Cathedral, 257},
	{d2enum.RegionAct2Town, 301},
	// {d2enum.RegionAct2Harem, 353},
	{d2enum.RegionAct3Town, 529},
	{d2enum.RegionAct3Jungle, 574},
	{d2enum.RegionAct4Town, 797},
	{d2enum.RegonAct5Town, 863},
	{d2enum.RegionAct5IceCaves, 1038},
	{d2enum.RegionAct5Siege, 879},
	{d2enum.RegionAct5Lava, 105},
	{d2enum.RegionAct5Barricade, 880},
}

func (v *MapEngineTest) LoadRegionByIndex(n int) {
	if n == 0 {
		v.mapEngine.GenerateAct1Overworld()
		return
	}
	region := regions[n-1]

	v.mapEngine = d2mapengine.CreateMapEngine(v.gameState, v.soundManager, v.fileProvider) // necessary for map name update
	v.mapEngine.OffsetY = 0
	v.mapEngine.OffsetX = 0
	v.mapEngine.GenerateMap(region.regionType, region.levelPreset)
}

func (v *MapEngineTest) Load() []func() {
	// TODO: Game seed comes from the game state object

	v.soundManager.PlayBGM("")
	return []func(){
		func() {
			v.mapEngine = d2mapengine.CreateMapEngine(v.gameState, v.soundManager, v.fileProvider)

			v.LoadRegionByIndex(v.currentRegion)
			// v.mapEngine.GenerateAct1Overworld()
			// v.mapEngine.GenerateMap(d2enum.RegionAct1Tristram, 300)
			// v.mapEngine.GenerateMap(d2enum.RegionAct1Cathedral, 257)
			//v.mapEngine.GenerateMap(d2enum.RegionAct2Town, 301)
			//v.mapEngine.GenerateMap(d2enum.RegionAct2Harem, 353) // Crashes on dcc load
			//v.mapEngine.GenerateMap(d2enum.RegionAct3Town, 529)
			//v.mapEngine.GenerateMap(d2enum.RegionAct3Jungle, 574)
			//v.mapEngine.GenerateMap(d2enum.RegionAct4Town, 797) // Broken height of large objects
			//v.mapEngine.GenerateMap(d2enum.RegonAct5Town, 863) // Completely broken!!
			//v.mapEngine.GenerateMap(d2enum.RegionAct5IceCaves, 1038) // Completely broken!
			//v.mapEngine.GenerateMap(d2enum.RegionAct5Siege, 879) // Completely broken!
			//v.mapEngine.GenerateMap(d2enum.RegionAct5Lava, 1057) // Broken
			//v.mapEngine.GenerateMap(d2enum.RegionAct5Barricade, 880) // Broken

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
	ebitenutil.DebugPrintAt(screen, line, 5, 5)
	ebitenutil.DebugPrintAt(screen, "Map: "+curRegion.Region.LevelType.Name, 5, 17)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%v [%v]", curRegion.Region.RegionPath, v.currentRegion), 5, 29)
	ebitenutil.DebugPrintAt(screen, "N - next map, P - previous map", 5, 41)
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

	if inpututil.IsKeyJustPressed(ebiten.KeyF7) {
		if v.mapEngine.ShowTiles < 2 {
			v.mapEngine.ShowTiles++
		} else {
			v.mapEngine.ShowTiles = 0
		}
	}

	if v.uiManager.KeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	if v.uiManager.KeyPressed(ebiten.KeyN) && !v.keyLocked {
		v.currentRegion++
		if v.currentRegion == len(regions) {
			v.currentRegion = 0
		}
		v.keyLocked = true
		fmt.Println("---")
		v.sceneProvider.SetNextScene(v)
		return
	}

	if v.uiManager.KeyPressed(ebiten.KeyP) && !v.keyLocked {
		v.currentRegion--
		if v.currentRegion == -1 {
			v.currentRegion = len(regions) - 1
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
