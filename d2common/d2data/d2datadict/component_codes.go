package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
	"log"
)

// ComponentCodeRecord represents a single row from compcode.txt
type ComponentCodeRecord struct {
	Component string
	Code      string
}

// ComponentCodes is a lookup table for DCC Animation Component Subtype,
// it links hardcoded data with the txt files
var ComponentCodes map[string]*ComponentCodeRecord //nolint:gochecknoglobals // Currently global by design

// LoadComponentCodes loads components code records from compcode.txt
func LoadComponentCodes(file []byte) {
	ComponentCodes = make(map[string]*ComponentCodeRecord)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		record := &ComponentCodeRecord{
			Component: d.String("component"),
			Code:      d.String("code"),
		}
		ComponentCodes[record.Component] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d ComponentCode records", len(ComponentCodes))
}
