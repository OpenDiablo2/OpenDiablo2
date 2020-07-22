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

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
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
	"github.com/pkg/profile"
	"golang.org/x/image/colornames"
	"gopkg.in/alecthomas/kingpin.v2"
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
	inputManager      d2interface.InputManager
	terminal          d2interface.Terminal
	scriptEngine      *d2script.ScriptEngine
	audio             d2interface.AudioProvider
	renderer          d2interface.Renderer
	tAllocSamples     *ring.Ring
}

type bindTerminalEntry struct {
	name        string
	description string
	action      interface{}
}

const (
	defaultFPS      = 0.04 // 1/25
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
	renderer d2interface.Renderer) *App {
	result := &App{
		gitBranch:     gitBranch,
		gitCommit:     gitCommit,
		inputManager:  inputManager,
		terminal:      terminal,
		scriptEngine:  scriptEngine,
		audio:         audio,
		renderer:      renderer,
		tAllocSamples: createZeroedRing(nSamplesTAlloc),
	}

	return result
}

// Run executes the application and kicks off the entire game process
func (p *App) Run() error {
	profileOption := kingpin.Flag("profile", "Profiles the program, one of (cpu, mem, block, goroutine, trace, thread, mutex)").String()
	kingpin.Parse()

	if len(*profileOption) > 0 {
		profiler := enableProfiler(*profileOption)
		if profiler != nil {
			defer profiler.Stop()
		}
	}

	windowTitle := fmt.Sprintf("OpenDiablo2 (%s)", p.gitBranch)
	// If we fail to initialize, we will show the error screen
	if err := p.initialize(); err != nil {
		if gameErr := p.renderer.Run(updateInitError, 800, 600, windowTitle); gameErr != nil {
			return gameErr
		}

		return err
	}

	p.ToMainMenu()

	if p.gitBranch == "" {
		p.gitBranch = "Local Build"
	}

	d2common.SetBuildInfo(p.gitBranch, p.gitCommit)

	if err := p.renderer.Run(p.update, 800, 600, windowTitle); err != nil {
		return err
	}

	return nil
}

func (p *App) initialize() error {
	p.timeScale = 1.0
	p.lastTime = d2common.Now()
	p.lastScreenAdvance = p.lastTime

	p.renderer.SetWindowIcon("d2logo.png")
	p.terminal.BindLogger()

	terminalActions := [...]bindTerminalEntry{
		{"dumpheap", "dumps the heap to pprof/heap.pprof", p.dumpHeap},
		{"fullscreen", "toggles fullscreen", p.toggleFullScreen},
		{"capframe", "captures a still frame", p.captureFrame},
		{"capgifstart", "captures an animation (start)", p.startAnimationCapture},
		{"capgifstop", "captures an animation (stop)", p.stopAnimationCapture},
		{"vsync", "toggles vsync", p.toggleVsync},
		{"fps", "toggle fps counter", p.toggleFpsCounter},
		{"timescale", "set scalar for elapsed time", p.setTimeScale},
		{"quit", "exits the game", p.quitGame},
		{"screen-gui", "enters the gui playground screen", p.enterGuiPlayground},
		{"js", "eval JS scripts", p.evalJS},
	}

	for idx := range terminalActions {
		action := &terminalActions[idx]

		if err := p.terminal.BindAction(action.name, action.description, action.action); err != nil {
			log.Fatal(err)
		}
	}

	if err := d2asset.Initialize(p.renderer, p.terminal); err != nil {
		return err
	}

	if err := d2gui.Initialize(p.inputManager); err != nil {
		return err
	}

	config := d2config.Config
	p.audio.SetVolumes(config.BgmVolume, config.SfxVolume)

	if err := p.loadDataDict(); err != nil {
		return err
	}

	if err := p.loadStrings(); err != nil {
		return err
	}

	d2inventory.LoadHeroObjects()

	d2ui.Initialize(p.inputManager, p.audio)

	return nil
}

func (p *App) loadStrings() error {
	tablePaths := []string{
		d2resource.PatchStringTable,
		d2resource.ExpansionStringTable,
		d2resource.StringTable,
	}

	for _, tablePath := range tablePaths {
		data, err := d2asset.LoadFile(tablePath)
		if err != nil {
			return err
		}

		d2common.LoadTextDictionary(data)
	}

	return nil
}

func (p *App) loadDataDict() error {
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
		{d2resource.Misc, d2datadict.LoadMiscItems},
		{d2resource.UniqueItems, d2datadict.LoadUniqueItems},
		{d2resource.Missiles, d2datadict.LoadMissiles},
		{d2resource.SoundSettings, d2datadict.LoadSounds},
		{d2resource.AnimationData, d2data.LoadAnimationData},
		{d2resource.MonStats, d2datadict.LoadMonStats},
		{d2resource.MonStats2, d2datadict.LoadMonStats2},
		{d2resource.MonPreset, d2datadict.LoadMonPresets},
		{d2resource.MagicPrefix, d2datadict.LoadMagicPrefix},
		{d2resource.MagicSuffix, d2datadict.LoadMagicSuffix},
		{d2resource.ItemStatCost, d2datadict.LoadItemStatCosts},
		{d2resource.CharStats, d2datadict.LoadCharStats},
		{d2resource.Hireling, d2datadict.LoadHireling},
		{d2resource.Experience, d2datadict.LoadExperienceBreakpoints},
		{d2resource.Gems, d2datadict.LoadGems},
		{d2resource.DifficultyLevels, d2datadict.LoadDifficultyLevels},
		{d2resource.AutoMap, d2datadict.LoadAutoMaps},
		{d2resource.LevelDetails, d2datadict.LoadLevelDetails},
		{d2resource.LevelMaze, d2datadict.LoadLevelMazeDetails},
		{d2resource.LevelSubstitutions, d2datadict.LoadLevelSubstitutions},
		{d2resource.CubeRecipes, d2datadict.LoadCubeRecipes},
		{d2resource.SuperUniques, d2datadict.LoadSuperUniques},
		{d2resource.Inventory, d2datadict.LoadInventory},
		{d2resource.Skills, d2datadict.LoadSkills},
		{d2resource.Properties, d2datadict.LoadProperties},
		{d2resource.SkillDesc, d2datadict.LoadSkillDescriptions},
	}

	d2datadict.InitObjectRecords()

	for _, entry := range entries {
		data, err := d2asset.LoadFile(entry.path)
		if err != nil {
			return err
		}

		entry.loader(data)
	}

	return nil
}

func (p *App) renderDebug(target d2interface.Surface) error {
	if !p.showFPS {
		return nil
	}

	vsyncEnabled := p.renderer.GetVSyncEnabled()
	fps := p.renderer.CurrentFPS()
	cx, cy := p.renderer.GetCursorPos()

	target.PushTranslation(5, 565)
	target.DrawTextf("vsync:" + strconv.FormatBool(vsyncEnabled) + "\nFPS:" + strconv.Itoa(int(fps)))
	target.Pop()

	var m runtime.MemStats

	runtime.ReadMemStats(&m)
	target.PushTranslation(680, 0)
	target.DrawTextf("Alloc    " + strconv.FormatInt(int64(m.Alloc)/bytesToMegabyte, 10))
	target.PushTranslation(0, 16)
	target.DrawTextf("TAlloc/s " + strconv.FormatFloat(p.allocRate(m.TotalAlloc, fps), 'f', 2, 64))
	target.PushTranslation(0, 16)
	target.DrawTextf("Pause    " + strconv.FormatInt(int64(m.PauseTotalNs/bytesToMegabyte), 10))
	target.PushTranslation(0, 16)
	target.DrawTextf("HeapSys  " + strconv.FormatInt(int64(m.HeapSys/bytesToMegabyte), 10))
	target.PushTranslation(0, 16)
	target.DrawTextf("NumGC    " + strconv.FormatInt(int64(m.NumGC), 10))
	target.PushTranslation(0, 16)
	target.DrawTextf("Coords   " + strconv.FormatInt(int64(cx), 10) + "," + strconv.FormatInt(int64(cy), 10))
	target.PopN(debugPopN)

	return nil
}

func (p *App) renderCapture(target d2interface.Surface) error {
	cleanupCapture := func() {
		p.captureState = captureStateNone
		p.capturePath = ""
		p.captureFrames = nil
	}

	switch p.captureState {
	case captureStateFrame:
		defer cleanupCapture()

		fp, err := os.Create(p.capturePath)
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

		log.Printf("saved frame to %s", p.capturePath)
	case captureStateGif:
		screenshot := target.Screenshot()
		p.captureFrames = append(p.captureFrames, screenshot)
	case captureStateNone:
		if len(p.captureFrames) > 0 {
			defer cleanupCapture()

			fp, err := os.Create(p.capturePath)
			if err != nil {
				return err
			}

			defer func() {
				if err := fp.Close(); err != nil {
					log.Fatal(err)
				}
			}()

			var (
				framesTotal  = len(p.captureFrames)
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
						if err := gif.Encode(&buffer, p.captureFrames[j], nil); err != nil {
							panic(err)
						}

						framePal, err := gif.Decode(&buffer)
						if err != nil {
							panic(err)
						}

						framesPal[j] = framePal.(*image.Paletted)
						frameDelays[j] = 5
					}
				}(i, d2common.MinInt(i+framesPerCPU, framesTotal))
			}

			waitGroup.Wait()

			if err := gif.EncodeAll(fp, &gif.GIF{Image: framesPal, Delay: frameDelays}); err != nil {
				return err
			}

			log.Printf("saved animation to %s", p.capturePath)
		}
	}

	return nil
}

func (p *App) render(target d2interface.Surface) error {
	if err := d2screen.Render(target); err != nil {
		return err
	}

	d2ui.Render(target)

	if err := d2gui.Render(target); err != nil {
		return err
	}

	if err := p.renderDebug(target); err != nil {
		return err
	}

	if err := p.renderCapture(target); err != nil {
		return err
	}

	if err := p.terminal.Render(target); err != nil {
		return err
	}

	return nil
}

func (p *App) advance(elapsed, current float64) error {
	elapsedLastScreenAdvance := (current - p.lastScreenAdvance) * p.timeScale

	if elapsedLastScreenAdvance > defaultFPS {
		p.lastScreenAdvance = current

		if err := d2screen.Advance(elapsedLastScreenAdvance); err != nil {
			return err
		}
	}

	d2ui.Advance(elapsed)

	if err := p.inputManager.Advance(elapsed, current); err != nil {
		return err
	}

	if err := d2gui.Advance(elapsed); err != nil {
		return err
	}

	if err := p.terminal.Advance(elapsed); err != nil {
		return err
	}

	return nil
}

func (p *App) update(target d2interface.Surface) error {
	currentTime := d2common.Now()
	elapsedTime := (currentTime - p.lastTime) * p.timeScale
	p.lastTime = currentTime

	if err := p.advance(elapsedTime, currentTime); err != nil {
		return err
	}

	if err := p.render(target); err != nil {
		return err
	}

	if target.GetDepth() > 0 {
		return errors.New("detected surface stack leak")
	}

	return nil
}

func (p *App) allocRate(totalAlloc uint64, fps float64) float64 {
	p.tAllocSamples.Value = totalAlloc
	p.tAllocSamples = p.tAllocSamples.Next()
	deltaAllocPerFrame := float64(totalAlloc-p.tAllocSamples.Value.(uint64)) / nSamplesTAlloc

	return deltaAllocPerFrame * fps / bytesToMegabyte
}

func (p *App) dumpHeap() {
	if err := os.Mkdir("./pprof/", 0750); err != nil {
		log.Fatal(err)
	}

	fileOut, _ := os.Create("./pprof/heap.pprof")

	if err := pprof.WriteHeapProfile(fileOut); err != nil {
		log.Fatal(err)
	}

	if err := fileOut.Close(); err != nil {
		log.Fatal(err)
	}
}

func (p *App) evalJS(code string) {
	val, err := p.scriptEngine.Eval(code)
	if err != nil {
		p.terminal.OutputErrorf("%s", err)
		return
	}

	log.Printf("%s", val)
}

func (p *App) toggleFullScreen() {
	fullscreen := !p.renderer.IsFullScreen()
	p.renderer.SetFullScreen(fullscreen)
	p.terminal.OutputInfof("fullscreen is now: %v", fullscreen)
}

func (p *App) captureFrame(path string) {
	p.captureState = captureStateFrame
	p.capturePath = path
	p.captureFrames = nil
}

func (p *App) startAnimationCapture(path string) {
	p.captureState = captureStateGif
	p.capturePath = path
	p.captureFrames = nil
}

func (p *App) stopAnimationCapture() {
	p.captureState = captureStateNone
}

func (p *App) toggleVsync() {
	vsync := !p.renderer.GetVSyncEnabled()
	p.renderer.SetVSyncEnabled(vsync)
	p.terminal.OutputInfof("vsync is now: %v", vsync)
}

func (p *App) toggleFpsCounter() {
	p.showFPS = !p.showFPS
	p.terminal.OutputInfof("fps counter is now: %v", p.showFPS)
}

func (p *App) setTimeScale(timeScale float64) {
	if timeScale <= 0 {
		p.terminal.OutputErrorf("invalid time scale value")
	} else {
		p.terminal.OutputInfof("timescale changed from %f to %f", p.timeScale, timeScale)
		p.timeScale = timeScale
	}
}

func (p *App) quitGame() {
	os.Exit(0)
}

func (p *App) enterGuiPlayground() {
	d2screen.SetNextScreen(d2gamescreen.CreateGuiTestMain(p.renderer))
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

	width, height := target.GetSize()

	target.PushTranslation(width/5, height/2)
	target.DrawTextf(`Could not find the MPQ files in the directory: 
		%s\nPlease put the files and re-run the game.`, d2config.Config.MpqPath)

	return nil
}

func (a *App) ToMainMenu() {
	mainMenu := d2gamescreen.CreateMainMenu(a, a.renderer, a.inputManager, a.audio)
	mainMenu.SetScreenMode(d2gamescreen.ScreenModeMainMenu)
	d2screen.SetNextScreen(mainMenu)
}

func (a *App) ToSelectHero(connType d2clientconnectiontype.ClientConnectionType, host string) {
	selectHero := d2gamescreen.CreateSelectHeroClass(a, a.renderer, a.audio, connType, host)
	d2screen.SetNextScreen(selectHero)
}

func (a *App) ToCreateGame(filePath string, connType d2clientconnectiontype.ClientConnectionType, host string) {
	gameClient, _ := d2client.Create(connType, a.scriptEngine)

	if err := gameClient.Open(host, filePath); err != nil {
		// TODO an error screen should be shown in this case
		fmt.Printf("can not connect to the host: %s", host)
	}

	d2screen.SetNextScreen(d2gamescreen.CreateGame(a, a.renderer, a.inputManager, a.audio, gameClient, a.terminal))
}

func (a *App) ToCharacterSelect(connType d2clientconnectiontype.ClientConnectionType, connHost string) {
	characterSelect := d2gamescreen.CreateCharacterSelect(a, a.renderer, a.inputManager, a.audio, connType, connHost)
	d2screen.SetNextScreen(characterSelect)
}

func (a *App) ToMapEngineTest(region int, level int) {
	met := d2gamescreen.CreateMapEngineTest(0, 1, a.terminal, a.renderer, a.inputManager)
	d2screen.SetNextScreen(met)
}

func (a *App) ToCredits() {
	d2screen.SetNextScreen(d2gamescreen.CreateCredits(a, a.renderer))
}
