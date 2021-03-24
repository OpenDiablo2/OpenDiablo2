package d2app

import (
	"fmt"
	"path/filepath"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset/types"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2animdata"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
)

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
	a.initTerminalCommands()

	gui, err := d2gui.CreateGuiManager(a.asset, *a.Options.LogLevel, a.inputManager)
	if err != nil {
		return err
	}

	a.guiManager = gui

	a.screen = d2screen.NewScreenManager(a.ui, *a.Options.LogLevel, a.guiManager)

	a.audio.SetVolumes(a.config.BgmVolume, a.config.SfxVolume)

	if err := a.loadStrings(); err != nil {
		return err
	}

	a.ui.Initialize()

	return nil
}

const (
	fmtErrSourceNotFound = `file not found: %q

Please check your config file at %q

Also, verify that the MPQ files exist at %q

Capitalization in the file name matters.
`
)

func (a *App) initConfig(config *d2config.Configuration) error {
	a.config = config

	for _, mpqName := range a.config.MpqLoadOrder {
		cleanDir := filepath.Clean(a.config.MpqPath)
		srcPath := filepath.Join(cleanDir, mpqName)

		err := a.asset.AddSource(srcPath, types.AssetSourceMPQ)
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

	animData, err := d2animdata.Load(animDataBytes)
	if err != nil {
		a.Error(err.Error())
	}

	a.Infof("Loaded %d animation data records", animData.GetRecordsCount())

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
