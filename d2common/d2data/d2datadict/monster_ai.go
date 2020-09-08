package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// MonsterAIRecord represents a single row from monai.txt
type MonsterAIRecord struct {
	AI string
}

// MonsterAI holds the MonsterAIRecords, The monai.txt file is a lookup table for unit AI codes
var MonsterAI map[string]*MonsterAIRecord //nolint:gochecknoglobals // Currently global by design

// LoadMonsterAI loads MonsterAIRecords from monai.txt
func LoadMonsterAI(file []byte) {
	MonsterAI = make(map[string]*MonsterAIRecord)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		record := &MonsterAIRecord{
			AI: d.String("AI"),
		}
		MonsterAI[record.AI] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonsterAI records", len(MonsterAI))
}
