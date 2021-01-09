package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func levelMazeDetailsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(LevelMazeDetails)

	for d.Next() {
		record := &LevelMazeDetailsRecord{
			Name:              d.String("Name"),
			LevelID:           d.Number("Level"),
			NumRoomsNormal:    d.Number("Rooms"),
			NumRoomsNightmare: d.Number("Rooms(N)"),
			NumRoomsHell:      d.Number("Rooms(H)"),
			SizeX:             d.Number("SizeX"),
			SizeY:             d.Number("SizeY"),
		}
		records[record.LevelID] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d LevelMazeDetails records", len(records))

	r.Level.Maze = records

	return nil
}
