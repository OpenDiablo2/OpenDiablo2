package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func lowQualityLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(LowQualities, 0)

	for d.Next() {
		record := &LowQualityRecord{
			Name: d.String("Hireling Description"),
		}

		records = append(records, record)
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Item.LowQualityPrefixes = records

	r.Logger.Infof("Loaded %d Low Item Quality records", len(records))

	return nil
}
