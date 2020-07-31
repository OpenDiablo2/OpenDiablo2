package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

//MonsterEquipmentRecord represents a single line in monequip.txt
//Information gathered from [https://d2mods.info/forum/kb/viewarticle?a=365]
type MonsterEquipmentRecord struct {
	//Name of monster, pointer to MonStats.txt
	Monster string

	//If true, monster is created by level, otherwise created by skill
	OnInit bool

	//Not written in description, only appear on monsters with OnInit false
	//Level of skill for which this equipment row can be used?
	Level int

	//Code of item
	Item1 string

	//Slot of equipped item
	Location1 string

	//Item quality
	Mod1 int

	//Ditto, 3 items maximum
	Item2     string
	Location2 string
	Mod2      int

	Item3     string
	Location3 string
	Mod3      int
}

//MonsterEquipment stores the MonsterEquipmentRecords
var MonsterEquipment map[string][]*MonsterEquipmentRecord //nolint:gochecknoglobals // Currently global by design

//LoadMonsterEquipment loads MonsterEquipmentRecords into MonsterEquipment
func LoadMonsterEquipment(file []byte) {
	MonsterEquipment = make(map[string][]*MonsterEquipmentRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &MonsterEquipmentRecord{
			Monster:   d.String("monster"),
			OnInit:    d.Bool("oninit"),
			Level:     d.Number("level"),
			Item1:     d.String("item1"),
			Location1: d.String("loc1"),
			Mod1:      d.Number("mod1"),
			Item2:     d.String("item2"),
			Location2: d.String("loc2"),
			Mod2:      d.Number("mod2"),
			Item3:     d.String("item3"),
			Location3: d.String("loc3"),
			Mod3:      d.Number("mod3"),
		}

		if _, ok := MonsterEquipment[record.Monster]; !ok {
			MonsterEquipment[record.Monster] = make([]*MonsterEquipmentRecord, 0)
		}
		MonsterEquipment[record.Monster] = append(MonsterEquipment[record.Monster], record)
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
