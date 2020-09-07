package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// MonsterPlacementRecord represents a line from MonPlace.txt.
type MonsterPlacementRecord string

// MonsterPlacements stores the MonsterPlacementRecords.
var MonsterPlacements []MonsterPlacementRecord //nolint:gochecknoglobals // Currently global by design

// LoadMonsterPlacements loads the MonsterPlacementRecords into MonsterPlacements.
func LoadMonsterPlacements(file []byte) {
	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		MonsterPlacements = append(MonsterPlacements, MonsterPlacementRecord(d.String("code")))
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonsterPlacement records", len(MonsterPlacements))
}
