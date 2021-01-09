package d2records

import (
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func itemQualityLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(ItemQualities)

	for d.Next() {
		qual := &ItemQualityRecord{
			NumMods:   d.Number("nummods"),
			Mod1Code:  d.String("mod1code"),
			Mod1Param: d.Number("mod1param"),
			Mod1Min:   d.Number("mod1min"),
			Mod1Max:   d.Number("mod1max"),
			Mod2Code:  d.String("mod2code"),
			Mod2Param: d.Number("mod2param"),
			Mod2Min:   d.Number("mod2min"),
			Mod2Max:   d.Number("mod2max"),
			Armor:     d.Bool("armor"),
			Weapon:    d.Bool("weapon"),
			Shield:    d.Bool("shield"),
			Thrown:    d.Bool("thrown"),
			Scepter:   d.Bool("scepter"),
			Wand:      d.Bool("wand"),
			Staff:     d.Bool("staff"),
			Bow:       d.Bool("bow"),
			Boots:     d.Bool("boots"),
			Gloves:    d.Bool("gloves"),
			Belt:      d.Bool("belt"),
			Level:     d.Number("level"),
			Multiply:  d.Number("multiply"),
			Add:       d.Number("add"),
		}

		records[strconv.Itoa(len(records))] = qual
	}

	if d.Err != nil {
		return d.Err
	}

	r.Item.Quality = records

	r.Logger.Infof("Loaded %d ItemQualities records", len(records))

	return nil
}
