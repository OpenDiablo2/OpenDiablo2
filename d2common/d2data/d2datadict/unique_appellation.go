package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type UniqueAppellationRecord struct {
	// The title
	Name string

	// The following booleans assign the title for a specific monster type
	MonType1  bool
	MonType2  bool
	MonType3  bool
	MonType4  bool
	MonType5  bool
	MonType6  bool
	MonType7  bool
	MonType8  bool
	MonType9  bool
	MonType10 bool
	MonType11 bool
	MonType12 bool
	MonType13 bool
	MonType14 bool
	MonType15 bool
	MonType16 bool
	MonType17 bool
	MonType18 bool
	MonType19 bool
	MonType20 bool
	MonType21 bool
	MonType22 bool
	MonType23 bool
	MonType24 bool
	MonType25 bool
	MonType26 bool
	MonType27 bool
	MonType28 bool
	MonType29 bool
	MonType30 bool
	MonType31 bool
	MonType32 bool
	MonType33 bool
	MonType34 bool
	MonType35 bool
	MonType36 bool
}

var UniqueAppellations map[string]*UniqueAppellationRecord

func LoadUniqueAppellations(file []byte) {
	UniqueAppellations = make(map[string]*UniqueAppellationRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &UniqueAppellationRecord{
			Name:      d.String("Name"),
			MonType1:  d.Bool("MonType1"),
			MonType2:  d.Bool("MonType2"),
			MonType3:  d.Bool("MonType3"),
			MonType4:  d.Bool("MonType4"),
			MonType5:  d.Bool("MonType5"),
			MonType6:  d.Bool("MonType6"),
			MonType7:  d.Bool("MonType7"),
			MonType8:  d.Bool("MonType8"),
			MonType9:  d.Bool("MonType9"),
			MonType10: d.Bool("MonType10"),
			MonType11: d.Bool("MonType11"),
			MonType12: d.Bool("MonType12"),
			MonType13: d.Bool("MonType13"),
			MonType14: d.Bool("MonType14"),
			MonType15: d.Bool("MonType15"),
			MonType16: d.Bool("MonType16"),
			MonType17: d.Bool("MonType17"),
			MonType18: d.Bool("MonType18"),
			MonType19: d.Bool("MonType19"),
			MonType20: d.Bool("MonType20"),
			MonType21: d.Bool("MonType21"),
			MonType22: d.Bool("MonType22"),
			MonType23: d.Bool("MonType23"),
			MonType24: d.Bool("MonType24"),
			MonType25: d.Bool("MonType25"),
			MonType26: d.Bool("MonType26"),
			MonType27: d.Bool("MonType27"),
			MonType28: d.Bool("MonType28"),
			MonType29: d.Bool("MonType29"),
			MonType30: d.Bool("MonType30"),
			MonType31: d.Bool("MonType31"),
			MonType32: d.Bool("MonType32"),
			MonType33: d.Bool("MonType33"),
			MonType34: d.Bool("MonType34"),
			MonType35: d.Bool("MonType35"),
			MonType36: d.Bool("MonType36"),
		}
		UniqueAppellations[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d UniqueAppellation records", len(UniqueAppellations))
}
