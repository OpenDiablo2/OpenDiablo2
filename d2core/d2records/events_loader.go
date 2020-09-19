package d2records

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadEvents loads all of the event records from events.txt
func eventsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(Events, 0)

	for d.Next() {
		record := &EventRecord{
			Event: d.String("event"),
		}

		records = append(records, record)
	}

	if d.Err != nil {
		return d.Err
	}

	log.Printf("Loaded %d Event records", len(records))

	r.Character.Events = records

	return nil
}
