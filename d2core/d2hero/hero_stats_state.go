package d2hero

type HeroStatsState struct{
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
	Stamina int // only MaxStamina is saved, Stamina gets reset on entering world
	NextLevelExp int
}