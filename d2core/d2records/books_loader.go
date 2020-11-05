package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func booksLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(Books)

	for d.Next() {
		record := &BooksRecord{
			Name:            d.String("Name"),
			Namco:           d.String("Namco"),
			Completed:       d.String("Completed"),
			ScrollSpellCode: d.String("ScrollSpellCode"),
			BookSpellCode:   d.String("BooksSpellCode"),
			Pspell:          d.Number("pSpell"),
			SpellIcon:       d.Number("SpellIcon"),
			ScrollSkill:     d.String("ScrollSkill"),
			BookSkill:       d.String("BookSkill"),
			BaseCost:        d.Number("BaseCost"),
			CostPerCharge:   d.Number("CostPerCharge"),
		}
		records[record.Namco] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Logger.Infof("Loaded %d book items", len(records))

	r.Item.Books = records

	return nil
}
