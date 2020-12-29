package d2records

// UniqueItems stores all of the UniqueItemRecords
type UniqueItems map[string]*UniqueItemRecord

// UniqueItemRecord is a representation of a row from uniqueitems.txt
type UniqueItemRecord struct {
	Properties [12]*UniqueItemProperty

	Name                  string
	Code                  string // three letter code, points to a record in Weapons, Armor, or Misc
	TypeDescription       string
	UberDescription       string
	CharacterGfxTransform string // palette shift applied to this items gfx when held and when
	// on the ground (flippy). Points to a record in Colors.txt
	InventoryGfxTransform string // palette shift applied to the inventory gfx
	FlippyFile            string // if non-empty, overrides the base item's dropped gfx
	InventoryFile         string // if non-empty, overrides the base item's inventory gfx
	DropSound             string // if non-empty, overrides the base item's drop sound
	UseSound              string // if non-empty, overrides the sound played when item is used

	Version int // 0 = classic pre 1.07, 1 = 1.07-1.11, 100 = expansion
	Rarity  int // 1-255, higher is more common (ironically...)
	Level   int // item's level, can only be dropped by monsters / recipes / vendors / objects of this level or higher
	// otherwise they would drop a rare item with enhanced durability
	RequiredLevel  int // character must have this level to use this item
	CostMultiplier int // base price is multiplied by this when sold, repaired or bought
	CostAdd        int // after multiplied by above, this much is added to the price
	DropSfxFrame   int // if non-empty, overrides the base item's frame at which the drop sound plays

	Enabled bool // if false, this record won't be loaded (should always be true...)
	Ladder  bool // if true, can only be found on ladder and not in single player / tcp/ip
	NoLimit bool // if true, can drop more than once per game
	// (if false, can only drop once per game; if it would drop,
	//	 instead a rare item with enhanced durability drops)
	SingleCopy bool // if true, player can only hold one of these. can't trade it or pick it up
}

// UniqueItemProperty is describes a property of a unique item
type UniqueItemProperty = PropertyDescriptor
