package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
	"log"
)

// MonModeRecord is a representation of a single row of Monmode.txt
type MonModeRecord struct {
	Name  string
	Token string
	Code  string
}

// MonModes stores all of the GemsRecords
var MonModes map[string]*MonModeRecord //nolint:gochecknoglobals // Currently global by design, only written once

// LoadMonModes loads gem records into a map[string]*MonModeRecord
func LoadMonModes(file []byte) {
	MonModes = make(map[string]*MonModeRecord)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		record := &MonModeRecord{
			Name:  d.String("name"),
			Token: d.String("token"),
			Code:  d.String("code"),
		}
		MonModes[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonMode records", len(MonModes))
}
