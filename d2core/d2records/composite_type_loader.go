package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func compositeTypeLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(CompositeTypes)

	for d.Next() {
		record := &CompositeTypeRecord{
			Name:  d.String("Name"),
			Token: d.String("Token"),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Animation.Token.Composite = records

	r.Logger.Infof("Loaded %d Composite Type records", len(records))

	return nil
}
