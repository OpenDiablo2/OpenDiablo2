package d2records

import (
	"fmt"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

const (
	numRareSuffixInclude = 7
	fmtRareSuffixInclude = "itype%d"

	numRareSuffixExclude = 4
	fmtRareSuffixExclude = "etype%d"
)

func rareItemSuffixLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make([]*RareItemSuffixRecord, 0)

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

		records = append(records, record)
	}

	if d.Err != nil {
		return d.Err
	}

	log.Printf("Loaded %d RareSuffix records", len(records))

	r.Item.Rare.Suffix = records

	return nil
}
