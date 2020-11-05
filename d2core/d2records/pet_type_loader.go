package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func petTypesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(PetTypes)

	for d.Next() {
		record := &PetTypeRecord{
			Name:      d.String("pet type"),
			ID:        d.Number("idx"),
			GroupID:   d.Number("group"),
			BaseMax:   d.Number("basemax"),
			Warp:      d.Bool("warp"),
			Range:     d.Bool("range"),
			PartySend: d.Bool("partysend"),
			Unsummon:  d.Bool("unsummon"),
			Automap:   d.Bool("automap"),
			IconName:  d.String("name"),
			DrawHP:    d.Bool("drawhp"),
			IconType:  d.Number("icontype"),
			BaseIcon:  d.String("baseicon"),
			MClass1:   d.Number("mclass1"),
			MIcon1:    d.String("micon1"),
			MClass2:   d.Number("mclass2"),
			MIcon2:    d.String("micon2"),
			MClass3:   d.Number("mclass3"),
			MIcon3:    d.String("micon3"),
			MClass4:   d.Number("mclass4"),
			MIcon4:    d.String("micon4"),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d PetType records", len(records))

	r.PetTypes = records

	return nil
}
