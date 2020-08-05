package d2interface

// PaletteManager is responsible for loading palettes
type PaletteManager interface {
	Cacher
	LoadPalette(palettePath string) (Palette, error)
}
