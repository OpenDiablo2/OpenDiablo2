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
	Name      string
	Equipment []*monEquip
	Level     int
	OnInit    bool
}

type monEquip struct {
	// Code of item, probably from ItemCommonRecords
	Code string

	// Location the body location of the item
	Location string

	// Quality of the item
	Quality int
}
