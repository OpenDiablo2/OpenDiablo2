package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func rareItemSuffixLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records, err := rareItemAffixLoader(d)
	if err != nil {
		return err
	}

	r.Logger.Infof("Loaded %d RareSuffix records", len(records))

	r.Item.Rare.Suffix = records

	return nil
}
