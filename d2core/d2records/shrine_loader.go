package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func shrineLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(map[string]*ShrineRecord)

	for d.Next() {
		record := &ShrineRecord{
			ShrineType:       d.String("Shrine Type"),
			ShrineName:       d.String("Shrine name"),
			Effect:           d.String("Effect"),
			Code:             d.Number("Code"),
			Arg0:             d.Number("Arg0"),
			Arg1:             d.Number("Arg1"),
			DurationFrames:   d.Number("Duration in frames"),
			ResetTimeMinutes: d.Number("reset time in minutes"),
			Rarity:           d.Number("rarity"),
			EffectClass:      d.Number("effectclass"),
			LevelMin:         d.Number("LevelMin"),
		}

		records[record.ShrineName] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Object.Shrines = records

	r.Logger.Infof("Loaded %d shrines", len(records))

	return nil
}
