package d2asset

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2term"
)

var singleton *assetManager

func Initialize() error {
	if singleton != nil {
		return ErrHasInit
	}

	config, err := d2config.Get()
	if err != nil {
		return err
	}

	var (
		archiveManager   = createArchiveManager(config)
		fileManager      = createFileManager(config, archiveManager)
		paletteManager   = createPaletteManager()
		animationManager = createAnimationManager()
	)

	singleton = &assetManager{
		archiveManager,
		fileManager,
		paletteManager,
		animationManager,
	}

	d2term.BindAction("assetspam", "display verbose asset manager logs", func(verbose bool) {
		if verbose {
			d2term.OutputInfo("asset manager verbose logging enabled")
		} else {
			d2term.OutputInfo("asset manager verbose logging disabled")
		}

		archiveManager.cache.SetVerbose(verbose)
		fileManager.cache.SetVerbose(verbose)
		paletteManager.cache.SetVerbose(verbose)
		animationManager.cache.SetVerbose(verbose)
	})

	d2term.BindAction("assetstat", "display asset manager cache statistics", func() {
		d2term.OutputInfo("archive cache: %f%%", float64(archiveManager.cache.GetWeight())/float64(archiveManager.cache.GetBudget())*100.0)
		d2term.OutputInfo("file cache: %f%%", float64(fileManager.cache.GetWeight())/float64(fileManager.cache.GetBudget())*100.0)
		d2term.OutputInfo("palette cache: %f%%", float64(paletteManager.cache.GetWeight())/float64(paletteManager.cache.GetBudget())*100.0)
		d2term.OutputInfo("animation cache: %f%%", float64(animationManager.cache.GetWeight())/float64(animationManager.cache.GetBudget())*100.0)
	})

	d2term.BindAction("assetclear", "clear asset manager cache", func() {
		archiveManager.cache.Clear()
		fileManager.cache.Clear()
		paletteManager.cache.Clear()
		animationManager.cache.Clear()
	})

	return nil
}

func Shutdown() {
	singleton = nil
}

func LoadArchive(archivePath string) (*d2mpq.MPQ, error) {
	if singleton == nil {
		return nil, ErrNoInit
	}

	return singleton.archiveManager.loadArchive(archivePath)
}

func LoadFile(filePath string) ([]byte, error) {
	if singleton == nil {
		return nil, ErrNoInit
	}

	data, err := singleton.fileManager.loadFile(filePath)
	if err != nil {
		log.Printf("error loading file %s (%v)", filePath, err.Error())
	}

	return data, err
}

func FileExists(filePath string) (bool, error) {
	if singleton == nil {
		return false, ErrNoInit
	}

	return singleton.fileManager.fileExists(filePath)
}

func LoadAnimation(animationPath, palettePath string) (*Animation, error) {
	return LoadAnimationWithTransparency(animationPath, palettePath, 255)
}

func LoadAnimationWithTransparency(animationPath, palettePath string, transparency int) (*Animation, error) {
	if singleton == nil {
		return nil, ErrNoInit
	}

	return singleton.animationManager.loadAnimation(animationPath, palettePath, transparency)
}

func LoadComposite(object *d2datadict.ObjectLookupRecord, palettePath string) (*Composite, error) {
	return CreateComposite(object, palettePath), nil
}
