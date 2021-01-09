package d2hero

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"
)

// HeroStatsState is a serializable state of hero stats.
type HeroStatsState struct {
	Level      int `json:"level"`
	Experience int `json:"experience"`

	Vitality  int `json:"vitality"`
	Energy    int `json:"energy"`
	Strength  int `json:"strength"`
	Dexterity int `json:"dexterity"`

	AttackRating  int `json:"attackRating"`
	DefenseRating int `json:"defenseRating"`

	MaxStamina int `json:"maxStamina"`
	Health     int `json:"health"`
	MaxHealth  int `json:"maxHealth"`
	Mana       int `json:"mana"`
	MaxMana    int `json:"maxMana"`

	FireResistance      int `json:"fireResistance"`
	ColdResistance      int `json:"coldResistance"`
	LightningResistance int `json:"lightningResistance"`
	PoisonResistance    int `json:"poisonResistance"`

	// values which are not saved/loaded(computed)
	Stamina      float64 `json:"-"` // only MaxStamina is saved, Stamina gets reset on entering world
	NextLevelExp int     `json:"-"`
}

// CreateHeroStatsState generates a running state from a hero stats.
func (f *HeroStateFactory) CreateHeroStatsState(heroClass d2enum.Hero, classStats *d2records.CharStatsRecord) *HeroStatsState {
	result := HeroStatsState{
		Level:        1,
		Experience:   0,
		NextLevelExp: f.asset.Records.GetExperienceBreakpoint(heroClass, 1),
		Strength:     classStats.InitStr,
		Dexterity:    classStats.InitDex,
		Vitality:     classStats.InitVit,
		Energy:       classStats.InitEne,

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
