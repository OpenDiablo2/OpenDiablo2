package d2records

const (
	numMonProps  = 6
	fmtProp      = "prop%d%s"
	fmtChance    = "chance%d%s"
	fmtPar       = "par%d%s"
	fmtMin       = "min%d%s"
	fmtMax       = "max%d%s"
	fmtNormal    = ""
	fmtNightmare = " (N)"
	fmtHell      = " (H)"
)

// MonsterProperties stores all of the MonPropRecords
type MonsterProperties map[string]*MonPropRecord

// MonPropRecord is a representation of a single row of monprop.txt
type MonPropRecord struct {
	ID string

	Properties struct {
		Normal    [numMonProps]*MonProp
		Nightmare [numMonProps]*MonProp
		Hell      [numMonProps]*MonProp
	}
}

// MonProp is a monster property
type MonProp struct {
	Code   string
	Param  string
	Chance int
	Min    int
	Max    int
}
