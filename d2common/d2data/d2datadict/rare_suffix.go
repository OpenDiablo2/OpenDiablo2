package d2datadict

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
	"log"
)

const (
	numRareSuffixInclude = 7
	fmtRareSuffixInclude = "itype%d"

	numRareSuffixExclude = 4
	fmtRareSuffixExclude = "etype%d"
)

// RareItemSuffixRecord is a name suffix for rare items (items with more than 2 affixes)
type RareItemSuffixRecord struct {
	Name          string
	IncludedTypes []string
	ExcludedTypes []string
}

// RareSuffixes is where all RareItemSuffixRecords are stored
var RareSuffixes []*RareItemSuffixRecord // nolint:gochecknoglobals // global by design

// LoadRareItemSuffixRecords loads the rare item suffix records from raresuffix.txt
func LoadRareItemSuffixRecords(file []byte) {
	d := d2txt.LoadDataDictionary(file)

	RareSuffixes = make([]*RareItemSuffixRecord, 0)

	for d.Next() {
		record := &RareItemSuffixRecord{
			Name:          d.String("name"),
			IncludedTypes: make([]string, 0),
			ExcludedTypes: make([]string, 0),
		}

		for idx := 1; idx <= numRareSuffixInclude; idx++ {
			column := fmt.Sprintf(fmtRareSuffixInclude, idx)
			if typeCode := d.String(column); typeCode != "" {
				record.IncludedTypes = append(record.IncludedTypes, typeCode)
			}
		}

		for idx := 1; idx <= numRareSuffixExclude; idx++ {
			column := fmt.Sprintf(fmtRareSuffixExclude, idx)
			if typeCode := d.String(column); typeCode != "" {
				record.ExcludedTypes = append(record.ExcludedTypes, typeCode)
			}
		}

		RareSuffixes = append(RareSuffixes, record)
	}

	log.Printf("Loaded %d RareSuffix records", len(RareSuffixes))
}
