package d2records

import (
	"fmt"

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

// LoadSetRecords loads set records from sets.txt
//nolint:funlen // doesn't make sense to split
func setLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(Sets)

	for d.Next() {
		record := &SetRecord{
			Key:            d.String("index"),
			StringTableKey: d.String("name"),
			Version:        d.Number("version"),
			Level:          d.Number("level"),
			Properties: struct {
				PartialA []*SetProperty
				PartialB []*SetProperty
				Full     []*SetProperty
			}{
				PartialA: make([]*SetProperty, 0),
				PartialB: make([]*SetProperty, 0),
				Full:     make([]*SetProperty, 0),
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

				propA := &SetProperty{
					Code:      codeA,
					Parameter: d.String(paramColumn),
					Min:       d.Number(minColumn),
					Max:       d.Number(maxColumn),
				}

				record.Properties.PartialA = append(record.Properties.PartialA, propA)
			}

			if codeB := d.String(columnB); codeB != "" {
				paramColumn := fmt.Sprintf(fmtPropParam, setPartialToken, num, setPartialTokenB)
				minColumn := fmt.Sprintf(fmtPropMin, setPartialToken, num, setPartialTokenB)
				maxColumn := fmt.Sprintf(fmtPropMax, setPartialToken, num, setPartialTokenB)

				propB := &SetProperty{
					Code:      codeB,
					Parameter: d.String(paramColumn),
					Min:       d.Number(minColumn),
					Max:       d.Number(maxColumn),
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
				prop := &SetProperty{
					Code:      code,
					Parameter: d.String(paramColumn),
					Min:       d.Number(minColumn),
					Max:       d.Number(maxColumn),
				}

				record.Properties.Full = append(record.Properties.Full, prop)
			}
		}

		records[record.Key] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Item.Sets = records

	r.Logger.Infof("Loaded %d records records", len(records))

	return nil
}
