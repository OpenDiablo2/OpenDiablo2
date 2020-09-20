package d2hero

import (
	"encoding/json"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"
)

// HeroSkill stores additional payload for a skill of a hero.
type HeroSkill struct {
	*d2records.SkillRecord
	*d2records.SkillDescriptionRecord
	SkillPoints int
	shallow     *shallowHeroSkill
}

// An auxilary struct which only stores the ID of the SkillRecord, instead of the whole SkillRecord and SkillDescrptionRecord.
type shallowHeroSkill struct {
	SkillID     int `json:"skillId"`
	SkillPoints int `json:"skillPoints"`
}

// MarshalJSON overrides the default logic used when the HeroSkill is serialized to a byte array.
func (hs *HeroSkill) MarshalJSON() ([]byte, error) {
	// only serialize the ID instead of the whole SkillRecord object.
	shallow := shallowHeroSkill{
		SkillID:     hs.SkillRecord.ID,
		SkillPoints: hs.SkillPoints,
	}

	bytes, err := json.Marshal(shallow)
	if err != nil {
		log.Fatalln(err)
	}

	return bytes, err
}

// UnmarshalJSON overrides the default logic used when the HeroSkill is deserialized from a byte array.
func (hs *HeroSkill) UnmarshalJSON(data []byte) error {
	shallow := &shallowHeroSkill{}
	if err := json.Unmarshal(data, shallow); err != nil {
		return err
	}

	hs.shallow = shallow

	return nil
}
