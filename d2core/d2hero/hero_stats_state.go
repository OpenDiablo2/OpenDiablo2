package d2hero

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"
)

// HeroStatsState is a serializable state of hero stats.
type HeroStatsState struct {
	Level      int `json:"level"`
	Experience int `json:"experience"`

	Strength  int `json:"strength"`
	Energy    int `json:"energy"`
	Dexterity int `json:"dexterity"`
	Vitality  int `json:"vitality"`
	// there are stats and skills points remaining to add.
	StatsPoints int `json:"statsPoints"`
	SkillPoints int `json:"skillPoints"`

	Health     int     `json:"health"`
	MaxHealth  int     `json:"maxHealth"`
	Mana       int     `json:"mana"`
	MaxMana    int     `json:"maxMana"`
	Stamina    float64 `json:"-"` // only MaxStamina is saved, Stamina gets reset on entering world
	MaxStamina int     `json:"maxStamina"`

	// values which are not saved/loaded(computed)
	NextLevelExp int `json:"-"`
}

// CreateHeroStatsState generates a running state from a hero stats.
func (f *HeroStateFactory) CreateHeroStatsState(heroClass d2enum.Hero, classStats *d2records.CharStatRecord) *HeroStatsState {
	result := HeroStatsState{
		Level:        1,
		Experience:   0,
		NextLevelExp: f.asset.Records.GetExperienceBreakpoint(heroClass, 1),
		Strength:     classStats.InitStr,
		Dexterity:    classStats.InitDex,
		Vitality:     classStats.InitVit,
		Energy:       classStats.InitEne,
		StatsPoints:  0,
		SkillPoints:  0,

		MaxHealth:  classStats.InitVit * classStats.LifePerVit,
		MaxMana:    classStats.InitEne * classStats.ManaPerEne,
		MaxStamina: classStats.InitStamina,
		// https://github.com/OpenDiablo2/OpenDiablo2/issues/814
	}

	result.Mana = result.MaxMana
	result.Health = result.MaxHealth
	result.Stamina = float64(result.MaxStamina)

	return &result
}
