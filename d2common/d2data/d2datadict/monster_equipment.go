package d2datadict

import (
	"fmt"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

const (
	numMonEquippedItems = 3
	fmtLocation         = "loc%d"
	fmtQuality          = "mod%d"
	fmtCode             = "item%d"
)

// MonsterEquipmentRecord represents a single line in monequip.txt
// Information gathered from [https://d2mods.info/forum/kb/viewarticle?a=365]
type MonsterEquipmentRecord struct {
	// Name of monster, pointer to MonStats.txt
	Name string

	// If true, monster is created by level, otherwise created by skill
	OnInit bool

	// Not written in description, only appear on monsters with OnInit false,
	// Level of skill for which this equipment row can be used?
	Level int

	Equipment []*monEquip
}

type monEquip struct {
	// Code of item, probably from ItemCommonRecords
	Code string

	// Location the body location of the item
	Location string

	// Quality of the item
	Quality int
}

// MonsterEquipment stores the MonsterEquipmentRecords
var MonsterEquipment map[string][]*MonsterEquipmentRecord //nolint:gochecknoglobals // Currently global by design

// LoadMonsterEquipment loads MonsterEquipmentRecords into MonsterEquipment
func LoadMonsterEquipment(file []byte) {
	MonsterEquipment = make(map[string][]*MonsterEquipmentRecord)

	d := d2common.LoadDataDictionary(file)
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

		if _, ok := MonsterEquipment[record.Name]; !ok {
			MonsterEquipment[record.Name] = make([]*MonsterEquipmentRecord, 0)
		}

		MonsterEquipment[record.Name] = append(MonsterEquipment[record.Name], record)
	}

	if d.Err != nil {
		panic(d.Err)
	}

	length := 0
	for k := range MonsterEquipment {
		length += len(MonsterEquipment[k])
	}

	log.Printf("Loaded %d MonsterEquipment records", length)
}
