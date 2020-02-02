package main

import (
	"log"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2config"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2gamescene"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	ebiten2 "github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio/ebiten"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render/ebiten"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2scene"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2term"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"

	"github.com/OpenDiablo2/OpenDiablo2/d2game"
)

// GitBranch is set by the CI build process to the name of the branch
var GitBranch string

// GitCommit is set by the CI build process to the commit hash
var GitCommit string

var region = kingpin.Arg("region", "Region type id").Int()
var preset = kingpin.Arg("preset", "Level preset").Int()

func main() {
	if len(GitBranch) == 0 {
		GitBranch = "Local Build"
		GitCommit = ""
	}
	d2common.SetBuildInfo(GitBranch, GitCommit)
	log.SetFlags(log.Lshortfile)
	log.Println("OpenDiablo2 - Open source Diablo 2 engine")
	if len(GitBranch) == 0 {
		GitBranch = "Local Build"
		GitCommit = ""
	}

	err := initializeEverything()
	if err != nil {
		log.Fatal(err)
		return
	}

	kingpin.Parse()
	if *region == 0 {
		d2scene.SetNextScene(d2gamescene.CreateMainMenu())
	} else {
		d2scene.SetNextScene(d2gamescene.CreateMapEngineTest(*region, *preset))
	}
	err = d2game.Run(GitBranch)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func loadTextDictionary() bool {
	var fileData []byte
	var err error

	toLoad := []string{
		d2resource.PatchStringTable,
		d2resource.ExpansionStringTable,
		d2resource.StringTable,
	}

	for _, item := range toLoad {
		fileData, err = d2asset.LoadFile(item)
		if err != nil {
			log.Fatal(err)
			return false
		}
		d2common.LoadDictionary(fileData)
	}
	log.Printf("Loaded %d entries from the string table", d2common.GetDictionaryEntryCount())
	return true
}

func loadPalettes() bool {
	for _, pal := range []string{
		"act1", "act2", "act3", "act4", "act5", "endgame", "endgame2", "fechar", "loading",
		"menu0", "menu1", "menu2", "menu3", "menu4", "sky", "static", "trademark", "units",
	} {
		filePath := `data\global\palette\` + pal + `\pal.dat`
		paletteType := d2enum.PaletteType(pal)
		file, _ := d2asset.LoadFile(filePath)
		d2datadict.LoadPalette(paletteType, file)
	}
	log.Printf("Loaded %d palettes", len(d2datadict.Palettes))
	return true
}

func initializeEverything() error {
	var err error

	err = d2config.Initialize()
	if err != nil {
		log.Fatal(err)
		return err
	}

	renderer, err := ebiten.CreateRenderer()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = d2render.Initialize(renderer)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = d2input.Initialize()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = d2term.Initialize()
	if err != nil {
		log.Fatal(err)
		return err
	}
	d2term.BindLogger()

	err = d2asset.Initialize()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = d2render.SetWindowIcon("d2logo.png")
	if err != nil {
		log.Fatal(err)
		return err
	}

	config, err := d2config.Get()
	if err != nil {
		log.Fatal(err)
		return err
	}

	var audioProvider *ebiten2.EbitenAudioProvider
	audioProvider, err = ebiten2.CreateAudio()
	err = d2audio.Initialize(audioProvider)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = d2audio.SetVolumes(config.BgmVolume, config.SfxVolume)
	if err != nil {
		log.Fatal(err)
		return err
	}

	d2mpq.InitializeCryptoBuffer()

	settings, _ := d2config.Get()
	d2resource.LanguageCode = settings.Language

	var file []byte

	loadPalettes()

	loadTextDictionary()

	file, err = d2asset.LoadFile(d2resource.LevelType)
	if err != nil {
		return err
	}
	d2datadict.LoadLevelTypes(file)

	file, err = d2asset.LoadFile(d2resource.LevelPreset)
	if err != nil {
		return err
	}
	d2datadict.LoadLevelPresets(file)

	file, err = d2asset.LoadFile(d2resource.LevelWarp)
	if err != nil {
		return err
	}
	d2datadict.LoadLevelWarps(file)

	file, err = d2asset.LoadFile(d2resource.ObjectType)
	if err != nil {
		return err
	}
	d2datadict.LoadObjectTypes(file)

	file, err = d2asset.LoadFile(d2resource.ObjectDetails)
	if err != nil {
		return err
	}
	d2datadict.LoadObjects(file)

	file, err = d2asset.LoadFile(d2resource.Weapons)
	if err != nil {
		return err
	}
	d2datadict.LoadWeapons(file)

	file, err = d2asset.LoadFile(d2resource.Armor)
	if err != nil {
		return err
	}
	d2datadict.LoadArmors(file)

	file, err = d2asset.LoadFile(d2resource.Misc)
	if err != nil {
		return err
	}
	d2datadict.LoadMiscItems(file)

	file, err = d2asset.LoadFile(d2resource.UniqueItems)
	if err != nil {
		return err
	}
	d2datadict.LoadUniqueItems(file)

	file, err = d2asset.LoadFile(d2resource.Missiles)
	if err != nil {
		return err
	}
	d2datadict.LoadMissiles(file)

	file, err = d2asset.LoadFile(d2resource.SoundSettings)
	if err != nil {
		return err
	}
	d2datadict.LoadSounds(file)

	file, err = d2asset.LoadFile(d2resource.AnimationData)
	if err != nil {
		return err
	}
	d2data.LoadAnimationData(file)

	file, err = d2asset.LoadFile(d2resource.MonStats)
	if err != nil {
		return err
	}
	d2datadict.LoadMonStats(file)

	animation, _ := d2asset.LoadAnimation(d2resource.LoadingScreen, d2resource.PaletteLoading)
	loadingSprite, _ := d2ui.LoadSprite(animation)
	loadingSpriteSizeX, loadingSpriteSizeY := loadingSprite.GetCurrentFrameSize()
	loadingSprite.SetPosition(400-(loadingSpriteSizeX/2), 300+(loadingSpriteSizeY/2))
	err = d2game.Initialize(loadingSprite)
	if err != nil {
		return err
	}

	animation, _ = d2asset.LoadAnimation(d2resource.CursorDefault, d2resource.PaletteUnits)
	cursorSprite, _ := d2ui.LoadSprite(animation)
	d2ui.Initialize(cursorSprite)

	err = d2term.BindAction("timescale", "set scalar for elapsed time", func(scale float64) {
		if scale <= 0 {
			err2 := d2term.OutputError("invalid time scale value")
			if err != nil {
				log.Fatal(err2)
			}
		} else {
			err2 := d2term.OutputInfo("timescale changed from %f to %f", d2game.GetTimeScale(), scale)
			if err != nil {
				log.Fatal(err2)
			}
			d2game.SetTimeScale(scale)
		}
	})
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
