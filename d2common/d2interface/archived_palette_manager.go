package d2interface

type ArchivedPaletteManager interface {
	LoadPalette(palettePath string) (Palette, error)
}
