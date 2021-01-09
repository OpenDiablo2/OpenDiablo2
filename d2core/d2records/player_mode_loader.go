package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadPlayerModes loads PlayerModeRecords into PlayerModes
func playerModesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(PlayerModes)

	for d.Next() {
		record := &PlayerModeRecord{
			Name:  d.String("Name"),
			Token: d.String("Token"),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Character.Modes = records

	r.Logger.Infof("Loaded %d PlayerMode records", len(records))

	return nil
}
