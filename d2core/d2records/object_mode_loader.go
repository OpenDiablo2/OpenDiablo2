package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func objectModesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(ObjectModes)

	for d.Next() {
		record := &ObjectModeRecord{
			Name:  d.String("Name"),
			Token: d.String("Token"),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Object.Modes = records

	r.Logger.Infof("Loaded %d ObjectMode records", len(records))

	return nil
}
