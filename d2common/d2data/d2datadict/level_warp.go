package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// LevelWarpRecord is a representation of a row from lvlwarp.txt
// it describes the warp graphics offsets and dimensions for levels
type LevelWarpRecord struct {
	Name       string
	ID         int
	SelectX    int
	SelectY    int
	SelectDX   int
	SelectDY   int
	ExitWalkX  int
	ExitWalkY  int
	OffsetX    int
	OffsetY    int
	LitVersion bool
	Tiles      int
	Direction  string
}

// LevelWarps loaded from txt records
//nolint:gochecknoglobals // Currently global by design, only written once
var LevelWarps map[int]*LevelWarpRecord

// LoadLevelWarps loads LevelWarpRecord's from text file data
func LoadLevelWarps(file []byte) {
	LevelWarps = make(map[int]*LevelWarpRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &LevelWarpRecord{
			Name:       d.String("Name"),
			ID:         d.Number("Id"),
			SelectX:    d.Number("SelectX"),
			SelectY:    d.Number("SelectY"),
			SelectDX:   d.Number("SelectDX"),
			SelectDY:   d.Number("SelectDY"),
			ExitWalkX:  d.Number("ExitWalkX"),
			ExitWalkY:  d.Number("ExitWalkY"),
			OffsetX:    d.Number("OffsetX"),
			OffsetY:    d.Number("OffsetY"),
			LitVersion: d.Bool("LitVersion"),
			Tiles:      d.Number("Tiles"),
			Direction:  d.String("Direction"),
		}
		LevelWarps[record.ID] = record
	}

	log.Printf("Loaded %d level warps", len(LevelWarps))
}
