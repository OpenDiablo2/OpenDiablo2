package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

//PetTypeRecord represents a single line in PetType.txt
//The information has been gathered from [https://d2mods.info/forum/kb/viewarticle?a=355]
type PetTypeRecord struct {
	//Name of the pet type, reffered by "pettype" in skills.txt
	PetType string

	//ID number of the pet type
	Idx int

	//ID number of the group this pet belongs to
	Group int

	BaseMax int

	//If true, the pet warps with the player, otherwise it dies
	Warp bool

	//If true, and Warp is false, the pet only die if the distance between the player
	//and the pet exceeds 41 sub-tiles.
	Range bool

	//Unknown
	PartySend bool

	//If true, can be unsummoned
	Unsummon bool

	//If true, pet is displayed on the automap
	Automap bool

	//String file for the text under the pet icon
	Name string

	//If true, the pet's HP will be displayed under the icon
	DrawHP bool

	//Pet icon type
	IconType int

	//.dc6 file for the pet's icon, located in /data/global/ui/hireables
	BaseIcon string

	//Alternative pet index from monstats.txt
	MClass1 int

	//Alternative pet icon .dc6 file
	MIcon1 string

	//ditto, there can be four alternatives
	MClass2 int
	MIcon2  string
	MClass3 int
	MIcon3  string
	MClass4 int
	MIcon4  string

	//EOL int, not loaded
}

//PetTypes stores the PetTypeRecords
var PetTypes map[string]*PetTypeRecord //nolint:gochecknoglobals // Currently global by design

//LoadPetTypes loads PetTypeRecords into PetTypes
func LoadPetTypes(file []byte) {
	PetTypes = make(map[string]*PetTypeRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &PetTypeRecord{
			PetType:   d.String("pet type"),
			Idx:       d.Number("idx"),
			Group:     d.Number("group"),
			BaseMax:   d.Number("basemax"),
			Warp:      d.Bool("warp"),
			Range:     d.Bool("range"),
			PartySend: d.Bool("partysend"),
			Unsummon:  d.Bool("unsummon"),
			Automap:   d.Bool("automap"),
			Name:      d.String("name"),
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
		PetTypes[record.PetType] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d PetType records", len(PetTypes))
}
