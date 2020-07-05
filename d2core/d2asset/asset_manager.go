package d2asset

import (
	"errors"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2cof"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

type assetManager struct {
	app 					d2interface.App
	archiveManager          d2interface.ArchiveManager
	archivedFileManager     d2interface.ArchivedFileManager
	paletteManager          d2interface.ArchivedPaletteManager
	paletteTransformManager *paletteTransformManager
	animationManager        d2interface.ArchivedAnimationManager
	fontManager             d2interface.ArchivedFontManager
}

// BindApp binds to the OpenDiablo2 app
func (am *assetManager) BindApp(app d2interface.App) error {
	if am.app != nil {
		return errors.New("asset manager already bound to an app instance")
	}

	am.app = app

	return nil
}

func (am *assetManager) Initialize() error {
	term, err := am.app.Terminal()
	if err != nil {
		return err
	}

	renderer, err := am.app.Renderer()
	if err != nil {
		return err
	}

	var (
		config                  = d2config.Get()
		archiveManager          = createArchiveManager(config)
		archivedFileManager     = createFileManager(config, archiveManager)
		paletteManager          = createPaletteManager()
		// paletteTransformManager = createPaletteTransformManager()
		animationManager        = createAnimationManager(renderer)
		fontManager             = createFontManager()
	)

	singleton.Bind(archiveManager)
	singleton.Bind(archivedFileManager)
	singleton.Bind(paletteManager)
	// singleton.Bind(paletteTransformManager) todo
	singleton.Bind(animationManager)
	singleton.Bind(fontManager)

	if err := term.BindAction("assetspam", "display verbose asset manager logs", func(verbose bool) {
		if verbose {
			term.OutputInfof("asset manager verbose logging enabled")
		} else {
			term.OutputInfof("asset manager verbose logging disabled")
		}

		archiveManager.GetCache().SetVerbose(verbose)
		archivedFileManager.GetCache().SetVerbose(verbose)
		paletteManager.GetCache().SetVerbose(verbose)
		am.paletteTransformManager.cache.SetVerbose(verbose)
		animationManager.cache.SetVerbose(verbose)
	}); err != nil {
		return err
	}

	if err := term.BindAction("assetstat", "display asset manager cache statistics", func() {
		var cacheStatistics = func(c d2interface.Cache) float64 {
			const percent = 100.0
			return float64(c.GetWeight()) / float64(c.GetBudget()) * percent
		}

		term.OutputInfof("archive cache: %f", cacheStatistics(archiveManager.GetCache()))
		term.OutputInfof("file cache: %f", cacheStatistics(archivedFileManager.GetCache()))
		term.OutputInfof("palette cache: %f", cacheStatistics(paletteManager.GetCache()))
		term.OutputInfof("palette transform cache: %f", cacheStatistics(am.paletteTransformManager.
			cache))
		term.OutputInfof("animation cache: %f", cacheStatistics(animationManager.cache))
		term.OutputInfof("font cache: %f", cacheStatistics(fontManager.GetCache()))
	}); err != nil {
		return err
	}

	if err := term.BindAction("assetclear", "clear asset manager cache", func() {
		archiveManager.ClearCache()
		archivedFileManager.GetCache().Clear()
		paletteManager.ClearCache()
		am.paletteTransformManager.cache.Clear()
		animationManager.ClearCache()
		fontManager.ClearCache()
	}); err != nil {
		return err
	}

	return nil
}

// Bind a subordinate asset manager
func (am *assetManager) Bind(manager d2interface.AssetManagerSubordinate) error {
	switch manager.(type) {
	case d2interface.ArchiveManager:
		if err := manager.Bind(am); err != nil {
			return nil
		}

		am.archiveManager = manager.(d2interface.ArchiveManager)

	case d2interface.ArchivedFileManager:
		if err := manager.Bind(am); err != nil {
			return nil
		}

		am.archivedFileManager = manager.(d2interface.ArchivedFileManager)

	case d2interface.ArchivedPaletteManager:
		if err := manager.Bind(am); err != nil {
			return nil
		}

		am.paletteManager = manager.(d2interface.ArchivedPaletteManager)

	case *paletteTransformManager:
		if err := manager.Bind(am); err == nil {
			am.paletteTransformManager = manager.(*paletteTransformManager)
		}
	case d2interface.ArchivedAnimationManager:
		if err := manager.Bind(am); err != nil {
			return err
		}

		am.animationManager = manager.(d2interface.ArchivedAnimationManager)

	case d2interface.ArchivedFontManager:
		if err := manager.Bind(am); err != nil {
			return err
		}

		am.fontManager = manager.(d2interface.ArchivedFontManager)

	}

	return nil
}

// LoadFileStream streams an MPQ file from a source file path
func (am *assetManager) LoadFileStream(filePath string) (d2interface.ArchiveDataStream, error) {
	data, err := singleton.archivedFileManager.LoadFileStream(filePath)
	if err != nil {
		log.Printf("error loading file stream %s (%v)", filePath, err.Error())
	}

	return data, err
}

// LoadFile loads an entire file from a source file path as a []byte
func (am *assetManager) LoadFile(filePath string) ([]byte, error) {
	data, err := singleton.archivedFileManager.LoadFile(filePath)
	if err != nil {
		log.Printf("error loading file %s (%v)", filePath, err.Error())
	}

	return data, err
}

// FileExists checks if a file exists on the underlying file system at the given file path.
func (am *assetManager) FileExists(filePath string) (bool, error) {
	return singleton.archivedFileManager.FileExists(filePath)
}

// LoadAnimation loads an animation by its resource path and its palette path
func (am *assetManager) LoadAnimation(animationPath, palettePath string) (d2interface.Animation, error) {
	return am.LoadAnimationWithTransparency(animationPath, palettePath, 255)
}

// LoadAnimationWithTransparency loads an animation by its resource path and its palette path with a given transparency value
func (am *assetManager) LoadAnimationWithTransparency(animationPath, palettePath string,
	transparency int) (d2interface.Animation, error) {
	return singleton.animationManager.LoadAnimation(animationPath, palettePath, transparency)
}

// LoadComposite creates a composite object from a ObjectLookupRecord and palettePath describing it
func (am *assetManager) LoadComposite(baseType d2enum.ObjectType, token,
	palettePath string) (d2interface.CompositeAnimation, error) {
	return CreateComposite(baseType, token, palettePath), nil
}

// LoadFont loads a font the resource files
func (am *assetManager) LoadFont(tablePath, spritePath, palettePath string) (d2interface.Font, error) {
	return singleton.fontManager.LoadFont(tablePath, spritePath, palettePath)
}

// LoadPalette loads a palette from a given palette path
func (am *assetManager) LoadPalette(palettePath string) (d2interface.Palette, error) {
	return singleton.paletteManager.LoadPalette(palettePath)
}

func (am *assetManager) loadDC6(dc6Path string) (*d2dc6.DC6, error) {
	dc6Data, err := am.LoadFile(dc6Path)
	if err != nil {
		return nil, err
	}

	dc6, err := d2dc6.Load(dc6Data)
	if err != nil {
		return nil, err
	}

	return dc6, nil
}

func (am *assetManager) loadDCC(dccPath string) (*d2dcc.DCC, error) {
	dccData, err := am.LoadFile(dccPath)
	if err != nil {
		return nil, err
	}

	return d2dcc.Load(dccData)
}

func (am *assetManager) loadCOF(cofPath string) (*d2cof.COF, error) {
	cofData, err := am.LoadFile(cofPath)
	if err != nil {
		return nil, err
	}

	return d2cof.Load(cofData)
}
