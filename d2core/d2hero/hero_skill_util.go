package d2hero

import "github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"

// HydrateSkills will load the SkillRecord & SkillDescriptionRecord from the asset manager, using the skill ID.
// This is done to avoid serializing the whole record data of HeroSkill to a game save or network packets.
// We cant do this while unmarshalling because there is no reference to the asset manager.
func HydrateSkills(skills map[int]*HeroSkill, asset *d2asset.AssetManager) {
	for skillID := range skills {
		heroSkill := skills[skillID]

		// TODO: figure out why these are nil sometimes
		if heroSkill == nil {
			continue
		}

		heroSkill.SkillRecord = asset.Records.Skill.Details[skillID]
		heroSkill.SkillDescriptionRecord = asset.Records.Skill.Descriptions[heroSkill.SkillRecord.Skilldesc]
		heroSkill.SkillPoints = skills[skillID].SkillPoints
	}
}
