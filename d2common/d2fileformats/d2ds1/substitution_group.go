package d2ds1

// SubstitutionGroup represents a substitution group in a DS1 file.
type SubstitutionGroup struct {
	TileX         int32
	TileY         int32
	WidthInTiles  int32
	HeightInTiles int32
	Unknown       int32
}
