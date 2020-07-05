package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"log"
)

var singleton *assetManager

// Initialize creates and assigns all necessary dependencies for the assetManager top-level functions to work correctly
func Initialize(renderer d2interface.Renderer, term d2interface.Terminal) (d2interface.
	AssetManager, error) {
	var (
		config                  = d2config.Get()
		archiveManager          = createArchiveManager(config)
		archivedFileManager     = createFileManager(config, archiveManager)
		paletteManager          = createPaletteManager()
		paletteTransformManager = createPaletteTransformManager()
		animationManager        = createAnimationManager(renderer)
		fontManager             = createFontManager()
	)

	singleton = &assetManager{
		paletteTransformManager: paletteTransformManager,
	}
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
		paletteTransformManager.cache.SetVerbose(verbose)
		animationManager.cache.SetVerbose(verbose)
	}); err != nil {
		return nil, err
	}

	if err := term.BindAction("assetstat", "display asset manager cache statistics", func() {
		var cacheStatistics = func(c d2interface.Cache) float64 {
			const percent = 100.0
			return float64(c.GetWeight()) / float64(c.GetBudget()) * percent
		}

		term.OutputInfof("archive cache: %f", cacheStatistics(archiveManager.GetCache()))
		term.OutputInfof("file cache: %f", cacheStatistics(archivedFileManager.GetCache()))
		term.OutputInfof("palette cache: %f", cacheStatistics(paletteManager.GetCache()))
		term.OutputInfof("palette transform cache: %f", cacheStatistics(paletteTransformManager.cache))
		term.OutputInfof("animation cache: %f", cacheStatistics(animationManager.cache))
		term.OutputInfof("font cache: %f", cacheStatistics(fontManager.GetCache()))
	}); err != nil {
		return nil, err
	}

	if err := term.BindAction("assetclear", "clear asset manager cache", func() {
		archiveManager.ClearCache()
		archivedFileManager.GetCache().Clear()
		paletteManager.ClearCache()
		paletteTransformManager.cache.Clear()
		animationManager.ClearCache()
		fontManager.ClearCache()
	}); err != nil {
		return nil, err
	}

	return singleton, nil
}

// TODO remove all of these funcs that use the singleton
// they are only here to prop up the other packages that
// dont have a reference to the asset manager yet

// LoadComposite creates a composite object from a ObjectLookupRecord and palettePath describing it
func LoadComposite(baseType d2enum.ObjectType, token, palettePath string) (d2interface.CompositeAnimation, error) {
	return CreateComposite(baseType, token, palettePath), nil
}

// LoadFile loads an entire file from a source file path as a []byte
func LoadFile(filePath string) ([]byte, error) {
	data, err := singleton.archivedFileManager.LoadFile(filePath)
	if err != nil {
		log.Printf("error loading file %s (%v)", filePath, err.Error())
	}

	return data, err
}

// FileExists checks if a file exists on the underlying file system at the given file path.
func FileExists(filePath string) (bool, error) {
	return singleton.archivedFileManager.FileExists(filePath)
}

// LoadAnimation loads an animation by its resource path and its palette path
func LoadAnimation(animationPath, palettePath string) (d2interface.Animation, error) {
	return singleton.LoadAnimationWithTransparency(animationPath, palettePath, 255)
}

// LoadAnimationWithTransparency loads an animation by its resource path and its palette path with a given transparency value
func LoadAnimationWithTransparency(animationPath, palettePath string,
	transparency int) (d2interface.Animation, error) {
	return singleton.animationManager.LoadAnimation(animationPath, palettePath, transparency)
}

// LoadFont loads a font the resource files
func LoadFont(tablePath, spritePath, palettePath string) (d2interface.Font, error) {
	return singleton.fontManager.LoadFont(tablePath, spritePath, palettePath)
}

// LoadPalette loads a palette from a given palette path
func LoadPalette(palettePath string) (d2interface.Palette, error) {
	return singleton.paletteManager.LoadPalette(palettePath)
}
