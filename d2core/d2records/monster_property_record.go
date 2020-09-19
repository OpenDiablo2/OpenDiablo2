package d2records

const (
	NumMonProps  = 6
	FmtProp      = "prop%d%s"
	FmtChance    = "chance%d%s"
	FmtPar       = "par%d%s"
	FmtMin       = "min%d%s"
	FmtMax       = "max%d%s"
	FmtNormal    = ""
	FmtNightmare = " (N)"
	FmtHell      = " (H)"
)

// MonsterProperties stores all of the MonPropRecords
type MonsterProperties map[string]*MonPropRecord

// MonPropRecord is a representation of a single row of monprop.txt
type MonPropRecord struct {
	ID string

	Properties struct {
		Normal    [NumMonProps]*MonProp
		Nightmare [NumMonProps]*MonProp
		Hell      [NumMonProps]*MonProp
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
