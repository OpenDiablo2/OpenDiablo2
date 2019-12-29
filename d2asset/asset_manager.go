package d2asset

import (
	"errors"
	"log"

	"github.com/OpenDiablo2/D2Shared/d2data/d2cof"
	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"
	"github.com/OpenDiablo2/D2Shared/d2data/d2dc6"
	"github.com/OpenDiablo2/D2Shared/d2data/d2dcc"
	"github.com/OpenDiablo2/D2Shared/d2data/d2mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2corecommon"
	"github.com/OpenDiablo2/OpenDiablo2/d2term"
)

const (
	// In megabytes
	ArchiveBudget = 1024 * 1024 * 512
	FileBudget    = 1024 * 1024 * 32

	// In counts
	PaletteBudget   = 64
	AnimationBudget = 64
)

var (
	ErrHasInit error = errors.New("asset system is already initialized")
	ErrNoInit  error = errors.New("asset system is not initialized")
)

type assetManager struct {
	archiveManager   *archiveManager
	fileManager      *fileManager
	paletteManager   *paletteManager
	animationManager *animationManager
}

var singleton *assetManager

func Initialize(config *d2corecommon.Configuration) error {
	if singleton != nil {
		return ErrHasInit
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

		archiveManager.cache.verbose = verbose
		fileManager.cache.verbose = verbose
		paletteManager.cache.verbose = verbose
		animationManager.cache.verbose = verbose
	})

	d2term.BindAction("assetstat", "display asset manager cache statistics", func() {
		d2term.OutputInfo("archive cache: %f%%", float64(archiveManager.cache.weight)/float64(archiveManager.cache.budget)*100.0)
		d2term.OutputInfo("file cache: %f%%", float64(fileManager.cache.weight)/float64(fileManager.cache.budget)*100.0)
		d2term.OutputInfo("palette cache: %f%%", float64(paletteManager.cache.weight)/float64(paletteManager.cache.budget)*100.0)
		d2term.OutputInfo("animation cache: %f%%", float64(animationManager.cache.weight)/float64(animationManager.cache.budget)*100.0)
	})

	d2term.BindAction("assetclear", "clear asset manager cache", func() {
		archiveManager.cache.clear()
		fileManager.cache.clear()
		paletteManager.cache.clear()
		animationManager.cache.clear()
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
	return createComposite(object, palettePath), nil
}

func loadPalette(palettePath string) (*d2datadict.PaletteRec, error) {
	if singleton == nil {
		return nil, ErrNoInit
	}

	return singleton.paletteManager.loadPalette(palettePath)
}

func loadDC6(dc6Path, palettePath string) (*d2dc6.DC6File, error) {
	dc6Data, err := LoadFile(dc6Path)
	if err != nil {
		return nil, err
	}

	paletteData, err := loadPalette(palettePath)
	if err != nil {
		return nil, err
	}

	dc6, err := d2dc6.LoadDC6(dc6Data, *paletteData)
	if err != nil {
		return nil, err
	}

	return &dc6, nil
}

func loadDCC(dccPath string) (*d2dcc.DCC, error) {
	dccData, err := LoadFile(dccPath)
	if err != nil {
		return nil, err
	}

	return d2dcc.LoadDCC(dccData)
}

func loadCOF(cofPath string) (*d2cof.COF, error) {
	cofData, err := LoadFile(cofPath)
	if err != nil {
		return nil, err
	}

	return d2cof.LoadCOF(cofData)
}
