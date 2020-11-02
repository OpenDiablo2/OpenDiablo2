package d2records

import (
	"fmt"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadObjectGroups loads the ObjectGroupRecords into ObjectGroups.
func objectGroupsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(ObjectGroups)

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
		records[record.Offset] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d ObjectGroup records", len(records))

	return nil
}

func createMembers(d *d2txt.DataDictionary, shrinesOrWells bool) *[objectsGroupSize]ObjectGroupMember {
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
