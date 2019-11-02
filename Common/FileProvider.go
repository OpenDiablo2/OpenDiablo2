package Common

import "github.com/OpenDiablo2/OpenDiablo2/PaletteDefs"

// FileProvider is an instance that can provide different types of files
type FileProvider interface {
	LoadFile(fileName string) []byte
	LoadSprite(fileName string, palette PaletteDefs.PaletteType) *Sprite
}
