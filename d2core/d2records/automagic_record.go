package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// AutoMagic has all of the AutoMagicRecords, used for generating magic properties for spawned items
type AutoMagic []*AutoMagicRecord

// AutoMagicRecord describes rules for automatically generating magic properties when spawning
// items
type AutoMagicRecord struct {
	// IncludeItemCodes
	// itype 1 to itype7
	// "Include Type" fields. You need to place item codes in any of these columns to allow that item
	// to receive mods from this row. See the note below.
	IncludeItemCodes [7]string

	// ModCode
	// They're the Property codes from Properties.txt.
	// These determine the actual properties which make up this autoprefix.
	// Each autoprefix can include up to three modifiers.
	ModCode [3]string

	// ExcludeItemCodes
	// etype 1 to etype3
	// 'Exclude type' . This field prevents certain mods from spawning on specific item codes.
	ExcludeItemCodes [3]string

	// ModParam, min, max
	// Parameter, min, and max values for the property
	ModParam [3]int
	ModMin   [3]int
	ModMax   [3]int

	// Name
	// String Comment Blizzard lists the equivalent prefix/affix here.
	// You can use what ever you wish here though. Handy for keeping track of groups.
	Name string

	// Version
	// it needs to be set to 0 if the prefix\affix you want to create or edit is going to be a
	// classic-only item ( with "classic" we mean "non-expansion" mode,
	// which you can toggle on and off when creating a character) or set to 100 if it's going to be
	// available in Expansion. This field is important,
	// as Items with " version" set to 100 will NOT be generated in Classic Diablo II.
	Version int

	// MinSpawnLevel
	// this field accepts numeric values and specifies the minimum level from which this autoprefix
	// can spawn. The column in question can be combined with the following maxlevel: to effectively
	// control groups of automods,
	// because you can use this field to combine multiple rows so that the autoprefixes are assigned
	// based on the level of the treasure drop [see below].
	MinSpawnLevel int

	// MaxSpawnLevel
	// this field accepts numeric values and specifies the maximum level beyond which the automod
	// stop spawning.
	MaxSpawnLevel int

	// LevelRequirement
	// It is the level requirement for this autoprefix.
	// This value is added to the Level Requirement of the item generated with this mod.
	LevelRequirement int

	// Class
	// the class type
	Class d2enum.Hero

	// ClassLevelRequirement
	// If class is set, this should allow a separate level requirement for this class.
	// This is a polite thing to do,
	// as other classes gain no advantage from class specific modifiers.
	// I am uncertain that this actually works.
	ClassLevelRequirement int

	// Frequency
	// For autoprefix groups, it states the chance to spawn this specific group member vs others.
	// Higher numbers means the automod will be more common. The 1.
	// 09 version file guide has some formuae relateing to this.
	Frequency int

	// Group
	// This field accepts numeric values and groups all the lines with the same values,
	// which are treated as a group. Only one autoprefix per group can be chosen,
	// and groups are influenced by levelreq, classlevelreq and frequency The 1.
	// 09 version file guide has a very nice tutorial about how to set up groups.
	// NOTE: The group number must also be entered in the 'auto prefix' column of each entry in
	// Weapons.txt or Armor.txt in order for the property to appear.
	Group int

	// PaletteTransform
	// If transform is set to 1 then the item will be colored with the chosen color code,
	// taken from Colors.txt
	PaletteTransform int

	// CostDivide
	// Numeric value that acts as divisor for the item price.
	CostDivide int

	// CostMultiply
	// Numeric value that acts as multiplier for the item price.
	CostMultiply int

	// CostAdd
	// Numeric value that acts as a flat sum added to the item price.
	CostAdd int

	// Spawnable
	// It is a boolean type field, and states if this autoprefix can actually spawn in the game.
	// You can disable this row by setting it to 0 , or enable it by setting it to 1
	Spawnable bool

	// SpawnOnRare
	// It decides whether this autoprefix spawns on rare quality items or not.
	// You can prevent that from happening by setting it to 0 , or you can allow it by setting it to 1
	SpawnOnRare bool

	// transform
	// It is a boolean value whichallows the colorization of the items.
	Transform bool
}
