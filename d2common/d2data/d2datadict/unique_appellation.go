package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type UniqueAppellationRecord struct {
	// The title
	Name string
}

var UniqueAppellations map[string]*UniqueAppellationRecord

func LoadUniqueAppellations(file []byte) {
	UniqueAppellations = make(map[string]*UniqueAppellationRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &UniqueAppellationRecord{
			Name:      d.String("Name"),
		}
		UniqueAppellations[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d UniqueAppellation records", len(UniqueAppellations))
}
