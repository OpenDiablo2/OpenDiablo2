package d2records

// Books stores all of the BookRecords
type Books map[string]*BookRecord

// BookRecord is a representation of a row from books.txt
type BookRecord struct {
	Name            string
	Namco           string
	Completed       string
	ScrollSpellCode string
	BookSpellCode   string
	ScrollSkill     string
	BookSkill       string
	Pspell          int
	SpellIcon       int
	BaseCost        int
	CostPerCharge   int
}
