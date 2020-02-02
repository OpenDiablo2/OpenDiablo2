package d2gamescene

import (
	"math"
	"os"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gamestate"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2scene"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map"
)

type RegionSpec struct {
	regionType       d2enum.RegionIdType
	startPresetIndex int
	endPresetIndex   int
	extra            []int
}

var regions = []RegionSpec{
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
	gameState *d2gamestate.GameState
	mapEngine *d2map.MapEngine

	//TODO: this is region specific properties, should be refactored for multi-region rendering
	currentRegion int
	levelPreset   int
	fileIndex     int
	regionSpec    RegionSpec
	filesCount    int
	debugVisLevel int
}

func CreateMapEngineTest(currentRegion int, levelPreset int) *MapEngineTest {
	result := &MapEngineTest{
		currentRegion: currentRegion,
		levelPreset:   levelPreset,
		fileIndex:     0,
		regionSpec:    RegionSpec{},
		filesCount:    0,
	}
	result.gameState = d2gamestate.CreateTestGameState()
	return result
}

func (met *MapEngineTest) LoadRegionByIndex(n int, levelPreset, fileIndex int) {
	for _, spec := range regions {
		if spec.regionType == d2enum.RegionIdType(n) {
			met.regionSpec = spec
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
			met.levelPreset = levelPreset
		}
	}

	if n == 0 {
		met.mapEngine.GenerateAct1Overworld()
	} else {
		met.mapEngine = d2map.CreateMapEngine(met.gameState) // necessary for map name update
		met.mapEngine.GenerateMap(d2enum.RegionIdType(n), levelPreset, fileIndex)
	}

	met.mapEngine.MoveCameraTo(met.mapEngine.WorldToOrtho(met.mapEngine.GetCenterPosition()))
}

func (met *MapEngineTest) Load() []func() {
	// TODO: Game seed comes from the game state object
	d2input.BindHandler(met)
	d2audio.PlayBGM("")
	return []func(){
		func() {
			met.mapEngine = d2map.CreateMapEngine(met.gameState)
			met.LoadRegionByIndex(met.currentRegion, met.levelPreset, met.fileIndex)
		},
	}
}

func (met *MapEngineTest) Unload() {
	d2input.UnbindHandler(met)
}

func (met *MapEngineTest) Render(screen d2render.Surface) {
	met.mapEngine.Render(screen)

	screenX, screenY, _ := d2render.GetCursorPos()
	worldX, worldY := met.mapEngine.ScreenToWorld(screenX, screenY)
	//subtileX := int(math.Ceil(math.Mod(worldX*10, 10))) / 2
	//subtileY := int(math.Ceil(math.Mod(worldY*10, 10))) / 2

	curRegion := met.mapEngine.GetRegionAtTile(int(worldX), int(worldY))
	if curRegion == nil {
		return
	}

	tileRect := curRegion.GetTileRect()

	levelFilesToPick := make([]string, 0)
	fileIndex := met.fileIndex
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
	if met.fileIndex == -1 {
		met.fileIndex = fileIndex
	}
	met.filesCount = len(levelFilesToPick)

	tileX := int(math.Floor(worldX)) - tileRect.Left
	tileY := int(math.Floor(worldY)) - tileRect.Top

	subtileX := int((worldX - float64(int(worldX))) * 5)
	subtileY := int((worldY - float64(int(worldY))) * 5)

	regionWidth, regionHeight := curRegion.GetTileSize()
	if tileX >= 0 && tileY >= 0 && tileX < regionWidth && tileY < regionHeight {
		tile := curRegion.GetTile(tileX, tileY)
		screen.PushTranslation(5, 5)
		screen.DrawText("%d, %d (Tile %d.%d, %d.%d)", screenX, screenY, tileX, subtileX, tileY, subtileY)
		screen.PushTranslation(0, 16)
		screen.DrawText("Map: " + curRegion.GetLevelType().Name)
		screen.PushTranslation(0, 16)
		screen.DrawText("%v: %v/%v [%v, %v]", regionPath, fileIndex+1, met.filesCount, met.currentRegion, met.levelPreset)
		screen.PushTranslation(0, 16)
		screen.DrawText("N - next region, P - previous region")
		screen.PushTranslation(0, 16)
		screen.DrawText("Shift+N - next preset, Shift+P - previous preset")
		screen.PushTranslation(0, 16)
		screen.DrawText("Ctrl+N - next file, Ctrl+P - previous file")
		screen.PushTranslation(0, 16)
		popN := 7
		if len(tile.Floors) > 0 {
			screen.PushTranslation(0, 16)
			screen.DrawText("Floors:")
			screen.PushTranslation(16, 0)
			for idx, floor := range tile.Floors {
				popN++
				screen.PushTranslation(0, 16)
				tileData := curRegion.GetTileData(int32(floor.Style), int32(floor.Sequence), d2enum.Floor)
				tileSubAttrs := d2dt1.SubTileFlags{}
				if tileData != nil {
					tileSubAttrs = *tileData.GetSubTileFlags(subtileX, subtileY)
				}
				screen.DrawText("Floor %v: [ANI:%t] %s", idx, floor.Animated, tileSubAttrs.DebugString())

			}
			screen.PushTranslation(-16, 0)
			popN += 3
		}
		if len(tile.Walls) > 0 {
			screen.PushTranslation(0, 16)
			screen.DrawText("Walls:")
			screen.PushTranslation(16, 0)
			for idx, wall := range tile.Walls {
				popN++
				screen.PushTranslation(0, 16)
				tileData := curRegion.GetTileData(int32(wall.Style), int32(wall.Sequence), d2enum.Floor)
				tileSubAttrs := d2dt1.SubTileFlags{}
				if tileData != nil {
					tileSubAttrs = *tileData.GetSubTileFlags(subtileX, subtileY)
				}
				screen.DrawText("Wall %v: [HID:%t] %s", idx, wall.Hidden, tileSubAttrs.DebugString())

			}
			screen.PushTranslation(-16, 0)
			popN += 3
		}
		if len(tile.Walls) > 0 {
			screen.PushTranslation(0, 16)
			screen.DrawText("Shadows:")
			screen.PushTranslation(16, 0)
			for idx, shadow := range tile.Shadows {
				popN++
				screen.PushTranslation(0, 16)
				tileData := curRegion.GetTileData(int32(shadow.Style), int32(shadow.Sequence), d2enum.Floor)
				tileSubAttrs := d2dt1.SubTileFlags{}
				if tileData != nil {
					tileSubAttrs = *tileData.GetSubTileFlags(subtileX, subtileY)
				}
				screen.DrawText("Wall %v: [HID:%t] %s", idx, shadow.Hidden, tileSubAttrs.DebugString())

			}
			screen.PushTranslation(-16, 0)
			popN += 3
		}
		screen.PopN(popN)
	}
}

func (met *MapEngineTest) Advance(tickTime float64) {
	met.mapEngine.Advance(tickTime)
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
			d2scene.SetNextScene(met)
		} else if event.KeyMod == d2input.KeyModShift {
			met.levelPreset = increment(met.levelPreset, met.regionSpec.startPresetIndex, met.regionSpec.endPresetIndex)
			d2scene.SetNextScene(met)
		} else {
			met.currentRegion = increment(met.currentRegion, 0, len(regions))
			d2scene.SetNextScene(met)
		}

		return true
	}

	if event.Key == d2input.KeyP {
		if event.KeyMod == d2input.KeyModControl {
			met.fileIndex = decrement(met.fileIndex, 0, met.filesCount-1)
			d2scene.SetNextScene(met)
		} else if event.KeyMod == d2input.KeyModShift {
			met.levelPreset = decrement(met.levelPreset, met.regionSpec.startPresetIndex, met.regionSpec.endPresetIndex)
			d2scene.SetNextScene(met)
		} else {
			met.currentRegion = decrement(met.currentRegion, 0, len(regions))
			d2scene.SetNextScene(met)
		}

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
