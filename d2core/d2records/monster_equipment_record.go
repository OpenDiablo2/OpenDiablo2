package d2records

const (
	numMonEquippedItems = 3
	fmtLocation         = "loc%d"
	fmtQuality          = "mod%d"
	fmtCode             = "item%d"
)

// MonsterEquipment stores the MonsterEquipmentRecords
type MonsterEquipment map[string][]*MonsterEquipmentRecord

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
