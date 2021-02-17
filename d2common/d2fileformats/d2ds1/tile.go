package d2ds1

// Tile represents a tile record in a DS1 file.
type Tile struct {
	Floors        []Floor        // Collection of floor records
	Walls         []Wall         // Collection of wall records
	Shadows       []Shadow       // Collection of shadow records
	Substitutions []Substitution // Collection of substitutions
}

func makeDefaultTile() Tile {
	return Tile{
		Floors:        []Floor{{}},
		Walls:         []Wall{{}},
		Shadows:       []Shadow{{}},
		Substitutions: []Substitution{{}},
	}
}
