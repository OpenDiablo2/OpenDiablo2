package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadMonsterAI loads MonsterAIRecords from monai.txt
func monsterAiLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(MonsterAI)

	for d.Next() {
		record := &MonsterAIRecord{
			AI: d.String("AI"),
		}
		records[record.AI] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d MonsterAI records", len(records))

	r.Monster.AI = records

	return nil
}
