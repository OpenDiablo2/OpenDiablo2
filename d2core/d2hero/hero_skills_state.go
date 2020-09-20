package d2hero

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"

type HeroSkillsState map[int] *HeroSkill

// CreateHeroSkillsState will assemble the hero skills from the class stats record.
func CreateHeroSkillsState(classStats *d2datadict.CharStatsRecord) *HeroSkillsState {
	baseSkills := HeroSkillsState{}

	for idx := range classStats.BaseSkill {
		skillName := &classStats.BaseSkill[idx]
		skillRecord := d2datadict.GetSkillByName(*skillName)

		baseSkills[skillRecord.ID] = &HeroSkill{SkillPoints: 1, SkillRecord: skillRecord}
	}

	return &baseSkills
}
