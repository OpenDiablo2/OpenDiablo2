package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// RunesRecord is a representation of a single row of runes.txt. It defines
// runewords available in the game.
type RunesRecord struct {
	Name     string
	RuneName string // More of a note - the actual name should be read from the TBL files.
	Complete bool   // An enabled/disabled flag. Only "Complete" runewords work in game.
	Server   bool   // Marks a runeword as only available on ladder, not single player or tcp/ip.

	// IType1-6 include item types into the list of item types that this runeword works in.
	IType1 string
	IType2 string
	IType3 string
	IType4 string
	IType5 string
	IType6 string

	// EType1-3 exclude item types from the list of item types that this runeword works in.
	EType1 string
	EType2 string
	EType3 string

	Runes string // The runes as they would appear in the finished runeword. Note field only.

	// Rune1-6 are ID pointers from Misc.txt. The fields control what runes are
	// required to make the rune word and in what order they are to be socketed.
	Rune1 string
	Rune2 string
	Rune3 string
	Rune4 string
	Rune5 string
	Rune6 string

	T1Code1  string
	T1Param1 string
	T1Min1   int
	T1Max1   int

	T1Code2  string
	T1Param2 string
	T1Min2   int
	T1Max2   int

	T1Code3  string
	T1Param3 string
	T1Min3   int
	T1Max3   int

	T1Code4  string
	T1Param4 string
	T1Min4   int
	T1Max4   int

	T1Code5  string
	T1Param5 string
	T1Min5   int
	T1Max5   int

	T1Code6  string
	T1Param6 string
	T1Min6   int
	T1Max6   int

	T1Code7  string
	T1Param7 string
	T1Min7   int
	T1Max7   int

	// EOL int // not loaded
}

// Runes stores all of the RunesRecords
var Runes map[string]*RunesRecord //nolint:gochecknoglobals // Currently global by design, only written once

// LoadRunes loads runes records into a map[string]*RunesRecord
func LoadRunes(file []byte) {
	Runes = make(map[string]*RunesRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &RunesRecord{
			Name:     d.String("name"),
			RuneName: d.String("Rune Name"),
			Complete: d.Bool("complete"),
			Server:   d.Bool("server"),
			IType1:   d.String("itype1"),
			IType2:   d.String("itype2"),
			IType3:   d.String("itype3"),
			IType4:   d.String("itype4"),
			IType5:   d.String("itype5"),
			IType6:   d.String("itype6"),
			EType1:   d.String("etype1"),
			EType2:   d.String("etype2"),
			EType3:   d.String("etype3"),
			Runes:    d.String("*runes"),
			Rune1:    d.String("Rune1"),
			Rune2:    d.String("Rune2"),
			Rune3:    d.String("Rune3"),
			Rune4:    d.String("Rune4"),
			Rune5:    d.String("Rune5"),
			Rune6:    d.String("Rune6"),
			T1Code1:  d.String("T1Code1"),
			T1Param1: d.String("T1Param1"),
			T1Min1:   d.Number("T1Min1"),
			T1Max1:   d.Number("T1Max1"),
			T1Code2:  d.String("T1Code2"),
			T1Param2: d.String("T1Param2"),
			T1Min2:   d.Number("T1Min2"),
			T1Max2:   d.Number("T1Max2"),
			T1Code3:  d.String("T1Code3"),
			T1Param3: d.String("T1Param3"),
			T1Min3:   d.Number("T1Min3"),
			T1Max3:   d.Number("T1Max3"),
			T1Code4:  d.String("T1Code4"),
			T1Param4: d.String("T1Param4"),
			T1Min4:   d.Number("T1Min4"),
			T1Max4:   d.Number("T1Max4"),
			T1Code5:  d.String("T1Code5"),
			T1Param5: d.String("T1Param5"),
			T1Min5:   d.Number("T1Min5"),
			T1Max5:   d.Number("T1Max5"),
			T1Code6:  d.String("T1Code6"),
			T1Param6: d.String("T1Param6"),
			T1Min6:   d.Number("T1Min6"),
			T1Max6:   d.Number("T1Max6"),
			T1Code7:  d.String("T1Code7"),
			T1Param7: d.String("T1Param7"),
			T1Min7:   d.Number("T1Min7"),
			T1Max7:   d.Number("T1Max7"),
		}
		Runes[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d Runes records", len(Runes))
}
