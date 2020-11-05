package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadElemTypes loads ElemTypeRecords into ElemTypes
func elemTypesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(ElemTypes)

	for d.Next() {
		record := &ElemTypeRecord{
			ElemType: d.String("Elemental Type"),
			Code:     d.String("Code"),
		}
		records[record.ElemType] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d ElemType records", len(records))

	r.ElemTypes = records

	return nil
}
