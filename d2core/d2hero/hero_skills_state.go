package d2hero

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"

// HeroSkillsState hold all spells that a hero has.
type HeroSkillsState map[int]*HeroSkill

// CreateHeroSkillsState will assemble the hero skills from the class stats record.
func (f *HeroStateFactory) CreateHeroSkillsState(classStats *d2datadict.CharStatsRecord) *HeroSkillsState {
	baseSkills := HeroSkillsState{}

	for idx := range classStats.BaseSkill {
		skillName := &classStats.BaseSkill[idx]
		if len(*skillName) == 0 {
			continue
		}

		skillRecord := d2datadict.GetSkillByName(*skillName)
		baseSkills[skillRecord.ID] = &HeroSkill{SkillPoints: 1, SkillRecord: skillRecord}
	}

	skillRecord := d2datadict.GetSkillByName("Attack")
	baseSkills[skillRecord.ID] = &HeroSkill{SkillPoints: 1, SkillRecord: skillRecord}

	return &baseSkills
}
