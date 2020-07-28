package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// MonTypeRecord is a representation of a single row of MonType.txt.
type MonTypeRecord struct {
	Type   string
	Equiv1 string
	Equiv2 string
	Equiv3 string
	// StrSing is the string displayed for the singular form (Skeleton), note
	// that this is unused in the original engine, since the only modifier
	// display code that accesses MonType uses StrPlur.
	StrSing   string
	StrPlural string
	// EOL int // unused
}

// MonTypes stores all of the MonTypeRecords
var MonTypes map[string]*MonTypeRecord //nolint:gochecknoglobals // Currently global by design, only written once

// LoadMonTypes loads MonType records into a map[string]*MonTypeRecord
func LoadMonTypes(file []byte) {
	MonTypes = make(map[string]*MonTypeRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &MonTypeRecord{
			Type:      d.String("type"),
			Equiv1:    d.String("equiv1"),
			Equiv2:    d.String("equiv2"),
			Equiv3:    d.String("equiv3"),
			StrSing:   d.String("strsing"),
			StrPlural: d.String("strplur"),
		}
		MonTypes[record.Type] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonType records", len(MonTypes))
}
