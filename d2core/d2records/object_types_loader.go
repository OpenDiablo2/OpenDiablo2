package d2records

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

const (
	nameSize  = 32
	tokenSize = 20
)

// LoadObjectTypes loads ObjectTypeRecords from objtype.txt
func objectTypesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(ObjectTypes, 0)

	for d.Next() {
		record := ObjectTypeRecord{
			Name:  d.String("Name"),
			Token: d.String("Token"),
		}

		records = append(records, record)
	}

	if d.Err != nil {
		return d.Err
	}

	log.Printf("Loaded %d object types", len(records))

	r.Object.Types = records

	return nil
}
