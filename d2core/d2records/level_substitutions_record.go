package d2records

// LevelSubstitutions stores all of the LevelSubstitutionRecords
type LevelSubstitutions map[int]*LevelSubstitutionRecord

// LevelSubstitutionRecord is a representation of a row from lvlsub.txt
// these records are parameters for levels and describe substitution rules
type LevelSubstitutionRecord struct {
	File         string
	Name         string
	ChanceSpawn4 int
	GridMax3     int
	BorderType   int
	GridSize     int
	Mask         int
	ChanceSpawn0 int
	ChanceSpawn1 int
	ChanceSpawn2 int
	ChanceSpawn3 int
	ID           int
	ChanceFloor0 int
	ChanceFloor1 int
	ChanceFloor2 int
	ChanceFloor3 int
	ChanceFloor4 int
	GridMax0     int
	GridMax1     int
	GridMax2     int
	GridMax4     int
	IsExpansion  bool
}
