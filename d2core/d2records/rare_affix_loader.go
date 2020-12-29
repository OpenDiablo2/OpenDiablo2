package d2records

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

const (
	numRareAffixInclude = 7
	fmtRareAffixInclude = "itype%d"

	numRareAffixExclude = 4
	fmtRareAffixExclude = "etype%d"
)

func rareItemAffixLoader(d *d2txt.DataDictionary) ([]*RareItemAffix, error) {
	records := make([]*RareItemAffix, 0)

	for d.Next() {
		record := &RareItemPrefixRecord{
			Name:          d.String("name"),
			IncludedTypes: make([]string, 0),
			ExcludedTypes: make([]string, 0),
		}

		for idx := 1; idx <= numRareAffixInclude; idx++ {
			column := fmt.Sprintf(fmtRareAffixInclude, idx)
			if typeCode := d.String(column); typeCode != "" {
				record.IncludedTypes = append(record.IncludedTypes, typeCode)
			}
		}

		for idx := 1; idx <= numRareAffixExclude; idx++ {
			column := fmt.Sprintf(fmtRareAffixExclude, idx)
			if typeCode := d.String(column); typeCode != "" {
				record.ExcludedTypes = append(record.ExcludedTypes, typeCode)
			}
		}

		records = append(records, record)
	}

	return records, d.Err
}
