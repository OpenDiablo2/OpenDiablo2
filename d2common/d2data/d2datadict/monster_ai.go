package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// The monai.txt file is a lookup table for unit AI codes
type MonsterAIRecord struct {
	AI string

	AIPackage1 string
	AIPackage2 string
	AIPackage3 string
	AIPackage4 string
	AIPackage5 string
	AIPackage6 string
	AIPackage7 string
	AIPackage8 string
}

var MonsterAI map[string]*MonsterAIRecord

func LoadMonsterAI(file []byte) {
	MonsterAI = make(map[string]*MonsterAIRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &MonsterAIRecord{
			AI:         d.String("AI"),
			AIPackage1: d.String("*aip1"),
			AIPackage2: d.String("*aip2"),
			AIPackage3: d.String("*aip3"),
			AIPackage4: d.String("*aip4"),
			AIPackage5: d.String("*aip5"),
			AIPackage6: d.String("*aip6"),
			AIPackage7: d.String("*aip7"),
			AIPackage8: d.String("*aip8"),
		}
		MonsterAI[record.AI] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonsterAI records", len(MonsterAI))

}
