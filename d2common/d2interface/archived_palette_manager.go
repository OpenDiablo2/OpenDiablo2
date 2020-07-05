package d2interface

type ArchivedPaletteManager interface {
	Cacher
	AssetManagerSubordinate
	LoadPalette(palettePath string) (Palette, error)
}
