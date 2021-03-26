// Package d2app contains the OpenDiablo2 application shell
package d2app

import (
	"bytes"
	"container/ring"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset/types"

	"github.com/pkg/profile"
	"golang.org/x/image/colornames"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
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
	Debug    *bool
	profiler *string
	Server   *d2networking.ServerOptions
	LogLevel *d2util.LogLevel
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
	runtime.LockOSThread()

	logger := d2util.NewLogger()
	logger.SetPrefix(appLoggerPrefix)

	app := &App{
		Logger:    logger,
		gitBranch: gitBranch,
		gitCommit: gitCommit,
		Options: &Options{
			Server: &d2networking.ServerOptions{},
		},
	}
	app.Infof("OpenDiablo2 - Open source Diablo 2 engine")

	app.parseArguments()

	app.SetLevel(*app.Options.LogLevel)

	app.asset, app.errorMessage = d2asset.NewAssetManager(*app.Options.LogLevel)

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

	srvErr := d2networking.StartDedicatedServer(a.asset, srvChanIn, srvChanLog, *a.Options.LogLevel, maxPlayers)
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

	audio := ebiten2.CreateAudio(*a.Options.LogLevel, a.asset)

	inputManager := d2input.NewInputManager()

	term, err := d2term.New(inputManager)
	if err != nil {
		return err
	}

	scriptEngine := d2script.CreateScriptEngine()

	uiManager := d2ui.NewUIManager(a.asset, renderer, inputManager, *a.Options.LogLevel, audio)

	a.inputManager = inputManager
	a.terminal = term
	a.scriptEngine = scriptEngine
	a.audio = audio
	a.ui = uiManager
	a.tAllocSamples = createZeroedRing(nSamplesTAlloc)

	return nil
}

func (a *App) parseArguments() {
	const (
		descProfile = "Profiles the program,\none of (cpu, mem, block, goroutine, trace, thread, mutex)"
		descPlayers = "Sets the number of max players for the dedicated server"
		descLogging = "Enables verbose logging. Log levels will include those below it.\n" +
			" 0 disables log messages\n" +
			" 1 shows fatal\n" +
			" 2 shows error\n" +
			" 3 shows warning\n" +
			" 4 shows info\n" +
			" 5 shows debug\n"
	)

	a.Options.profiler = flag.String("profile", "", descProfile)
	a.Options.Server.Dedicated = flag.Bool("dedicated", false, "Starts a dedicated server")
	a.Options.Server.MaxPlayers = flag.Int("players", 0, descPlayers)
	a.Options.LogLevel = flag.Int("l", d2util.LogLevelDefault, descLogging)
	showVersion := flag.Bool("v", false, "Show version")
	showHelp := flag.Bool("h", false, "Show help")

	flag.Usage = func() {
		fmt.Printf("usage: %s [<flags>]\n\nFlags:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *a.Options.LogLevel >= d2util.LogLevelUnspecified {
		*a.Options.LogLevel = d2util.LogLevelDefault
	}

	if *showVersion {
		a.Infof("version: OpenDiablo2 (%s %s)", a.gitBranch, a.gitCommit)
		os.Exit(0)
	}

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}
}

// LoadConfig loads the OpenDiablo2 config file
func (a *App) LoadConfig() (*d2config.Configuration, error) {
	// by now the, the loader has initialized and added our config dirs as sources...
	configBaseName := filepath.Base(d2config.DefaultConfigPath())

	configAsset, _ := a.asset.LoadAsset(configBaseName)

	config := &d2config.Configuration{}
	config.SetPath(d2config.DefaultConfigPath())

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

	a.Infof("loaded configuration file from %s", config.Path())

	return config, nil
}

// Run executes the application and kicks off the entire game process
func (a *App) Run() (err error) {
	// add our possible config directories
	_ = a.asset.AddSource(filepath.Dir(d2config.LocalConfigPath()), types.AssetSourceFileSystem)
	_ = a.asset.AddSource(filepath.Dir(d2config.DefaultConfigPath()), types.AssetSourceFileSystem)

	if a.config, err = a.LoadConfig(); err != nil {
		return err
	}

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

func (a *App) doCaptureFrame(target d2interface.Surface) error {
	fp, err := os.Create(a.capturePath)
	if err != nil {
		a.terminal.Errorf("failed to create %q", a.capturePath)
		return err
	}

	screenshot := target.Screenshot()
	if err := png.Encode(fp, screenshot); err != nil {
		return err
	}

	if err := fp.Close(); err != nil {
		a.terminal.Errorf("failed to create %q", a.capturePath)
		return nil
	}

	a.terminal.Infof("saved frame to %s", a.capturePath)

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

	a.Infof("saved animation to %s", a.capturePath)

	return nil
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
		*a.Options.LogLevel, errorMessageOptional...)
	if err != nil {
		a.Error(err.Error())
		return
	}

	a.screen.SetNextScreen(mainMenu)
}

// ToSelectHero forces the game to transition to the Select Hero (create character) screen
func (a *App) ToSelectHero(connType d2clientconnectiontype.ClientConnectionType, host string) {
	selectHero, err := d2gamescreen.CreateSelectHeroClass(a, a.asset, a.renderer, a.audio, a.ui, connType, *a.Options.LogLevel, host)
	if err != nil {
		a.Error(err.Error())
		return
	}

	a.screen.SetNextScreen(selectHero)
}

// ToCreateGame forces the game to transition to the Create Game screen
func (a *App) ToCreateGame(filePath string, connType d2clientconnectiontype.ClientConnectionType, host string) {
	gameClient, err := d2client.Create(connType, a.asset, *a.Options.LogLevel, a.scriptEngine)
	if err != nil {
		a.Error(err.Error())
	}

	if gameClient == nil {
		a.Error("could not create client")
		return
	}

	if err = gameClient.Open(host, filePath); err != nil {
		errorMessage := fmt.Sprintf("can not connect to the host: %s", host)
		a.Error(errorMessage)
		a.ToMainMenu(errorMessage)
	} else {
		game, err := d2gamescreen.CreateGame(
			a, a.asset, a.ui, a.renderer, a.inputManager, a.audio, gameClient, a.terminal, *a.Options.LogLevel, a.guiManager,
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
		a.audio, a.ui, connType, *a.Options.LogLevel, connHost)
	if err != nil {
		a.Errorf("unable to create character select screen: %s", err)
	}

	a.screen.SetNextScreen(characterSelect)
}

// ToMapEngineTest forces the game to transition to the map engine test screen
func (a *App) ToMapEngineTest(region, level int) {
	met, err := d2gamescreen.CreateMapEngineTest(region, level, a.asset, a.terminal, a.renderer, a.inputManager, a.audio,
		*a.Options.LogLevel, a.screen)
	if err != nil {
		a.Error(err.Error())
		return
	}

	a.screen.SetNextScreen(met)
}

// ToCredits forces the game to transition to the credits screen
func (a *App) ToCredits() {
	a.screen.SetNextScreen(d2gamescreen.CreateCredits(a, a.asset, a.renderer, *a.Options.LogLevel, a.ui))
}

// ToCinematics forces the game to transition to the cinematics menu
func (a *App) ToCinematics() {
	a.screen.SetNextScreen(d2gamescreen.CreateCinematics(a, a.asset, a.renderer, a.audio, *a.Options.LogLevel, a.ui))
}
