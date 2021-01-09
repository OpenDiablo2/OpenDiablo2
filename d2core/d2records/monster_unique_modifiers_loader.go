package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func monsterUniqModifiersLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(MonsterUniqueModifiers)
	constants := make([]int, 0)

	for d.Next() {
		record := &MonUModRecord{
			Name:          d.String("uniquemod"),
			ID:            d.Number("id"),
			Enabled:       d.Bool("enabled"),
			ExpansionOnly: d.Number("version") == expansionCode,
			Xfer:          d.Bool("xfer"),
			Champion:      d.Bool("champion"),
			FPick:         d.Number("fpick"),
			Exclude1:      d.String("exclude1"),
			Exclude2:      d.String("exclude2"),
			PickFrequencies: struct {
				Normal    *PickFreq
				Nightmare *PickFreq
				Hell      *PickFreq
			}{
				Normal: &PickFreq{
					Champion: d.Number("cpick"),
					Unique:   d.Number("upick"),
				},
				Nightmare: &PickFreq{
					Champion: d.Number("cpick (N)"),
					Unique:   d.Number("upick (N)"),
				},
				Hell: &PickFreq{
					Champion: d.Number("cpick (H)"),
					Unique:   d.Number("upick (H)"),
				},
			},
		}

		records[record.Name] = record

		constants = append(constants, d.Number("constants"))
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d MonsterUniqueModifier records", len(records))

	r.Monster.Unique.Mods = records
	r.Monster.Unique.Constants = constants

	return nil
}
