package d2records

// SetItems holds all of the SetItemRecords
type SetItems map[string]*SetItemRecord

// SetItemRecord represents a set item
type SetItemRecord struct {
	Properties                [numPropertiesOnSetItem]*SetItemProperty
	SetPropertiesLevel2       [numBonusPropertiesOnSetItem]*SetItemProperty
	SetPropertiesLevel1       [numBonusPropertiesOnSetItem]*SetItemProperty
	InvFile                   string
	ItemCode                  string
	UseSound                  string
	DropSound                 string
	SetKey                    string
	FlippyFile                string
	SetItemKey                string
	InventoryPaletteTransform int
	CharacterPaletteTransform int
	DropSfxFrame              int
	CostMult                  int
	CostAdd                   int
	AddFn                     int
	QualityLevel              int
	Rarity                    int
	RequiredLevel             int
}
