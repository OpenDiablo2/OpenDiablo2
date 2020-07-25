package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type SetRecord struct {
	// index
	// String key linked to by the set field in SetItems.
	// txt - used to tie all of the set's items to the same set.
	Key string

	// name
	// String key to item's name in a .tbl file.
	StringTableKey string

	// Version 0 for vanilla, 100 for LoD expansion
	Version int

	// Level
	// set level, perhaps intended as a minimum level for partial or full attributes to appear
	// (reference only, not loaded into game).
	Level int

	// PCodeA, PCodeB -- PCode2a,PCode2b to PCode5a,PCode5b
	// An ID pointer of a property from Properties.txt,
	// these columns control each of the five pairs of different partial set modifiers a set item can
	// grant you at most.
	PCodeA [4]string
	PCodeB [4]string

	// PParamA, PParamB -- PParam2[a|b] to PParam5[a|b]
	// The parameter passed on to the associated property, this is used to pass skill IDs, state IDs,
	// monster IDs, montype IDs and the like on to the properties that require them,
	// these fields support calculations.
	PParamA [4]int
	PParamB [4]int

	// PMinA, PMaxA, PMinB, PMaxB -- P[Min|Max]2[a|b] to P[Min|Max]5[a|b]
	// Minimum value to assign to the associated property.
	// Certain properties have special interpretations based on stat encoding (e.g.
	// chance-to-cast and charged skills). See the File Guides for Properties.txt and ItemStatCost.
	// txt for further details.
	PMinA [4]int
	PMaxA [4]int
	PMinB [4]int
	PMaxB [4]int

	// FCode -- FCode1 to FCode8
	// An ID pointer of a property from Properties.txt,
	// these columns control each of the eight different full set modifiers a set item can grant you
	// at most.
	FCode [8]string

	// FParam -- FParam1 to FParam8
	// The parameter passed on to the associated property, this is used to pass skill IDs, state IDs,
	// monster IDs, montype IDs and the like on to the properties that require them,
	// these fields support calculations.
	FParam [8]int

	// FMin -- FMin1 to FMin8
	// Minimum value to assign to the associated property.
	// Certain properties have special interpretations based on stat encoding (e.g.
	// chance-to-cast and charged skills). See the File Guides for Properties.txt and ItemStatCost.
	// txt for further details.
	FMin [8]int

	// FMax -- FMax1 to FMax8
	// Maximum value to assign to the associated property.
	// Certain properties have special interpretations based on stat encoding (e.g.
	// chance-to-cast and charged skills). See the File Guides for Properties.txt and ItemStatCost.
	// txt for further details.
	FMax [8]int
}

// SetRecords contain the set records from sets.txt
var SetRecords map[string]*SetRecord

// LoadSetRecords loads set records from sets.txt
func LoadSets(file []byte) {
	SetRecords = make(map[string]*SetRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &SetRecord{
			Key:            d.String("index"),
			StringTableKey: d.String("name"),
			Version:        d.Number("version"),
			Level:          d.Number("level"),
			PCodeA:          [4]string{
				d.String("PCode2a"),
				d.String("PCode3a"),
				d.String("PCode4a"),
				d.String("PCode5a"),
			},
			PCodeB:          [4]string{
				d.String("PCode2b"),
				d.String("PCode3b"),
				d.String("PCode4b"),
				d.String("PCode5b"),
			},
			PParamA:          [4]int{
				d.Number("PParam2a"),
				d.Number("PParam3a"),
				d.Number("PParam4a"),
				d.Number("PParam5a"),
			},
			PParamB:          [4]int{
				d.Number("PParam2b"),
				d.Number("PParam3b"),
				d.Number("PParam4b"),
				d.Number("PParam5b"),
			},
			PMinA:          [4]int{
				d.Number("PMin2a"),
				d.Number("PMin3a"),
				d.Number("PMin4a"),
				d.Number("PMin5a"),
			},
			PMinB:          [4]int{
				d.Number("PMin2b"),
				d.Number("PMin3b"),
				d.Number("PMin4b"),
				d.Number("PMin5b"),
			},
			PMaxA:          [4]int{
				d.Number("PMax2a"),
				d.Number("PMax3a"),
				d.Number("PMax4a"),
				d.Number("PMax5a"),
			},
			PMaxB:          [4]int{
				d.Number("PMax2b"),
				d.Number("PMax3b"),
				d.Number("PMax4b"),
				d.Number("PMax5b"),
			},
			FCode:          [8]string{
				d.String("FCode1"),
				d.String("FCode2"),
				d.String("FCode3"),
				d.String("FCode4"),
				d.String("FCode5"),
				d.String("FCode6"),
				d.String("FCode7"),
				d.String("FCode9"),
			},
			FParam:          [8]int{
				d.Number("FParam1"),
				d.Number("FParam2"),
				d.Number("FParam3"),
				d.Number("FParam4"),
				d.Number("FParam5"),
				d.Number("FParam6"),
				d.Number("FParam7"),
				d.Number("FParam9"),
			},
			FMin:          [8]int{
				d.Number("FMin1"),
				d.Number("FMin2"),
				d.Number("FMin3"),
				d.Number("FMin4"),
				d.Number("FMin5"),
				d.Number("FMin6"),
				d.Number("FMin7"),
				d.Number("FMin9"),
			},
			FMax:          [8]int{
				d.Number("FMax1"),
				d.Number("FMax2"),
				d.Number("FMax3"),
				d.Number("FMax4"),
				d.Number("FMax5"),
				d.Number("FMax6"),
				d.Number("FMax7"),
				d.Number("FMax9"),
			},
		}

		SetRecords[record.Key] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d Sets records", len(SetRecords))
}
