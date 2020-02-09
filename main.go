package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2gamescene"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	ebiten2 "github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio/ebiten"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render/ebiten"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2scene"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2term"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

// GitBranch is set by the CI build process to the name of the branch
var GitBranch string

// GitCommit is set by the CI build process to the commit hash
var GitCommit string

var singleton struct {
	lastTime  float64
	showFPS   bool
	timeScale float64
}

func main() {
	if len(GitBranch) == 0 {
		GitBranch = "Local Build"
	}

	d2common.SetBuildInfo(GitBranch, GitCommit)

	region := kingpin.Arg("region", "Region type id").Int()
	preset := kingpin.Arg("preset", "Level preset").Int()
	kingpin.Parse()

	log.SetFlags(log.Lshortfile)
	log.Println("OpenDiablo2 - Open source Diablo 2 engine")

	if err := initialize(); err != nil {
		log.Fatal(err)
	}

	if *region == 0 {
		d2scene.SetNextScene(d2gamescene.CreateMainMenu())
	} else {
		d2scene.SetNextScene(d2gamescene.CreateMapEngineTest(*region, *preset))
	}

	windowTitle := fmt.Sprintf("OpenDiablo2 (%s)", GitBranch)
	if err := d2render.Run(update, 800, 600, windowTitle); err != nil {
		log.Fatal(err)
	}
}

func initialize() error {
	singleton.timeScale = 1.0
	singleton.lastTime = d2common.Now()

	if err := d2config.Initialize(); err != nil {
		return err
	}

	config, err := d2config.Get()
	if err != nil {
		return err
	}
	d2resource.LanguageCode = config.Language

	renderer, err := ebiten.CreateRenderer()
	if err != nil {
		return err
	}

	if err := d2render.Initialize(renderer); err != nil {
		return err
	}

	if err := d2render.SetWindowIcon("d2logo.png"); err != nil {
		return err
	}

	if err := d2asset.Initialize(); err != nil {
		return err
	}

	if err := d2input.Initialize(); err != nil {
		return err
	}

	if err := d2gui.Initialize(); err != nil {
		return err
	}

	if err := d2term.Initialize(); err != nil {
		return err
	}

	d2term.BindLogger()
	d2term.BindAction("fullscreen", "toggles fullscreen", func() {
		fullscreen, err := d2render.IsFullScreen()
		if err == nil {
			fullscreen = !fullscreen
			d2render.SetFullScreen(fullscreen)
			d2term.OutputInfo("fullscreen is now: %v", fullscreen)
		} else {
			d2term.OutputError(err.Error())
		}
	})
	d2term.BindAction("vsync", "toggles vsync", func() {
		vsync, err := d2render.GetVSyncEnabled()
		if err == nil {
			vsync = !vsync
			d2render.SetVSyncEnabled(vsync)
			d2term.OutputInfo("vsync is now: %v", vsync)
		} else {
			d2term.OutputError(err.Error())
		}
	})
	d2term.BindAction("fps", "toggle fps counter", func() {
		singleton.showFPS = !singleton.showFPS
		d2term.OutputInfo("fps counter is now: %v", singleton.showFPS)
	})
	d2term.BindAction("timescale", "set scalar for elapsed time", func(timeScale float64) {
		if timeScale <= 0 {
			d2term.OutputError("invalid time scale value")
		} else {
			d2term.OutputInfo("timescale changed from %f to %f", singleton.timeScale, timeScale)
			singleton.timeScale = timeScale
		}
	})
	d2term.BindAction("quit", "exits the game", func() {
		os.Exit(0)
	})

	audioProvider, err := ebiten2.CreateAudio()
	if err != nil {
		return err
	}

	if err := d2audio.Initialize(audioProvider); err != nil {
		return err
	}

	if err := d2audio.SetVolumes(config.BgmVolume, config.SfxVolume); err != nil {
		return err
	}

	if err := loadDataDict(); err != nil {
		return err
	}

	if err := loadPalettes(); err != nil {
		return err
	}

	if err := loadStrings(); err != nil {
		return err
	}

	d2ui.Initialize()

	return nil
}

func update(target d2render.Surface) error {
	currentTime := d2common.Now()
	elapsedTime := (currentTime - singleton.lastTime) * singleton.timeScale
	singleton.lastTime = currentTime

	if err := advance(elapsedTime); err != nil {
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

func advance(elapsed float64) error {
	if err := d2scene.Advance(elapsed); err != nil {
		return err
	}

	d2ui.Advance(elapsed)

	if err := d2input.Advance(elapsed); err != nil {
		return err
	}

	if err := d2gui.Advance(elapsed); err != nil {
		return err
	}

	if err := d2term.Advance(elapsed); err != nil {
		return err
	}

	return nil
}

func render(target d2render.Surface) error {
	if err := d2scene.Render(target); err != nil {
		return err
	}

	d2ui.Render(target)

	if err := d2gui.Render(target); err != nil {
		return err
	}

	if err := d2term.Render(target); err != nil {
		return err
	}

	if err := renderDebug(target); err != nil {
		return err
	}

	return nil
}

func renderDebug(target d2render.Surface) error {
	if singleton.showFPS {
		vsyncEnabled, err := d2render.GetVSyncEnabled()
		if err != nil {
			return err
		}

		fps, err := d2render.CurrentFPS()
		if err != nil {
			return err
		}

		target.PushTranslation(5, 565)
		target.DrawText("vsync:" + strconv.FormatBool(vsyncEnabled) + "\nFPS:" + strconv.Itoa(int(fps)))
		target.Pop()

		cx, cy, err := d2render.GetCursorPos()
		if err != nil {
			return err
		}

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

func loadPalettes() error {
	palNames := []string{
		"act1",
		"act2",
		"act3",
		"act4",
		"act5",
		"endgame",
		"endgame2",
		"fechar",
		"loading",
		"menu0",
		"menu1",
		"menu2",
		"menu3",
		"menu4",
		"sky",
		"static",
		"trademark",
		"units",
	}

	for _, pal := range palNames {
		file, err := d2asset.LoadFile(`data\global\palette\` + pal + `\pal.dat`)
		if err != nil {
			return err
		}

		d2datadict.LoadPalette(d2enum.PaletteType(pal), file)
	}

	log.Printf("Loaded %d palettes", len(d2datadict.Palettes))
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
	}

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

		d2common.LoadDictionary(data)
	}

	return nil
}
