package d2gamescreen

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapgen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2maprenderer"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
)

const (
	subtilesPerTile = 5
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

// CreateMapEngineTest creates the Map Engine Test screen and returns a pointer to it
func CreateMapEngineTest(currentRegion,
	levelPreset int,
	asset *d2asset.AssetManager,
	term d2interface.Terminal,
	renderer d2interface.Renderer,
	inputManager d2interface.InputManager,
	audioProvider d2interface.AudioProvider,
	l d2util.LogLevel,
	screen *d2screen.ScreenManager,
) (*MapEngineTest, error) {
	heroStateFactory, err := d2hero.NewHeroStateFactory(asset)
	if err != nil {
		return nil, err
	}

	mapEngineTest := &MapEngineTest{
		currentRegion:      currentRegion,
		levelPreset:        levelPreset,
		fileIndex:          0,
		regionSpec:         regionSpec{},
		filesCount:         0,
		asset:              asset,
		terminal:           term,
		renderer:           renderer,
		inputManager:       inputManager,
		audioProvider:      audioProvider,
		screen:             screen,
		playerStateFactory: heroStateFactory,
		logLevel:           l,
	}

	mapEngineTest.playerState = heroStateFactory.CreateTestGameState()

	mapEngineTest.Logger = d2util.NewLogger()
	mapEngineTest.Logger.SetLevel(l)
	mapEngineTest.Logger.SetPrefix(logPrefix)

	return mapEngineTest, nil
}

// MapEngineTest represents the MapEngineTest screen
type MapEngineTest struct {
	asset              *d2asset.AssetManager
	playerStateFactory *d2hero.HeroStateFactory
	playerState        *d2hero.HeroState
	mapEngine          *d2mapengine.MapEngine
	mapGen             *d2mapgen.MapGenerator
	mapRenderer        *d2maprenderer.MapRenderer
	terminal           d2interface.Terminal
	renderer           d2interface.Renderer
	inputManager       d2interface.InputManager
	audioProvider      d2interface.AudioProvider
	screen             *d2screen.ScreenManager

	lastMouseX, lastMouseY int
	selX, selY             int
	selectedTile           *d2mapengine.MapTile

	// https://github.com/OpenDiablo2/OpenDiablo2/issues/806
	currentRegion int
	levelPreset   int
	fileIndex     int
	regionSpec    regionSpec
	filesCount    int

	*d2util.Logger
	logLevel d2util.LogLevel
}

func (met *MapEngineTest) loadRegionByIndex(n, levelPreset, fileIndex int) {
	met.Infof("Loaded region: Type(%d) LevelPreset(%d) FileIndex(%d)", n, levelPreset, fileIndex)
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

	mapGen, _ := d2mapgen.NewMapGenerator(met.asset, met.logLevel, met.mapEngine)
	met.mapGen = mapGen

	if n == 0 {
		met.mapEngine.SetSeed(time.Now().UnixNano())
		met.mapGen.GenerateAct1Overworld()
	} else {
		met.mapEngine = d2mapengine.CreateMapEngine(met.logLevel, met.asset) // necessary for map name update
		met.mapEngine.SetSeed(time.Now().UnixNano())
		met.mapEngine.GenerateMap(d2enum.RegionIdType(n), levelPreset, fileIndex)
	}

	met.mapRenderer.SetMapEngine(met.mapEngine)
	position := d2vector.NewPosition(met.mapRenderer.WorldToOrtho(met.mapEngine.GetCenterPosition()))
	met.mapRenderer.SetCameraPosition(&position)

	musicDef := d2resource.GetMusicDef(met.regionSpec.regionType)

	met.audioProvider.PlayBGM(musicDef.MusicFile)
}

// OnLoad loads the resources for the Map Engine Test screen
func (met *MapEngineTest) OnLoad(loading d2screen.LoadingState) {
	if err := met.inputManager.BindHandler(met); err != nil {
		met.Error("could not add MapEngineTest as event handler")
	}

	loading.Progress(twentyPercent)

	met.mapEngine = d2mapengine.CreateMapEngine(met.logLevel, met.asset)

	loading.Progress(fiftyPercent)

	met.mapRenderer = d2maprenderer.CreateMapRenderer(met.asset, met.renderer, met.mapEngine,
		met.terminal, met.logLevel, 0.0, 0.0)

	loading.Progress(seventyPercent)
	met.loadRegionByIndex(met.currentRegion, met.levelPreset, met.fileIndex)
}

// OnUnload releases the resources for the Map Engine Test screen
func (met *MapEngineTest) OnUnload() error {
	//  https://github.com/OpenDiablo2/OpenDiablo2/issues/792
	if err := met.inputManager.UnbindHandler(met); err != nil {
		return err
	}

	return nil
}

const (
	lineSmallOffsetY  = 12
	lineNormalOffsetY = 16
	lineSmallIndentX  = 10
	lineNormalIndentX = 15
	lineBigIndentX    = 170 // distance between text columns
)

// Render renders the Map Engine Test screen
func (met *MapEngineTest) Render(screen d2interface.Surface) {
	met.mapRenderer.Render(screen)

	screen.PushTranslation(0, lineNormalOffsetY)
	defer screen.Pop()

	screen.DrawTextf("N - next region, P - previous region")

	screen.PushTranslation(0, lineNormalOffsetY)
	defer screen.Pop()

	screen.DrawTextf("Shift+N - next preset, Shift+P - previous preset")

	screen.PushTranslation(0, lineNormalOffsetY)
	defer screen.Pop()

	screen.DrawTextf("Ctrl+N - next file, Ctrl+P - previous file")

	screen.PushTranslation(0, lineNormalOffsetY)
	defer screen.Pop()

	screen.DrawTextf("Left click selects tile, right click unselects")

	screen.PushTranslation(0, lineNormalOffsetY)
	defer screen.Pop()

	met.renderTileInfo(screen)
}

func (met *MapEngineTest) renderTileInfo(screen d2interface.Surface) {
	if met.selectedTile == nil {
		screen.PushTranslation(lineNormalIndentX, lineNormalOffsetY)
		defer screen.Pop()

		screen.DrawTextf("No tile selected")
	} else {
		screen.PushTranslation(lineSmallIndentX, lineNormalOffsetY)
		defer screen.Pop()
		screen.PushTranslation(0, lineNormalOffsetY) // extra vspace
		defer screen.Pop()

		screen.DrawTextf("Tile %v,%v", met.selX, met.selY)

		screen.PushTranslation(lineNormalIndentX, lineNormalOffsetY)
		defer screen.Pop()

		screen.DrawTextf("Walls")

		tpop := 0
		for _, wall := range met.selectedTile.Components.Walls {
			screen.PushTranslation(0, lineSmallOffsetY)
			tpop++
			tmpString := fmt.Sprintf("%#v", wall)
			stringSlice := strings.Split(tmpString, " ")
			tmp2 := strings.Split(stringSlice[0], "{")
			stringSlice[0] = tmp2[1]
			for _, str := range stringSlice {
				screen.PushTranslation(0, lineSmallOffsetY)
				tpop++
				screen.DrawTextf(str)
			}
		}

		screen.PopN(tpop)

		screen.PushTranslation(lineBigIndentX, 0)
		defer screen.Pop()

		screen.DrawTextf("Floors")

		tpop = 0
		for _, floor := range met.selectedTile.Components.Floors {
			screen.PushTranslation(0, lineSmallOffsetY)
			tpop++
			tmpString := fmt.Sprintf("%#v", floor)
			stringSlice := strings.Split(tmpString, " ")
			tmp2 := strings.Split(stringSlice[0], "{")
			stringSlice[0] = tmp2[1]
			for _, str := range stringSlice {
				screen.PushTranslation(0, lineSmallOffsetY)
				tpop++
				screen.DrawTextf(str)
			}
		}
		screen.PopN(tpop)

		screen.PushTranslation(lineBigIndentX, 0)
		defer screen.Pop()

		screen.DrawTextf("Shadows")

		tpop = 0
		for _, shadow := range met.selectedTile.Components.Shadows {
			screen.PushTranslation(0, lineSmallOffsetY)
			tpop++
			tmpString := fmt.Sprintf("%#v", shadow)
			stringSlice := strings.Split(tmpString, " ")
			tmp2 := strings.Split(stringSlice[0], "{")
			stringSlice[0] = tmp2[1]
			for _, str := range stringSlice {
				screen.PushTranslation(0, lineSmallOffsetY)
				tpop++
				screen.DrawTextf(str)
			}
		}
		screen.PopN(tpop)

		screen.PushTranslation(lineBigIndentX, 0)
		defer screen.Pop()

		screen.DrawTextf("Substitutions")

		tpop = 0
		for _, subst := range met.selectedTile.Components.Substitutions {
			screen.PushTranslation(0, lineSmallOffsetY)
			tpop++
			tmpString := fmt.Sprintf("%#v", subst)
			stringSlice := strings.Split(tmpString, " ")
			tmp2 := strings.Split(stringSlice[0], "{")
			stringSlice[0] = tmp2[1]
			for _, str := range stringSlice {
				screen.PushTranslation(0, lineSmallOffsetY)
				tpop++
				screen.DrawTextf(str)
			}
		}
		screen.PopN(tpop)
	}
}

// OnMouseMove is the mouse move handler
func (met *MapEngineTest) OnMouseMove(event d2interface.MouseMoveEvent) bool {
	mx, my := event.X(), event.Y()
	met.lastMouseX = mx
	met.lastMouseY = my

	return false
}

// OnMouseButtonDown handles mouse button down events
func (met *MapEngineTest) OnMouseButtonDown(event d2interface.MouseEvent) bool {
	if event.Button() == d2enum.MouseButtonLeft {
		met.handleLeftClick()

		return true
	}

	if event.Button() == d2enum.MouseButtonRight {
		met.selectedTile = nil

		return true
	}

	return false
}

// OnMouseButtonRepeat handles repeated mouse clicks
func (met *MapEngineTest) OnMouseButtonRepeat(event d2interface.MouseEvent) bool {
	if event.Button() == d2enum.MouseButtonLeft {
		met.handleLeftClick()

		return true
	}

	return false
}

func (met *MapEngineTest) handleLeftClick() {
	px, py := met.mapRenderer.ScreenToWorld(met.lastMouseX, met.lastMouseY)
	met.selX = int(px)
	met.selY = int(py)
	met.selectedTile = met.mapEngine.TileAt(int(px), int(py))

	camVect := met.mapRenderer.Camera.GetPosition().Vector

	halfScreenWidth, halfScreenHeight := screenWidth>>1, screenHeight>>1

	x := float64(met.lastMouseX-halfScreenWidth) / subtilesPerTile
	y := float64(met.lastMouseY-halfScreenHeight) / subtilesPerTile

	targetPosition := d2vector.NewPositionTile(x, y)
	targetPosition.Add(&camVect)

	met.mapRenderer.SetCameraTarget(&targetPosition)
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
			met.screen.SetNextScreen(met)
		case d2enum.KeyModShift:
			met.levelPreset = increment(met.levelPreset, met.regionSpec.startPresetIndex, met.regionSpec.endPresetIndex)
			met.screen.SetNextScreen(met)
		default:
			met.currentRegion = increment(met.currentRegion, 0, len(getRegions()))
			met.screen.SetNextScreen(met)
		}

		return true
	}

	if event.Key() == d2enum.KeyP {
		switch event.KeyMod() {
		case d2enum.KeyModControl:
			met.fileIndex--
			met.screen.SetNextScreen(met)
		case d2enum.KeyModShift:
			met.levelPreset = decrement(met.levelPreset, met.regionSpec.startPresetIndex, met.regionSpec.endPresetIndex)
			met.screen.SetNextScreen(met)
		default:
			met.currentRegion = decrement(met.currentRegion, 0, len(getRegions()))
			met.screen.SetNextScreen(met)
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
