package d2datadict

import (
	"fmt"
	"log"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

const (
	objectsGroupSize     = 7
	memberDensityMin     = 0
	memberDensityMax     = 125
	memberProbabilityMin = 0
	memberProbabilityMax = 100
	expansionDataMarker  = "EXPANSION"
)

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

// ObjectGroups stores the ObjectGroupRecords.
var ObjectGroups map[int]*ObjectGroupRecord //nolint:gochecknoglobals // Currently global by design.

// LoadObjectGroups loads the ObjectGroupRecords into ObjectGroups.
func LoadObjectGroups(file []byte) {
	ObjectGroups = make(map[int]*ObjectGroupRecord)
	d := d2common.LoadDataDictionary(file)

	for d.Next() {
		groupName := d.String("GroupName")
		if groupName == expansionDataMarker {
			continue
		}

		shrines, wells := d.Bool("Shrines"), d.Bool("Wells")
		record := &ObjectGroupRecord{
			GroupName: groupName,
			Offset:    d.Number("Offset"),
			Members:   createMembers(d, shrines || wells),
			Shrines:   shrines,
			Wells:     wells,
		}
		ObjectGroups[record.Offset] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d ObjectGroup records", len(ObjectGroups))
}

func createMembers(d *d2common.DataDictionary, shrinesOrWells bool) *[objectsGroupSize]ObjectGroupMember {
	var members [objectsGroupSize]ObjectGroupMember

	for i := 0; i < objectsGroupSize; i++ {
		suffix := strconv.Itoa(i)
		members[i].ID = d.Number("ID" + suffix)

		members[i].Density = d.Number("DENSITY" + suffix)
		if members[i].Density < memberDensityMin || members[i].Density > memberDensityMax {
			panic(fmt.Sprintf("Invalid object group member density: %v, in group: %v",
				members[i].Density, d.String("GroupName"))) // Vanilla crashes when density is over 125.
		}

		if shrinesOrWells && members[i].Density != 0 {
			panic(fmt.Sprintf("Shrine and well object groups must have densities set to 0, in group: %v", d.String("GroupName")))
		}

		members[i].Probability = d.Number("PROB" + suffix)
		if members[i].Probability < memberProbabilityMin || members[i].Probability > memberProbabilityMax {
			panic(fmt.Sprintf("Invalid object group member probability: %v, in group: %v",
				members[i].Probability, d.String("GroupName")))
		}
	}

	return &members
}
