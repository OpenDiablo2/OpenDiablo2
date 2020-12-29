package d2records

// The information has been gathered from [https:// d2mods.info/forum/kb/viewarticle?a=355]

// PetTypes stores the PetTypeRecords
type PetTypes map[string]*PetTypeRecord

// PetTypeRecord represents a single line in PetType.txt
type PetTypeRecord struct {
	// Name of the pet type, refferred by "pettype" in skills.txt
	Name string

	// Name text under the pet icon
	IconName string

	// .dc6 file for the pet's icon, located in /data/global/ui/hireables
	BaseIcon string

	// Alternative pet icon .dc6 file
	MIcon1 string
	MIcon2 string
	MIcon3 string
	MIcon4 string

	// ID number of the pet type
	ID int

	// GroupID number of the group this pet belongs to
	GroupID int

	// BaseMax unknown what this does...
	BaseMax int

	// Pet icon type
	IconType int

	// Alternative pet index from monstats.txt
	MClass1 int
	MClass2 int
	MClass3 int
	MClass4 int

	// Warp warps with the player, otherwise it dies
	Warp bool

	// Range the pet only die if the distance between the player  and the pet exceeds 41 sub-tiles.
	Range bool

	// Unknown
	PartySend bool

	// Unsummon whether the pet can be unsummoned
	Unsummon bool

	// Automap whether the pet is displayed on the automap
	Automap bool

	// If true, the pet's HP will be displayed under the icon
	DrawHP bool
}
