package d2records

// Hirelings stores hireling (mercenary) records
type Hirelings []*HirelingRecord

// HirelingRecord is a representation of rows in hireling.txt
// these records describe mercenaries
type HirelingRecord struct {
	Hireling        string
	SubType         string
	ID              int
	Class           int
	Act             int
	Difficulty      int
	Level           int
	Seller          int
	NameFirst       string
	NameLast        string
	Gold            int
	ExpPerLvl       int
	HP              int
	HPPerLvl        int
	Defense         int
	DefPerLvl       int
	Str             int
	StrPerLvl       int
	Dex             int
	DexPerLvl       int
	AR              int
	ARPerLvl        int
	Share           int
	DmgMin          int
	DmgMax          int
	DmgPerLvl       int
	Resist          int
	ResistPerLvl    int
	WType1          string
	WType2          string
	HireDesc        string
	DefaultChance   int
	Skill1          string
	Mode1           int
	Chance1         int
	ChancePerLevel1 int
	Level1          int
	LvlPerLvl1      int
	Skill2          string
	Mode2           int
	Chance2         int
	ChancePerLevel2 int
	Level2          int
	LvlPerLvl2      int
	Skill3          string
	Mode3           int
	Chance3         int
	ChancePerLevel3 int
	Level3          int
	LvlPerLvl3      int
	Skill4          string
	Mode4           int
	Chance4         int
	ChancePerLevel4 int
	Level4          int
	LvlPerLvl4      int
	Skill5          string
	Mode5           int
	Chance5         int
	ChancePerLevel5 int
	Level5          int
	LvlPerLvl5      int
	Skill6          string
	Mode6           int
	Chance6         int
	ChancePerLevel6 int
	Level6          int
	LvlPerLvl6      int
	Head            int
	Torso           int
	Weapon          int
	Shield          int
}
