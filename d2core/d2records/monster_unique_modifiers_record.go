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
	PickFrequencies struct {
		Normal    *PickFreq
		Nightmare *PickFreq
		Hell      *PickFreq
	}
	Name          string
	Exclude1      string
	Exclude2      string
	FPick         int
	ID            int
	Xfer          bool
	Champion      bool
	Enabled       bool
	ExpansionOnly bool
}

// PickFreq restricts the range of modifiers that can spawn ...
type PickFreq struct {
	// ...on champion monsters
	Champion int

	// ... on unique monsters
	Unique int
}
