package d2datadict

import (
	"log"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// QualityRecord represents a single row of QualityItems.txt, which controls
// properties for superior quality items
type QualityRecord struct {
	NumMods   int
	Mod1Code  string
	Mod1Param int
	Mod1Min   int
	Mod1Max   int
	Mod2Code  string
	Mod2Param int
	Mod2Min   int
	Mod2Max   int

	// The following fields determine this row's applicability to
	// categories of item.
	Armor   bool
	Weapon  bool
	Shield  bool
	Thrown  bool
	Scepter bool
	Wand    bool
	Staff   bool
	Bow     bool
	Boots   bool
	Gloves  bool
	Belt    bool

	Level    int
	Multiply int
	Add      int
}

// QualityItems stores all of the QualityRecords
var QualityItems map[string]*QualityRecord //nolint:gochecknoglobals // Currently global by design, only written once

// LoadQualityItems loads QualityItem records into a map[string]*QualityRecord
func LoadQualityItems(file []byte) {
	QualityItems = make(map[string]*QualityRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		qual := &QualityRecord{
			NumMods:   d.Number("nummods"),
			Mod1Code:  d.String("mod1code"),
			Mod1Param: d.Number("mod1param"),
			Mod1Min:   d.Number("mod1min"),
			Mod1Max:   d.Number("mod1max"),
			Mod2Code:  d.String("mod2code"),
			Mod2Param: d.Number("mod2param"),
			Mod2Min:   d.Number("mod2min"),
			Mod2Max:   d.Number("mod2max"),
			Armor:     d.Bool("armor"),
			Weapon:    d.Bool("weapon"),
			Shield:    d.Bool("shield"),
			Thrown:    d.Bool("thrown"),
			Scepter:   d.Bool("scepter"),
			Wand:      d.Bool("wand"),
			Staff:     d.Bool("staff"),
			Bow:       d.Bool("bow"),
			Boots:     d.Bool("boots"),
			Gloves:    d.Bool("gloves"),
			Belt:      d.Bool("belt"),
			Level:     d.Number("level"),
			Multiply:  d.Number("multiply"),
			Add:       d.Number("add"),
		}

		QualityItems[strconv.Itoa(len(QualityItems))] = qual
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d QualityItems records", len(QualityItems))
}
