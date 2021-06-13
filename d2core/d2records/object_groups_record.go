package d2records

const (
	objectsGroupSize     = 7
	memberDensityMin     = 0
	memberDensityMax     = 125
	memberProbabilityMin = 0
	memberProbabilityMax = 100
	expansionDataMarker  = "EXPANSION"
)

// ObjectGroups stores the ObjectGroupRecords.
type ObjectGroups map[int]*ObjectGroupRecord

// ObjectGroupRecord represents a single line in objgroup.txt.
// Information has been gathered from [https://d2mods.info/forum/kb/viewarticle?a=394].
type ObjectGroupRecord struct {
	Members   *[objectsGroupSize]ObjectGroupMember
	GroupName string
	Offset    int
	Shrines   bool
	Wells     bool
}

// ObjectGroupMember represents a member of an object group.
type ObjectGroupMember struct {
	// ID is the ID of the object.
	ID int

	// Density is how densely the level is filled with the object.
	// Must be below 125.
	Density int

	// Probability is the probability of this particular object being spawned.
	// The value is a percentage.
	Probability int
}
