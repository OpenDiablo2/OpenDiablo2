package d2interface

// FileProvider is an instance that can provide different types of files
type FileProvider interface {
	LoadFile(fileName string) []byte
	//LoadSprite(fileName string, palette enums.PaletteType) *d2render.Sprite
}
