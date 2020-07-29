package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// BooksRecord is a representation of a row from books.txt
type BooksRecord struct {
	Name            string
	Namco           string // The displayed name, where the string prefix is "Tome"
	Completed       string
	ScrollSpellCode string
	BookSpellCode   string
	pSpell          int
	SpellIcon       int
	ScrollSkill     string
	BookSkill       string
	BaseCost        int
	CostPerCharge   int
}

// Books stores all of the BooksRecords
var Books map[string]*BooksRecord //nolint:gochecknoglobals // Currently global by design, only written once

// LoadBooks loads Books records into a map[string]*BooksRecord
func LoadBooks(file []byte) {
	Books = make(map[string]*BooksRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &BooksRecord{
			Name:            d.String("Name"),
			Namco:           d.String("Namco"),
			Completed:       d.String("Completed"),
			ScrollSpellCode: d.String("ScrollSpellCode"),
			BookSpellCode:   d.String("BooksSpellCode"),
			pSpell:          d.Number("pSpell"),
			SpellIcon:       d.Number("SpellIcon"),
			ScrollSkill:     d.String("ScrollSkill"),
			BookSkill:       d.String("BookSkill"),
			BaseCost:        d.Number("BaseCost"),
			CostPerCharge:   d.Number("CostPerCharge"),
		}
		Books[record.Namco] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d book items", len(Books))
}
