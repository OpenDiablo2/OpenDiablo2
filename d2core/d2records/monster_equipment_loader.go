package d2records

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadMonsterEquipment loads MonsterEquipmentRecords into MonsterEquipment
func monsterEquipmentLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(map[string][]*MonsterEquipmentRecord)

	for d.Next() {
		record := &MonsterEquipmentRecord{
			Name:      d.String("monster"),
			OnInit:    d.Bool("oninit"),
			Level:     d.Number("level"),
			Equipment: make([]*monEquip, 0),
		}

		for idx := 0; idx < numMonEquippedItems; idx++ {
			num := idx + 1
			code := d.String(fmt.Sprintf(fmtCode, num))
			location := d.String(fmt.Sprintf(fmtLocation, num))
			quality := d.Number(fmt.Sprintf(fmtQuality, num))

			if code == "" {
				continue
			}

			equip := &monEquip{code, location, quality}

			record.Equipment = append(record.Equipment, equip)
		}

		if _, ok := records[record.Name]; !ok {
			records[record.Name] = make([]*MonsterEquipmentRecord, 0)
		}

		records[record.Name] = append(records[record.Name], record)
	}

	if d.Err != nil {
		return d.Err
	}

	length := 0
	for k := range records {
		length += len(records[k])
	}

	r.Logger.Infof("Loaded %d MonsterEquipment records", length)

	r.Monster.Equipment = records

	return nil
}
