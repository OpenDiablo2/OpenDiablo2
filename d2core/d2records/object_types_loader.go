package d2records

import (
	"log"
	"strings"

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
			Name:  sanitizeObjectString(d.String("Name")),
			Token: sanitizeObjectString(d.String("Token")),
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

func sanitizeObjectString(str string) string {
	result := strings.TrimSpace(strings.ReplaceAll(str, string(byte(0)), ""))
	result = strings.ToLower(result)

	return result
}
