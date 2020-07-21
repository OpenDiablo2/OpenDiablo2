package d2ds1

// TileRecord represents a tile record in a DS1 file.
type TileRecord struct {
	Floors        []FloorShadowRecord  // Collection of floor records
	Walls         []WallRecord         // Collection of wall records
	Shadows       []FloorShadowRecord  // Collection of shadow records
	Substitutions []SubstitutionRecord // Collection of substitutions
}
