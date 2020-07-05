package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dat"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var singleton *assetManager

// Initialize creates and assigns all necessary dependencies for the assetManager top-level functions to work correctly
func Initialize(renderer d2interface.Renderer,
	term d2interface.Terminal) error {
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
		archiveManager,
		archivedFileManager,
		paletteManager,
		paletteTransformManager,
		animationManager,
		fontManager,
	}

	if err := term.BindAction("assetspam", "display verbose asset manager logs", func(verbose bool) {
		if verbose {
			term.OutputInfof("asset manager verbose logging enabled")
		} else {
			term.OutputInfof("asset manager verbose logging disabled")
		}

		archiveManager.GetCache().SetVerbose(verbose)
		archivedFileManager.GetCache().SetVerbose(verbose)
		paletteManager.cache.SetVerbose(verbose)
		paletteTransformManager.cache.SetVerbose(verbose)
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
		term.OutputInfof("palette cache: %f", cacheStatistics(paletteManager.cache))
		term.OutputInfof("palette transform cache: %f", cacheStatistics(paletteTransformManager.cache))
		term.OutputInfof("animation cache: %f", cacheStatistics(animationManager.cache))
		term.OutputInfof("font cache: %f", cacheStatistics(fontManager.cache))
	}); err != nil {
		return err
	}

	if err := term.BindAction("assetclear", "clear asset manager cache", func() {
		archiveManager.ClearCache()
		archivedFileManager.GetCache().Clear()
		paletteManager.cache.Clear()
		paletteTransformManager.cache.Clear()
		animationManager.cache.Clear()
		fontManager.cache.Clear()
	}); err != nil {
		return err
	}

	return nil
}

// LoadFileStream streams an MPQ file from a source file path
func LoadFileStream(filePath string) (d2interface.ArchiveDataStream, error) {
	data, err := singleton.archivedFileManager.LoadFileStream(filePath)
	if err != nil {
		log.Printf("error loading file stream %s (%v)", filePath, err.Error())
	}

	return data, err
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
func LoadAnimation(animationPath, palettePath string) (*Animation, error) {
	return LoadAnimationWithTransparency(animationPath, palettePath, 255)
}

// LoadAnimationWithTransparency loads an animation by its resource path and its palette path with a given transparency value
func LoadAnimationWithTransparency(animationPath, palettePath string, transparency int) (*Animation, error) {
	return singleton.animationManager.loadAnimation(animationPath, palettePath, transparency)
}

// LoadComposite creates a composite object from a ObjectLookupRecord and palettePath describing it
func LoadComposite(baseType d2enum.ObjectType, token, palettePath string) (*Composite, error) {
	return CreateComposite(baseType, token, palettePath), nil
}

// LoadFont loads a font the resource files
func LoadFont(tablePath, spritePath, palettePath string) (*Font, error) {
	return singleton.fontManager.loadFont(tablePath, spritePath, palettePath)
}

// LoadPalette loads a palette from a given palette path
func LoadPalette(palettePath string) (*d2dat.DATPalette, error) {
	return singleton.paletteManager.loadPalette(palettePath)
}
