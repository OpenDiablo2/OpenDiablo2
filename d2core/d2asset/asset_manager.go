package d2asset

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2cof"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// AssetManager loads files and game objects
type AssetManager struct {
	archiveManager          d2interface.ArchiveManager
	archivedFileManager     d2interface.FileManager
	paletteManager          d2interface.PaletteManager
	paletteTransformManager *paletteTransformManager
	animationManager        d2interface.AnimationManager
	fontManager             d2interface.FontManager
}

// LoadFileStream streams an MPQ file from a source file path
func (am *AssetManager) LoadFileStream(filePath string) (d2interface.ArchiveDataStream, error) {
	data, err := am.archivedFileManager.LoadFileStream(filePath)
	if err != nil {
		log.Printf("error loading file stream %s (%v)", filePath, err.Error())
	}

	return data, err
}

// LoadFile loads an entire file from a source file path as a []byte
func (am *AssetManager) LoadFile(filePath string) ([]byte, error) {
	data, err := am.archivedFileManager.LoadFile(filePath)
	if err != nil {
		log.Printf("error loading file %s (%v)", filePath, err.Error())
	}

	return data, err
}

// FileExists checks if a file exists on the underlying file system at the given file path.
func (am *AssetManager) FileExists(filePath string) (bool, error) {
	return am.archivedFileManager.FileExists(filePath)
}

// LoadAnimation loads an animation by its resource path and its palette path
func (am *AssetManager) LoadAnimation(animationPath, palettePath string) (d2interface.Animation, error) {
	return am.LoadAnimationWithEffect(animationPath, palettePath, d2enum.DrawEffectNone)
}

// LoadAnimationWithEffect loads an animation by its resource path and its palette path with a given transparency value
func (am *AssetManager) LoadAnimationWithEffect(animationPath, palettePath string,
	drawEffect d2enum.DrawEffect) (d2interface.Animation, error) {
	return am.animationManager.LoadAnimation(animationPath, palettePath, drawEffect)
}

// LoadComposite creates a composite object from a ObjectLookupRecord and palettePath describing it
func (am *AssetManager) LoadComposite(baseType d2enum.ObjectType, token, palettePath string) (*Composite, error) {
	return CreateComposite(baseType, token, palettePath), nil
}

// LoadFont loads a font the resource files
func (am *AssetManager) LoadFont(tablePath, spritePath, palettePath string) (d2interface.Font, error) {
	return am.fontManager.LoadFont(tablePath, spritePath, palettePath)
}

// LoadPalette loads a palette from a given palette path
func (am *AssetManager) LoadPalette(palettePath string) (d2interface.Palette, error) {
	return am.paletteManager.LoadPalette(palettePath)
}

func loadDC6(dc6Path string) (*d2dc6.DC6, error) {
	dc6Data, err := Singleton.LoadFile(dc6Path)
	if err != nil {
		return nil, err
	}

	dc6, err := d2dc6.Load(dc6Data)
	if err != nil {
		return nil, err
	}

	return dc6, nil
}

func loadDCC(dccPath string) (*d2dcc.DCC, error) {
	dccData, err := Singleton.LoadFile(dccPath)
	if err != nil {
		return nil, err
	}

	return d2dcc.Load(dccData)
}

func loadCOF(cofPath string) (*d2cof.COF, error) {
	cofData, err := Singleton.LoadFile(cofPath)
	if err != nil {
		return nil, err
	}

	return d2cof.Load(cofData)
}

func (am *AssetManager) BindTerminalCommands(term d2interface.Terminal) error {
	if err := term.BindAction("assetspam", "display verbose asset manager logs", func(verbose bool) {
		if verbose {
			term.OutputInfof("asset manager verbose logging enabled")
		} else {
			term.OutputInfof("asset manager verbose logging disabled")
		}

		am.archiveManager.GetCache().SetVerbose(verbose)
		am.archivedFileManager.GetCache().SetVerbose(verbose)
		am.paletteManager.GetCache().SetVerbose(verbose)
		am.paletteTransformManager.cache.SetVerbose(verbose)
		am.animationManager.GetCache().SetVerbose(verbose)
	}); err != nil {
		return err
	}

	if err := term.BindAction("assetstat", "display asset manager cache statistics", func() {
		var cacheStatistics = func(c d2interface.Cache) float64 {
			const percent = 100.0
			return float64(c.GetWeight()) / float64(c.GetBudget()) * percent
		}

		term.OutputInfof("archive cache: %f", cacheStatistics(am.archiveManager.GetCache()))
		term.OutputInfof("file cache: %f", cacheStatistics(am.archivedFileManager.GetCache()))
		term.OutputInfof("palette cache: %f", cacheStatistics(am.paletteManager.GetCache()))
		term.OutputInfof("palette transform cache: %f", cacheStatistics(am.paletteTransformManager.
			cache))
		term.OutputInfof("animation cache: %f", cacheStatistics(am.animationManager.GetCache()))
		term.OutputInfof("font cache: %f", cacheStatistics(am.fontManager.GetCache()))
	}); err != nil {
		return err
	}

	if err := term.BindAction("assetclear", "clear asset manager cache", func() {
		am.archiveManager.ClearCache()
		am.archivedFileManager.GetCache().Clear()
		am.paletteManager.ClearCache()
		am.paletteTransformManager.cache.Clear()
		am.animationManager.ClearCache()
		am.fontManager.ClearCache()
	}); err != nil {
		return err
	}

	return nil
}
