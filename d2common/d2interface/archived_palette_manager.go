package d2interface

type ArchivedPaletteManager interface {
	Cacher
	LoadPalette(palettePath string) (Palette, error)
}
