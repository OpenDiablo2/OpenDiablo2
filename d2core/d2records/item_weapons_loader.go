package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// LoadWeapons loads weapon records
func weaponsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records, err := loadCommonItems(d, d2enum.InventoryItemTypeWeapon)
	if err != nil {
		return err
	}

	r.Logger.Infof("Loaded %d weapons", len(records))

	r.Item.Weapons = records

	return nil
}
