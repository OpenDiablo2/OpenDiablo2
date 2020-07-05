package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"log"
)

var singleton *assetManager

// Create creates and assigns all necessary dependencies for the assetManager top-level functions to work correctly
func Create() (d2interface.AssetManager, error) {
	singleton = &assetManager{
		paletteTransformManager: createPaletteTransformManager(),
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
