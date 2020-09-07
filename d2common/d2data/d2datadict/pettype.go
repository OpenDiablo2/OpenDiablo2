package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
	"log"
)

// PetTypeRecord represents a single line in PetType.txt
// The information has been gathered from [https:// d2mods.info/forum/kb/viewarticle?a=355]
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

// PetTypes stores the PetTypeRecords
var PetTypes map[string]*PetTypeRecord // nolint:gochecknoglobals // Currently global by design

// LoadPetTypes loads PetTypeRecords into PetTypes
func LoadPetTypes(file []byte) {
	PetTypes = make(map[string]*PetTypeRecord)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		record := &PetTypeRecord{
			Name:      d.String("pet type"),
			ID:        d.Number("idx"),
			GroupID:   d.Number("group"),
			BaseMax:   d.Number("basemax"),
			Warp:      d.Bool("warp"),
			Range:     d.Bool("range"),
			PartySend: d.Bool("partysend"),
			Unsummon:  d.Bool("unsummon"),
			Automap:   d.Bool("automap"),
			IconName:  d.String("name"),
			DrawHP:    d.Bool("drawhp"),
			IconType:  d.Number("icontype"),
			BaseIcon:  d.String("baseicon"),
			MClass1:   d.Number("mclass1"),
			MIcon1:    d.String("micon1"),
			MClass2:   d.Number("mclass2"),
			MIcon2:    d.String("micon2"),
			MClass3:   d.Number("mclass3"),
			MIcon3:    d.String("micon3"),
			MClass4:   d.Number("mclass4"),
			MIcon4:    d.String("micon4"),
		}
		PetTypes[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d PetType records", len(PetTypes))
}
