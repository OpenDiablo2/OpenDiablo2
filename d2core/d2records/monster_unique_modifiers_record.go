package d2records

// See [https://d2mods.info/forum/kb/viewarticle?a=161] for more info

// MonsterUniqueModifiers stores the MonsterUniqueModifierRecords
type MonsterUniqueModifiers map[string]*MonUModRecord

// MonsterUniqueModifierConstants contains constants from monumod.txt,
// can be accessed with indices from d2enum.MonUModConstIndex
type MonsterUniqueModifierConstants []int

// MonUModRecord represents a single line in monumod.txt
// Information gathered from [https://d2mods.info/forum/kb/viewarticle?a=161]
type MonUModRecord struct {
	// Name of modifer, not used by other files
	Name string

	// ID of the modifier,
	// the Mod fields of SuperUniqueRecord refer to these ID's
	ID int

	// Enabled boolean for whether this modifier can be applied
	Enabled bool

	// ExpansionOnly boolean for whether this modifier can only be applied in an expansion game.
	// In the file, the value 100 represents expansion only
	ExpansionOnly bool

	// If true, "Minion" will be displayed below the life bar of minions of
	// the monster with this modifier
	Xfer bool

	// Champion boolean, only usable by champion monsters
	Champion bool

	// FPick Unknown
	FPick int

	// Exclude1 monster type code that cannot have this modifier
	Exclude1 string

	// Exclude2 monster type code that cannot have this modifier
	Exclude2 string

	PickFrequencies struct {
		Normal    *PickFreq
		Nightmare *PickFreq
		Hell      *PickFreq
	}
}

// PickFreq restricts the range of modifiers that can spawn ...
type PickFreq struct {
	// ...on champion monsters
	Champion int

	// ... on unique monsters
	Unique int
}
