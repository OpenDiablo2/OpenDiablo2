package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
	"log"
)

// ElemTypeRecord represents a single line in ElemType.txt
type ElemTypeRecord struct {
	// ElemType Elemental damage type name
	ElemType string

	// Code Elemental damage type code
	Code string
}

// ElemTypes stores the ElemTypeRecords
var ElemTypes map[string]*ElemTypeRecord //nolint:gochecknoglobals // Currently global by design

// LoadElemTypes loads ElemTypeRecords into ElemTypes
func LoadElemTypes(file []byte) {
	ElemTypes = make(map[string]*ElemTypeRecord)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		record := &ElemTypeRecord{
			ElemType: d.String("Elemental Type"),
			Code:     d.String("Code"),
		}
		ElemTypes[record.ElemType] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d ElemType records", len(ElemTypes))
}
