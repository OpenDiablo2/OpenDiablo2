package d2gamescreen

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapgen"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2maprenderer"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
)

type regionSpec struct {
	regionType       d2enum.RegionIdType
	startPresetIndex int
	endPresetIndex   int
	extra            []int
}

var regions = []regionSpec{
	// Act I
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

	// Act II
	{d2enum.RegionAct2Town, 301, 301, []int{}},
	{d2enum.RegionAct2Sewer, 302, 352, []int{}},
	{d2enum.RegionAct2Harem, 353, 357, []int{}},
	{d2enum.RegionAct2Basement, 358, 361, []int{}},
	{d2enum.RegionAct2Desert, 362, 413, []int{}},
	{d2enum.RegionAct2Tomb, 414, 481, []int{}},
	{d2enum.RegionAct2Lair, 482, 509, []int{}},
	{d2enum.RegionAct2Arcane, 510, 528, []int{}},

	// Act III
	{d2enum.RegionAct3Town, 529, 529, []int{}},
	{d2enum.RegionAct3Jungle, 530, 604, []int{}},
	{d2enum.RegionAct3Kurast, 605, 658, []int{
		748, 749, 750, 751, 752, 753, 754, 755, 756, 757, 758, 759, 760, 761, 762, 763, 764, 765, 766, 767, 768, 769,
		770, 771, 772, 773, 774, 775, 776, 777, 778, 779, 780, 781, 782, 783, 784, 785, 786, 787, 788, 789, 790, 791,
		792, 793, 794, 795, 796,
	}},
	{d2enum.RegionAct3Spider, 659, 664, []int{}},
	{d2enum.RegionAct3Dungeon, 665, 704, []int{}},
	{d2enum.RegionAct3Sewer, 705, 747, []int{}},

	// Act IV
	{d2enum.RegionAct4Town, 797, 798, []int{}},
	{d2enum.RegionAct4Mesa, 799, 835, []int{}},
	{d2enum.RegionAct4Lava, 836, 862, []int{}},

	// Act V -- broken or wrong order
	{d2enum.RegonAct5Town, 863, 864, []int{}},
	{d2enum.RegionAct5Siege, 865, 879, []int{}},
	{d2enum.RegionAct5Barricade, 880, 1002, []int{}},
	{d2enum.RegionAct5IceCaves, 1003, 1041, []int{}},
	{d2enum.RegionAct5Temple, 1042, 1052, []int{}},
	{d2enum.RegionAct5Baal, 1059, 1090, []int{}},
	{d2enum.RegionAct5Lava, 1053, 1058, []int{}},
}

type MapEngineTest struct {
	gameState   *d2player.PlayerState
	mapEngine   *d2mapengine.MapEngine
	mapRenderer *d2maprenderer.MapRenderer
	terminal    d2interface.Terminal
	renderer    d2interface.Renderer

	//TODO: this is region specific properties, should be refactored for multi-region rendering
	currentRegion int
	levelPreset   int
	fileIndex     int
	regionSpec    regionSpec
	filesCount    int
	debugVisLevel int
}

// CreateMapEngineTest creates the Map Engine Test screen and returns a pointer to it
func CreateMapEngineTest(currentRegion, levelPreset int, term d2interface.Terminal, renderer d2interface.Renderer) *MapEngineTest {
	result := &MapEngineTest{
		currentRegion: currentRegion,
		levelPreset:   levelPreset,
		fileIndex:     0,
		regionSpec:    regionSpec{},
		filesCount:    0,
		terminal:      term,
		renderer:      renderer,
	}
	result.gameState = d2player.CreateTestGameState()

	return result
}

func (met *MapEngineTest) loadRegionByIndex(n int, levelPreset, fileIndex int) {
	log.Printf("Loaded region: Type(%d) LevelPreset(%d) FileIndex(%d)", n, levelPreset, fileIndex)
	d2maprenderer.InvalidateImageCache()

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
		met.mapEngine.SetSeed(time.Now().UnixNano())
		d2mapgen.GenerateAct1Overworld(met.mapEngine)
	} else {
		met.mapEngine = d2mapengine.CreateMapEngine() // necessary for map name update
		met.mapEngine.SetSeed(time.Now().UnixNano())
		met.mapEngine.GenerateMap(d2enum.RegionIdType(n), levelPreset, fileIndex, true)
		met.mapEngine.RegenerateWalkPaths()
	}

	met.mapRenderer.SetMapEngine(met.mapEngine)
	met.mapRenderer.MoveCameraTo(met.mapRenderer.WorldToOrtho(met.mapEngine.GetCenterPosition()))
}

// OnLoad loads the resources for the Map Engine Test screen
func (met *MapEngineTest) OnLoad(loading d2screen.LoadingState) {
	if err := d2input.BindHandler(met); err != nil {
		fmt.Printf("could not add MapEngineTest as event handler")
	}

	loading.Progress(0.2)

	met.mapEngine = d2mapengine.CreateMapEngine()

	loading.Progress(0.5)

	met.mapRenderer = d2maprenderer.CreateMapRenderer(met.renderer, met.mapEngine, met.terminal)

	loading.Progress(0.7)
	met.loadRegionByIndex(met.currentRegion, met.levelPreset, met.fileIndex)
}

// OnUnload releases the resources for the Map Engine Test screen
func (met *MapEngineTest) OnUnload() error {
	if err := d2input.UnbindHandler(met); err != nil {
		return err
	}

	return nil
}

// Render renders the Map Engine Test screen
func (met *MapEngineTest) Render(screen d2interface.Surface) error {
	met.mapRenderer.Render(screen)

	//
	//levelFilesToPick := make([]string, 0)
	//fileIndex := met.fileIndex
	//levelPreset := curRegion.LevelPreset()
	//regionPath := curRegion.RegionPath()
	//for n, fileRecord := range levelPreset.Files {
	//	if len(fileRecord) == 0 || fileRecord == "" || fileRecord == "0" {
	//		continue
	//	}
	//	levelFilesToPick = append(levelFilesToPick, fileRecord)
	//	if fileRecord == regionPath {
	//		fileIndex = n
	//	}
	//}
	//if met.fileIndex == -1 {
	//	met.fileIndex = fileIndex
	//}
	//met.filesCount = len(levelFilesToPick)
	//
	//
	//regionWidth, regionHeight := curRegion.GetTileSize()
	//if tileX >= 0 && tileY >= 0 && tileX < regionWidth && tileY < regionHeight {
	//	tile := curRegion.Tile(tileX, tileY)
	//	screen.PushTranslation(5, 5)
	//	screen.DrawText("%d, %d (Tile %d.%d, %d.%d)", screenX, screenY, tileX, subtileX, tileY, subtileY)
	//	screen.PushTranslation(0, 16)
	//	screen.DrawText("Map: " + curRegion.LevelType().Name)
	//	screen.PushTranslation(0, 16)
	//	screen.DrawText("%v: %v/%v [%v, %v]", regionPath, fileIndex+1, met.filesCount, met.currentRegion, met.levelPreset)
	//	screen.PushTranslation(0, 16)
	//	screen.DrawText("N - next region, P - previous region")
	//	screen.PushTranslation(0, 16)
	//	screen.DrawText("Shift+N - next preset, Shift+P - previous preset")
	//	screen.PushTranslation(0, 16)
	//	screen.DrawText("Ctrl+N - next file, Ctrl+P - previous file")
	//	screen.PushTranslation(0, 16)
	//	popN := 7
	//	if len(tile.Floors) > 0 {
	//		screen.PushTranslation(0, 16)
	//		screen.DrawText("Floors:")
	//		screen.PushTranslation(16, 0)
	//		for idx, floor := range tile.Floors {
	//			popN++
	//			screen.PushTranslation(0, 16)
	//			tileData := curRegion.TileData(int32(floor.Style), int32(floor.Sequence), d2enum.Floor)
	//			tileSubAttrs := d2dt1.SubTileFlags{}
	//			if tileData != nil {
	//				tileSubAttrs = *tileData.GetSubTileFlags(subtileX, subtileY)
	//			}
	//			screen.DrawText("Floor %v: [ANI:%t] %s", idx, floor.Animated, tileSubAttrs.DebugString())
	//
	//		}
	//		screen.PushTranslation(-16, 0)
	//		popN += 3
	//	}
	//	if len(tile.Walls) > 0 {
	//		screen.PushTranslation(0, 16)
	//		screen.DrawText("Walls:")
	//		screen.PushTranslation(16, 0)
	//		for idx, wall := range tile.Walls {
	//			popN++
	//			screen.PushTranslation(0, 16)
	//			tileData := curRegion.TileData(int32(wall.Style), int32(wall.Sequence), d2enum.Floor)
	//			tileSubAttrs := d2dt1.SubTileFlags{}
	//			if tileData != nil {
	//				tileSubAttrs = *tileData.GetSubTileFlags(subtileX, subtileY)
	//			}
	//			screen.DrawText("Wall %v: [HID:%t] %s", idx, wall.Hidden, tileSubAttrs.DebugString())
	//
	//		}
	//		screen.PushTranslation(-16, 0)
	//		popN += 3
	//	}
	//	if len(tile.Walls) > 0 {
	//		screen.PushTranslation(0, 16)
	//		screen.DrawText("Shadows:")
	//		screen.PushTranslation(16, 0)
	//		for idx, shadow := range tile.Shadows {
	//			popN++
	//			screen.PushTranslation(0, 16)
	//			tileData := curRegion.TileData(int32(shadow.Style), int32(shadow.Sequence), d2enum.Floor)
	//			tileSubAttrs := d2dt1.SubTileFlags{}
	//			if tileData != nil {
	//				tileSubAttrs = *tileData.GetSubTileFlags(subtileX, subtileY)
	//			}
	//			screen.DrawText("Wall %v: [HID:%t] %s", idx, shadow.Hidden, tileSubAttrs.DebugString())
	//
	//		}
	//		screen.PushTranslation(-16, 0)
	//		popN += 3
	//	}
	//	screen.PopN(popN)
	//}

	return nil
}

// Advance runs the update logic on the Map Engine Test screen
func (met *MapEngineTest) Advance(tickTime float64) error {
	met.mapEngine.Advance(tickTime)
	met.mapRenderer.Advance(tickTime)

	return nil
}

// OnKeyRepeat is called to handle repeated key presses
func (met *MapEngineTest) OnKeyRepeat(event d2interface.KeyEvent) bool {
	var moveSpeed float64 = 8
	if event.KeyMod() == d2interface.KeyModShift {
		moveSpeed *= 2
	}

	if event.Key() == d2interface.KeyDown {
		met.mapRenderer.MoveCameraBy(0, moveSpeed)
		return true
	}

	if event.Key() == d2interface.KeyUp {
		met.mapRenderer.MoveCameraBy(0, -moveSpeed)
		return true
	}

	if event.Key() == d2interface.KeyRight {
		met.mapRenderer.MoveCameraBy(moveSpeed, 0)
		return true
	}

	if event.Key() == d2interface.KeyLeft {
		met.mapRenderer.MoveCameraBy(-moveSpeed, 0)
		return true
	}

	return false
}

// OnKeyDown defines the actions of the Map Engine Test screen when a key is pressed
func (met *MapEngineTest) OnKeyDown(event d2interface.KeyEvent) bool {
	if event.Key() == d2interface.KeyEscape {
		os.Exit(0)
		return true
	}

	if event.Key() == d2interface.KeyN {
		switch event.KeyMod() {
		case d2interface.KeyModControl:
			met.fileIndex++
			d2screen.SetNextScreen(met)
		case d2interface.KeyModShift:
			met.levelPreset = increment(met.levelPreset, met.regionSpec.startPresetIndex, met.regionSpec.endPresetIndex)
			d2screen.SetNextScreen(met)
		default:
			met.currentRegion = increment(met.currentRegion, 0, len(regions))
			d2screen.SetNextScreen(met)
		}

		return true
	}

	if event.Key() == d2interface.KeyP {
		switch event.KeyMod() {
		case d2interface.KeyModControl:
			met.fileIndex--
			d2screen.SetNextScreen(met)
		case d2interface.KeyModShift:
			met.levelPreset = decrement(met.levelPreset, met.regionSpec.startPresetIndex, met.regionSpec.endPresetIndex)
			d2screen.SetNextScreen(met)
		default:
			met.currentRegion = decrement(met.currentRegion, 0, len(regions))
			d2screen.SetNextScreen(met)
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
