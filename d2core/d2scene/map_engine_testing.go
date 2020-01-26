package d2scene

import (
	"math"
	"os"

	"github.com/OpenDiablo2/OpenDiablo2/d2core"
	"github.com/OpenDiablo2/OpenDiablo2/d2input"

	"github.com/OpenDiablo2/OpenDiablo2/d2corecommon/d2coreinterface"

	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2surface"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2ui"
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
	sceneProvider d2coreinterface.SceneProvider
	gameState     *d2core.GameState
	mapEngine     *d2mapengine.MapEngine

	//TODO: this is region specific properties, should be refactored for multi-region rendering
	currentRegion int
	levelPreset   int
	fileIndex     int
	regionSpec    RegionSpec
	filesCount    int
	debugVisLevel int
}

func CreateMapEngineTest(sceneProvider d2coreinterface.SceneProvider, uiManager *d2ui.Manager, soundManager *d2audio.Manager, currentRegion int, levelPreset int) *MapEngineTest {
	result := &MapEngineTest{
		uiManager:     uiManager,
		soundManager:  soundManager,
		sceneProvider: sceneProvider,
		currentRegion: currentRegion,
		levelPreset:   levelPreset,
		fileIndex:     0,
		regionSpec:    RegionSpec{},
		filesCount:    0,
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
	} else {
		v.mapEngine = d2mapengine.CreateMapEngine(v.gameState, v.soundManager) // necessary for map name update
		v.mapEngine.GenerateMap(d2enum.RegionIdType(n), levelPreset, fileIndex)
	}

	v.mapEngine.MoveCameraTo(v.mapEngine.WorldToOrtho(v.mapEngine.GetCenterPosition()))
}

func (v *MapEngineTest) Load() []func() {
	// TODO: Game seed comes from the game state object
	d2input.BindHandler(v)
	v.soundManager.PlayBGM("")
	return []func(){
		func() {
			v.mapEngine = d2mapengine.CreateMapEngine(v.gameState, v.soundManager)
			v.LoadRegionByIndex(v.currentRegion, v.levelPreset, v.fileIndex)
		},
	}
}

func (v *MapEngineTest) Unload() {
	d2input.UnbindHandler(v)
}

func (v *MapEngineTest) Render(screen *d2surface.Surface) {
	v.mapEngine.Render(screen)
	screenX := v.uiManager.CursorX
	screenY := v.uiManager.CursorY
	worldX, worldY := v.mapEngine.ScreenToWorld(screenX, screenY)
	subtileX := int(math.Ceil(math.Mod((worldX*10), 10))) / 2
	subtileY := int(math.Ceil(math.Mod((worldY*10), 10))) / 2
	curRegion := v.mapEngine.GetRegionAtTile(int(worldX), int(worldY))
	if curRegion == nil {
		return
	}

	tileRect := curRegion.GetTileRect()

	levelFilesToPick := make([]string, 0)
	fileIndex := v.fileIndex
	levelPreset := curRegion.GetLevelPreset()
	regionPath := curRegion.GetPath()
	for n, fileRecord := range levelPreset.Files {
		if len(fileRecord) == 0 || fileRecord == "" || fileRecord == "0" {
			continue
		}
		levelFilesToPick = append(levelFilesToPick, fileRecord)
		if fileRecord == regionPath {
			fileIndex = n
		}
	}
	if v.fileIndex == -1 {
		v.fileIndex = fileIndex
	}
	v.filesCount = len(levelFilesToPick)

	screen.PushTranslation(5, 5)
	screen.DrawText("%d, %d (Tile %d.%d, %d.%d)", screenX, screenY, int(math.Floor(worldX))-tileRect.Left, subtileX, int(math.Floor(worldY))-tileRect.Top, subtileY)
	screen.PushTranslation(0, 16)
	screen.DrawText("Map: " + curRegion.GetLevelType().Name)
	screen.PushTranslation(0, 16)
	screen.DrawText("%v: %v/%v [%v, %v]", regionPath, fileIndex+1, v.filesCount, v.currentRegion, v.levelPreset)
	screen.PushTranslation(0, 16)
	screen.DrawText("N - next region, P - previous region")
	screen.PushTranslation(0, 16)
	screen.DrawText("Shift+N - next preset, Shift+P - previous preset")
	screen.PushTranslation(0, 16)
	screen.DrawText("Ctrl+N - next file, Ctrl+P - previous file")
	screen.PopN(6)
}

func (v *MapEngineTest) Advance(tickTime float64) {
	v.mapEngine.Advance(tickTime)
}

func (met *MapEngineTest) OnKeyRepeat(event d2input.KeyEvent) bool {
	var moveSpeed float64 = 8
	if event.KeyMod == d2input.KeyModShift {
		moveSpeed *= 2
	}

	if event.Key == d2input.KeyDown {
		met.mapEngine.MoveCameraBy(0, moveSpeed)
		return true
	}

	if event.Key == d2input.KeyUp {
		met.mapEngine.MoveCameraBy(0, -moveSpeed)
		return true
	}

	if event.Key == d2input.KeyRight {
		met.mapEngine.MoveCameraBy(moveSpeed, 0)
		return true
	}

	if event.Key == d2input.KeyLeft {
		met.mapEngine.MoveCameraBy(-moveSpeed, 0)
		return true
	}

	return false
}

func (met *MapEngineTest) OnKeyDown(event d2input.KeyEvent) bool {
	if event.Key == d2input.KeyEscape {
		os.Exit(0)
		return true
	}

	if event.Key == d2input.KeyN {
		if event.KeyMod == d2input.KeyModControl {
			met.fileIndex = increment(met.fileIndex, 0, met.filesCount-1)
			met.sceneProvider.SetNextScene(met)
		} else if event.KeyMod == d2input.KeyModShift {
			met.levelPreset = increment(met.levelPreset, met.regionSpec.startPresetIndex, met.regionSpec.endPresetIndex)
			met.sceneProvider.SetNextScene(met)
		} else {
			met.currentRegion = increment(met.currentRegion, 0, len(regions))
			met.sceneProvider.SetNextScene(met)
		}

		return true
	}

	if event.Key == d2input.KeyP {
		if event.KeyMod == d2input.KeyModControl {
			met.fileIndex = decrement(met.fileIndex, 0, met.filesCount-1)
			met.sceneProvider.SetNextScene(met)
		} else if event.KeyMod == d2input.KeyModShift {
			met.levelPreset = decrement(met.levelPreset, met.regionSpec.startPresetIndex, met.regionSpec.endPresetIndex)
			met.sceneProvider.SetNextScene(met)
		} else {
			met.currentRegion = decrement(met.currentRegion, 0, len(regions))
			met.sceneProvider.SetNextScene(met)
		}

		return true
	}

	if event.Key == d2input.KeyF7 {
		if met.debugVisLevel < 2 {
			met.debugVisLevel++
		} else {
			met.debugVisLevel = 0
		}

		met.mapEngine.SetDebugVisLevel(met.debugVisLevel)
		return true
	}

	return false
}

func increment(v, min, max int) int {
	v++
	if v > max {
		return min
	}
	return v
}

func decrement(v, min, max int) int {
	v--
	if v < min {
		return max
	}
	return v
}
