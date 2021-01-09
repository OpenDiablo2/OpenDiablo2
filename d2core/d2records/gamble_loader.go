package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func gambleLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(Gamble)

	for d.Next() {
		record := &GambleRecord{
			Name: d.String("name"),
			Code: d.String("code"),
		}
		records[record.Name] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d gamble records", len(records))

	r.Gamble = records

	return nil
}
