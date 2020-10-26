package d2records

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func rareItemPrefixLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records, err := rareItemAffixLoader(d)
	if err != nil {
		return err
	}

	r.Item.Rare.Prefix = records

	log.Printf("Loaded %d RarePrefix records", len(records))

	return nil
}
