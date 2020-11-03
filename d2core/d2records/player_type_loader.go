package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func playerTypeLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(PlayerTypes)

	for d.Next() {
		record := &PlayerTypeRecord{
			Name:  d.String("Player Class"),
			Token: d.String("Token"),
		}

		if record.Name == expansionString {
			continue
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Logger.Infof("Loaded %d PlayerType records", len(records))

	r.Animation.Token.Player = records

	return nil
}
