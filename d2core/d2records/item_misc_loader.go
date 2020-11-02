package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// LoadMiscItems loads ItemCommonRecords from misc.txt
func miscItemsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records, err := loadCommonItems(d, d2enum.InventoryItemTypeItem)
	if err != nil {
		return err
	}

	r.Logger.Infof("Loaded %d misc items", len(records))

	r.Item.Misc = records

	return nil
}
