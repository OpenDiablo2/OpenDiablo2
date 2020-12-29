package d2records

// MonsterLevels stores the MonsterLevelRecords
type MonsterLevels map[int]*MonsterLevelRecord

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
