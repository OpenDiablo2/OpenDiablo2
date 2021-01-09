package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func uniqueAppellationsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(UniqueAppellations)

	for d.Next() {
		record := &UniqueAppellationRecord{
			Name: d.String("Name"),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Monster.Unique.Appellations = records

	r.Logger.Infof("Loaded %d UniqueAppellation records", len(records))

	return nil
}
