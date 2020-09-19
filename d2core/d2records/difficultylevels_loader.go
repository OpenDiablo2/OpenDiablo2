package d2records

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func difficultyLevelsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(DifficultyLevels)

	for d.Next() {
		record := &DifficultyLevelRecord{
			Name:                   d.String("Name"),
			ResistancePenalty:      d.Number("ResistPenalty"),
			DeathExperiencePenalty: d.Number("DeathExpPenalty"),
			DropChanceLow:          d.Number("UberCodeOddsNormal"),
			DropChanceNormal:       d.Number("UberCodeOddsNormal"),
			DropChanceSuperior:     d.Number("UberCodeOddsNormal"),
			DropChanceExceptional:  d.Number("UberCodeOddsNormal"),
			DropChanceMagic:        d.Number("UberCodeOddsGood"),
			DropChanceRare:         d.Number("UberCodeOddsGood"),
			DropChanceSet:          d.Number("UberCodeOddsGood"),
			DropChanceUnique:       d.Number("UberCodeOddsGood"),
			MonsterSkillBonus:      d.Number("MonsterSkillBonus"),
			MonsterColdDivisor:     d.Number("MonsterColdDivisor"),
			MonsterFreezeDivisor:   d.Number("MonsterFreezeDivisor"),
			AiCurseDivisor:         d.Number("AiCurseDivisor"),
			LifeStealDivisor:       d.Number("LifeStealDivisor"),
			ManaStealDivisor:       d.Number("ManaStealDivisor"),
		}
		records[record.Name] = record
	}

	if d.Err != nil {
		return d.Err
	}

	log.Printf("Loaded %d DifficultyLevel records", len(records))

	r.DifficultyLevels = records

	return nil
}
