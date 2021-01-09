package d2records

import (
	"fmt"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

//	Name
//	MonType1
//	MonType2
//	MonType3
//	MonType4
//	MonType5
//	MonType6
//	MonType7
//	MonType8
//	MonType9
//	MonType10
//	MonType11
//	MonType12
//	MonType13
//	MonType14
//	MonType15
//	MonType16
//	MonType17
//	MonType18
//	MonType19
//	MonType20
//	MonType21
//	MonType22
//	MonType23
//	MonType24
//	MonType25
//	MonType26
//	MonType27
//	MonType28
//	MonType29
//	MonType30
//	MonType31
//	MonType32
//	MonType33
//	MonType34
//	MonType35
//	MonType36

const (
	numMonsterTypes      = 36
	fmtMonsterTypeColumn = "MonType%d"
)

func uniqueMonsterPrefixLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records, err := uniqueMonsterAffixCommonLoader(d)
	if err != nil {
		return err
	}

	r.Monster.Name.Prefix = records

	r.Logger.Infof("Loaded %d unique monster prefix records", len(records))

	return nil
}

func uniqueMonsterSuffixLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records, err := uniqueMonsterAffixCommonLoader(d)
	if err != nil {
		return err
	}

	r.Monster.Name.Suffix = records

	r.Logger.Infof("Loaded %d unique monster suffix records", len(records))

	return nil
}

func uniqueMonsterAffixCommonLoader(d *d2txt.DataDictionary) (UniqueMonsterAffixes, error) {
	records := make(UniqueMonsterAffixes)

	for d.Next() {
		record := &UniqueMonsterAffixRecord{
			StringTableKey:   d.String("Name"),
			MonsterTypeFlags: akara.NewBitSet(),
		}

		for idx := 0; idx < numMonsterTypes; idx++ {
			bit := d.Number(fmt.Sprintf(fmtMonsterTypeColumn, idx)) > 0
			record.MonsterTypeFlags.Set(idx, bit)
		}

		records[record.StringTableKey] = record
	}

	return records, d.Err
}
