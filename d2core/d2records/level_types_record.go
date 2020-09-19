package d2records

// LevelTypes stores all of the LevelTypeRecords
type LevelTypes []*LevelTypeRecord

// LevelTypeRecord is a representation of a row from lvltype.txt
// the fields describe what ds1 files a level uses
type LevelTypeRecord struct {
	Files     [32]string
	Name      string
	ID        int
	Act       int
	Beta      bool
	Expansion bool
}
