package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type EventRecord struct {
	Event string
}

var Events map[string]*EventRecord

func LoadEvents(file []byte) {
	Events = make(map[string]*EventRecord)

	d := d2common.LoadDataDictionary(file)
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
