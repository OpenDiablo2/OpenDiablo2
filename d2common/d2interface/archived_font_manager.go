package d2interface

// ArchivedFontManager manages fonts that are in archives being
// managed by the ArchiveManager
type ArchivedFontManager interface {
	Cacher
	AssetManagerSubordinate
	LoadFont(tablePath, spritePath, palettePath string) (Font, error)
}
