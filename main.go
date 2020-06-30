package main

import (
	"bytes"
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

	ebiten_input "github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/ebiten"

	ebiten2 "github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio/ebiten"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render/ebiten"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2term"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2gamescreen"
	"github.com/OpenDiablo2/OpenDiablo2/d2script"
	"github.com/pkg/profile"
	"gopkg.in/alecthomas/kingpin.v2"
)

// GitBranch is set by the CI build process to the name of the branch
var GitBranch string

// GitCommit is set by the CI build process to the commit hash
var GitCommit string

type captureState int

const (
	captureStateNone captureState = iota
	captureStateFrame
	captureStateGif
)

var singleton struct {
	lastTime          float64
	lastScreenAdvance float64
	showFPS           bool
	timeScale         float64

	captureState  captureState
	capturePath   string
	captureFrames []*image.RGBA
}

var terminal_hack d2interface.Terminal // we need to make this available inside of advance/update/render

func main() {
	region := kingpin.Arg("region", "Region type id").Int()
	preset := kingpin.Arg("preset", "Level preset").Int()
	profileOption := kingpin.Flag("profile", "Profiles the program, one of (cpu, mem, block, goroutine, trace, thread, mutex)").String()

	kingpin.Parse()

	log.SetFlags(log.Lshortfile)
	log.Println("OpenDiablo2 - Open source Diablo 2 engine")

	// Initialize our providers
	audioProvider, err := ebiten2.CreateAudio()
	if err != nil {
		panic(err)
	}

	d2input.Initialize(ebiten_input.InputService{}) // TODO d2input singleton must be init before d2term
	term, err := d2term.Initialize()
	terminal_hack = term // needs to be used in advance, no easy way for that right now
	if err != nil {
		log.Fatal(err)
	}

	if err := initialize(audioProvider, term); err != nil {
		if os.IsNotExist(err) {
			run(updateInitError)
		}
		log.Fatal(err)
	}

	if len(*profileOption) > 0 {
		profiler := enableProfiler(*profileOption)
		if profiler != nil {
			defer profiler.Stop()
		}
	}

	if *region == 0 {
		d2screen.SetNextScreen(d2gamescreen.CreateMainMenu(audioProvider, term))
	} else {
		d2screen.SetNextScreen(d2gamescreen.CreateMapEngineTest(*region, *preset, term))
	}

	run(update)
}

func initialize(audioProvider d2interface.AudioProvider, term d2interface.Terminal) error {
	singleton.timeScale = 1.0
	singleton.lastTime = d2common.Now()
	singleton.lastScreenAdvance = singleton.lastTime

	if err := d2config.Load(); err != nil {
		return err
	}

	config := d2config.Get()
	d2resource.LanguageCode = config.Language

	renderer, err := ebiten.CreateRenderer()
	if err != nil {
		return err
	}

	if err := d2render.Initialize(renderer); err != nil {
		return err
	}
	d2render.SetWindowIcon("d2logo.png")

	term.BindLogger()
	term.BindAction("dumpheap", "dumps the heap to pprof/heap.pprof", func() {
		os.Mkdir("./pprof/", 0755)
		fileOut, _ := os.Create("./pprof/heap.pprof")
		pprof.WriteHeapProfile(fileOut)
		fileOut.Close()
	})
	term.BindAction("fullscreen", "toggles fullscreen", func() {
		fullscreen := !d2render.IsFullScreen()
		d2render.SetFullScreen(fullscreen)
		term.OutputInfo("fullscreen is now: %v", fullscreen)
	})
	term.BindAction("capframe", "captures a still frame", func(path string) {
		singleton.captureState = captureStateFrame
		singleton.capturePath = path
		singleton.captureFrames = nil
	})
	term.BindAction("capgifstart", "captures an animation (start)", func(path string) {
		singleton.captureState = captureStateGif
		singleton.capturePath = path
		singleton.captureFrames = nil
	})
	term.BindAction("capgifstop", "captures an animation (stop)", func() {
		singleton.captureState = captureStateNone
	})
	term.BindAction("vsync", "toggles vsync", func() {
		vsync := !d2render.GetVSyncEnabled()
		d2render.SetVSyncEnabled(vsync)
		term.OutputInfo("vsync is now: %v", vsync)
	})
	term.BindAction("fps", "toggle fps counter", func() {
		singleton.showFPS = !singleton.showFPS
		term.OutputInfo("fps counter is now: %v", singleton.showFPS)
	})
	term.BindAction("timescale", "set scalar for elapsed time", func(timeScale float64) {
		if timeScale <= 0 {
			term.OutputError("invalid time scale value")
		} else {
			term.OutputInfo("timescale changed from %f to %f", singleton.timeScale, timeScale)
			singleton.timeScale = timeScale
		}
	})
	term.BindAction("quit", "exits the game", func() {
		os.Exit(0)
	})
	term.BindAction("screen-gui", "enters the gui playground screen", func() {
		d2screen.SetNextScreen(d2gamescreen.CreateGuiTestMain())
	})

	if err := d2asset.Initialize(term); err != nil {
		return err
	}

	if err := d2gui.Initialize(); err != nil {
		return err
	}

	audioProvider.SetVolumes(config.BgmVolume, config.SfxVolume)

	if err := loadDataDict(); err != nil {
		return err
	}

	if err := loadStrings(); err != nil {
		return err
	}

	d2inventory.LoadHeroObjects()

	d2ui.Initialize(audioProvider)

	d2script.CreateScriptEngine()

	return nil
}

func run(updateFunc func(d2interface.Surface) error) {
	if len(GitBranch) == 0 {
		GitBranch = "Local Build"
	}
	d2common.SetBuildInfo(GitBranch, GitCommit)
	windowTitle := fmt.Sprintf("OpenDiablo2 (%s)", GitBranch)
	if err := d2render.Run(updateFunc, 800, 600, windowTitle); err != nil {
		log.Fatal(err)
	}
}

func update(target d2interface.Surface) error {
	currentTime := d2common.Now()
	elapsedTime := (currentTime - singleton.lastTime) * singleton.timeScale
	singleton.lastTime = currentTime

	if err := advance(elapsedTime, currentTime); err != nil {
		return err
	}

	if err := render(target); err != nil {
		return err
	}

	if target.GetDepth() > 0 {
		return errors.New("detected surface stack leak")
	}

	return nil
}

func updateInitError(target d2interface.Surface) error {
	width, height := target.GetSize()
	target.PushTranslation(width/5, height/2)
	target.DrawText("Could not find the MPQ files in the directory: %s\nPlease put the files and re-run the game.", d2config.Get().MpqPath)
	return nil
}

const FPS_25 = 0.04 // 1/25

func advance(elapsed, current float64) error {
	elapsedLastScreenAdvance := (current - singleton.lastScreenAdvance) * singleton.timeScale

	if elapsedLastScreenAdvance > FPS_25 {
		singleton.lastScreenAdvance = current
		if err := d2screen.Advance(elapsedLastScreenAdvance); err != nil {
			return err
		}
	}

	d2ui.Advance(elapsed)

	if err := d2input.Advance(elapsed); err != nil {
		return err
	}

	if err := d2gui.Advance(elapsed); err != nil {
		return err
	}

	if err := terminal_hack.Advance(elapsed); err != nil {
		return err
	}

	return nil
}

func render(target d2interface.Surface) error {
	if err := d2screen.Render(target); err != nil {
		return err
	}

	d2ui.Render(target)

	if err := d2gui.Render(target); err != nil {
		return err
	}

	if err := renderDebug(target); err != nil {
		return err
	}

	if err := renderCapture(target); err != nil {
		return err
	}

	if err := terminal_hack.Render(target); err != nil {
		return err
	}

	return nil
}

func renderCapture(target d2interface.Surface) error {
	cleanupCapture := func() {
		singleton.captureState = captureStateNone
		singleton.capturePath = ""
		singleton.captureFrames = nil
	}

	switch singleton.captureState {
	case captureStateFrame:
		defer cleanupCapture()

		fp, err := os.Create(singleton.capturePath)
		if err != nil {
			return err
		}

		defer fp.Close()

		screenshot := target.Screenshot()
		if err := png.Encode(fp, screenshot); err != nil {
			return err
		}

		log.Printf("saved frame to %s", singleton.capturePath)
	case captureStateGif:
		screenshot := target.Screenshot()
		singleton.captureFrames = append(singleton.captureFrames, screenshot)
	case captureStateNone:
		if len(singleton.captureFrames) > 0 {
			defer cleanupCapture()

			fp, err := os.Create(singleton.capturePath)
			if err != nil {
				return err
			}

			defer fp.Close()

			var (
				framesTotal  = len(singleton.captureFrames)
				framesPal    = make([]*image.Paletted, framesTotal)
				frameDelays  = make([]int, framesTotal)
				framesPerCpu = framesTotal / runtime.NumCPU()
			)

			var waitGroup sync.WaitGroup
			for i := 0; i < framesTotal; i += framesPerCpu {
				waitGroup.Add(1)
				go func(start, end int) {
					defer waitGroup.Done()

					for j := start; j < end; j++ {
						var buffer bytes.Buffer
						if err := gif.Encode(&buffer, singleton.captureFrames[j], nil); err != nil {
							panic(err)
						}

						framePal, err := gif.Decode(&buffer)
						if err != nil {
							panic(err)
						}

						framesPal[j] = framePal.(*image.Paletted)
						frameDelays[j] = 5
					}
				}(i, d2common.MinInt(i+framesPerCpu, framesTotal))
			}

			waitGroup.Wait()

			if err := gif.EncodeAll(fp, &gif.GIF{Image: framesPal, Delay: frameDelays}); err != nil {
				return err
			}

			log.Printf("saved animation to %s", singleton.capturePath)
		}
	}
	return nil
}

func renderDebug(target d2interface.Surface) error {
	if singleton.showFPS {
		vsyncEnabled := d2render.GetVSyncEnabled()
		fps := d2render.CurrentFPS()
		cx, cy := d2render.GetCursorPos()

		target.PushTranslation(5, 565)
		target.DrawText("vsync:" + strconv.FormatBool(vsyncEnabled) + "\nFPS:" + strconv.Itoa(int(fps)))
		target.Pop()

		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		target.PushTranslation(680, 0)
		target.DrawText("Alloc   " + strconv.FormatInt(int64(m.Alloc)/1024/1024, 10))
		target.PushTranslation(0, 16)
		target.DrawText("Pause   " + strconv.FormatInt(int64(m.PauseTotalNs/1024/1024), 10))
		target.PushTranslation(0, 16)
		target.DrawText("HeapSys " + strconv.FormatInt(int64(m.HeapSys/1024/1024), 10))
		target.PushTranslation(0, 16)
		target.DrawText("NumGC   " + strconv.FormatInt(int64(m.NumGC), 10))
		target.PushTranslation(0, 16)
		target.DrawText("Coords  " + strconv.FormatInt(int64(cx), 10) + "," + strconv.FormatInt(int64(cy), 10))
		target.PopN(5)
	}

	return nil
}

func loadDataDict() error {
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

func loadStrings() error {
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
