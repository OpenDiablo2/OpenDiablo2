package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func weaponClassesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(WeaponClasses)

	for d.Next() {
		record := &WeaponClassRecord{
			Name:  d.String("Weapon Class"),
			Token: d.String("Code"),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Animation.Token.Weapon = records

	r.Logger.Infof("Loaded %d WeaponClass records", len(records))

	return nil
}
