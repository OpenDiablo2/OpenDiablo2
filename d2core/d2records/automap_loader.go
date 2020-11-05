package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func autoMapLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(AutoMaps, 0)

	var frameFields = []string{"Cel1", "Cel2", "Cel3", "Cel4"}

	for d.Next() {
		record := &AutoMapRecord{
			LevelName: d.String("LevelName"),
			TileName:  d.String("TileName"),

			Style:         d.Number("Style"),
			StartSequence: d.Number("StartSequence"),
			EndSequence:   d.Number("EndSequence"),

			// Note: aren't useful see the AutoMapRecord struct.
			//Type1: d.String("Type1"),
			//Type2: d.String("Type2"),
			//Type3: d.String("Type3"),
			//Type4: d.String("Type4"),
		}
		record.Frames = make([]int, len(frameFields))

		for i := range frameFields {
			record.Frames[i] = d.Number(frameFields[i])
		}

		records = append(records, record)
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d AutoMapRecord records", len(records))

	r.Level.AutoMaps = records

	return nil
}
