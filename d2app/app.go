// Package d2app contains the OpenDiablo2 application shell
package d2app

import (
	"bytes"
	"container/ring"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync"

	"github.com/pkg/profile"
	"golang.org/x/image/colornames"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2tbl"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2gamescreen"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2script"
)

// these are used for debug print info
const (
	fpsX, fpsY         = 5, 565
	memInfoX, memInfoY = 670, 5
	debugLineHeight    = 16
	errMsgPadding      = 20
)

// App represents the main application for the engine
type App struct {
	lastTime          float64
	lastScreenAdvance float64
	showFPS           bool
	timeScale         float64
	captureState      captureState
	capturePath       string
	captureFrames     []*image.RGBA
	gitBranch         string
	gitCommit         string
	asset             *d2asset.AssetManager
	inputManager      d2interface.InputManager
	terminal          d2interface.Terminal
	scriptEngine      *d2script.ScriptEngine
	audio             d2interface.AudioProvider
	renderer          d2interface.Renderer
	screen            *d2screen.ScreenManager
	ui                *d2ui.UIManager
	tAllocSamples     *ring.Ring
	guiManager        *d2gui.GuiManager
}

type bindTerminalEntry struct {
	name        string
	description string
	action      interface{}
}

const (
	bytesToMegabyte = 1024 * 1024
	nSamplesTAlloc  = 100
	debugPopN       = 6
)

// Create creates a new instance of the application
func Create(gitBranch, gitCommit string,
	inputManager d2interface.InputManager,
	terminal d2interface.Terminal,
	scriptEngine *d2script.ScriptEngine,
	audio d2interface.AudioProvider,
	renderer d2interface.Renderer,
	asset *d2asset.AssetManager,
) *App {
	uiManager := d2ui.NewUIManager(asset, renderer, inputManager, audio)

	result := &App{
		gitBranch:     gitBranch,
		gitCommit:     gitCommit,
		inputManager:  inputManager,
		terminal:      terminal,
		scriptEngine:  scriptEngine,
		audio:         audio,
		renderer:      renderer,
		ui:            uiManager,
		asset:         asset,
		tAllocSamples: createZeroedRing(nSamplesTAlloc),
	}

	if result.gitBranch == "" {
		result.gitBranch = "Local Build"
	}

	return result
}

// Run executes the application and kicks off the entire game process
func (a *App) Run() error {
	profileOption := kingpin.Flag("profile", "Profiles the program, one of (cpu, mem, block, goroutine, trace, thread, mutex)").String()
	kingpin.Parse()

	if len(*profileOption) > 0 {
		profiler := enableProfiler(*profileOption)
		if profiler != nil {
			defer profiler.Stop()
		}
	}

	windowTitle := fmt.Sprintf("OpenDiablo2 (%s)", a.gitBranch)
	// If we fail to initialize, we will show the error screen
	if err := a.initialize(); err != nil {
		if gameErr := a.renderer.Run(updateInitError, 800, 600, windowTitle); gameErr != nil {
			return gameErr
		}

		return err
	}

	a.ToMainMenu()

	if err := a.renderer.Run(a.update, 800, 600, windowTitle); err != nil {
		return err
	}

	return nil
}

func (a *App) initialize() error {
	a.timeScale = 1.0
	a.lastTime = d2util.Now()
	a.lastScreenAdvance = a.lastTime

	a.renderer.SetWindowIcon("d2logo.png")
	a.terminal.BindLogger()

	terminalActions := [...]bindTerminalEntry{
		{"dumpheap", "dumps the heap to pprof/heap.pprof", a.dumpHeap},
		{"fullscreen", "toggles fullscreen", a.toggleFullScreen},
		{"capframe", "captures a still frame", a.setupCaptureFrame},
		{"capgifstart", "captures an animation (start)", a.startAnimationCapture},
		{"capgifstop", "captures an animation (stop)", a.stopAnimationCapture},
		{"vsync", "toggles vsync", a.toggleVsync},
		{"fps", "toggle fps counter", a.toggleFpsCounter},
		{"timescale", "set scalar for elapsed time", a.setTimeScale},
		{"quit", "exits the game", a.quitGame},
		{"screen-gui", "enters the gui playground screen", a.enterGuiPlayground},
		{"js", "eval JS scripts", a.evalJS},
	}

	for idx := range terminalActions {
		action := &terminalActions[idx]

		if err := a.terminal.BindAction(action.name, action.description, action.action); err != nil {
			log.Fatal(err)
		}
	}

	var err error

	a.guiManager, err = d2gui.CreateGuiManager(a.asset, a.inputManager)
	if err != nil {
		return err
	}

	a.screen = d2screen.NewScreenManager(a.ui, a.guiManager)

	config := d2config.Config
	a.audio.SetVolumes(config.BgmVolume, config.SfxVolume)

	if err := a.loadDataDict(); err != nil {
		return err
	}

	if err := a.loadStrings(); err != nil {
		return err
	}

	d2inventory.LoadHeroObjects()

	a.ui.Initialize()

	return nil
}

func (a *App) loadStrings() error {
	tablePaths := []string{
		d2resource.PatchStringTable,
		d2resource.ExpansionStringTable,
		d2resource.StringTable,
	}

	for _, tablePath := range tablePaths {
		data, err := a.asset.LoadFile(tablePath)
		if err != nil {
			return err
		}

		d2tbl.LoadTextDictionary(data)
	}

	return nil
}

func (a *App) loadDataDict() error {
	entries := []struct {
		path   string
		loader func(data []byte)
	}{
		{d2resource.LevelType, d2datadict.LoadLevelTypes},
		{d2resource.LevelPreset, d2datadict.LoadLevelPresets},
		{d2resource.LevelWarp, d2datadict.LoadLevelWarps},
		{d2resource.ObjectType, d2datadict.LoadObjectTypes},
		{d2resource.ObjectDetails, d2datadict.LoadObjects},
		{d2resource.Weapons, d2datadict.LoadWeapons},
		{d2resource.Armor, d2datadict.LoadArmors},
		{d2resource.Books, d2datadict.LoadBooks},
		{d2resource.Misc, d2datadict.LoadMiscItems},
		{d2resource.UniqueItems, d2datadict.LoadUniqueItems},
		{d2resource.Missiles, d2datadict.LoadMissiles},
		{d2resource.SoundSettings, d2datadict.LoadSounds},
		{d2resource.AnimationData, d2data.LoadAnimationData},
		{d2resource.MonStats, d2datadict.LoadMonStats},
		{d2resource.MonStats2, d2datadict.LoadMonStats2},
		{d2resource.MonPreset, d2datadict.LoadMonPresets},
		{d2resource.MonProp, d2datadict.LoadMonProps},
		{d2resource.MonType, d2datadict.LoadMonTypes},
		{d2resource.MonMode, d2datadict.LoadMonModes},
		{d2resource.MagicPrefix, d2datadict.LoadMagicPrefix},
		{d2resource.MagicSuffix, d2datadict.LoadMagicSuffix},
		{d2resource.ItemStatCost, d2datadict.LoadItemStatCosts},
		{d2resource.ItemRatio, d2datadict.LoadItemRatios},
		{d2resource.Overlays, d2datadict.LoadOverlays},
		{d2resource.CharStats, d2datadict.LoadCharStats},
		{d2resource.Hireling, d2datadict.LoadHireling},
		{d2resource.Experience, d2datadict.LoadExperienceBreakpoints},
		{d2resource.Gems, d2datadict.LoadGems},
		{d2resource.QualityItems, d2datadict.LoadQualityItems},
		{d2resource.Runes, d2datadict.LoadRunewords},
		{d2resource.DifficultyLevels, d2datadict.LoadDifficultyLevels},
		{d2resource.AutoMap, d2datadict.LoadAutoMaps},
		{d2resource.LevelDetails, d2datadict.LoadLevelDetails},
		{d2resource.LevelMaze, d2datadict.LoadLevelMazeDetails},
		{d2resource.LevelSubstitutions, d2datadict.LoadLevelSubstitutions},
		{d2resource.CubeRecipes, d2datadict.LoadCubeRecipes},
		{d2resource.SuperUniques, d2datadict.LoadSuperUniques},
		{d2resource.Inventory, d2datadict.LoadInventory},
		{d2resource.Skills, d2datadict.LoadSkills},
		{d2resource.SkillCalc, d2datadict.LoadSkillCalculations},
		{d2resource.MissileCalc, d2datadict.LoadMissileCalculations},
		{d2resource.Properties, d2datadict.LoadProperties},
		{d2resource.SkillDesc, d2datadict.LoadSkillDescriptions},
		{d2resource.ItemTypes, d2datadict.LoadItemTypes},
		{d2resource.BodyLocations, d2datadict.LoadBodyLocations},
		{d2resource.Sets, d2datadict.LoadSetRecords},
		{d2resource.SetItems, d2datadict.LoadSetItems},
		{d2resource.AutoMagic, d2datadict.LoadAutoMagicRecords},
		{d2resource.TreasureClass, d2datadict.LoadTreasureClassRecords},
		{d2resource.States, d2datadict.LoadStates},
		{d2resource.SoundEnvirons, d2datadict.LoadSoundEnvirons},
		{d2resource.Shrines, d2datadict.LoadShrines},
		{d2resource.ElemType, d2datadict.LoadElemTypes},
		{d2resource.PlrMode, d2datadict.LoadPlrModes},
		{d2resource.PetType, d2datadict.LoadPetTypes},
		{d2resource.NPC, d2datadict.LoadNPCs},
		{d2resource.MonsterUniqueModifier, d2datadict.LoadMonsterUniqueModifiers},
		{d2resource.MonsterEquipment, d2datadict.LoadMonsterEquipment},
		{d2resource.UniqueAppellation, d2datadict.LoadUniqueAppellations},
		{d2resource.MonsterLevel, d2datadict.LoadMonsterLevels},
		{d2resource.MonsterSound, d2datadict.LoadMonsterSounds},
		{d2resource.MonsterSequence, d2datadict.LoadMonsterSequences},
		{d2resource.PlayerClass, d2datadict.LoadPlayerClasses},
		{d2resource.MonsterPlacement, d2datadict.LoadMonsterPlacements},
		{d2resource.ObjectGroup, d2datadict.LoadObjectGroups},
		{d2resource.CompCode, d2datadict.LoadComponentCodes},
		{d2resource.MonsterAI, d2datadict.LoadMonsterAI},
		{d2resource.RarePrefix, d2datadict.LoadRareItemPrefixRecords},
		{d2resource.RareSuffix, d2datadict.LoadRareItemSuffixRecords},
	}

	d2datadict.InitObjectRecords()

	for _, entry := range entries {
		data, err := a.asset.LoadFile(entry.path)
		if err != nil {
			return err
		}

		entry.loader(data)
	}

	d2datadict.LoadItemEquivalencies() // depends on ItemCommon and ItemTypes

	return nil
}

func (a *App) renderDebug(target d2interface.Surface) error {
	if !a.showFPS {
		return nil
	}

	vsyncEnabled := a.renderer.GetVSyncEnabled()
	fps := a.renderer.CurrentFPS()
	cx, cy := a.renderer.GetCursorPos()

	target.PushTranslation(fpsX, fpsY)
	target.DrawTextf("vsync:" + strconv.FormatBool(vsyncEnabled) + "\nFPS:" + strconv.Itoa(int(fps)))
	target.Pop()

	var m runtime.MemStats

	runtime.ReadMemStats(&m)
	target.PushTranslation(memInfoX, memInfoY)
	target.DrawTextf("Alloc    " + strconv.FormatInt(int64(m.Alloc)/bytesToMegabyte, 10))
	target.PushTranslation(0, debugLineHeight)
	target.DrawTextf("TAlloc/s " + strconv.FormatFloat(a.allocRate(m.TotalAlloc, fps), 'f', 2, 64))
	target.PushTranslation(0, debugLineHeight)
	target.DrawTextf("Pause    " + strconv.FormatInt(int64(m.PauseTotalNs/bytesToMegabyte), 10))
	target.PushTranslation(0, debugLineHeight)
	target.DrawTextf("HeapSys  " + strconv.FormatInt(int64(m.HeapSys/bytesToMegabyte), 10))
	target.PushTranslation(0, debugLineHeight)
	target.DrawTextf("NumGC    " + strconv.FormatInt(int64(m.NumGC), 10))
	target.PushTranslation(0, debugLineHeight)
	target.DrawTextf("Coords   " + strconv.FormatInt(int64(cx), 10) + "," + strconv.FormatInt(int64(cy), 10))
	target.PopN(debugPopN)

	return nil
}

func (a *App) renderCapture(target d2interface.Surface) error {
	cleanupCapture := func() {
		a.captureState = captureStateNone
		a.capturePath = ""
		a.captureFrames = nil
	}

	switch a.captureState {
	case captureStateFrame:
		defer cleanupCapture()

		if err := a.doCaptureFrame(target); err != nil {
			return err
		}
	case captureStateGif:
		a.doCaptureGif(target)
	case captureStateNone:
		if len(a.captureFrames) > 0 {
			defer cleanupCapture()

			if err := a.convertFramesToGif(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *App) render(target d2interface.Surface) error {
	if err := a.screen.Render(target); err != nil {
		return err
	}

	a.ui.Render(target)

	if err := a.guiManager.Render(target); err != nil {
		return err
	}

	if err := a.renderDebug(target); err != nil {
		return err
	}

	if err := a.renderCapture(target); err != nil {
		return err
	}

	if err := a.terminal.Render(target); err != nil {
		return err
	}

	return nil
}

func (a *App) advance(elapsed, elapsedUnscaled, current float64) error {
	elapsedLastScreenAdvance := (current - a.lastScreenAdvance) * a.timeScale

	a.lastScreenAdvance = current

	if err := a.screen.Advance(elapsedLastScreenAdvance); err != nil {
		return err
	}

	a.ui.Advance(elapsed)

	if err := a.inputManager.Advance(elapsed, current); err != nil {
		return err
	}

	if err := a.guiManager.Advance(elapsed); err != nil {
		return err
	}

	if err := a.terminal.Advance(elapsedUnscaled); err != nil {
		return err
	}

	return nil
}

func (a *App) update(target d2interface.Surface) error {
	currentTime := d2util.Now()
	elapsedTimeUnscaled := currentTime - a.lastTime
	elapsedTime := elapsedTimeUnscaled * a.timeScale
	a.lastTime = currentTime

	if err := a.advance(elapsedTime, elapsedTimeUnscaled, currentTime); err != nil {
		return err
	}

	if err := a.render(target); err != nil {
		return err
	}

	if target.GetDepth() > 0 {
		return errors.New("detected surface stack leak")
	}

	return nil
}

func (a *App) allocRate(totalAlloc uint64, fps float64) float64 {
	a.tAllocSamples.Value = totalAlloc
	a.tAllocSamples = a.tAllocSamples.Next()
	deltaAllocPerFrame := float64(totalAlloc-a.tAllocSamples.Value.(uint64)) / nSamplesTAlloc

	return deltaAllocPerFrame * fps / bytesToMegabyte
}

func (a *App) dumpHeap() {
	if _, err := os.Stat("./pprof/"); os.IsNotExist(err) {
		if err := os.Mkdir("./pprof/", 0750); err != nil {
			log.Fatal(err)
		}
	}

	fileOut, _ := os.Create("./pprof/heap.pprof")

	if err := pprof.WriteHeapProfile(fileOut); err != nil {
		log.Fatal(err)
	}

	if err := fileOut.Close(); err != nil {
		log.Fatal(err)
	}
}

func (a *App) evalJS(code string) {
	val, err := a.scriptEngine.Eval(code)
	if err != nil {
		a.terminal.OutputErrorf("%s", err)
		return
	}

	log.Printf("%s", val)
}

func (a *App) toggleFullScreen() {
	fullscreen := !a.renderer.IsFullScreen()
	a.renderer.SetFullScreen(fullscreen)
	a.terminal.OutputInfof("fullscreen is now: %v", fullscreen)
}

func (a *App) setupCaptureFrame(path string) {
	a.captureState = captureStateFrame
	a.capturePath = path
	a.captureFrames = nil
}

func (a *App) doCaptureFrame(target d2interface.Surface) error {
	fp, err := os.Create(a.capturePath)
	if err != nil {
		return err
	}

	defer func() {
		if err := fp.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	screenshot := target.Screenshot()
	if err := png.Encode(fp, screenshot); err != nil {
		return err
	}

	log.Printf("saved frame to %s", a.capturePath)

	return nil
}

func (a *App) doCaptureGif(target d2interface.Surface) {
	screenshot := target.Screenshot()
	a.captureFrames = append(a.captureFrames, screenshot)
}

func (a *App) convertFramesToGif() error {
	fp, err := os.Create(a.capturePath)
	if err != nil {
		return err
	}

	defer func() {
		if err := fp.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var (
		framesTotal  = len(a.captureFrames)
		framesPal    = make([]*image.Paletted, framesTotal)
		frameDelays  = make([]int, framesTotal)
		framesPerCPU = framesTotal / runtime.NumCPU()
	)

	var waitGroup sync.WaitGroup

	for i := 0; i < framesTotal; i += framesPerCPU {
		waitGroup.Add(1)

		go func(start, end int) {
			defer waitGroup.Done()

			for j := start; j < end; j++ {
				var buffer bytes.Buffer
				if err := gif.Encode(&buffer, a.captureFrames[j], nil); err != nil {
					panic(err)
				}

				framePal, err := gif.Decode(&buffer)
				if err != nil {
					panic(err)
				}

				framesPal[j] = framePal.(*image.Paletted)
				frameDelays[j] = 5
			}
		}(i, d2math.MinInt(i+framesPerCPU, framesTotal))
	}

	waitGroup.Wait()

	if err := gif.EncodeAll(fp, &gif.GIF{Image: framesPal, Delay: frameDelays}); err != nil {
		return err
	}

	log.Printf("saved animation to %s", a.capturePath)

	return nil
}

func (a *App) startAnimationCapture(path string) {
	a.captureState = captureStateGif
	a.capturePath = path
	a.captureFrames = nil
}

func (a *App) stopAnimationCapture() {
	a.captureState = captureStateNone
}

func (a *App) toggleVsync() {
	vsync := !a.renderer.GetVSyncEnabled()
	a.renderer.SetVSyncEnabled(vsync)
	a.terminal.OutputInfof("vsync is now: %v", vsync)
}

func (a *App) toggleFpsCounter() {
	a.showFPS = !a.showFPS
	a.terminal.OutputInfof("fps counter is now: %v", a.showFPS)
}

func (a *App) setTimeScale(timeScale float64) {
	if timeScale <= 0 {
		a.terminal.OutputErrorf("invalid time scale value")
	} else {
		a.terminal.OutputInfof("timescale changed from %f to %f", a.timeScale, timeScale)
		a.timeScale = timeScale
	}
}

func (a *App) quitGame() {
	os.Exit(0)
}

func (a *App) enterGuiPlayground() {
	a.screen.SetNextScreen(d2gamescreen.CreateGuiTestMain(a.renderer, a.guiManager, a.asset))
}

func createZeroedRing(n int) *ring.Ring {
	r := ring.New(n)
	for i := 0; i < n; i++ {
		r.Value = uint64(0)
		r = r.Next()
	}

	return r
}

func enableProfiler(profileOption string) interface{ Stop() } {
	var options []func(*profile.Profile)

	switch strings.ToLower(strings.Trim(profileOption, " ")) {
	case "cpu":
		log.Printf("CPU profiling is enabled.")

		options = append(options, profile.CPUProfile)
	case "mem":
		log.Printf("Memory profiling is enabled.")

		options = append(options, profile.MemProfile)
	case "block":
		log.Printf("Block profiling is enabled.")

		options = append(options, profile.BlockProfile)
	case "goroutine":
		log.Printf("Goroutine profiling is enabled.")

		options = append(options, profile.GoroutineProfile)
	case "trace":
		log.Printf("Trace profiling is enabled.")

		options = append(options, profile.TraceProfile)
	case "thread":
		log.Printf("Thread creation profiling is enabled.")

		options = append(options, profile.ThreadcreationProfile)
	case "mutex":
		log.Printf("Mutex profiling is enabled.")

		options = append(options, profile.MutexProfile)
	}

	options = append(options, profile.ProfilePath("./pprof/"))

	if len(options) > 1 {
		return profile.Start(options...)
	}

	return nil
}

func updateInitError(target d2interface.Surface) error {
	_ = target.Clear(colornames.Darkred)

	target.PushTranslation(errMsgPadding, errMsgPadding)
	target.DrawTextf(`Could not find the MPQ files in the directory: 
		%s\nPlease put the files and re-run the game.`, d2config.Config.MpqPath)

	return nil
}

// ToMainMenu forces the game to transition to the Main Menu
func (a *App) ToMainMenu() {
	buildInfo := d2gamescreen.BuildInfo{Branch: a.gitBranch, Commit: a.gitCommit}

	mainMenu := d2gamescreen.CreateMainMenu(a, a.asset, a.renderer, a.inputManager, a.audio, a.ui,
		buildInfo)

	a.screen.SetNextScreen(mainMenu)
}

// ToSelectHero forces the game to transition to the Select Hero (create character) screen
func (a *App) ToSelectHero(connType d2clientconnectiontype.ClientConnectionType, host string) {
	selectHero := d2gamescreen.CreateSelectHeroClass(a, a.asset, a.renderer, a.audio, a.ui,
		connType, host)
	a.screen.SetNextScreen(selectHero)
}

// ToCreateGame forces the game to transition to the Create Game screen
func (a *App) ToCreateGame(filePath string, connType d2clientconnectiontype.ClientConnectionType, host string) {
	gameClient, _ := d2client.Create(connType, a.asset, a.scriptEngine)

	if err := gameClient.Open(host, filePath); err != nil {
		// TODO an error screen should be shown in this case
		fmt.Printf("can not connect to the host: %s", host)
	}

	a.screen.SetNextScreen(d2gamescreen.CreateGame(a, a.asset, a.ui, a.renderer, a.inputManager,
		a.audio, gameClient, a.terminal, a.guiManager))
}

// ToCharacterSelect forces the game to transition to the Character Select (load character) screen
func (a *App) ToCharacterSelect(connType d2clientconnectiontype.ClientConnectionType, connHost string) {
	characterSelect := d2gamescreen.CreateCharacterSelect(a, a.asset, a.renderer, a.inputManager,
		a.audio, a.ui, connType, connHost)

	a.screen.SetNextScreen(characterSelect)
}

// ToMapEngineTest forces the game to transition to the map engine test screen
func (a *App) ToMapEngineTest(region, level int) {
	met := d2gamescreen.CreateMapEngineTest(region, level, a.asset, a.terminal, a.renderer,
		a.inputManager,
		a.audio, a.screen)
	a.screen.SetNextScreen(met)
}

// ToCredits forces the game to transition to the credits screen
func (a *App) ToCredits() {
	a.screen.SetNextScreen(d2gamescreen.CreateCredits(a, a.asset, a.renderer, a.ui))
}
