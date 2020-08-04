package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// UniqueAppellationRecord described the extra suffix of a unique monster name
type UniqueAppellationRecord struct {
	// The title
	Name string
}

// UniqueAppellations contains all of the UniqueAppellationRecords
//nolint:gochecknoglobals // Currently global by design
var UniqueAppellations map[string]*UniqueAppellationRecord

// LoadUniqueAppellations loads UniqueAppellationRecords from UniqueAppelation.txt
func LoadUniqueAppellations(file []byte) {
	UniqueAppellations = make(map[string]*UniqueAppellationRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &UniqueAppellationRecord{
			Name: d.String("Name"),
		}
		UniqueAppellations[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d UniqueAppellation records", len(UniqueAppellations))
}
