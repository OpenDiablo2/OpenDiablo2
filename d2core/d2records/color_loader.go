package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func colorsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(Colors)

	for d.Next() {
		record := &ColorRecord{
			TransformColor: d.String("Transform Color"),
			Code:           d.String("Code"),
		}

		records[record.TransformColor] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Colors = records

	r.Logger.Infof("Loaded %d Color records", len(records))

	return nil
}
