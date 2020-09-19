package d2records

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func uniqueItemsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(UniqueItems)

	for d.Next() {
		record := &UniqueItemRecord{
			Name:    d.String("index"),
			Version: d.Number("version"),
			Enabled: d.Number("enabled") == 1,

			Ladder:  d.Number("ladder") == 1,
			Rarity:  d.Number("rarity"),
			NoLimit: d.Number("nolimit") == 1,

			Level:         d.Number("lvl"),
			RequiredLevel: d.Number("lvl req"),
			Code:          d.String("code"),

			TypeDescription: d.String("*type"),
			UberDescription: d.String("*uber"),
			SingleCopy:      d.Number("carry1") == 1,
			CostMultiplier:  d.Number("cost mult"),
			CostAdd:         d.Number("cost add"),

			CharacterGfxTransform: d.String("chrtransform"),
			InventoryGfxTransform: d.String("invtransform"),
			FlippyFile:            d.String("flippyfile"),
			InventoryFile:         d.String("invfile"),

			DropSound:    d.String("dropsound"),
			DropSfxFrame: d.Number("dropsfxframe"),
			UseSound:     d.String("usesound"),

			Properties: [12]UniqueItemProperty{
				UniqueItemProperty{
					Code:      d.String("prop1"),
					Parameter: d.String("par1"),
					Min:       d.Number("min1"),
					Max:       d.Number("max1"),
				},
				UniqueItemProperty{
					Code:      d.String("prop2"),
					Parameter: d.String("par2"),
					Min:       d.Number("min2"),
					Max:       d.Number("max2"),
				},
				UniqueItemProperty{
					Code:      d.String("prop3"),
					Parameter: d.String("par3"),
					Min:       d.Number("min3"),
					Max:       d.Number("max3"),
				},
				UniqueItemProperty{
					Code:      d.String("prop4"),
					Parameter: d.String("par4"),
					Min:       d.Number("min4"),
					Max:       d.Number("max4"),
				},

				UniqueItemProperty{
					Code:      d.String("prop5"),
					Parameter: d.String("par5"),
					Min:       d.Number("min5"),
					Max:       d.Number("max5"),
				},
				UniqueItemProperty{
					Code:      d.String("prop6"),
					Parameter: d.String("par6"),
					Min:       d.Number("min6"),
					Max:       d.Number("max6"),
				},
				UniqueItemProperty{
					Code:      d.String("prop7"),
					Parameter: d.String("par7"),
					Min:       d.Number("min7"),
					Max:       d.Number("max7"),
				},
				UniqueItemProperty{
					Code:      d.String("prop8"),
					Parameter: d.String("par8"),
					Min:       d.Number("min8"),
					Max:       d.Number("max8"),
				},

				UniqueItemProperty{
					Code:      d.String("prop9"),
					Parameter: d.String("par9"),
					Min:       d.Number("min9"),
					Max:       d.Number("max9"),
				},
				UniqueItemProperty{
					Code:      d.String("prop10"),
					Parameter: d.String("par10"),
					Min:       d.Number("min10"),
					Max:       d.Number("max10"),
				},
				UniqueItemProperty{
					Code:      d.String("prop11"),
					Parameter: d.String("par11"),
					Min:       d.Number("min11"),
					Max:       d.Number("max11"),
				},
				UniqueItemProperty{
					Code:      d.String("prop12"),
					Parameter: d.String("par12"),
					Min:       d.Number("min12"),
					Max:       d.Number("max12"),
				},
			},
		}

		if record.Name == "" {
			continue
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Item.Unique = records

	log.Printf("Loaded %d unique items", len(records))

	return nil
}
