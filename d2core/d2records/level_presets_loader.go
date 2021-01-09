package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadLevelPresets loads level presets from text file
func levelPresetLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(LevelPresets)

	for d.Next() {
		record := LevelPresetRecord{
			Name:         d.String("Name"),
			DefinitionID: d.Number("Def"),
			LevelID:      d.Number("LevelId"),
			Populate:     d.Number("Populate") == 1,
			Logicals:     d.Number("Logicals") == 1,
			Outdoors:     d.Number("Outdoors") == 1,
			Animate:      d.Number("Animate") == 1,
			KillEdge:     d.Number("KillEdge") == 1,
			FillBlanks:   d.Number("FillBlanks") == 1,
			SizeX:        d.Number("SizeX"),
			SizeY:        d.Number("SizeY"),
			AutoMap:      d.Number("AutoMap") == 1,
			Scan:         d.Number("Scan") == 1,
			Pops:         d.Number("Pops"),
			PopPad:       d.Number("PopPad"),
			FileCount:    d.Number("Files"),
			Files: [6]string{
				d.String("File1"),
				d.String("File2"),
				d.String("File3"),
				d.String("File4"),
				d.String("File5"),
				d.String("File6"),
			},
			Dt1Mask:   uint(d.Number("Dt1Mask")),
			Beta:      d.Number("Beta") == 1,
			Expansion: d.Number("Expansion") == 1,
		}

		records[record.DefinitionID] = record
	}

	r.Logger.Infof("Loaded %d level presets", len(records))

	if d.Err != nil {
		return d.Err
	}

	r.Level.Presets = records

	return nil
}
