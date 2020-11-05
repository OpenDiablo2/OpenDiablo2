package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func hirelingDescriptionLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(HirelingDescriptions)

	for d.Next() {
		record := &HirelingDescriptionRecord{
			Name:  d.String("Hireling Description"),
			Token: d.String("Code"),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Hireling.Descriptions = records

	r.Logger.Infof("Loaded %d Hireling Descriptions records", len(records))

	return nil
}
