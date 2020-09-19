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
	// GroupName is the name of the group.
	GroupName string

	// Offset is the ID of the group, referred to by Levels.txt.
	Offset int

	// Members are the objects in the group.
	Members *[objectsGroupSize]ObjectGroupMember

	// Shrines determines whether this is a group of shrines.
	// Note: for shrine groups, densities must be set to 0.
	Shrines bool

	// Wells determines whether this is a group of wells.
	// Note: for wells groups, densities must be set to 0.
	Wells bool
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
