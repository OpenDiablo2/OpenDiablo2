package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// MonsterLevelRecord represents a single row in monlvl.txt
type MonsterLevelRecord struct {

	// The level
	Level int

	// Values for Battle.net
	BattleNet monsterDifficultyLevels

	// Values for ladder/single player/lan
	Ladder monsterDifficultyLevels
}

type monsterDifficultyLevels struct {
	Normal    monsterLevelValues
	Nightmare monsterLevelValues
	Hell      monsterLevelValues
}

type monsterLevelValues struct {
	// DefenseRating AC is calcuated as (MonLvl.txt Ac * Monstats.txt AC) / 100)
	DefenseRating int // also known as armor class

	// ToHit influences ToHit values for both attacks
	// (MonLvl.txt TH * Monstats.txt A1TH
	// and MonLvl.txt TH * Monstats.txt A2TH) / 100
	AttackRating int

	// Hitpoints, influences both minimum and maximum HP
	// (MonLvl.txt HP * Monstats.txt minHP) / 100
	// and MonLvl.txt HP * Monstats.txt maxHP) / 100
	Hitpoints int

	// Damage, influences minimum and maximum damage for both attacks
	// MonLvl.txt DM * Monstats.txt A1MinD) / 100
	// and MonLvl.txt DM * Monstats.txt A1MaxD) / 100
	// and MonLvl.txt DM * Monstats.txt A2MinD) / 100
	// and MonLvl.txt DM * Monstats.txt A2MaxD) / 100
	Damage int

	// Experience points,
	// the formula is (MonLvl.txt XP * Monstats.txt Exp) / 100
	Experience int
}

// MonsterLevels stores the MonsterLevelRecords
var MonsterLevels map[int]*MonsterLevelRecord //nolint:gochecknoglobals // Currently global by design

// LoadMonsterLevels loads LoadMonsterLevelRecords into MonsterLevels
func LoadMonsterLevels(file []byte) {
	MonsterLevels = make(map[int]*MonsterLevelRecord)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		record := &MonsterLevelRecord{
			Level: d.Number("Level"),
			BattleNet: monsterDifficultyLevels{
				Normal: monsterLevelValues{
					Hitpoints:  d.Number("HP"),
					Damage:     d.Number("DM"),
					Experience: d.Number("XP"),
				},
				Nightmare: monsterLevelValues{
					Hitpoints:  d.Number("HP(N)"),
					Damage:     d.Number("DM(N)"),
					Experience: d.Number("XP(N)"),
				},
				Hell: monsterLevelValues{
					Hitpoints:  d.Number("HP(H)"),
					Damage:     d.Number("DM(H)"),
					Experience: d.Number("XP(H)"),
				},
			},
			Ladder: monsterDifficultyLevels{
				Normal: monsterLevelValues{
					Hitpoints:  d.Number("L-HP"),
					Damage:     d.Number("L-DM"),
					Experience: d.Number("L-XP"),
				},
				Nightmare: monsterLevelValues{
					Hitpoints:  d.Number("L-HP(N)"),
					Damage:     d.Number("L-DM(N)"),
					Experience: d.Number("L-XP(N)"),
				},
				Hell: monsterLevelValues{
					Hitpoints:  d.Number("L-HP(H)"),
					Damage:     d.Number("L-DM(H)"),
					Experience: d.Number("L-XP(H)"),
				},
			},
		}
		MonsterLevels[record.Level] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonsterLevel records", len(MonsterLevels))
}
