package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func componentCodesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(ComponentCodes)

	for d.Next() {
		record := &ComponentCodeRecord{
			Component: d.String("component"),
			Code:      d.String("code"),
		}
		records[record.Component] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d ComponentCode records", len(records))

	r.ComponentCodes = records

	return nil
}
