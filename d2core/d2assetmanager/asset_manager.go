package d2assetmanager

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2config"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2archivemanager"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2filemanager"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2cof"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2term"
)

const (
	AnimationBudget = 64
)

var (
	ErrHasInit error = errors.New("asset system is already initialized")
	ErrNoInit  error = errors.New("asset system is not initialized")
)

type AssetManager struct {
	archiveManager *d2archivemanager.ArchiveManager
	fileManager    *d2filemanager.FileManager
	paletteManager *PaletteManager
	cache          *d2common.Cache
}

var singleton *AssetManager

func Initialize() error {
	if singleton != nil {
		return ErrHasInit
	}
	config, _ := d2config.Get()
	var (
		archiveManager = d2archivemanager.CreateArchiveManager(config)
		fileManager    = d2filemanager.CreateFileManager(config, archiveManager)
		paletteManager = CreatePaletteManager()
		//animationManager = d2animationmanager.CreateAnimationManager()
	)

	singleton = &AssetManager{
		archiveManager,
		fileManager,
		paletteManager,
		nil,
	}
	singleton.cache = d2common.CreateCache(AnimationBudget)

	d2term.BindAction("assetspam", "display verbose asset manager logs", func(verbose bool) {
		if verbose {
			d2term.OutputInfo("asset manager verbose logging enabled")
		} else {
			d2term.OutputInfo("asset manager verbose logging disabled")
		}

		archiveManager.SetCacheVerbose(verbose)
		fileManager.SetCacheVerbose(verbose)
		paletteManager.SetCacheVerbose(verbose)
	})

	d2term.BindAction("assetstat", "display asset manager cache statistics", func() {
		d2term.OutputInfo("archive cache: %f%%", float64(archiveManager.GetCacheWeight())/float64(archiveManager.GetCacheBudget())*100.0)
		d2term.OutputInfo("file cache: %f%%", float64(fileManager.GetCacheWeight())/float64(fileManager.GetCacheBudget())*100.0)
		d2term.OutputInfo("palette cache: %f%%", float64(paletteManager.GetCacheWeight())/float64(paletteManager.GetCacheBudget())*100.0)
		//d2term.OutputInfo("animation cache: %f%%", float64(GetCacheWeight())/float64(GetCacheBudget())*100.0)
	})

	d2term.BindAction("assetclear", "clear asset manager cache", func() {
		archiveManager.ClearCache()
		fileManager.ClearCache()
		paletteManager.ClearCache()
		//am.ClearCache()
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

	return singleton.archiveManager.LoadArchive(archivePath)
}

func LoadFile(filePath string) ([]byte, error) {
	if singleton == nil {
		return nil, ErrNoInit
	}

	data, err := singleton.fileManager.LoadFile(filePath)
	if err != nil {
		log.Printf("error loading file %s (%v)", filePath, err.Error())
	}

	return data, err
}

func FileExists(filePath string) (bool, error) {
	if singleton == nil {
		return false, ErrNoInit
	}

	return singleton.fileManager.FileExists(filePath)
}

func LoadAnimation(animationPath, palettePath string) (*d2render.Animation, error) {
	return LoadAnimationWithTransparency(animationPath, palettePath, 255)
}

func LoadAnimationWithTransparency(animationPath, palettePath string, transparency int) (*d2render.Animation, error) {
	if singleton == nil {
		return nil, ErrNoInit
	}

	return singleton.LoadAnimation(animationPath, palettePath, transparency)
}

func LoadComposite(object *d2datadict.ObjectLookupRecord, palettePath string) (*Composite, error) {
	return CreateComposite(object, palettePath), nil
}

func loadPalette(palettePath string) (*d2datadict.PaletteRec, error) {
	if singleton == nil {
		return nil, ErrNoInit
	}

	return singleton.paletteManager.LoadPalette(palettePath)
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

func (am *AssetManager) SetCacheVerbose(verbose bool) {
	am.cache.SetCacheVerbose(verbose)
}

func (am *AssetManager) ClearCache() {
	am.cache.Clear()
}

func (am *AssetManager) GetCacheWeight() int {
	return am.cache.GetWeight()
}

func (am *AssetManager) GetCacheBudget() int {
	return am.cache.GetBudget()
}

func (am *AssetManager) LoadAnimation(animationPath, palettePath string, transparency int) (*d2render.Animation, error) {
	cachePath := fmt.Sprintf("%s;%s;%d", animationPath, palettePath, transparency)
	if animation, found := am.cache.Retrieve(cachePath); found {
		return animation.(*d2render.Animation).Clone(), nil
	}

	var animation *d2render.Animation
	switch strings.ToLower(filepath.Ext(animationPath)) {
	case ".dc6":
		dc6, err := loadDC6(animationPath, palettePath)
		if err != nil {
			return nil, err
		}

		animation, err = d2render.CreateAnimationFromDC6(dc6)
		if err != nil {
			return nil, err
		}
	case ".dcc":
		dcc, err := loadDCC(animationPath)
		if err != nil {
			return nil, err
		}

		palette, err := loadPalette(palettePath)
		if err != nil {
			return nil, err
		}

		animation, err = d2render.CreateAnimationFromDCC(dcc, palette, transparency)
		if err != nil {
			return nil, err
		}

	default:
		return nil, errors.New("unknown animation format")
	}

	if err := am.cache.Insert(cachePath, animation.Clone(), 1); err != nil {
		return nil, err
	}

	return animation, nil
}
