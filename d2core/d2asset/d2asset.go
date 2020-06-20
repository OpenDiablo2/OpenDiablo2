/*
Package d2asset has behaviors to load and save assets from disk.
*/
package d2asset

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dat"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
)

var singleton *assetManager

// Initialize creates and assigns all necessary dependencies for the assetManager top-level functions to work correctly
func Initialize(term d2interface.Terminal) error {
	var (
		config                  = d2config.Get()
		archiveManager          = createArchiveManager(&config)
		fileManager             = createFileManager(&config, archiveManager)
		paletteManager          = createPaletteManager()
		paletteTransformManager = createPaletteTransformManager()
		animationManager        = createAnimationManager()
		fontManager             = createFontManager()
	)

	singleton = &assetManager{
		archiveManager,
		fileManager,
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

		archiveManager.cache.SetVerbose(verbose)
		fileManager.cache.SetVerbose(verbose)
		paletteManager.cache.SetVerbose(verbose)
		paletteTransformManager.cache.SetVerbose(verbose)
		animationManager.cache.SetVerbose(verbose)
	}); err != nil {
		return err
	}

	if err := term.BindAction("assetstat", "display asset manager cache statistics", func() {
		type cache interface {
			GetWeight() int
			GetBudget() int
		}

		var cacheStatistics = func(c cache) float64 {
			const percent = 100.0
			return float64(c.GetWeight()) / float64(c.GetBudget()) * percent
		}

		term.OutputInfof("archive cache: %f", cacheStatistics(archiveManager.cache))
		term.OutputInfof("file cache: %f", cacheStatistics(fileManager.cache))
		term.OutputInfof("palette cache: %f", cacheStatistics(paletteManager.cache))
		term.OutputInfof("palette transform cache: %f", cacheStatistics(paletteTransformManager.cache))
		term.OutputInfof("animation cache: %f", cacheStatistics(animationManager.cache))
		term.OutputInfof("font cache: %f", cacheStatistics(fontManager.cache))
	}); err != nil {
		return err
	}

	if err := term.BindAction("assetclear", "clear asset manager cache", func() {
		archiveManager.cache.Clear()
		fileManager.cache.Clear()
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
func LoadFileStream(filePath string) (*d2mpq.MpqDataStream, error) {
	data, err := singleton.fileManager.loadFileStream(filePath)
	if err != nil {
		log.Printf("error loading file stream %s (%v)", filePath, err.Error())
	}

	return data, err
}

// LoadFile loads an entire file from a source file path as a []byte
func LoadFile(filePath string) ([]byte, error) {
	data, err := singleton.fileManager.loadFile(filePath)
	if err != nil {
		log.Printf("error loading file %s (%v)", filePath, err.Error())
	}

	return data, err
}

// FileExists checks if a file exists on the underlying file system at the given file path.
func FileExists(filePath string) (bool, error) {
	return singleton.fileManager.fileExists(filePath)
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
func LoadComposite(object *d2datadict.ObjectLookupRecord, palettePath string) (*Composite, error) {
	return CreateComposite(object, palettePath), nil
}

// LoadFont loads a font the resource files
func LoadFont(tablePath, spritePath, palettePath string) (*Font, error) {
	return singleton.fontManager.loadFont(tablePath, spritePath, palettePath)
}

// LoadPalette loads a palette from a given palette path
func LoadPalette(palettePath string) (*d2dat.DATPalette, error) {
	return singleton.paletteManager.loadPalette(palettePath)
}
