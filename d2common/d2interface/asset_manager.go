package d2interface

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

type AssetManager interface {
	AppComponent
	LoadFileStream(filePath string) (ArchiveDataStream, error)
	LoadFile(filePath string) ([]byte, error)
	FileExists(filePath string) (bool, error)
	LoadAnimation(animationPath, palettePath string) (Animation, error)
	LoadAnimationWithEffect(animationPath, palettePath string, effect d2enum.DrawEffect) (Animation, error)
	LoadComposite(baseType d2enum.ObjectType, token, palettePath string) (CompositeAnimation, error)
	LoadFont(tablePath, spritePath, palettePath string) (Font, error)
	LoadPalette(palettePath string) (Palette, error)
	Bind(manager AssetManagerSubordinate) error
}

type AssetManagerSubordinate interface {
	Bind(AssetManager) error
}
