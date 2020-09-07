package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
	"log"
)

// BodyLocationRecord describes a body location that items can be equipped to
type BodyLocationRecord struct {
	Name string
	Code string
}

// BodyLocations contains the body location records
//nolint:gochecknoglobals // Currently global by design, only written once
var BodyLocations map[string]*BodyLocationRecord

// LoadBodyLocations loads body locations from
func LoadBodyLocations(file []byte) {
	BodyLocations = make(map[string]*BodyLocationRecord)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		location := &BodyLocationRecord{
			Name: d.String("Name"),
			Code: d.String("Code"),
		}
		BodyLocations[location.Code] = location
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d Body Location records", len(BodyLocations))
}
