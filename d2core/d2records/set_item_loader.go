package d2records

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

const (
	numPropertiesOnSetItem      = 9
	numBonusPropertiesOnSetItem = 5
	bonusToken1                 = "a"
	bonusToken2                 = "b"
	propCodeFmt                 = "prop%d"
	propParamFmt                = "par%d"
	propMinFmt                  = "min%d"
	propMaxFmt                  = "max%d"
	bonusCodeFmt                = "aprop%d%s"
	bonusParamFmt               = "apar%d%s"
	bonusMinFmt                 = "amin%d%s"
	bonusMaxFmt                 = "amax%d%s"
)

// SetItemProperty is describes a property of a set item
type SetItemProperty = PropertyDescriptor

// Loadrecords loads all of the SetItemRecords from records.txt
func setItemLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(map[string]*SetItemRecord)

	for d.Next() {
		record := &SetItemRecord{
			SetItemKey:                d.String("index"),
			SetKey:                    d.String("set"),
			ItemCode:                  d.String("item"),
			Rarity:                    d.Number("rarity"),
			QualityLevel:              d.Number("lvl"),
			RequiredLevel:             d.Number("lvl req"),
			CharacterPaletteTransform: d.Number("chrtransform"),
			InventoryPaletteTransform: d.Number("invtransform"),
			InvFile:                   d.String("invfile"),
			FlippyFile:                d.String("flippyfile"),
			DropSound:                 d.String("dropsound"),
			DropSfxFrame:              d.Number("dropsfxframe"),
			UseSound:                  d.String("usesound"),
			CostMult:                  d.Number("cost mult"),
			CostAdd:                   d.Number("cost add"),
			AddFn:                     d.Number("add func"),
		}

		// normal properties
		props := [numPropertiesOnSetItem]*SetItemProperty{}

		for idx := 0; idx < numPropertiesOnSetItem; idx++ {
			num := idx + 1
			props[idx] = &SetItemProperty{
				d.String(fmt.Sprintf(propCodeFmt, num)),
				d.String(fmt.Sprintf(propParamFmt, num)),
				d.Number(fmt.Sprintf(propMinFmt, num)),
				d.Number(fmt.Sprintf(propMaxFmt, num)),
			}
		}

		// set bonus properties
		bonus1 := [numBonusPropertiesOnSetItem]*SetItemProperty{}
		bonus2 := [numBonusPropertiesOnSetItem]*SetItemProperty{}

		for idx := 0; idx < numBonusPropertiesOnSetItem; idx++ {
			num := idx + 1

			bonus1[idx] = &SetItemProperty{
				d.String(fmt.Sprintf(bonusCodeFmt, num, bonusToken1)),
				d.String(fmt.Sprintf(bonusParamFmt, num, bonusToken1)),
				d.Number(fmt.Sprintf(bonusMinFmt, num, bonusToken1)),
				d.Number(fmt.Sprintf(bonusMaxFmt, num, bonusToken1)),
			}

			bonus2[idx] = &SetItemProperty{
				d.String(fmt.Sprintf(bonusCodeFmt, num, bonusToken2)),
				d.String(fmt.Sprintf(bonusParamFmt, num, bonusToken2)),
				d.Number(fmt.Sprintf(bonusMinFmt, num, bonusToken2)),
				d.Number(fmt.Sprintf(bonusMaxFmt, num, bonusToken2)),
			}
		}

		record.Properties = props

		records[record.SetItemKey] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Item.SetItems = records

	r.Logger.Infof("Loaded %d SetItem records", len(records))

	return nil
}
