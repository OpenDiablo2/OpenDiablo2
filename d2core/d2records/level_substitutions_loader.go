package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func levelSubstitutionsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(LevelSubstitutions)

	for d.Next() {
		record := &LevelSubstitutionRecord{
			Name:         d.String("Name"),
			ID:           d.Number("Type"),
			File:         d.String("File"),
			IsExpansion:  d.Number("Expansion") > 0,
			BorderType:   d.Number("BordType"),
			GridSize:     d.Number("GridSize"),
			Mask:         d.Number("Dt1Mask"),
			ChanceSpawn0: d.Number("Prob0"),
			ChanceSpawn1: d.Number("Prob1"),
			ChanceSpawn2: d.Number("Prob2"),
			ChanceSpawn3: d.Number("Prob3"),
			ChanceSpawn4: d.Number("Prob4"),
			ChanceFloor0: d.Number("Trials0"),
			ChanceFloor1: d.Number("Trials1"),
			ChanceFloor2: d.Number("Trials2"),
			ChanceFloor3: d.Number("Trials3"),
			ChanceFloor4: d.Number("Trials4"),
			GridMax0:     d.Number("Max0"),
			GridMax1:     d.Number("Max1"),
			GridMax2:     d.Number("Max2"),
			GridMax3:     d.Number("Max3"),
			GridMax4:     d.Number("Max4"),
		}

		records[record.ID] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d LevelSubstitution records", len(records))

	r.Level.Sub = records

	return nil
}
