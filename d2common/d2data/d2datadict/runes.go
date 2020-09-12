package d2datadict

import (
	"fmt"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

const (
	numRunewordTypeInclude = 6
	numRunewordTypeExclude = 3
	numRunewordMaxSockets  = 6
	numRunewordProperties  = 7
)

// format strings
const (
	fmtTypeInclude       = "itype%d"
	fmtTypeExclude       = "etype%d"
	fmtRuneStr           = "Rune%d"
	fmtRunewordPropCode  = "T1Code%d"
	fmtRunewordPropParam = "T1Param%d"
	fmtRunewordPropMin   = "T1Min%d"
	fmtRunewordPropMax   = "T1Max%d"
)

// RunesRecord is a representation of a single row of runes.txt. It defines
// runewords available in the game.
type RunesRecord struct {
	Name     string
	RuneName string // More of a note - the actual name should be read from the TBL files.
	Complete bool   // An enabled/disabled flag. Only "Complete" runewords work in game.
	Server   bool   // Marks a runeword as only available on ladder, not single player or tcp/ip.

	// The item types for includsion/exclusion for this runeword record
	ItemTypes struct {
		Include []string
		Exclude []string
	}

	// Runes slice of ID pointers from Misc.txt, controls what runes are
	// required to make the rune word and in what order they are to be socketed.
	Runes []string

	Properties []*runewordProperty
}

type runewordProperty struct {
	// Code is the property code
	Code string

	// Param is either string or int, parameter for the property
	Param string

	// Min is the minimum value for the property
	Min int

	// Max is the maximum value for the property
	Max int
}

// Runewords stores all of the RunesRecords
var Runewords map[string]*RunesRecord //nolint:gochecknoglobals // Currently global by design, only written once

// LoadRunewords loads runes records into a map[string]*RunesRecord
func LoadRunewords(file []byte) {
	Runewords = make(map[string]*RunesRecord)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		record := &RunesRecord{
			Name:     d.String("name"),
			RuneName: d.String("Rune Name"),
			Complete: d.Bool("complete"),
			Server:   d.Bool("server"),
			ItemTypes: struct {
				Include []string
				Exclude []string
			}{
				Include: make([]string, 0),
				Exclude: make([]string, 0),
			},
			Runes:      make([]string, 0),
			Properties: make([]*runewordProperty, 0),
		}

		for idx := 0; idx < numRunewordTypeInclude; idx++ {
			column := fmt.Sprintf(fmtTypeInclude, idx+1)
			if code := d.String(column); code != "" {
				record.ItemTypes.Include = append(record.ItemTypes.Include, code)
			}
		}

		for idx := 0; idx < numRunewordTypeExclude; idx++ {
			column := fmt.Sprintf(fmtTypeExclude, idx+1)
			if code := d.String(column); code != "" {
				record.ItemTypes.Exclude = append(record.ItemTypes.Exclude, code)
			}
		}

		for idx := 0; idx < numRunewordMaxSockets; idx++ {
			column := fmt.Sprintf(fmtRuneStr, idx+1)
			if code := d.String(column); code != "" {
				record.Runes = append(record.Runes, code)
			}
		}

		for idx := 0; idx < numRunewordProperties; idx++ {
			codeColumn := fmt.Sprintf(fmtRunewordPropCode, idx+1)
			if code := codeColumn; code != "" {
				prop := &runewordProperty{
					code,
					d.String(fmt.Sprintf(fmtRunewordPropParam, idx+1)),
					d.Number(fmt.Sprintf(fmtRunewordPropMin, idx+1)),
					d.Number(fmt.Sprintf(fmtRunewordPropMax, idx+1)),
				}

				record.Properties = append(record.Properties, prop)
			}
		}

		Runewords[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d Runewords records", len(Runewords))
}
