package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

func armorLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	if r.Item.Armors != nil {
		return nil // already loaded
	}

	records, err := loadCommonItems(d, d2enum.InventoryItemTypeArmor)
	if err != nil {
		return err
	}

	r.Logger.Infof("Loaded %d armors", len(records))

	r.Item.Armors = records

	return nil
}
