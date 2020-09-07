package d2datadict

import (
	"fmt"
	"log"
	
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

const (
	numPartialSetProperties = 4
	numFullSetProperties    = 8
	fmtPropCode             = "%sCode%d%s"
	fmtPropParam            = "%sParam%d%s"
	fmtPropMin              = "%sMin%d%s"
	fmtPropMax              = "%sMax%d%s"
	partialIdxOffset        = 2
	setPartialToken         = "P"
	setPartialTokenA        = "a"
	setPartialTokenB        = "b"
	setFullToken            = "F"
)

// SetRecord describes the set bonus for a group of set items
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

	// Properties contains the partial and full set bonus properties.
	Properties struct {
		PartialA []*setProperty
		PartialB []*setProperty
		Full     []*setProperty
	}
}

type setProperty struct {
	// Code is an ID pointer of a property from Properties.txt,
	// these columns control each of the eight different full set modifiers a set item can grant you
	// at most.
	Code string

	// Param is the passed on to the associated property, this is used to pass skill IDs, state IDs,
	// monster IDs, montype IDs and the like on to the properties that require them,
	// these fields support calculations.
	Param string

	// Min value to assign to the associated property.
	// Certain properties have special interpretations based on stat encoding (e.g.
	// chance-to-cast and charged skills). See the File Guides for Properties.txt and ItemStatCost.
	// txt for further details.
	Min int

	// Max value to assign to the associated property.
	// Certain properties have special interpretations based on stat encoding (e.g.
	// chance-to-cast and charged skills). See the File Guides for Properties.txt and ItemStatCost.
	// txt for further details.
	Max int
}

// SetRecords contain the set records from sets.txt
var SetRecords map[string]*SetRecord //nolint:gochecknoglobals // Currently global by design

// LoadSetRecords loads set records from sets.txt
func LoadSetRecords(file []byte) { //nolint:funlen // doesn't make sense to split
	SetRecords = make(map[string]*SetRecord)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		record := &SetRecord{
			Key:            d.String("index"),
			StringTableKey: d.String("name"),
			Version:        d.Number("version"),
			Level:          d.Number("level"),
			Properties: struct {
				PartialA []*setProperty
				PartialB []*setProperty
				Full     []*setProperty
			}{
				PartialA: make([]*setProperty, 0),
				PartialB: make([]*setProperty, 0),
				Full:     make([]*setProperty, 0),
			},
		}

		// for partial properties 2a thru 5b
		for idx := 0; idx < numPartialSetProperties; idx++ {
			num := idx + partialIdxOffset // needs to be 2,3,4,5
			columnA := fmt.Sprintf(fmtPropCode, setPartialToken, num, setPartialTokenA)
			columnB := fmt.Sprintf(fmtPropCode, setPartialToken, num, setPartialTokenB)

			if codeA := d.String(columnA); codeA != "" {
				paramColumn := fmt.Sprintf(fmtPropParam, setPartialToken, num, setPartialTokenA)
				minColumn := fmt.Sprintf(fmtPropMin, setPartialToken, num, setPartialTokenA)
				maxColumn := fmt.Sprintf(fmtPropMax, setPartialToken, num, setPartialTokenA)

				propA := &setProperty{
					Code:  codeA,
					Param: d.String(paramColumn),
					Min:   d.Number(minColumn),
					Max:   d.Number(maxColumn),
				}

				record.Properties.PartialA = append(record.Properties.PartialA, propA)
			}

			if codeB := d.String(columnB); codeB != "" {
				paramColumn := fmt.Sprintf(fmtPropParam, setPartialToken, num, setPartialTokenB)
				minColumn := fmt.Sprintf(fmtPropMin, setPartialToken, num, setPartialTokenB)
				maxColumn := fmt.Sprintf(fmtPropMax, setPartialToken, num, setPartialTokenB)

				propB := &setProperty{
					Code:  codeB,
					Param: d.String(paramColumn),
					Min:   d.Number(minColumn),
					Max:   d.Number(maxColumn),
				}

				record.Properties.PartialB = append(record.Properties.PartialB, propB)
			}
		}

		for idx := 0; idx < numFullSetProperties; idx++ {
			num := idx + 1
			codeColumn := fmt.Sprintf(fmtPropCode, setFullToken, num, "")
			paramColumn := fmt.Sprintf(fmtPropParam, setFullToken, num, "")
			minColumn := fmt.Sprintf(fmtPropMin, setFullToken, num, "")
			maxColumn := fmt.Sprintf(fmtPropMax, setFullToken, num, "")

			if code := d.String(codeColumn); code != "" {
				prop := &setProperty{
					Code:  code,
					Param: d.String(paramColumn),
					Min:   d.Number(minColumn),
					Max:   d.Number(maxColumn),
				}

				record.Properties.Full = append(record.Properties.Full, prop)
			}
		}

		SetRecords[record.Key] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d Sets records", len(SetRecords))
}
