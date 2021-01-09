package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadMonsterPlacements loads the MonsterPlacementRecords into MonsterPlacements.
func monsterPlacementsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(MonsterPlacements, 0)

	for d.Next() {
		records = append(records, MonsterPlacementRecord(d.String("code")))
	}

	if d.Err != nil {
		return d.Err
	}

	r.Monster.Placements = records

	r.Logger.Infof("Loaded %d MonsterPlacement records", len(records))

	return nil
}
