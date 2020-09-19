package d2records

// Shrines contains the Unique Appellations
type Shrines map[string]*ShrineRecord

// ShrineRecord is a representation of a row from shrines.txt
type ShrineRecord struct {
	ShrineType       string // None, Recharge, Booster, or Magic
	ShrineName       string // Name of the Shrine
	Effect           string // Effect on the player
	Code             int    // Unique identifier
	Arg0             int    // ? (0-400)
	Arg1             int    // ? (0-2000)
	DurationFrames   int    // How long the shrine lasts in frames
	ResetTimeMinutes int    // How many minutes until the shrine resets?
	Rarity           int    // 1-3
	EffectClass      int    // 0-4
	LevelMin         int    // 0-32
}
