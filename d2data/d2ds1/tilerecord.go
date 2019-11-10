package d2ds1

type TileRecord struct {
	Floors        []FloorShadowRecord
	Walls         []WallRecord
	Shadows       []FloorShadowRecord
	Substitutions []SubstitutionRecord
}
