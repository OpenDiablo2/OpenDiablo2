package d2hero

import (
	"encoding/json"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"
)

// HeroSkill stores additional payload for a skill of a hero.
type HeroSkill struct {
	*d2records.SkillRecord
	*d2records.SkillDescriptionRecord
	SkillPoints int
	Shallow     *shallowHeroSkill
}

// An auxiliary struct which only stores the ID of the SkillRecord, instead of the whole SkillRecord
// and SkillDescrptionRecord.
type shallowHeroSkill struct {
	SkillID     int `json:"skillId"`
	SkillPoints int `json:"skillPoints"`
}

// MarshalJSON overrides the default logic used when the HeroSkill is serialized to a byte array.
func (hs *HeroSkill) MarshalJSON() ([]byte, error) {
	// only serialize the Shallow object instead of the SkillRecord & SkillDescriptionRecord
	bytes, err := json.Marshal(hs.Shallow)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// UnmarshalJSON overrides the default logic used when the HeroSkill is deserialized from a byte array.
func (hs *HeroSkill) UnmarshalJSON(data []byte) error {
	shallow := &shallowHeroSkill{}
	if err := json.Unmarshal(data, shallow); err != nil {
		return err
	}

	hs.Shallow = shallow

	return nil
}
