package d2records

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadPlayerClasses loads the PlayerClassRecords into PlayerClasses
func playerClassLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(PlayerClasses)

	for d.Next() {
		record := &PlayerClassRecord{
			Name: d.String("Player Class"),
			Code: d.String("Code"),
		}

		if record.Name == expansionString {
			continue
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	if d.Err != nil {
		return d.Err
	}

	log.Printf("Loaded %d PlayerClass records", len(records))

	r.Character.Classes = records

	return nil
}
