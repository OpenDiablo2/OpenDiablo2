package d2asset

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
)

var Singleton *AssetManager //nolint:gochecknoglobals // Currently global by design

// Initialize creates and assigns all necessary dependencies for the AssetManager top-level functions to work correctly
func Initialize(renderer d2interface.Renderer,
	term d2interface.Terminal) error {
	var (
		archiveManager          = createArchiveManager(d2config.Config)
		archivedFileManager     = createFileManager(d2config.Config, archiveManager)
		paletteManager          = createPaletteManager()
		paletteTransformManager = createPaletteTransformManager()
		animationManager        = createAnimationManager(renderer)
		fontManager             = createFontManager()
	)

	Singleton = &AssetManager{
		archiveManager,
		archivedFileManager,
		paletteManager,
		paletteTransformManager,
		animationManager,
		fontManager,
	}

	if term != nil {
		return Singleton.BindTerminalCommands(term)
	}

	return nil
}

// LoadFileStream streams an MPQ file from a source file path
func LoadFileStream(filePath string) (d2interface.ArchiveDataStream, error) {
	data, err := Singleton.archivedFileManager.LoadFileStream(filePath)
	if err != nil {
		log.Printf("error loading file stream %s (%v)", filePath, err.Error())
	}

	return data, err
}

// LoadFile loads an entire file from a source file path as a []byte
func LoadFile(filePath string) ([]byte, error) {
	data, err := Singleton.archivedFileManager.LoadFile(filePath)
	if err != nil {
		log.Printf("error loading file %s (%v)", filePath, err.Error())
	}

	return data, err
}

// FileExists checks if a file exists on the underlying file system at the given file path.
func FileExists(filePath string) (bool, error) {
	return Singleton.archivedFileManager.FileExists(filePath)
}

// LoadAnimation loads an animation by its resource path and its palette path
func LoadAnimation(animationPath, palettePath string) (d2interface.Animation, error) {
	return LoadAnimationWithEffect(animationPath, palettePath, d2enum.DrawEffectNone)
}

// LoadAnimationWithEffect loads an animation by its resource path and its palette path with a given transparency value
func LoadAnimationWithEffect(animationPath, palettePath string,
	drawEffect d2enum.DrawEffect) (d2interface.Animation, error) {
	return Singleton.animationManager.LoadAnimation(animationPath, palettePath, drawEffect)
}

// LoadComposite creates a composite object from a ObjectLookupRecord and palettePath describing it
func LoadComposite(baseType d2enum.ObjectType, token, palettePath string) (*Composite, error) {
	return CreateComposite(baseType, token, palettePath), nil
}

// LoadFont loads a font the resource files
func LoadFont(tablePath, spritePath, palettePath string) (d2interface.Font, error) {
	return Singleton.fontManager.LoadFont(tablePath, spritePath, palettePath)
}

// LoadPalette loads a palette from a given palette path
func LoadPalette(palettePath string) (d2interface.Palette, error) {
	return Singleton.paletteManager.LoadPalette(palettePath)
}
