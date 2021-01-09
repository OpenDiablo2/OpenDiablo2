package d2records

import (
	"fmt"

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

// Loadrecords loads runes records into a map[string]*RunesRecord
func runewordLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(map[string]*RunesRecord)

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
			Properties: make([]*RunewordProperty, 0),
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
				prop := &RunewordProperty{
					code,
					d.String(fmt.Sprintf(fmtRunewordPropParam, idx+1)),
					d.Number(fmt.Sprintf(fmtRunewordPropMin, idx+1)),
					d.Number(fmt.Sprintf(fmtRunewordPropMax, idx+1)),
				}

				record.Properties = append(record.Properties, prop)
			}
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Item.Runewords = records

	r.Logger.Infof("Loaded %d records records", len(records))

	return nil
}
