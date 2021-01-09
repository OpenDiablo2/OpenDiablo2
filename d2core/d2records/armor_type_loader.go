package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func armorTypesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(ArmorTypes)

	for d.Next() {
		record := &ArmorTypeRecord{
			Name:  d.String("Name"),
			Token: d.String("Token"),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Animation.Token.Armor = records

	r.Logger.Infof("Loaded %d ArmorType records", len(records))

	return nil
}
