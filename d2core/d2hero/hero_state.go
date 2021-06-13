package d2hero

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
)

// HeroState stores the state of the player
type HeroState struct {
	Equipment  d2inventory.CharacterEquipment `json:"equipment"`
	Skills     map[int]*HeroSkill             `json:"skills"`
	Stats      *HeroStatsState                `json:"stats"`
	FilePath   string                         `json:"-"`
	HeroName   string                         `json:"heroName"`
	Act        int                            `json:"act"`
	HeroType   d2enum.Hero                    `json:"heroType"`
	X          float64                        `json:"x"`
	Y          float64                        `json:"y"`
	LeftSkill  int                            `json:"leftSkill"`
	RightSkill int                            `json:"rightSkill"`
	Gold       int                            `json:"Gold"`
	Difficulty d2enum.DifficultyType          `json:"difficulty"`
}
