package d2scene

import (
	"fmt"
	"math"
	"os"

	"github.com/OpenDiablo2/OpenDiablo2/d2helper"

	"github.com/OpenDiablo2/OpenDiablo2/d2core"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	_map "github.com/OpenDiablo2/OpenDiablo2/d2render/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2ui"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type RegionSpec struct {
	regionType       d2enum.RegionIdType
	startPresetIndex int
	endPresetIndex   int
	extra            []int
}

var regions []RegionSpec = []RegionSpec{
	//Act I
	{d2enum.RegionAct1Town, 1, 3, []int{}},
	{d2enum.RegionAct1Wilderness, 4, 52, []int{
		108,
		160, 161, 162, 163, 164,
	}},
	{d2enum.RegionAct1Cave, 53, 107, []int{}},
	{d2enum.RegionAct1Crypt, 109, 159, []int{}},
	{d2enum.RegionAct1Monestary, 165, 165, []int{}},
	{d2enum.RegionAct1Courtyard, 166, 166, []int{256}},
	{d2enum.RegionAct1Barracks, 167, 205, []int{}},
	{d2enum.RegionAct1Jail, 206, 255, []int{}},
	{d2enum.RegionAct1Cathedral, 257, 257, []int{}},
	{d2enum.RegionAct1Catacombs, 258, 299, []int{}},
	{d2enum.RegionAct1Tristram, 300, 300, []int{}},

	//Act II
	{d2enum.RegionAct2Town, 301, 301, []int{}},
	{d2enum.RegionAct2Sewer, 302, 352, []int{}},
	{d2enum.RegionAct2Harem, 353, 357, []int{}},
	{d2enum.RegionAct2Basement, 358, 361, []int{}},
	{d2enum.RegionAct2Desert, 362, 413, []int{}},
	{d2enum.RegionAct2Tomb, 414, 481, []int{}},
	{d2enum.RegionAct2Lair, 482, 509, []int{}},
	{d2enum.RegionAct2Arcane, 510, 528, []int{}},

	//Act III
	{d2enum.RegionAct3Town, 529, 529, []int{}},
	{d2enum.RegionAct3Jungle, 530, 604, []int{}},
	{d2enum.RegionAct3Kurast, 605, 658, []int{
		748, 749, 750, 751, 752, 753, 754,
		755, 756, 757, 758, 759, 760, 761, 762, 763, 764, 765, 766, 767, 768, 769, 770, 771, 772, 773, 774, 775, 776, 777, 778, 779, 780, 781, 782, 783, 784, 785, 786, 787, 788, 789, 790, 791, 792, 793, 794, 795, 796,
		//yeah, i know =(
	}},
	{d2enum.RegionAct3Spider, 659, 664, []int{}},
	{d2enum.RegionAct3Dungeon, 665, 704, []int{}},
	{d2enum.RegionAct3Sewer, 705, 747, []int{}},

	//Act IV
	{d2enum.RegionAct4Town, 797, 798, []int{}},
	{d2enum.RegionAct4Mesa, 799, 835, []int{}},
	{d2enum.RegionAct4Lava, 836, 862, []int{}},

	//Act V -- broken or wrong order
	{d2enum.RegonAct5Town, 863, 864, []int{}},
	{d2enum.RegionAct5Siege, 865, 879, []int{}},
	{d2enum.RegionAct5Barricade, 880, 1002, []int{}},
	{d2enum.RegionAct5IceCaves, 1003, 1041, []int{}},
	{d2enum.RegionAct5Temple, 1042, 1052, []int{}},
	{d2enum.RegionAct5Baal, 1059, 1090, []int{}},
	{d2enum.RegionAct5Lava, 1053, 1058, []int{}},

}

type MapEngineTest struct {
	uiManager     *d2ui.Manager
	soundManager  *d2audio.Manager
	fileProvider  d2interface.FileProvider
	sceneProvider d2interface.SceneProvider
	gameState     *d2core.GameState
	mapEngine     *_map.Engine
	currentRegion int
	levelPreset   int
	fileIndex     int
	keyLocked     bool
	regionSpec    RegionSpec
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
		levelPreset:   levelPreset,
		fileIndex:     -1,
		keyLocked:     false,
		regionSpec:    RegionSpec{},
	}
	result.gameState = d2core.CreateTestGameState()
	return result
}

func (v *MapEngineTest) LoadRegionByIndex(n int, levelPreset, fileIndex int) {
	for _, spec := range regions {
		if spec.regionType == d2enum.RegionIdType(n) {
			v.regionSpec = spec
			inExtra := false
			for _, e := range spec.extra {
				if e == levelPreset {
					inExtra = true
					break
				}
			}
			if !inExtra {
				if levelPreset < spec.startPresetIndex {
					levelPreset = spec.startPresetIndex
				}

				if levelPreset > spec.endPresetIndex {
					levelPreset = spec.endPresetIndex
				}
			}
			v.levelPreset = levelPreset
		}
	}

	if n == 0 {
		v.mapEngine.GenerateAct1Overworld()
		return
	}

	v.mapEngine = _map.CreateMapEngine(v.gameState, v.soundManager, v.fileProvider) // necessary for map name update
	v.mapEngine.OffsetY = 0
	v.mapEngine.OffsetX = 0
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

	if v.uiManager.KeyPressed(ebiten.KeyN) && !v.keyLocked && v.uiManager.KeyPressed(ebiten.KeyControl) {
		v.fileIndex++
		if v.fileIndex == 10 {
			v.fileIndex = 0
		}
		v.keyLocked = true
		v.sceneProvider.SetNextScene(v)
		return
	}

	if v.uiManager.KeyPressed(ebiten.KeyP) && !v.keyLocked && v.uiManager.KeyPressed(ebiten.KeyControl) {
		v.fileIndex--
		if v.fileIndex == 0 {
			v.fileIndex = 10
		}
		v.keyLocked = true
		v.sceneProvider.SetNextScene(v)
		return
	}

	if v.uiManager.KeyPressed(ebiten.KeyN) && !v.keyLocked && v.uiManager.KeyPressed(ebiten.KeyShift) {
		v.levelPreset++
		if v.levelPreset > v.regionSpec.endPresetIndex {
			v.levelPreset = v.regionSpec.startPresetIndex
		}
		v.keyLocked = true
		v.sceneProvider.SetNextScene(v)
		return
	}

	if v.uiManager.KeyPressed(ebiten.KeyP) && !v.keyLocked && v.uiManager.KeyPressed(ebiten.KeyShift) {
		v.levelPreset--
		if v.levelPreset < v.regionSpec.startPresetIndex {
			v.levelPreset = v.regionSpec.endPresetIndex
		}
		v.keyLocked = true
		v.sceneProvider.SetNextScene(v)
		return
	}

	if v.uiManager.KeyPressed(ebiten.KeyN) && !v.keyLocked {
		v.currentRegion++
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
