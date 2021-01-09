package d2records

// Books stores all of the BooksRecords
type Books map[string]*BooksRecord

// BooksRecord is a representation of a row from books.txt
type BooksRecord struct {
	Name            string
	Namco           string // The displayed name, where the string prefix is "Tome"
	Completed       string
	ScrollSpellCode string
	BookSpellCode   string
	Pspell          int
	SpellIcon       int
	ScrollSkill     string
	BookSkill       string
	BaseCost        int
	CostPerCharge   int
}
