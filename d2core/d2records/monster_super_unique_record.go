package d2records

// https://d2mods.info/forum/kb/viewarticle?a=162

// SuperUniques stores all of the SuperUniqueRecords
type SuperUniques map[string]*SuperUniqueRecord

// SuperUniqueRecord Defines the unique monsters and their properties.
// SuperUnique monsters are boss monsters which always appear at the same places
// and always have the same base special abilities
// with the addition of one or two extra ones per difficulty (Nightmare provides one extra ability, Hell provides two).
// Notable examples are enemies such as Corpsefire, Pindleskin or Nihlathak.
type SuperUniqueRecord struct {
	TreasureClassNormal    string
	Name                   string
	Class                  string
	HcIdx                  string
	MonSound               string
	UTransNightmare        string
	UTransNormal           string
	TreasureClassHell      string
	Key                    string
	TreasureClassNightmare string
	UTransHell             string
	Mod                    [3]int
	MaxGrp                 int
	MinGrp                 int
	Stacks                 bool
	AutoPosition           bool
	IsExpansion            bool
}
