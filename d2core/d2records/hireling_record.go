package d2records

// Hirelings stores hireling (mercenary) records
type Hirelings []*HirelingRecord

// HirelingRecord is a representation of rows in hireling.txt
// these records describe mercenaries
type HirelingRecord struct {
	WType1          string
	SubType         string
	Skill6          string
	Skill5          string
	Skill4          string
	Skill3          string
	Skill2          string
	Hireling        string
	NameFirst       string
	NameLast        string
	Skill1          string
	HireDesc        string
	WType2          string
	Level1          int
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
	HPPerLvl        int
	HP              int
	ExpPerLvl       int
	DefaultChance   int
	Gold            int
	Mode1           int
	Chance1         int
	ChancePerLevel1 int
	Seller          int
	LvlPerLvl1      int
	Level           int
	Mode2           int
	Chance2         int
	ChancePerLevel2 int
	Level2          int
	LvlPerLvl2      int
	Difficulty      int
	Mode3           int
	Chance3         int
	ChancePerLevel3 int
	Level3          int
	LvlPerLvl3      int
	Act             int
	Mode4           int
	Chance4         int
	ChancePerLevel4 int
	Level4          int
	LvlPerLvl4      int
	Class           int
	Mode5           int
	Chance5         int
	ChancePerLevel5 int
	Level5          int
	LvlPerLvl5      int
	ID              int
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
