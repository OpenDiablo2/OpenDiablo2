package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// The monai.txt file is a lookup table for unit AI codes
type MonsterAIRecord struct {
	AI string
}

var MonsterAI map[string]*MonsterAIRecord

func LoadMonsterAI(file []byte) {
	MonsterAI = make(map[string]*MonsterAIRecord)

	d := d2common.LoadDataDictionary(file)
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
