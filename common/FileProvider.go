package common

import "github.com/OpenDiablo2/OpenDiablo2/palettedefs"

// FileProvider is an instance that can provide different types of files
type FileProvider interface {
	LoadFile(fileName string) []byte
	LoadSprite(fileName string, palette palettedefs.PaletteType) *Sprite
}
