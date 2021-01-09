// Package d2app contains the OpenDiablo2 application shell
package d2app

import (
	"bytes"
	"container/ring"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data"

	"github.com/pkg/profile"
	"golang.org/x/image/colornames"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	ebiten2 "github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio/ebiten"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render/ebiten"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2term"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2gamescreen"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking"
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
	language          string
	charset           string
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
	config            *d2config.Configuration
	*d2util.Logger
	errorMessage error
	*Options
}

// Options is used to store all of the app options that can be set with arguments
type Options struct {
	printVersion *bool
	Debug        *bool
	profiler     *string
	Server       *d2networking.ServerOptions
	LogLevel     *d2util.LogLevel
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

const (
	appLoggerPrefix = "App"
)

// Create creates a new instance of the application
func Create(gitBranch, gitCommit string) *App {
	assetManager, assetError := d2asset.NewAssetManager()

	app := &App{
		gitBranch: gitBranch,
		gitCommit: gitCommit,
		asset:     assetManager,
		Options: &Options{
			Server: &d2networking.ServerOptions{},
		},
		errorMessage: assetError,
	}

	app.Logger = d2util.NewLogger()
	app.Logger.SetPrefix(appLoggerPrefix)
	app.Logger.SetLevel(d2util.LogLevelNone)

	return app
}

func updateNOOP() error {
	return nil
}

func (a *App) startDedicatedServer() error {
	min, max := d2networking.ServerMinPlayers, d2networking.ServerMaxPlayersDefault
	maxPlayers := d2math.ClampInt(*a.Options.Server.MaxPlayers, min, max)

	srvChanIn := make(chan int)
	srvChanLog := make(chan string)

	srvErr := d2networking.StartDedicatedServer(a.asset, srvChanIn, srvChanLog, a.config.LogLevel, maxPlayers)
	if srvErr != nil {
		return srvErr
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // This traps Control-c to safely shut down the server

	go func() {
		<-c
		srvChanIn <- d2networking.ServerEventStop
	}()

	for {
		for data := range srvChanLog {
			a.Info(data)
		}
	}
}

func (a *App) loadEngine() error {
	// Create our renderer
	renderer, err := ebiten.CreateRenderer(a.config)
	if err != nil {
		return err
	}

	a.renderer = renderer

	if a.errorMessage != nil {
		return a.renderer.Run(a.updateInitError, updateNOOP, 800, 600, "OpenDiablo2")
	}

	// if the log level was specified at the command line, use it
	logLevel := *a.Options.LogLevel
	if logLevel == d2util.LogLevelUnspecified {
		logLevel = a.config.LogLevel
	}

	a.asset.SetLogLevel(logLevel)

	audio := ebiten2.CreateAudio(a.config.LogLevel, a.asset)

	inputManager := d2input.NewInputManager()

	term, err := d2term.New(inputManager)
	if err != nil {
		return err
	}

	err = a.asset.BindTerminalCommands(term)
	if err != nil {
		return err
	}

	scriptEngine := d2script.CreateScriptEngine()

	uiManager := d2ui.NewUIManager(a.asset, renderer, inputManager, a.config.LogLevel, audio)

	a.inputManager = inputManager
	a.terminal = term
	a.scriptEngine = scriptEngine
	a.audio = audio
	a.ui = uiManager
	a.tAllocSamples = createZeroedRing(nSamplesTAlloc)

	if a.gitBranch == "" {
		a.gitBranch = "Local Build"
	}

	return nil
}

func (a *App) parseArguments() {
	const (
		versionArg   = "version"
		versionShort = 'v'
		versionDesc  = "Prints the version of the app"

		profilerArg  = "profile"
		profilerDesc = "Profiles the program, one of (cpu, mem, block, goroutine, trace, thread, mutex)"

		serverArg   = "dedicated"
		serverShort = 'd'
		serverDesc  = "Starts a dedicated server"

		playersArg  = "players"
		playersDesc = "Sets the number of max players for the dedicated server"

		loggingArg   = "loglevel"
		loggingShort = 'l'
		loggingDesc  = "Enables verbose logging. Log levels will include those below it. " +
			"0 disables log messages, " +
			"1 shows errors, " +
			"2 shows warnings, " +
			"3 shows info, " +
			"4 shows debug" +
			"5 uses value from config file (default)"
	)

	a.Options.profiler = kingpin.Flag(profilerArg, profilerDesc).String()
	a.Options.Server.Dedicated = kingpin.Flag(serverArg, serverDesc).Short(serverShort).Bool()
	a.Options.printVersion = kingpin.Flag(versionArg, versionDesc).Short(versionShort).Bool()
	a.Options.Server.MaxPlayers = kingpin.Flag(playersArg, playersDesc).Int()
	a.Options.LogLevel = kingpin.Flag(loggingArg, loggingDesc).
		Short(loggingShort).
		Default(strconv.Itoa(d2util.LogLevelUnspecified)).
		Int()

	kingpin.Parse()
}

// LoadConfig loads the OpenDiablo2 config file
func (a *App) LoadConfig() (*d2config.Configuration, error) {
	// by now the, the loader has initialized and added our config dirs as sources...
	configBaseName := filepath.Base(d2config.DefaultConfigPath())

	configAsset, _ := a.asset.LoadAsset(configBaseName)

	config := &d2config.Configuration{}

	// create the default if not found
	if configAsset == nil {
		config = d2config.DefaultConfig()

		fullPath := filepath.Join(config.Dir(), config.Base())
		config.SetPath(fullPath)

		a.Infof("creating default configuration file at %s...", fullPath)

		saveErr := config.Save()

		return config, saveErr
	}

	if err := json.NewDecoder(configAsset).Decode(config); err != nil {
		return nil, err
	}

	config.SetPath(filepath.Join(configAsset.Source().Path(), configAsset.Path()))

	a.Infof("loaded configuration file from %s", config.Path())

	return config, nil
}

// Run executes the application and kicks off the entire game process
func (a *App) Run() error {
	a.parseArguments()

	// add our possible config directories
	_, _ = a.asset.AddSource(filepath.Dir(d2config.LocalConfigPath()))
	_, _ = a.asset.AddSource(filepath.Dir(d2config.DefaultConfigPath()))

	config, err := a.LoadConfig()
	if err != nil {
		return err
	}

	a.config = config

	a.asset.SetLogLevel(config.LogLevel)

	// print version and exit if `--version` was supplied
	if *a.Options.printVersion {
		fmtVersion := "OpenDiablo2 (%s %s)"

		if a.gitBranch == "" {
			a.gitBranch = "local"
		}

		if a.gitCommit == "" {
			a.gitCommit = "build"
		}

		fmt.Printf(fmtVersion, a.gitBranch, a.gitCommit)
		os.Exit(0)
	}

	logLevel := *a.Options.LogLevel
	if logLevel == d2util.LogLevelUnspecified {
		logLevel = a.config.LogLevel
	}

	a.asset.SetLogLevel(logLevel)

	// start profiler if argument was supplied
	if len(*a.Options.profiler) > 0 {
		profiler := enableProfiler(*a.Options.profiler, a)
		if profiler != nil {
			defer profiler.Stop()
		}
	}

	// start the server if `--listen` option was supplied
	if *a.Options.Server.Dedicated {
		if err := a.startDedicatedServer(); err != nil {
			return err
		}
	}

	if err := a.loadEngine(); err != nil {
		a.renderer.ShowPanicScreen(err.Error())
		return err
	}

	windowTitle := fmt.Sprintf("OpenDiablo2 (%s)", a.gitBranch)

	// If we fail to initialize, we will show the error screen
	if err := a.initialize(); err != nil {
		if a.errorMessage == nil {
			a.errorMessage = err // if there was an error during init, don't clobber it
		}

		gameErr := a.renderer.Run(a.updateInitError, updateNOOP, 800, 600, windowTitle)
		if gameErr != nil {
			return gameErr
		}

		return err
	}

	a.ToMainMenu()

	if err := a.renderer.Run(a.update, a.advance, 800, 600, windowTitle); err != nil {
		return err
	}

	return nil
}

func (a *App) initialize() error {
	if err := a.initConfig(a.config); err != nil {
		return err
	}

	a.initLanguage()

	if err := a.initDataDictionaries(); err != nil {
		return err
	}

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
			a.Fatal(err.Error())
		}
	}

	gui, err := d2gui.CreateGuiManager(a.asset, a.config.LogLevel, a.inputManager)
	if err != nil {
		return err
	}

	a.guiManager = gui

	a.screen = d2screen.NewScreenManager(a.ui, a.config.LogLevel, a.guiManager)

	a.audio.SetVolumes(a.config.BgmVolume, a.config.SfxVolume)

	if err := a.loadStrings(); err != nil {
		return err
	}

	a.ui.Initialize()

	return nil
}

const (
	fmtErrSourceNotFound = `file not found: %s

Please check your config file at %s

Also, verify that the MPQ files exist at %s

Capitalization in the file name matters.
`
)

func (a *App) initConfig(config *d2config.Configuration) error {
	a.config = config

	for _, mpqName := range a.config.MpqLoadOrder {
		cleanDir := filepath.Clean(a.config.MpqPath)
		srcPath := filepath.Join(cleanDir, mpqName)

		_, err := a.asset.AddSource(srcPath)
		if err != nil {
			// nolint:stylecheck // we want a multiline error message here..
			return fmt.Errorf(fmtErrSourceNotFound, srcPath, a.config.Path(), a.config.MpqPath)
		}
	}

	return nil
}

func (a *App) initLanguage() {
	a.language = a.asset.LoadLanguage(d2resource.LocalLanguage)
	a.asset.Loader.SetLanguage(&a.language)

	a.charset = d2resource.GetFontCharset(a.language)
	a.asset.Loader.SetCharset(&a.charset)
}

func (a *App) initDataDictionaries() error {
	dictPaths := []string{
		d2resource.LevelType, d2resource.LevelPreset, d2resource.LevelWarp,
		d2resource.ObjectType, d2resource.ObjectDetails, d2resource.Weapons,
		d2resource.Armor, d2resource.Misc, d2resource.Books, d2resource.ItemTypes,
		d2resource.UniqueItems, d2resource.Missiles, d2resource.SoundSettings,
		d2resource.MonStats, d2resource.MonStats2, d2resource.MonPreset,
		d2resource.MonProp, d2resource.MonType, d2resource.MonMode,
		d2resource.MagicPrefix, d2resource.MagicSuffix, d2resource.ItemStatCost,
		d2resource.ItemRatio, d2resource.StorePage, d2resource.Overlays,
		d2resource.CharStats, d2resource.Hireling, d2resource.Experience,
		d2resource.Gems, d2resource.QualityItems, d2resource.Runes,
		d2resource.DifficultyLevels, d2resource.AutoMap, d2resource.LevelDetails,
		d2resource.LevelMaze, d2resource.LevelSubstitutions, d2resource.CubeRecipes,
		d2resource.SuperUniques, d2resource.Inventory, d2resource.Skills,
		d2resource.SkillCalc, d2resource.MissileCalc, d2resource.Properties,
		d2resource.SkillDesc, d2resource.BodyLocations, d2resource.Sets,
		d2resource.SetItems, d2resource.AutoMagic, d2resource.TreasureClass,
		d2resource.TreasureClassEx, d2resource.States, d2resource.SoundEnvirons,
		d2resource.Shrines, d2resource.ElemType, d2resource.PlrMode,
		d2resource.PetType, d2resource.NPC, d2resource.MonsterUniqueModifier,
		d2resource.MonsterEquipment, d2resource.UniqueAppellation, d2resource.MonsterLevel,
		d2resource.MonsterSound, d2resource.MonsterSequence, d2resource.PlayerClass,
		d2resource.MonsterPlacement, d2resource.ObjectGroup, d2resource.CompCode,
		d2resource.MonsterAI, d2resource.RarePrefix, d2resource.RareSuffix,
		d2resource.Events, d2resource.Colors, d2resource.ArmorType,
		d2resource.WeaponClass, d2resource.PlayerType, d2resource.Composite,
		d2resource.HitClass, d2resource.UniquePrefix, d2resource.UniqueSuffix,
		d2resource.CubeModifier, d2resource.CubeType, d2resource.HirelingDescription,
		d2resource.LowQualityItems,
	}

	a.Info("Initializing asset manager")

	for _, path := range dictPaths {
		err := a.asset.LoadRecords(path)
		if err != nil {
			return err
		}
	}

	err := a.initAnimationData(d2resource.AnimationData)
	if err != nil {
		return err
	}

	return nil
}

const (
	fmtLoadAnimData = "loading animation data from: %s"
)

func (a *App) initAnimationData(path string) error {
	animDataBytes, err := a.asset.LoadFile(path)
	if err != nil {
		return err
	}

	a.Debugf(fmtLoadAnimData, path)

	animData := d2data.LoadAnimationData(animDataBytes)

	a.Infof("Loaded %d animation data records", len(animData))

	a.asset.Records.Animation.Data = animData

	return nil
}

func (a *App) loadStrings() error {
	tablePaths := []string{
		d2resource.PatchStringTable,
		d2resource.ExpansionStringTable,
		d2resource.StringTable,
	}

	for _, tablePath := range tablePaths {
		_, err := a.asset.LoadStringTable(tablePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) renderDebug(target d2interface.Surface) {
	if !a.showFPS {
		return
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

func (a *App) render(target d2interface.Surface) {
	a.screen.Render(target)
	a.ui.Render(target)

	if err := a.guiManager.Render(target); err != nil {
		return
	}

	a.renderDebug(target)

	if err := a.renderCapture(target); err != nil {
		return
	}

	if err := a.terminal.Render(target); err != nil {
		return
	}
}

func (a *App) advance() error {
	current := d2util.Now()
	elapsedUnscaled := current - a.lastTime
	elapsed := elapsedUnscaled * a.timeScale

	a.lastTime = current

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
	a.render(target)

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
			a.Fatal(err.Error())
		}
	}

	fileOut, err := os.Create("./pprof/heap.pprof")
	if err != nil {
		a.Error(err.Error())
	}

	if err := pprof.WriteHeapProfile(fileOut); err != nil {
		a.Fatal(err.Error())
	}

	if err := fileOut.Close(); err != nil {
		a.Fatal(err.Error())
	}
}

func (a *App) evalJS(code string) {
	val, err := a.scriptEngine.Eval(code)
	if err != nil {
		a.terminal.OutputErrorf("%s", err)
		return
	}

	a.Info("%s" + val)
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
			a.Fatal(err.Error())
		}
	}()

	screenshot := target.Screenshot()
	if err := png.Encode(fp, screenshot); err != nil {
		return err
	}

	a.Info(fmt.Sprintf("saved frame to %s", a.capturePath))

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
			a.Fatal(err.Error())
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

	a.Info(fmt.Sprintf("saved animation to %s", a.capturePath))

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
	a.screen.SetNextScreen(d2gamescreen.CreateGuiTestMain(a.renderer, a.guiManager, a.config.LogLevel, a.asset))
}

func createZeroedRing(n int) *ring.Ring {
	r := ring.New(n)
	for i := 0; i < n; i++ {
		r.Value = uint64(0)
		r = r.Next()
	}

	return r
}

func enableProfiler(profileOption string, a *App) interface{ Stop() } {
	var options []func(*profile.Profile)

	switch strings.ToLower(strings.Trim(profileOption, " ")) {
	case "cpu":
		a.Logger.Debug("CPU profiling is enabled.")

		options = append(options, profile.CPUProfile)
	case "mem":
		a.Logger.Debug("Memory profiling is enabled.")

		options = append(options, profile.MemProfile)
	case "block":
		a.Logger.Debug("Block profiling is enabled.")

		options = append(options, profile.BlockProfile)
	case "goroutine":
		a.Logger.Debug("Goroutine profiling is enabled.")

		options = append(options, profile.GoroutineProfile)
	case "trace":
		a.Logger.Debug("Trace profiling is enabled.")

		options = append(options, profile.TraceProfile)
	case "thread":
		a.Logger.Debug("Thread creation profiling is enabled.")

		options = append(options, profile.ThreadcreationProfile)
	case "mutex":
		a.Logger.Debug("Mutex profiling is enabled.")

		options = append(options, profile.MutexProfile)
	}

	options = append(options, profile.ProfilePath("./pprof/"))

	if len(options) > 1 {
		return profile.Start(options...)
	}

	return nil
}

func (a *App) updateInitError(target d2interface.Surface) error {
	target.Clear(colornames.Darkred)
	target.PushTranslation(errMsgPadding, errMsgPadding)
	target.DrawTextf(a.errorMessage.Error())

	return nil
}

// ToMainMenu forces the game to transition to the Main Menu
func (a *App) ToMainMenu(errorMessageOptional ...string) {
	buildInfo := d2gamescreen.BuildInfo{Branch: a.gitBranch, Commit: a.gitCommit}

	mainMenu, err := d2gamescreen.CreateMainMenu(a, a.asset, a.renderer, a.inputManager, a.audio, a.ui, buildInfo,
		a.config.LogLevel, errorMessageOptional...)
	if err != nil {
		a.Error(err.Error())
		return
	}

	a.screen.SetNextScreen(mainMenu)
}

// ToSelectHero forces the game to transition to the Select Hero (create character) screen
func (a *App) ToSelectHero(connType d2clientconnectiontype.ClientConnectionType, host string) {
	selectHero, err := d2gamescreen.CreateSelectHeroClass(a, a.asset, a.renderer, a.audio, a.ui, connType, a.config.LogLevel, host)
	if err != nil {
		a.Error(err.Error())
		return
	}

	a.screen.SetNextScreen(selectHero)
}

// ToCreateGame forces the game to transition to the Create Game screen
func (a *App) ToCreateGame(filePath string, connType d2clientconnectiontype.ClientConnectionType, host string) {
	gameClient, err := d2client.Create(connType, a.asset, a.config.LogLevel, a.scriptEngine)
	if err != nil {
		a.Error(err.Error())
	}

	if err = gameClient.Open(host, filePath); err != nil {
		errorMessage := fmt.Sprintf("can not connect to the host: %s", host)
		fmt.Println(errorMessage)
		a.ToMainMenu(errorMessage)
	} else {
		game, err := d2gamescreen.CreateGame(
			a, a.asset, a.ui, a.renderer, a.inputManager, a.audio, gameClient, a.terminal, a.config.LogLevel, a.guiManager,
		)
		if err != nil {
			a.Error(err.Error())
		}

		a.screen.SetNextScreen(game)
	}
}

// ToCharacterSelect forces the game to transition to the Character Select (load character) screen
func (a *App) ToCharacterSelect(connType d2clientconnectiontype.ClientConnectionType, connHost string) {
	characterSelect, err := d2gamescreen.CreateCharacterSelect(a, a.asset, a.renderer, a.inputManager,
		a.audio, a.ui, connType, a.config.LogLevel, connHost)
	if err != nil {
		fmt.Printf("unable to create character select screen: %s", err)
	}

	a.screen.SetNextScreen(characterSelect)
}

// ToMapEngineTest forces the game to transition to the map engine test screen
func (a *App) ToMapEngineTest(region, level int) {
	met, err := d2gamescreen.CreateMapEngineTest(region, level, a.asset, a.terminal, a.renderer, a.inputManager, a.audio,
		a.config.LogLevel, a.screen)
	if err != nil {
		a.Error(err.Error())
		return
	}

	a.screen.SetNextScreen(met)
}

// ToCredits forces the game to transition to the credits screen
func (a *App) ToCredits() {
	a.screen.SetNextScreen(d2gamescreen.CreateCredits(a, a.asset, a.renderer, a.config.LogLevel, a.ui))
}

// ToCinematics forces the game to transition to the cinematics menu
func (a *App) ToCinematics() {
	a.screen.SetNextScreen(d2gamescreen.CreateCinematics(a, a.asset, a.renderer, a.audio, a.config.LogLevel, a.ui))
}
