package d2gamescreen

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapgen"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2maprenderer"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
)

type regionSpec struct {
	regionType       d2enum.RegionIdType
	startPresetIndex int
	endPresetIndex   int
	extra            []int
}

func getRegions() []regionSpec {
	return []regionSpec{
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
}

// MapEngineTest represents the MapEngineTest screen
type MapEngineTest struct {
	gameState     *d2player.PlayerState
	mapEngine     *d2mapengine.MapEngine
	mapRenderer   *d2maprenderer.MapRenderer
	terminal      d2interface.Terminal
	renderer      d2interface.Renderer
	inputManager  d2interface.InputManager
	audioProvider d2interface.AudioProvider

	lastMouseX, lastMouseY int
	selX, selY             int
	selectedTile           *d2mapengine.MapTile

	//TODO: this is region specific properties, should be refactored for multi-region rendering
	currentRegion int
	levelPreset   int
	fileIndex     int
	regionSpec    regionSpec
	filesCount    int
	debugVisLevel int
}

// CreateMapEngineTest creates the Map Engine Test screen and returns a pointer to it
func CreateMapEngineTest(currentRegion,
	levelPreset int,
	term d2interface.Terminal,
	renderer d2interface.Renderer,
	inputManager d2interface.InputManager,
	audioProvider d2interface.AudioProvider,
) *MapEngineTest {
	result := &MapEngineTest{
		currentRegion: currentRegion,
		levelPreset:   levelPreset,
		fileIndex:     0,
		regionSpec:    regionSpec{},
		filesCount:    0,
		terminal:      term,
		renderer:      renderer,
		inputManager:  inputManager,
		audioProvider: audioProvider,
	}
	result.gameState = d2player.CreateTestGameState()

	return result
}

func (met *MapEngineTest) loadRegionByIndex(n, levelPreset, fileIndex int) {
	log.Printf("Loaded region: Type(%d) LevelPreset(%d) FileIndex(%d)", n, levelPreset, fileIndex)
	met.mapRenderer.InvalidateImageCache()

	for _, spec := range getRegions() {
		if spec.regionType != d2enum.RegionIdType(n) {
			continue
		}

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

	if n == 0 {
		met.mapEngine.SetSeed(time.Now().UnixNano())
		d2mapgen.GenerateAct1Overworld(met.mapEngine)
	} else {
		met.mapEngine = d2mapengine.CreateMapEngine() // necessary for map name update
		met.mapEngine.SetSeed(time.Now().UnixNano())
		met.mapEngine.GenerateMap(d2enum.RegionIdType(n), levelPreset, fileIndex)
		//met.mapEngine.RegenerateWalkPaths()
	}

	met.mapRenderer.SetMapEngine(met.mapEngine)
	position := d2vector.NewPosition(met.mapRenderer.WorldToOrtho(met.mapEngine.GetCenterPosition()))
	met.mapRenderer.SetCameraPosition(&position)

	musicDef := d2common.GetMusicDef(met.regionSpec.regionType)

	met.audioProvider.PlayBGM(musicDef.MusicFile)
}

// OnLoad loads the resources for the Map Engine Test screen
func (met *MapEngineTest) OnLoad(loading d2screen.LoadingState) {
	if err := met.inputManager.BindHandler(met); err != nil {
		fmt.Printf("could not add MapEngineTest as event handler")
	}

	loading.Progress(twentyPercent)

	met.mapEngine = d2mapengine.CreateMapEngine()

	loading.Progress(fiftyPercent)

	met.mapRenderer = d2maprenderer.CreateMapRenderer(met.renderer, met.mapEngine, met.terminal, 0.0, 0.0)

	loading.Progress(seventyPercent)
	met.loadRegionByIndex(met.currentRegion, met.levelPreset, met.fileIndex)
}

// OnUnload releases the resources for the Map Engine Test screen
func (met *MapEngineTest) OnUnload() error {
	if err := met.inputManager.UnbindHandler(met); err != nil {
		return err
	}

	return nil
}

// Render renders the Map Engine Test screen
func (met *MapEngineTest) Render(screen d2interface.Surface) error {
	met.mapRenderer.Render(screen)

	screen.PushTranslation(0, 16)
	screen.DrawTextf("N - next region, P - previous region")
	screen.PushTranslation(0, 16)
	screen.DrawTextf("Shift+N - next preset, Shift+P - previous preset")
	screen.PushTranslation(0, 16)
	screen.DrawTextf("Ctrl+N - next file, Ctrl+P - previous file")
	screen.PushTranslation(0, 16)
	screen.DrawTextf("Left click selects tile, right click unselects")
	screen.PushTranslation(0, 16)

	popN := 5

	if met.selectedTile == nil {
		screen.PushTranslation(15, 16)
		popN++

		screen.DrawTextf("No tile selected")
	} else {
		screen.PushTranslation(10, 32)
		screen.DrawTextf("Tile %v,%v", met.selX, met.selY)

		screen.PushTranslation(15, 16)
		screen.DrawTextf("Walls")
		tpop := 0
		for _, wall := range met.selectedTile.Components.Walls {
			screen.PushTranslation(0, 12)
			tpop++
			tmpString := fmt.Sprintf("%#v", wall)
			stringSlice := strings.Split(tmpString, " ")
			tmp2 := strings.Split(stringSlice[0], "{")
			stringSlice[0] = tmp2[1]
			for _, str := range stringSlice {
				screen.PushTranslation(0, 12)
				tpop++
				screen.DrawTextf(str)
			}
		}
		screen.PopN(tpop)

		screen.PushTranslation(170, 0)
		screen.DrawTextf("Floors")
		tpop = 0
		for _, floor := range met.selectedTile.Components.Floors {
			screen.PushTranslation(0, 12)
			tpop++
			tmpString := fmt.Sprintf("%#v", floor)
			stringSlice := strings.Split(tmpString, " ")
			tmp2 := strings.Split(stringSlice[0], "{")
			stringSlice[0] = tmp2[1]
			for _, str := range stringSlice {
				screen.PushTranslation(0, 12)
				tpop++
				screen.DrawTextf(str)
			}
		}
		screen.PopN(tpop)

		tpop = 0
		screen.PushTranslation(170, 0)
		screen.DrawTextf("Shadows")
		for _, shadow := range met.selectedTile.Components.Shadows {
			screen.PushTranslation(0, 12)
			tpop++
			tmpString := fmt.Sprintf("%#v", shadow)
			stringSlice := strings.Split(tmpString, " ")
			tmp2 := strings.Split(stringSlice[0], "{")
			stringSlice[0] = tmp2[1]
			for _, str := range stringSlice {
				screen.PushTranslation(0, 12)
				tpop++
				screen.DrawTextf(str)
			}
		}
		screen.PopN(tpop)

		tpop = 0
		screen.PushTranslation(170, 0)
		screen.DrawTextf("Substitutions")
		for _, subst := range met.selectedTile.Components.Substitutions {
			screen.PushTranslation(0, 12)
			tpop++
			tmpString := fmt.Sprintf("%#v", subst)
			stringSlice := strings.Split(tmpString, " ")
			tmp2 := strings.Split(stringSlice[0], "{")
			stringSlice[0] = tmp2[1]
			for _, str := range stringSlice {
				screen.PushTranslation(0, 12)
				tpop++
				screen.DrawTextf(str)
			}
		}
		screen.PopN(tpop)

		popN += 5
	}

	screen.PopN(popN)

	return nil
}

func (met *MapEngineTest) OnMouseMove(event d2interface.MouseMoveEvent) bool {
	mx, my := event.X(), event.Y()
	met.lastMouseX = mx
	met.lastMouseY = my

	return false
}

func (met *MapEngineTest) OnMouseButtonDown(event d2interface.MouseEvent) bool {
	if event.Button() == d2enum.MouseButtonLeft {
		px, py := met.mapRenderer.ScreenToWorld(met.lastMouseX, met.lastMouseY)
		met.selX = int(px)
		met.selY = int(py)
		met.selectedTile = met.mapEngine.TileAt(int(px), int(py))

		camVect := met.mapRenderer.Camera.GetPosition().Vector

		x, y := float64(met.lastMouseX-400)/5, float64(met.lastMouseY-300)/5
		targetPosition := d2vector.NewPositionTile(x, y)
		targetPosition.Add(&camVect)

		met.mapRenderer.SetCameraTarget(&targetPosition)

		return true
	}

	if event.Button() == d2enum.MouseButtonRight {
		met.selectedTile = nil

		return true
	}

	return false
}

func (met *MapEngineTest) OnMouseButtonRepeat(event d2interface.MouseEvent) bool {
	if event.Button() == d2enum.MouseButtonLeft {
		px, py := met.mapRenderer.ScreenToWorld(met.lastMouseX, met.lastMouseY)
		met.selX = int(px)
		met.selY = int(py)
		met.selectedTile = met.mapEngine.TileAt(int(px), int(py))

		camVect := met.mapRenderer.Camera.GetPosition().Vector

		x, y := float64(met.lastMouseX-400)/5, float64(met.lastMouseY-300)/5
		targetPosition := d2vector.NewPositionTile(x, y)
		targetPosition.Add(&camVect)

		met.mapRenderer.SetCameraTarget(&targetPosition)

		return true
	}

	return false
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
	if event.KeyMod() == d2enum.KeyModShift {
		moveSpeed *= 2
	}

	if event.Key() == d2enum.KeyDown {
		v := d2vector.NewVector(0, moveSpeed)
		met.mapRenderer.MoveCameraTargetBy(v)

		return true
	}

	if event.Key() == d2enum.KeyUp {
		v := d2vector.NewVector(0, -moveSpeed)
		met.mapRenderer.MoveCameraTargetBy(v)

		return true
	}

	if event.Key() == d2enum.KeyRight {
		v := d2vector.NewVector(moveSpeed, 0)
		met.mapRenderer.MoveCameraTargetBy(v)

		return true
	}

	if event.Key() == d2enum.KeyLeft {
		v := d2vector.NewVector(-moveSpeed, 0)
		met.mapRenderer.MoveCameraTargetBy(v)

		return true
	}

	return false
}

// OnKeyDown defines the actions of the Map Engine Test screen when a key is pressed
func (met *MapEngineTest) OnKeyDown(event d2interface.KeyEvent) bool {
	if event.Key() == d2enum.KeyEscape {
		os.Exit(0)
		return true
	}

	if event.Key() == d2enum.KeyN {
		switch event.KeyMod() {
		case d2enum.KeyModControl:
			met.fileIndex++
			d2screen.SetNextScreen(met)
		case d2enum.KeyModShift:
			met.levelPreset = increment(met.levelPreset, met.regionSpec.startPresetIndex, met.regionSpec.endPresetIndex)
			d2screen.SetNextScreen(met)
		default:
			met.currentRegion = increment(met.currentRegion, 0, len(getRegions()))
			d2screen.SetNextScreen(met)
		}

		return true
	}

	if event.Key() == d2enum.KeyP {
		switch event.KeyMod() {
		case d2enum.KeyModControl:
			met.fileIndex--
			d2screen.SetNextScreen(met)
		case d2enum.KeyModShift:
			met.levelPreset = decrement(met.levelPreset, met.regionSpec.startPresetIndex, met.regionSpec.endPresetIndex)
			d2screen.SetNextScreen(met)
		default:
			met.currentRegion = decrement(met.currentRegion, 0, len(getRegions()))
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
