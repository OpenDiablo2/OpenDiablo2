package d2records

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func monsterTypesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(MonsterTypes)

	for d.Next() {
		record := &MonTypeRecord{
			Type:      d.String("type"),
			Equiv1:    d.String("equiv1"),
			Equiv2:    d.String("equiv2"),
			Equiv3:    d.String("equiv3"),
			StrSing:   d.String("strsing"),
			StrPlural: d.String("strplur"),
		}
		records[record.Type] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonType records", len(records))

	r.Monster.Types = records

	return nil
}
