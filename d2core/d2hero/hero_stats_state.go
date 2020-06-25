package d2hero

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

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
	Stamina      int // only MaxStamina is saved, Stamina gets reset on entering world
	NextLevelExp int
}

func CreateHeroStatsState(heroClass d2enum.Hero, classStats d2datadict.CharStatsRecord, level int, exp int) *HeroStatsState {
	result := HeroStatsState{
		Level:        level,
		Experience:   exp,
		NextLevelExp: d2datadict.GetExperienceBreakpoint(heroClass, 1), 
		Strength:     classStats.InitStr,
		Dexterity:    classStats.InitDex,
		Vitality:     classStats.InitVit,
		Energy:       classStats.InitEne,
		//TODO: proper formula for calculating health and mana
		Health:     classStats.InitVit * classStats.LifePerVit / 4,
		MaxHealth:  classStats.InitVit * classStats.LifePerVit / 4,
		Mana:       classStats.InitEne * classStats.ManaPerEne / 4,
		MaxMana:    classStats.InitEne * classStats.ManaPerEne / 4,
		Stamina:    classStats.InitStamina,
		MaxStamina: classStats.InitStamina,
		//TODO chance to hit, defense rating
	}

	//TODO: those are added only for demonstration purposes(to show that hp mana exp status bars and character stats panel get updated depending on current stats)
	result.Health /= 2
	result.Mana /= 3
	result.Experience = result.NextLevelExp / 3

	return &result
}
