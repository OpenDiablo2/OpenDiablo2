package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func autoMagicLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(AutoMagic, 0)

	charCodeMap := map[string]d2enum.Hero{
		"ama": d2enum.HeroAmazon,
		"ass": d2enum.HeroAssassin,
		"bar": d2enum.HeroBarbarian,
		"dru": d2enum.HeroDruid,
		"nec": d2enum.HeroNecromancer,
		"pal": d2enum.HeroPaladin,
		"sor": d2enum.HeroSorceress,
	}

	for d.Next() {
		record := &AutoMagicRecord{
			Name:                  d.String("Name"),
			Version:               d.Number("version"),
			Spawnable:             d.Number("spawnable") > 0,
			SpawnOnRare:           d.Number("rare") > 0,
			MinSpawnLevel:         d.Number("level"),
			MaxSpawnLevel:         d.Number("maxlevel"),
			LevelRequirement:      d.Number("levelreq"),
			Class:                 charCodeMap[d.String("class")],
			ClassLevelRequirement: d.Number("classlevelreq"),
			Frequency:             d.Number("frequency"),
			Group:                 d.Number("group"),
			ModCode: [3]string{
				d.String("mod1code"),
				d.String("mod2code"),
				d.String("mod3code"),
			},
			ModParam: [3]int{
				d.Number("mod1param"),
				d.Number("mod2param"),
				d.Number("mod3param"),
			},
			ModMin: [3]int{
				d.Number("mod1min"),
				d.Number("mod2min"),
				d.Number("mod3min"),
			},
			ModMax: [3]int{
				d.Number("mod1max"),
				d.Number("mod2max"),
				d.Number("mod3max"),
			},
			Transform:        d.Number("transform") > 0,
			PaletteTransform: d.Number("transformcolor"),
			IncludeItemCodes: [7]string{
				d.String("itype1"),
				d.String("itype2"),
				d.String("itype3"),
				d.String("itype4"),
				d.String("itype5"),
				d.String("itype6"),
				d.String("itype7"),
			},
			ExcludeItemCodes: [3]string{
				d.String("etype1"),
				d.String("etype2"),
				d.String("etype3"),
			},
			CostDivide:   d.Number("divide"),
			CostMultiply: d.Number("multiply"),
			CostAdd:      d.Number("add"),
		}

		records = append(records, record)
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d AutoMagic records", len(records))

	r.Item.AutoMagic = records

	return nil
}
