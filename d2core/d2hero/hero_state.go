package d2hero

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
)

// HeroState stores the state of the player
type HeroState struct {
	HeroName  string                         `json:"heroName"`
	HeroType  d2enum.Hero                    `json:"heroType"`
	HeroLevel int                            `json:"heroLevel"`
	Act       int                            `json:"act"`
	FilePath  string                         `json:"-"`
	Equipment d2inventory.CharacterEquipment `json:"equipment"`
	Stats     *HeroStatsState                `json:"stats"`
	Skills    map[int]*HeroSkill             `json:"skills"`
	X         float64                        `json:"x"`
	Y         float64                        `json:"y"`
}
