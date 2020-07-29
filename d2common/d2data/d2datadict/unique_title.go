package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// UniqueTitleRecord describes the different titles. Not listed on Phrozen Keep.
type UniqueTitleRecord struct {
	Namco string // Names such as Judge, Magistrate, and Count
}

// UniqueTitles contains the UniqueTitle records
//nolint:gochecknoglobals // Currently global by design, only written once
var UniqueTitles map[string]*UniqueTitleRecord

// LoadUniqueTitles loads UniqueTitles from the supplied file
func LoadUniqueTitles(file []byte) {
	UniqueTitles = make(map[string]*UniqueTitleRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &UniqueTitleRecord{
			Namco: d.String("Namco"),
		}
		UniqueTitles[record.Namco] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d UniqueTitle records", len(UniqueTitles))
}
