package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

//MonsterLevelRecord represents a single row in monlvl.txt
type MonsterLevelRecord struct {

	//The level
	Level int

	//Values for Battle.net
	BattleNet monsterLevelValues

	//Values for ladder/single player/lan
	Ladder monsterLevelValues
}

type monsterLevelValues struct {
	//Armor class
	//An AC is calcuated as (MonLvl.txt Ac * Monstats.txt AC) / 100)
	ArmorClass          int
	ArmorClassNightmare int
	ArmorClassHell      int

	//To hit, influences ToHit values for both attacks
	// (MonLvl.txt TH * Monstats.txt A1TH
	// and MonLvl.txt TH * Monstats.txt A2TH) / 100
	ToHit          int
	ToHitNightmare int
	ToHitHell      int

	//Hitpoints, influences both minimum and maximum HP
	//(MonLvl.txt HP * Monstats.txt minHP) / 100
	// and MonLvl.txt HP * Monstats.txt maxHP) / 100
	Hitpoints          int
	HitpointsNightmare int
	HitpointsHell      int

	//Damage, influences minimum and maximum damage for both attacks
	//MonLvl.txt DM * Monstats.txt A1MinD) / 100
	// and MonLvl.txt DM * Monstats.txt A1MaxD) / 100
	// and MonLvl.txt DM * Monstats.txt A2MinD) / 100
	// and MonLvl.txt DM * Monstats.txt A2MaxD) / 100
	Damage          int
	DamageNightmare int
	DamageHell      int

	//Experience points
	//The formula is (MonLvl.txt XP * Monstats.txt Exp) / 100
	Experience          int
	ExperienceNightmare int
	ExperienceHell      int
}

//MonsterLevels stores the MonsterLevelRecords
var MonsterLevels map[int]*MonsterLevelRecord

//LoadMonsterLevels loads LoadMonsterLevelRecords into MonsterLevels
func LoadMonsterLevels(file []byte) {
	MonsterLevels = make(map[int]*MonsterLevelRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &MonsterLevelRecord{
			Level: d.Number("Level"),
			BattleNet: monsterLevelValues{
				ArmorClass:          d.Number("AC"),
				ArmorClassNightmare: d.Number("AC(N)"),
				ArmorClassHell:      d.Number("AC(H)"),
				ToHit:               d.Number("TH"),
				ToHitNightmare:      d.Number("TH (N)"),
				ToHitHell:           d.Number("TH (H)"),
				Hitpoints:           d.Number("HP"),
				HitpointsNightmare:  d.Number("HP (N)"),
				HitpointsHell:       d.Number("HP (H)"),
				Damage:              d.Number("DM"),
				DamageNightmare:     d.Number("DM (N)"),
				DamageHell:          d.Number("DM (H)"),
				Experience:          d.Number("XP"),
				ExperienceNightmare: d.Number("XP (N)"),
				ExperienceHell:      d.Number("XP (H)"),
			},
			Ladder: monsterLevelValues{
				ArmorClass:          d.Number("L-AC"),
				ArmorClassNightmare: d.Number("L-AC(N)"),
				ArmorClassHell:      d.Number("L-AC(H)"),
				ToHit:               d.Number("L-TH"),
				ToHitNightmare:      d.Number("L-TH (N)"),
				ToHitHell:           d.Number("L-TH (H)"),
				Hitpoints:           d.Number("L-HP"),
				HitpointsNightmare:  d.Number("L-HP (N)"),
				HitpointsHell:       d.Number("L-HP (H)"),
				Damage:              d.Number("L-DM"),
				DamageNightmare:     d.Number("L-DM (N)"),
				DamageHell:          d.Number("L-DM (H)"),
				Experience:          d.Number("L-XP"),
				ExperienceNightmare: d.Number("L-XP (N)"),
				ExperienceHell:      d.Number("L-XP (H)"),
			},
		}
		MonsterLevels[record.Level] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonsterLevel records", len(MonsterLevels))
}
