package d2datadict

import (
	"fmt"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

const (
	numRarePrefixInclude = 7
	fmtRarePrefixInclude = "itype%d"

	numRarePrefixExclude = 4
	fmtRarePrefixExclude = "etype%d"
)

// RareItemPrefixRecord is a name prefix for rare items (items with more than 2 affixes)
type RareItemPrefixRecord struct {
	Name          string
	IncludedTypes []string
	ExcludedTypes []string
}

// RarePrefixes is where all RareItemPrefixRecords are stored
var RarePrefixes map[string]*RareItemPrefixRecord // nolint:gochecknoglobals // global by design

// LoadRareItemPrefixRecords loads the rare item prefix records from rareprefix.txt
func LoadRareItemPrefixRecords(file []byte) {
	d := d2common.LoadDataDictionary(file)

	RarePrefixes = make(map[string]*RareItemPrefixRecord)

	for d.Next() {
		record := &RareItemPrefixRecord{
			Name:          d.String("name"),
			IncludedTypes: make([]string, 0),
			ExcludedTypes: make([]string, 0),
		}

		for idx := 1; idx <= numRarePrefixInclude; idx++ {
			column := fmt.Sprintf(fmtRarePrefixInclude, idx)
			if typeCode := d.String(column); typeCode != "" {
				record.IncludedTypes = append(record.IncludedTypes, typeCode)
			}
		}

		for idx := 1; idx <= numRarePrefixExclude; idx++ {
			column := fmt.Sprintf(fmtRarePrefixExclude, idx)
			if typeCode := d.String(column); typeCode != "" {
				record.ExcludedTypes = append(record.ExcludedTypes, typeCode)
			}
		}

		RarePrefixes[record.Name] = record
	}

	log.Printf("Loaded %d RarePrefix records", len(RarePrefixes))
}
