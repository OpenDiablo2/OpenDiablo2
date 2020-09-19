package d2records

import (
	"fmt"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

const (
	numRarePrefixInclude = 7
	fmtRarePrefixInclude = "itype%d"

	numRarePrefixExclude = 4
	fmtRarePrefixExclude = "etype%d"
)

func rareItemPrefixLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(RarePrefixes, 0)

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

		records = append(records, record)
	}

	if d.Err != nil {
		return d.Err
	}

	r.Item.Rare.Prefix = records

	log.Printf("Loaded %d RarePrefix records", len(records))

	return nil
}
