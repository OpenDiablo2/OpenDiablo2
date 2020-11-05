package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func levelWarpsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(LevelWarps)

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
		records[record.ID] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d level warps", len(records))

	r.Level.Warp = records

	return nil
}
