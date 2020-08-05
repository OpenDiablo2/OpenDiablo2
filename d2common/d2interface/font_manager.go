package d2interface

// FontManager manages fonts that are in archives being
// managed by the ArchiveManager
type FontManager interface {
	Cacher
	LoadFont(tablePath, spritePath, palettePath string) (Font, error)
}
