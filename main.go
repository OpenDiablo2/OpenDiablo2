package main

import (
	"log"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2gamescene"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	ebiten2 "github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio/ebiten"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
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

func main() {
	if len(GitBranch) == 0 {
		d2common.SetBuildInfo("Local Build", "")
	} else {
		d2common.SetBuildInfo(GitBranch, GitCommit)
	}

	region := kingpin.Arg("region", "Region type id").Int()
	preset := kingpin.Arg("preset", "Level preset").Int()
	kingpin.Parse()

	log.SetFlags(log.Lshortfile)
	log.Println("OpenDiablo2 - Open source Diablo 2 engine")

	if err := initializeEverything(); err != nil {
		log.Fatal(err)
	}

	if *region == 0 {
		d2scene.SetNextScene(d2gamescene.CreateMainMenu())
	} else {
		d2scene.SetNextScene(d2gamescene.CreateMapEngineTest(*region, *preset))
	}

	if err := d2game.Run(GitBranch); err != nil {
		log.Fatal(err)
	}
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

	log.Printf("Loaded %d entries from the string table", d2common.GetDictionaryEntryCount())
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

func loadLoadingSprite() (*d2ui.Sprite, error) {
	animation, err := d2asset.LoadAnimation(d2resource.LoadingScreen, d2resource.PaletteLoading)
	if err != nil {
		return nil, err
	}

	loadingSprite, err := d2ui.LoadSprite(animation)
	if err != nil {
		return nil, err
	}

	loadingSpriteSizeX, loadingSpriteSizeY := loadingSprite.GetCurrentFrameSize()
	loadingSprite.SetPosition(400-(loadingSpriteSizeX/2), 300+(loadingSpriteSizeY/2))
	return loadingSprite, nil
}

func loadCursorSprite() (*d2ui.Sprite, error) {
	animation, err := d2asset.LoadAnimation(d2resource.CursorDefault, d2resource.PaletteUnits)
	if err != nil {
		return nil, err
	}

	cursorSprite, err := d2ui.LoadSprite(animation)
	if err != nil {
		return nil, err
	}

	return cursorSprite, nil
}

func initializeEverything() error {
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

	if err := d2input.Initialize(); err != nil {
		return err
	}

	if err := d2term.Initialize(); err != nil {
		return err
	}
	d2term.BindLogger()

	d2mpq.InitializeCryptoBuffer()
	if err := d2asset.Initialize(); err != nil {
		return err
	}

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

	cursorSprite, err := loadCursorSprite()
	if err != nil {
		return err
	}
	d2ui.Initialize(cursorSprite)

	loadingSprite, err := loadLoadingSprite()
	if err != nil {
		return err
	}
	d2game.Initialize(loadingSprite)

	return nil
}
