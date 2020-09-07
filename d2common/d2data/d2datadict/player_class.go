package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
	"log"
)

// PlayerClassRecord represents a single line from PlayerClass.txt
// Lookup table for class codes
type PlayerClassRecord struct {
	// Name of the player class
	Name string

	// Code for the player class
	Code string
}

// PlayerClasses stores the PlayerClassRecords
var PlayerClasses map[string]*PlayerClassRecord // nolint:gochecknoglobals // Currently global by design

// LoadPlayerClasses loads the PlayerClassRecords into PlayerClasses
func LoadPlayerClasses(file []byte) {
	PlayerClasses = make(map[string]*PlayerClassRecord)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		record := &PlayerClassRecord{
			Name: d.String("Player Class"),
			Code: d.String("Code"),
		}

		if record.Name == expansion {
			continue
		}

		PlayerClasses[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d PlayerClass records", len(PlayerClasses))
}
