package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
)

// Singleton is the single asst manager instance
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
	return Singleton.LoadFileStream(filePath)
}

// LoadFile loads an entire file from a source file path as a []byte
func LoadFile(filePath string) ([]byte, error) {
	return Singleton.LoadFile(filePath)
}

// FileExists checks if a file exists on the underlying file system at the given file path.
func FileExists(filePath string) (bool, error) {
	return Singleton.FileExists(filePath)
}

// LoadAnimation loads an animation by its resource path and its palette path
func LoadAnimation(animationPath, palettePath string) (d2interface.Animation, error) {
	return Singleton.LoadAnimation(animationPath, palettePath)
}

// LoadAnimationWithEffect loads an animation by its resource path and its palette path with a given transparency value
func LoadAnimationWithEffect(animationPath, palettePath string,
	drawEffect d2enum.DrawEffect) (d2interface.Animation, error) {
	return Singleton.LoadAnimationWithEffect(animationPath, palettePath, drawEffect)
}

// LoadComposite creates a composite object from a ObjectLookupRecord and palettePath describing it
func LoadComposite(baseType d2enum.ObjectType, token, palettePath string) (*Composite, error) {
	return Singleton.LoadComposite(baseType, token, palettePath)
}

// LoadFont loads a font the resource files
func LoadFont(tablePath, spritePath, palettePath string) (d2interface.Font, error) {
	return Singleton.LoadFont(tablePath, spritePath, palettePath)
}

// LoadPalette loads a palette from a given palette path
func LoadPalette(palettePath string) (d2interface.Palette, error) {
	return Singleton.LoadPalette(palettePath)
}
