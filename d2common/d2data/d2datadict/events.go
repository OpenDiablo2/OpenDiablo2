package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
	"log"
)

// EventRecord is a representation of a single row from events.txt
type EventRecord struct {
	Event string
}

// Events holds all of the event records from events.txt
var Events map[string]*EventRecord //nolint:gochecknoglobals // Currently global by design

// LoadEvents loads all of the event records from events.txt
func LoadEvents(file []byte) {
	Events = make(map[string]*EventRecord)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		record := &EventRecord{
			Event: d.String("event"),
		}
		Events[record.Event] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d Event records", len(Events))
}
